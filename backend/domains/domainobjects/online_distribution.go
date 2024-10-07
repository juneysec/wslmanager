package domainobjects

import "fmt"

// ディストリビューション
type OnlineDistribution struct {
	// 名前
	Name string `json:"name"`

	// 名前
	FriendlyName string `json:"friendlyName"`

	// インストール済み
	AlreadyInstalled bool `json:"alreadyInstalled"`
}

// インスタンス生成
func NewOnlineDistribution(name string, friendlyName string, alreadyInstalled bool) (*OnlineDistribution, error) {
	return &OnlineDistribution{
		Name:      name,
		FriendlyName: friendlyName,
		AlreadyInstalled: alreadyInstalled,
	}, nil
}

// 文字列化
func (p *OnlineDistribution) String() string {
	return fmt.Sprintf("online(\"%v\" \"%v\" \"%v\")", p.Name, p.FriendlyName, p.AlreadyInstalled)
}
