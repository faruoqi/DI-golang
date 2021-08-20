package controllers

import (
	"DI/services"
	"DI/utils"
	"net/http"
)

type Controller struct {
	l     utils.Logger
	logic services.Logic
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("In SayHello")
	userID := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

func NewController(l utils.Logger, logic services.Logic) Controller {
	return Controller{
		l:     l,
		logic: logic,
	}
}
