package web

import "net/http"


var homeTmpl = parseTmpl("home.tmpl")


type homeData struct {
	Session
}

func (h *Handler) renderHome(w http.ResponseWriter, data homeData, statusCode int, r *http.Request) {
	h.renderTmpl(w, homeTmpl, data, statusCode)
}


func (h *Handler) showHome(w http.ResponseWriter, r *http.Request) {
	
}