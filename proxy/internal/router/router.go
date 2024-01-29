package router

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"net/http/httputil"
	"net/url"
	"proxy/internal/controller"
)

func StRout(cn *controller.GeoHandle) *chi.Mux {

	r := chi.NewRouter()
	rp := NewReverseProxy("hugo", "1313")
	r.Use(rp.ReverseProxy)
	r.Post("/api/address/search", cn.SearchHandler)

	return r
}

type ReverseProxy struct {
	host string
	port string
}

func NewReverseProxy(host, port string) *ReverseProxy {
	return &ReverseProxy{
		host: host,
		port: port,
	}
}

func (rp *ReverseProxy) ReverseProxy(next http.Handler) http.Handler {

	target, _ := url.Parse(fmt.Sprintf("http://%s:%s", rp.host, rp.port))
	proxy := httputil.NewSingleHostReverseProxy(target)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//r.URL.Path = r.URL.Path[len("/api/"):]
		proxy.ServeHTTP(w, r)
	})
}
