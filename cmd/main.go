package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivan3bx/kpie"
)

func main() {
	h := kpie.NewHandler()

	r := gin.Default()
	configureAssets(r)
	configureTemplates(r)

	r.GET("/", h.Index)

	if r.Run("127.0.0.1:8080") != nil {
		panic("Unable to start server")
	}
}

func configureAssets(r *gin.Engine) {
	targetFS := kpie.AssetsFS()
	r.StaticFS("/assets", http.FS(targetFS))
}

func configureTemplates(r *gin.Engine) {
	if gin.Mode() == gin.DebugMode {
		r.LoadHTMLGlob("templates/*")
	} else {
		templ := template.Must(template.New("").ParseFS(kpie.TemplatesFS(), "*"))
		r.SetHTMLTemplate(templ)
	}
}
