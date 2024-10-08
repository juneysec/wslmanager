package http

import (
	"fmt"
	"os/exec"
	"path"
	"path/filepath"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Serve(port int, isDebug bool) {
	if !isDebug {
		gin.SetMode(gin.ReleaseMode)
		exec.Command("cmd.exe", "/C", "start", fmt.Sprintf("http://localhost:%d", port)).Start()
	}

	r := gin.Default()
	if gin.IsDebugging() {
		r.Use(cors.New(cors.Config{
			// 許可したいHTTPメソッドの一覧
			AllowMethods: []string{
				"POST",
				"GET",
				"OPTIONS",
				"PUT",
				"DELETE",
			},
			// 許可したいHTTPリクエストヘッダの一覧
			AllowHeaders: []string{
				"Access-Control-Allow-Headers",
				"Content-Type",
				"Content-Length",
				"Accept-Encoding",
				"X-CSRF-Token",
				"Authorization",
			},
			AllowOriginFunc: func(origin string) bool {
				// すべて許可
				return true
			},
			// preflight requestで許可した後の接続可能時間
			// https://godoc.org/github.com/gin-contrib/cors#Config の中のコメントに詳細あり
			MaxAge: 24 * time.Hour,
		}))
	}

	RegisterRoutes(r)

	r.Static("/assets", "./dist/assets")
	r.NoRoute(func(c *gin.Context) {
		_, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		//ディレクトリアクセス（ファイル名がない）かパスクエリ（拡張子がない）
		if file == "" || ext == "" {
			c.File("./dist/index.html")
		} else {
			fmt.Println(c.Request.RequestURI)
			fmt.Println("./dist" + c.Request.RequestURI)
			c.File("./dist" + c.Request.RequestURI)
		}
	})

	r.Run(fmt.Sprintf(":%d", port))
}
