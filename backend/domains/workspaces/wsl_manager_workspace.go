// WSL管理ツールのワークスペース
package workspaces

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"wslmanager/domains/domainobjects"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// WSL管理ツールのメインワークスペース
type WSLManagerWorkspace struct {
	Distributions []*domainobjects.Distribution
}

// WSL管理ツールのワークスペース生成
func NewWSLManagerWorkspace() (*WSLManagerWorkspace, error) {
	ret := &WSLManagerWorkspace{}
	err := ret.Fetch()

	if err != nil {
		return nil, err
	} else {
		return ret, err
	}
}

// ディストリビューションリストの読み込み。
//
// p.Distributions にディストリビューションリストが格納される
func (p *WSLManagerWorkspace) Fetch() error {
	p.Distributions = p.Distributions[:0]
	re, err := regexp.Compile(`(\*?)[\s]+([^\s]+)[\s]+([^\s]+)[\s]+([^\s]+)`)
	if err != nil {
		log.Println("failed to compile regexp:", err)
		return err
	}

	_, output, err := execUTF16("wsl.exe", "-l", "-v")
	if err != nil {
		log.Println("failed to call winexec:", err)
		return err
	}

	scanner := bufio.NewScanner(strings.NewReader(output))
	if scanner.Scan() { // ヘッダスキップ
		for scanner.Scan() {
			line := scanner.Text()

			matches := re.FindStringSubmatch(line)
			if len(matches) > 0 {
				distribution, err := domainobjects.NewDistribution(matches[1] == "*", matches[2], matches[3], matches[4])
				if err != nil {
					return err
				}

				p.Distributions = append(p.Distributions, distribution)
			}
		}
	}

	return p.fetchVHDPath()
}

func (p *WSLManagerWorkspace) Find(name string) (*domainobjects.Distribution, error) {
	idx := slices.IndexFunc(p.Distributions, func(d *domainobjects.Distribution) bool { return d.Name == name })
	if idx < 0 {
		return nil, fmt.Errorf("ディストリビューション[%s] は存在しません。", name)
	}

	return p.Distributions[idx], nil
}

func (p *WSLManagerWorkspace) Start(name string) error {
	exitCode, output, _ := execUTF16("wsl.exe", "-d", name, "-e", "echo", "hello")
	if exitCode != 0 {
		return fmt.Errorf("wsl.exe returns %v.\noutput is:\n%s", exitCode, output)
	}

	return nil
}

func (p *WSLManagerWorkspace) Stop(name string) error {
	exitCode, output, _ := execUTF16("wsl.exe", "-t", name)
	if exitCode != 0 {
		return fmt.Errorf("wsl.exe returns %v.\noutput is:\n%s", exitCode, output)
	}

	return nil
}

func (p *WSLManagerWorkspace) ExecShell(name string) error {
	exitCode, output, _ := execNoWaitUTF16("cmd", "/C", "start", "wsl.exe", "-d", name)
	if exitCode != 0 {
		return fmt.Errorf("wsl.exe returns %v.\noutput is:\n%s", exitCode, output)
	}

	return nil
}

func (p *WSLManagerWorkspace) Export(name string, path string) error {
	// 拡張子を取得。.tar ではない場合、.tar を付与する
	re := regexp.MustCompile(`(?i)tar\.gz$`)
	if !re.MatchString(path) {
		path += ".tar.gz"
	}

	exitCode, output, _ := execNoWaitUTF16("cmd", "/C", "start", "wsl.exe", "--export", name, path)
	if exitCode != 0 {
		return fmt.Errorf("wsl.exe returns %v.\noutput is:\n%s", exitCode, output)
	}

	return nil
}

func (p *WSLManagerWorkspace) Import(name string, vhdPath string, sourcePath string) error {
	// ソースファイル存在チェック
	if _, err := os.Stat(sourcePath); os.IsNotExist((err)) {
		return err
	}

	// VHD のパスが存在する場合
	if _, err := os.Stat(vhdPath); err == nil {
		return fmt.Errorf("VHDパス[%v]は既に存在します。", vhdPath)
	}

	exitCode, output, _ := execUTF16("cmd", "/C", "start", "wsl.exe", "--import", name, vhdPath, sourcePath, "--version", "2")
	if exitCode != 0 {
		return fmt.Errorf("wsl.exe returns %v.\noutput is:\n%s", exitCode, output)
	}

	return nil
}

func (p *WSLManagerWorkspace) SetDefault(name string) error {
	exitCode, output, err := execUTF16("wsl.exe", "--set-default", name)
	if err != nil {
		return err
	}

	if exitCode != 0 {
		return fmt.Errorf("wsl.exe returns %v.\noutput is:\n%s", exitCode, output)
	}

	return nil
}

func (p *WSLManagerWorkspace) MoveVHD(name string, path string) error {
	if folderExists(path) {
		if size, _ := getFolderSize(path); size > 0 {
			return fmt.Errorf("フォルダ[%v]はすでに存在します。", path)
		}
	}

	exitCode, output, _ := execUTF16("wsl.exe", "--manage", name, "--move", path)
	if exitCode != 0 {
		return fmt.Errorf("wsl.exe returns %d.\n%s", exitCode, output)
	}

	return nil
}

func (p *WSLManagerWorkspace) OpenVHD(name string) error {
	dist, err := p.Find(name)
	if err != nil {
		return err
	}

	// フォルダ存在チェック
	if !folderExists(dist.VhdPath) {
		return fmt.Errorf("VHDパス[%v]が存在しません。", dist.VhdPath)
	}

	execNoWaitUTF16("explorer.exe", dist.VhdPath)
	return nil
}

// ディストリビューションの削除
func (p *WSLManagerWorkspace) Unregister(name string) error {
	exitCode, output, err := execUTF16("wsl.exe", "--unregister", name)
	if err != nil {
		return err
	}

	if exitCode != 0 {
		return fmt.Errorf("wsl.exe returns %v.\noutput is:\n%s", exitCode, output)
	}

	return nil
}

// オンラインディストリビューションのリストを取得
func (p *WSLManagerWorkspace) GetOnlineDistributions() ([]*domainobjects.OnlineDistribution, error) {
	reHeader, _ := regexp.Compile(`(NAME)[\s]+(FRIENDLY NAME)`)
	re, _ := regexp.Compile(`([^\s]+)[\s]+(.+)`)

	_, output, err := execUTF16("wsl.exe", "-l", "-o")
	if err != nil {
		log.Println("failed to call winexec:", err)
		return nil, err
	}

	var retval []*domainobjects.OnlineDistribution
	scanner := bufio.NewScanner(strings.NewReader(output))

	for scanner.Scan() {
		line := scanner.Text()
		matches := reHeader.FindStringSubmatch(line)
		if len(matches) > 0 {
			break
		}
	}

	for scanner.Scan() {
		line := scanner.Text()

		matches := re.FindStringSubmatch(line)
		if len(matches) > 0 {
			distribution := matches[1]
			friendlyName := matches[2]
			alreadyInstalled := false

			if found, _ := p.Find(distribution); found != nil {
				alreadyInstalled = true
			}

			od, err := domainobjects.NewOnlineDistribution(distribution, friendlyName, alreadyInstalled)
			if err != nil {
				return nil, err
			}

			retval = append(retval, od)
		}
	}

	return retval, nil
}

func (p *WSLManagerWorkspace) InstallOnlineDistribution(name string) error {
	exitCode, output, err := execUTF16("cmd.exe", "/C", "start", "wsl.exe", "--install", name)
	if err != nil {
		return err
	}

	if exitCode != 0 {
		return fmt.Errorf("wsl.exe returns %v.\noutput is:\n%s", exitCode, output)
	}

	return nil
}

// VHDファイル情報の取得
func (p *WSLManagerWorkspace) fetchVHDPath() error {
	re, err := regexp.Compile(`([^\s]+)[\s]+([^\s]+)`)
	if err != nil {
		log.Println("failed to compile regexp:", err)
		return err
	}

	///_, output, err := execUTF8("powershell.exe", "-Command", "chcp 65001; Write-Host 'ほげ'")
	_, output, err := execUTF8("powershell.exe", "-Command", "chcp 65001; (Get-ChildItem -Path HKCU:\\Software\\Microsoft\\Windows\\CurrentVersion\\Lxss | Select-Object @{ Name = 'DistributionName'; Expression={$_.GetValue('DistributionName')} }, @{ Name = 'Path'; Expression={$_.GetValue('BasePath')} })")
	if err != nil {
		log.Println("failed to call winexec:", err)
		return err
	}

	scanner := bufio.NewScanner(strings.NewReader(output))
	if scanner.Scan() { // 1行目スキップ
		for scanner.Scan() {
			line := scanner.Text()

			matches := re.FindStringSubmatch(line)
			if len(matches) > 0 {
				distributionName := matches[1]
				vhdPath := matches[2]

				if folderExists(vhdPath) {
					if found, err := p.Find(distributionName); err == nil {
						found.VhdPath = vhdPath

						if size, err := getFolderSize(vhdPath); err == nil {
							found.VhdSize = size
						}
					}
				}
			}
		}
	}

	return nil
}

// フォルダの存在チェック
func folderExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}

// フォルダのファイルサイズ計算
func getFolderSize(path string) (int64, error) {
	var totalSize int64

	// WalkDirはディレクトリ内のすべてのファイルとフォルダを再帰的に処理します
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// ファイルであれば、そのサイズを合計に加える
		if !info.IsDir() {
			totalSize += info.Size()
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return totalSize, nil
}

// コマンド実行
// Windows の場合、StdOut が UTF-16 になるのでUTF-8に変換して返す。
func execUTF8(name string, arg ...string) (int, string, error) {
	cmd := exec.Command(name, arg...)

	output, err := cmd.Output()
	var exitError *exec.ExitError
	if err != nil && !errors.As(err, &exitError) {
		log.Printf("failed to call exec.Command(%v).Output: %v", name, err)
		return -1, "", err
	}

	utf16Reader := transform.NewReader(bytes.NewReader(output), unicode.UTF8.NewDecoder())
	var buf bytes.Buffer
	_, err = io.Copy(&buf, utf16Reader)
	if err != nil {
		log.Printf("failed to call io.Copy(): %v", err)
		return -1, "", err
	}

	var exitCode = -1
	if exitError == nil {
		exitCode = cmd.ProcessState.ExitCode()
	}

	return exitCode, buf.String(), nil
}

// コマンド実行
// Windows の場合、StdOut が UTF-16 になるのでUTF-8に変換して返す。
func execUTF16(name string, arg ...string) (int, string, error) {
	cmd := exec.Command(name, arg...)

	output, err := cmd.Output()
	var exitError *exec.ExitError
	if err != nil && !errors.As(err, &exitError) {
		log.Printf("failed to call exec.Command(%v).Output: %v", name, err)
		return -1, "", err
	}

	utf16Reader := transform.NewReader(bytes.NewReader(output), unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder())
	var buf bytes.Buffer
	_, err = io.Copy(&buf, utf16Reader)
	if err != nil {
		log.Printf("failed to call io.Copy(): %v", err)
		return -1, "", err
	}

	var exitCode = -1
	if exitError == nil {
		exitCode = cmd.ProcessState.ExitCode()
	}

	return exitCode, buf.String(), nil
}

// コマンド実行(終了を待機しない)
// Windows の場合、StdOut が UTF-16 になるのでUTF-8に変換して返す。
func execNoWaitUTF16(name string, arg ...string) (int, string, error) {
	cmd := exec.Command(name, arg...)
	err := cmd.Start()
	if err != nil {
		return -1, "実行に失敗しました。", err
	}

	return 0, "", nil
}
