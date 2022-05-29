package router

import (
	"blog-go/views"
	"net/http"
	_ "time"
)

func Router() {
	// 1. 页面 views
	// 2. 数据（json）
	// 3. 静态资源
	http.HandleFunc("/", views.HTML.Index)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))

}
