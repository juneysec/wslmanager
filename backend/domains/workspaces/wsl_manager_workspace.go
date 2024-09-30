// WSL管理ツールのワークスペース
package workspaces

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os/exec"
	"regexp"
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
	return &WSLManagerWorkspace{}, nil
}

// ディストリビューションリストの読み込み。
//
// p.Distributions にディストリビューションリストが格納される
func (p *WSLManagerWorkspace) Fetch() error {
	re, err := regexp.Compile(`(\*?)[\s]+([^\s]+)[\s]+([^\s]+)[\s]+([^\s]+)`)
	if err != nil {
		log.Println("failed to compile regexp:", err)
		return err
	}

	output, err := winexec("wsl.exe", "-l", "-v")
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

// コマンド実行
// Windows の場合、StdOut が UTF-16 になるのでUTF-8に変換して返す。
func winexec(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)

	output, err := cmd.Output()
	if err != nil {
		log.Printf("failed to call exec.Command(%v).Output: %v", name, err)
		return "", err
	}

	utf16Reader := transform.NewReader(bytes.NewReader(output), unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder())
	var buf bytes.Buffer
	_, err = io.Copy(&buf, utf16Reader)
	if err != nil {
		log.Printf("failed to call io.Copy(): %v", err)
		return "", err
	}

	return buf.String(), nil
}
