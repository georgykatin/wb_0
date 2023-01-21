package httpserver

import (
	"embed"
	"github.com/gin-gonic/gin"
	"io/fs"
	"log"
	"net/http"
)

func (serv *Server) RunServer() {

	go func() {
		serv.gin.LoadHTMLFiles("front/index.html")
		//OrderRoot := Folder(front.Content, "static")
		//serv.gin.Use(static.Serve("/wb", OrderRoot))
		serv.gin.GET("/wb", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index.html", gin.H{
				"title": "main page",
			})
		})
		serv.gin.GET("/test", func(c *gin.Context) {
			c.String(200, "test")
		})
		if err := serv.gin.Run(); err != nil {
			log.Printf("Error :%v", err)
		}
	}()
}

type FileSystem struct {
	http.FileSystem
}

func (f *FileSystem) Exist(path string) bool {
	_, err := f.Open(path)
	if err != nil {
		return false
	}
	return true
}

func Folder(FsEmbed embed.FS, target string) FileSystem {
	fsys, err := fs.Sub(FsEmbed, target)
	if err != nil {
		panic(err)
	}
	return FileSystem{
		FileSystem: http.FS(fsys),
	}
}
