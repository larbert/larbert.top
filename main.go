package main

import (
	"net/http"

	"github.com/larbert/mayflylog"
	"github.com/larbert/mayflyweb"
)

func main() {
	mayflylog.SetLevel(mayflylog.DebugLevel)

	s := mayflyweb.New()
	s.Get("/", func(ctx *mayflyweb.Context) {
		ctx.String(http.StatusOK, "hello %s!", "larbert")
	})

	err := s.Run(":80")
	if err != nil {
		mayflylog.Error("启动失败")
	}
}
