package main

import (
	"net/http"

	"github.com/Neirous/GoWeb/pkg/config"
	"github.com/Neirous/GoWeb/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	//创建chi路由器
	mux := chi.NewRouter()

	//中间件使用
	//用于恢复因 panic 而崩溃的 HTTP 请求
	mux.Use(middleware.Recoverer)
	//Nosurf 用于防止 CSRF 攻击
	mux.Use(Nosurf)

	mux.Use(SessionLoad)

	//HTTP请求处理
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
