package mux

import (
	"net/http"
)

// Mux 根据域名调用不同的handle
type Mux struct {
	allStd   map[string]http.Handler
	firstStd http.Handler
}

func New() Mux {
	return Mux{
		allStd: make(map[string]http.Handler),
	}
}

// AddStd 添加特定域名的handle
// 如果通过ip访问网站，使用第一次添加的handle
func (m *Mux) AddStd(domain string, handle http.Handler) {
	if m.firstStd == nil {
		m.firstStd = handle
	}
	m.allStd[domain] = handle
}

func (m Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 根据域名调用对应的handle
	if handle := m.allStd[r.Host]; handle != nil {
		handle.ServeHTTP(w, r)
		return
	}
	// 如果之前的vps搭建的是一个ip对应一个域名的网站，通过https://ip可以访问网站
	// 在变成vps搭建的是一个ip对应多个域名的网站后，通过下面的代码，
	// 配合先添加旧域名网站的handle，保留https://ip可以访问旧网站的行为
	m.firstStd.ServeHTTP(w, r)
}
