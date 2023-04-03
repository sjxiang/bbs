package web

import (
	"net/http"

	"github.com/sjxiang/bbs"
)


type Session struct {
	IsLoggedIn bool
	User       bbs.User
}



func (h *Handler) sessionFromReq(r *http.Request) Session {
	var out Session

	u, exist := h.Session.Values["user"]
	if !exist {
		return out
	} 

	out.IsLoggedIn = true
	out.User = u.(bbs.User)

	return out
}