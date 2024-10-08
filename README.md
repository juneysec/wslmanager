# WSL Manager

Windows で WSL を GUI で管理するツール。ブラウザで動作。  
デフォルトでローカルポート 7711 を使用

## 起動方法

wslmanager.exe を実行

パラメータ  
| 引数 | 説明 |
| -----| -------- |
| -debug | デバッグ(開発)モードで起動。gin がデバッグモード、ポートは常に 7711、ブラウザを起動しない |
| -port | 使用するポート番号。デフォルトは 7711。一度指定すると以降はそのポート番号で起動する($HOME/.wslmanager.json に保存) |
| -help | ヘルプ表示 |

## ビルド方法

### 必要なもの

- [VSCode](https://code.visualstudio.com/)
- [Docker Desktop](https://www.docker.com/get-started/)
- [Go](https://go.dev/)

### Go のライセンスリスト JSON ファイルの更新

[ライセンスリスト作成](./backend/docs/ラインセスリスト作成.md) を参考に作成  
frontend/vue/src/assets に保存する

```shell
go-licenses report . --template ./license.tpl > ../frontend/vue/src/assets/go-licenses.json
```

できたファイルは UTF-16 なので UTF-8 に変換する必要あり。

### フロントエンドのビルド

.\frontend\vue を VSCode で開いてリモートコンテナで開く。  
※run_vscode.cmd 実行で backend と frontend の両方を開ける。

VSCode のターミナルから下記を実行して、Node.js のライセンスファイル更新

```shell
npx license-checker --json > src/assets/node-js-licenses.json
```

ビルド

```shell
npm run build
```

dist にできる。backend/dist がマウントされていてそっちにできているはず。

### バックエンドのビルド

配布用のライセンスファイルを取得(バージョンは適宜変更)

```shell
$version = '1.0.0'
go-licenses save . --save_path="../release/WSLManager-$version/licenses"
go build -o ..\release\WSLManager-$version\WSLManager.exe
xcopy /s /y .\dist ..\release\WSLManager-$version\dist\
```
