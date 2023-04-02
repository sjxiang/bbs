package web

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/sjxiang/bbs"
)


type Handler struct {
	Logger  log.Logger	
	handler http.Handler
	once    sync.Once

	Service bbs.Service
}

func (h *Handler) init() {
	r := mux.NewRouter()
	
	r.HandleFunc("/login", h.showLogin).Methods("GET").Name("登录页面")
	r.HandleFunc("/login", h.login).Methods("POST").Name("请求登录")


	h.handler = r
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.once.Do(h.init)

	h.handler.ServeHTTP(w, r)
}	