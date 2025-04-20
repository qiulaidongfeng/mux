package mux

import (
	"net/http"
	"strings"
)

// Mux 根据域名调用不同的handle
type Mux struct {
	allStd      map[string]http.Handler
	firstStd    http.Handler
	firstDomain string
}

func New() *Mux {
	return &Mux{
		allStd: make(map[string]http.Handler),
	}
}

// Allow 报告域名是否有handle
func (m *Mux) Allow(domain string) bool {
	// 根据域名调用对应的handle
	if handle := m.allStd[domain]; handle != nil {
		return true
	}
	return strings.HasSuffix(domain, m.firstDomain)
}

// AddStd 添加特定域名的handle
// 如果通过ip访问网站，无响应
// 如果通过第一次添加的子域名访问网站，使用第一次添加的handle
func (m *Mux) AddStd(domain string, handle http.Handler) {
	if m.firstStd == nil {
		m.firstStd = handle
		m.firstDomain = domain
	}
	m.allStd[domain] = handle
}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 根据域名调用对应的handle
	if handle := m.allStd[r.Host]; handle != nil {
		handle.ServeHTTP(w, r)
		return
	}
	if strings.HasSuffix(r.Host, m.firstDomain) {
		m.firstStd.ServeHTTP(w, r)
	}
}
