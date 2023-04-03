package web

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/sjxiang/bbs"
)

var loginTmpl = parseTmpl("login.tmpl")


type loginData struct {
	Form url.Values
	Err  error
}


func (h *Handler) renderLogin(w http.ResponseWriter, data loginData, statusCode int) {
	h.renderTmpl(w, loginTmpl, data, statusCode)	
}


func (h *Handler) showLogin(w http.ResponseWriter, r *http.Request) {
	h.renderLogin(w, loginData{}, http.StatusOK)
}


func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	
	if err := r.ParseForm(); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	input := bbs.LoginInput{
		Email:    r.PostFormValue("email"),
		Username: formStr(r.PostForm, "username"),
	}

	user , err := h.Service.Login(ctx, input)
	if err != nil {
		h.Logger.Printf("could not login: %v\n", err)
		http.Error(w, fmt.Sprintf("internal server error"), http.StatusInternalServerError)
		return 
	}



	http.Redirect(w, r, "/", http.StatusFound)
}


func formStr(form url.Values, key string) *string {
	if !form.Has(key) {
		return nil
	}

	s := form.Get(key)
	return &s
}