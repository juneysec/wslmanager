// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * WSLManager API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 */

package schema

type RequestDistributionPost struct {

  // ディストリビューションの名前
  Name string `json:"name"`

  // VHDファイルの作成場所
  VhdPath string `json:"vhdPath,omitempty"`

  // インポート元のパス
  SourcePath string `json:"sourcePath,omitempty"`
}