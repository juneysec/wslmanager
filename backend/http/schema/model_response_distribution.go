// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * WSLManager API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 */

package schema

type ResponseDistribution struct {

  // デフォルトの場合true
  IsDefault bool `json:"isDefault"`

  // ディストリビューション名
  Name string `json:"name"`

  // 状態
  State string `json:"state"`

  // WSLのバージョン
  Version string `json:"version"`

  // VHDファイルのベースパス
  VhdPath string `json:"vhdPath,omitempty"`

  // VHDファイルのサイズ
  VhdSize float64 `json:"vhdSize,omitempty"`
}