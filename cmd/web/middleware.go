package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// Nosurf 是一个防止跨站请求伪造（CSRF）攻击的中间件。
// 它接收一个 http.Handler 并返回一个新的 http.Handler，
// 这个新 Handler 会自动为每个请求生成并验证 CSRF 令牌。

func Nosurf(next http.Handler) http.Handler {
	// 创建一个新的 nosurf CSRF 处理器，使用传入的处理器作为基础。
	csrfHandler := nosurf.New(next)

	// 设置 CSRF 处理器的 Cookie 配置
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,                 // 使 Cookie 仅能通过 HTTP(S) 协议传输，无法通过 JavaScript 访问，增加安全性。
		Path:     "/",                  // 设置 Cookie 的作用路径为网站根路径，适用于所有路径。
		Secure:   false,                // 仅在 HTTPS 下传输 Cookie，如果是 false 则可以在 HTTP 下传输（生产环境应设置为 true）。
		SameSite: http.SameSiteLaxMode, // 设定 Cookie 的 SameSite 属性为 Lax，防止跨站请求，保护用户数据。
	})
	// 返回设置好的 CSRF 处理器。
	return csrfHandler
}
