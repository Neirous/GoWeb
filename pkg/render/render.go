package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Neirous/GoWeb/pkg/config"
	"github.com/Neirous/GoWeb/pkg/models"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate 渲染指定的模板，并将其写入 HTTP 响应中，tmpl表示页面模板，layout表示该页面模板要组合什么布局模板
func RenderTemplate(w http.ResponseWriter, tmpl string, layout string, td *models.TemplateData) {

	var tc map[string]*template.Template
	if app.UseCache {
		// 使用已保存的缓存
		tc = app.TemplateCache
	} else {
		//重新解析模板
		tc, _ = CreateTemplateCache()
	}

	// 获取请求的页面模板
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from cache")
	}

	// 创建一个缓冲区以捕获模板执行结果
	buf := new(bytes.Buffer)

	// 执行模板，并指定要使用的布局模板
	err := t.ExecuteTemplate(buf, layout, td)
	if err != nil {
		log.Println("Error executing template:", err)
	}

	// 将渲染后的内容写入HTTP响应
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser:", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("../../templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		// 解析页面模板和基础布局模板
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// 加载所有布局模板，并将其解析到页面模板中
		layoutFiles, err := filepath.Glob("../../templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(layoutFiles) > 0 {
			ts, err = ts.ParseFiles(layoutFiles...)
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}
