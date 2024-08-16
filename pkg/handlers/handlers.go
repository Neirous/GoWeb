package handlers

import (
	"net/http"

	"github.com/Neirous/GoWeb/pkg/config"
	"github.com/Neirous/GoWeb/pkg/models"
	"github.com/Neirous/GoWeb/pkg/render"
)

// Repo 是用于处理器的全局存储库实例
var Repo *Repository

// Repository 是存储库类型，用于存储应用程序配置
type Repository struct {
	App *config.AppConfig
}

// NewRepo 创建一个新的存储库实例并返回
// 它将应用程序的配置传递给存储库，以便处理器可以使用这些配置
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers 设置全局存储库实例
// 这允许处理器访问存储库中的数据和方法
func NewHandlers(r *Repository) {
	Repo = r
}

// Home 处理器函数，用于处理首页请求
// 它会创建一个字符串映射，并将其与模板数据一起传递给渲染函数
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	// 创建一个字符串映射
	stringMap := make(map[string]string)
	stringMap["test"] = "hello world"
	stringMap["remote_ip"] = remoteIp

	// 调用 RenderTemplate 函数渲染 home.page.tmpl 模板
	render.RenderTemplate(w, "home.page.tmpl", "base", &models.TemplateData{
		StringMap: stringMap,
	})
}

// About 处理器函数，用于处理关于页面的请求
// 这里不传递任何额外数据，只渲染 about.page.tmpl 模板
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	// 创建一个字符串映射
	stringMap := make(map[string]string)
	stringMap["test"] = "hello world"
	stringMap["remote_ip"] = remoteIp
	render.RenderTemplate(w, "about.page.tmpl", "base", &models.TemplateData{
		StringMap: stringMap,
	})

}
