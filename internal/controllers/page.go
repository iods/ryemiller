package controllers

import "html/template"

var tpl = template.Must(template.ParseFiles(
	"web/template/base.gohtml",
	"web/template/index.gohtml",
	"web/template/partials/head.gohtml",
	"web/template/page/home.gohtml",
	"web/template/page/error.gohtml",
))

type indexView struct {
	PageView  string
	PageTitle string
}

type errorView struct {
	 PageView  string
	 PageTitle string
}