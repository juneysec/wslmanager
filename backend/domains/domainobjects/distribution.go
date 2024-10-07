// WSLのディストリビューション定義
package domainobjects

import "fmt"

// ディストリビューション
type Distribution struct {
	// デフォルトの場合 true
	IsDefault bool `json:"isDefault"`

	// 名前
	Name string `json:"name"`

	// 状態(running/stop)
	State string `json:"state"`

	// バージョン
	Version string `json:"version"`

	// VHDのパス
	VhdPath string `json:"vhdPath"`
}

// インスタンス生成
func NewDistribution(isDefault bool, name string, state string, version string) (*Distribution, error) {
	return &Distribution{
		IsDefault: isDefault,
		Name:      name,
		State:     state,
		Version:   version,
		VhdPath:   "",
	}, nil
}

// 文字列化
func (p *Distribution) String() string {
	defaultStr := " "
	if p.IsDefault {
		defaultStr = "*"
	}

	return fmt.Sprintf("%v %v %v %v %v", defaultStr, p.Name, p.State, p.Version, p.VhdPath)
}
