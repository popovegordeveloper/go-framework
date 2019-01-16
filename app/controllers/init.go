package controllers

import (
	"github.com/qor/render"
	"html/template"
	"net/http"
)

var Render *render.Render

func init() {
	Render = render.New(&render.Config{
		ViewPaths:     []string{"resources/views"},
		DefaultLayout: "app",
		FuncMapMaker: func(*render.Render, *http.Request, http.ResponseWriter) template.FuncMap {
			return template.FuncMap{}
		},
	})
}
