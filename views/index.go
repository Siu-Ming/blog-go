package views

import (
	"blog-go/common"
	"blog-go/service"
	"errors"
	"log"
	"net/http"
	"strconv"
	_ "time"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index

	// 分页
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败：", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员"))
		return
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	// 每页显示的数量
	pageSize := 10

	hr, err := service.GetAllIndexInfo(page, pageSize)
	if err != nil {
		log.Println("首页获取数据出错：", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员"))
	}

	index.WriteData(w, hr)
}
