package controllers

import (
	"github.com/larbert/mayfly"
	"net/http"
)

type IndexController struct {
}

func (c *IndexController) GetIndex(context *mayfly.Context) {
	context.HTML(http.StatusOK, "index.html", mayfly.H{})
}
