<!-- VSCode:Ctrl+k, v -->
# open-api-go-gin-server
Go の Gin ＋ OpenAPI Generator 

## コード生成部分のカスタマイズ

下記を適宜修正する

- src/src/main/java/com/my/company/codegen/GoCustomServerGenerator.java
- src/src/main/resources/go-custom-server/*.mustache

ビルド
```shell
cd ./src
mvn package -DskipTests=true
```


## コード生成

--type-mappings number=float64 を指定しないとfloat32 になってしまうので注意。  
json.Number を使うほうが良いかも。  
```shell
java -cp ./openapi-generator-cli.jar:./src/target/go-custom-server-openapi-generator-1.0.0.jar org.openapitools.codegen.OpenAPIGenerator generate -g go-custom-server -i ./src/openapi/openapi.yaml -o ./src/gen --type-mappings number=float64
```

## 参考
https://zenn.dev/ysk1to/books/248fad8cb34abe/viewer/a7fab2
