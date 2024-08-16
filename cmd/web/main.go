package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Neirous/GoWeb/pkg/config"
	"github.com/Neirous/GoWeb/pkg/handlers"
	"github.com/Neirous/GoWeb/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

// 创建应用配置
var app config.AppConfig

var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	//UseCache为false时说明此时为开发者模式，每次运行都会重新解析模板
	//置为true说明为生产模式，只有服务器启动时会解析一次，假如刷新网页
	//并不会重新解析模板，只会从缓存中获取模板进行渲染
	app.UseCache = false

	//创建模板缓存
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc

	//初始化模板渲染
	render.NewTemplates(&app)

	//创建存储库，并将应用配置传递给它
	repo := handlers.NewRepo(&app)
	//设置全局存储库
	handlers.NewHandlers(repo)

	//启动HTTP服务
	fmt.Printf(fmt.Sprintf("Starting application on port %s", portNumber))
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = serve.ListenAndServe()
	log.Fatal(err)
}
