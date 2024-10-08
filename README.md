# WSL Manager

Windows で WSL を GUI で管理するツール。ブラウザで動作。  
ローカルポート 8080 を使用

## 起動方法

wslmanager.exe を起動するだけ。

## ビルド方法

### 必要なもの

- [VSCode](https://code.visualstudio.com/)
- [Docker Desktop](https://www.docker.com/get-started/)
- [Go](https://go.dev/)


### Go のライセンスリストJSONファイルの更新

[ライセンスリスト作成](./backend/docs/ラインセスリスト作成.md) を参考に作成  
frontend/vue/src/assets に保存する

```shell
go-licenses report . --template ./license.tpl > ../frontend/vue/src/assets/go-licenses.json
```

できたファイルは UTF-16 なので UTF-8 に変換する必要あり。

### フロントエンドのビルド

.\frontend\vue を VSCode で開いてリモートコンテナを開く。  
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
go-licenses save . --save_path="./release/WSLManager-0.0.2/licenses"
go build -o .\release\WSLManager-0.0.2\WSLManager.exe
xcopy /s /y .\dist .\release\WSLManager-0.0.2\dist\
```
