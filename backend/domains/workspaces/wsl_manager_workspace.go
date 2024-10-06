// WSL管理ツールのワークスペース
package workspaces

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
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

	_, output, err := winexec("wsl.exe", "-l", "-v")
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

	return nil
}

func (p *WSLManagerWorkspace) Find(name string) (*domainobjects.Distribution, error) {
	idx := slices.IndexFunc(p.Distributions, func (d *domainobjects.Distribution) bool { return d.Name == name}) 
	if idx < 0 {
		return nil, fmt.Errorf("ディストリビューション[%s] は存在しません。", name)
	}

	return p.Distributions[idx], nil
}

func (p *WSLManagerWorkspace) Start(name string) error {
	exitCode, output, err := winexec("wsl.exe", "-d", name, "-e", "echo", "hello")
	if err != nil {
		return err
	}

	if exitCode != 0 {
		return fmt.Errorf("wsl.exe returns %v.\noutput is:\n%s", exitCode, output)
	}

	return nil
}

func (p *WSLManagerWorkspace) Stop(name string) error {
	exitCode, output, err := winexec("wsl.exe", "-t", name)
	if err != nil {
		return err
	}

	if exitCode != 0 {
		return fmt.Errorf("wsl.exe returns %v.\noutput is:\n%s", exitCode, output)
	}

	return nil
}

func (p *WSLManagerWorkspace) ExecShell(name string) error {
	exitCode, output, err := wincmd("cmd", "/C", "start", "wsl.exe", "-d", name)
	if err != nil {
		return err
	}

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

	exitCode, output, err := wincmd("cmd", "/C", "start", "wsl.exe", "--export", name, path)
	if err != nil {
		return err
	}

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

	exitCode, output, err := winexec("cmd", "/C", "start", "wsl.exe", "--import", name, vhdPath, sourcePath, "--version", "2")
	if err != nil {
		return err
	}

	if exitCode != 0 {
		return fmt.Errorf("wsl.exe returns %v.\noutput is:\n%s", exitCode, output)
	}
	
	return nil
}

func (p *WSLManagerWorkspace) SetDefault(name string) error {
	exitCode, output, err := winexec("wsl.exe", "--set-default", name)
	if err != nil {
		return err
	}

	if exitCode != 0 {
		return fmt.Errorf("wsl.exe returns %v.\noutput is:\n%s", exitCode, output)
	}

	return nil
}

func (p *WSLManagerWorkspace) Unregister(name string) error {
	exitCode, output, err := winexec("wsl.exe", "--unregister", name)
	if err != nil {
		return err
	}

	if exitCode != 0 {
		return fmt.Errorf("wsl.exe returns %v.\noutput is:\n%s", exitCode, output)
	}

	return nil
}

// コマンド実行
// Windows の場合、StdOut が UTF-16 になるのでUTF-8に変換して返す。
func winexec(name string, arg ...string) (int, string, error) {
	cmd := exec.Command(name, arg...)

	output, err := cmd.Output()
	if err != nil {
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

	return cmd.ProcessState.ExitCode(), buf.String(), nil
}

// コマンド実行
// Windows の場合、StdOut が UTF-16 になるのでUTF-8に変換して返す。
func wincmd(name string, arg ...string) (int, string, error) {
	cmd := exec.Command(name, arg...)
	err := cmd.Start()
	if err != nil {
		return -1, "実行に失敗しました。", err
	}

	return 0, "", nil
}
