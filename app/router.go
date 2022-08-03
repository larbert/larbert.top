package app

import (
	"github.com/larbert/mayfly"
	"larbert.top/app/controllers"
	"reflect"
	"strings"
)

func RegisterRoutes(e *mayfly.Engine) {
	f(e, &controllers.IndexController{})
}

func f(e *mayfly.Engine, c interface{}) {
	cv := reflect.ValueOf(c)

	ct := cv.Type()
	cn := ct.Elem().Name()
	for i := 0; i < ct.NumMethod(); i++ {
		cm := cv.Method(i)
		name := ct.Method(i).Name
		if strings.HasPrefix(name, `Get`) {
			path := name[3:]
			e.Get("/"+cn+"/"+path, cm.Interface().(func(*mayfly.Context)))
		}
	}
}
