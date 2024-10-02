# Node.js 開発用の Dockerfile, docker-compose.yml

主にWindowsのDocker Desktopで使用を想定。  
ユーザー名は developer  

src フォルダを作って bind する。  
下記バッチ起動でビルド
```cmd
00.build.cmd
```

下記バッチ起動でサービスアップ
```cmd
01.up.cmd
```

下記バッチ起動でシェル起動
```cmd
03.exec_shell.cmd
```

下記バッチ起動でOpenAPI のクライアント用スキーマファイル生成
```cmd
03.exec_shell.cmd
```
