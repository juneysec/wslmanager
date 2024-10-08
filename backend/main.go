package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"wslmanager/http"
)

type Config struct {
	Port int `json:"port"`
}

func readConfig() *Config {
	defaultConfig := Config { Port:7711 }

	// ホームディレクトリのパスを取得
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("ホームディレクトリを取得できませんでした:", err)
		return &defaultConfig
	}

	// 設定ファイルである .wslmanager.json のパスを組み立てる
	configPath := filepath.Join(homeDir, ".wslmanager.json")

	// ファイルを開く
	file, err := os.Open(configPath)
	if err != nil {
		return &defaultConfig
	}
	defer file.Close()

	// JSON データを読み込む
	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		fmt.Println(".wslmanager.json データの解析に失敗しました:", err)
		return &defaultConfig
	}

	return &config
}

func writeConfig(c *Config) {
	// ホームディレクトリのパスを取得
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("ホームディレクトリを取得できませんでした:", err)
		return
	}

	// 設定ファイルである .wslmanager.json のパスを組み立てる
	configPath := filepath.Join(homeDir, ".wslmanager.json")
	// ファイルを上書きするために再度開く（書き込み用）
	file, err := os.Create(configPath)
	if err != nil {
		fmt.Println(".wslmanager.json ファイルを作成できませんでした:", err)
		return
	}
	defer file.Close()

	// JSON データを書き込む
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // JSON を見やすく整形
	if err := encoder.Encode(c); err != nil {
		fmt.Println("JSON データの書き込みに失敗しました:", err)
		return
	}
}

func main() {
	config := readConfig()

	// コマンドライン引数解析
	isDebugPtr := flag.Bool("debug", false, "デバッグ(開発)モードで起動する場合に指定")
	portPtr := flag.Int("port", config.Port, "ポート番号")
	flag.Parse()

	fmt.Printf("isDebug: %v, port: %v", *isDebugPtr, *portPtr)

	port := *portPtr

	// リリースモードの場合はポートを保存
	if (!*isDebugPtr) {
		config.Port = *portPtr
		writeConfig(config)
	} else {
		// デバッグモードの場合は常に 7711 で起動
		port = 7711
	}

	http.Serve(port, *isDebugPtr)
}
