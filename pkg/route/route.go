package route

import (
	"github.com/julienschmidt/httprouter"
)

func Init() *httprouter.Router {
	r := httprouter.New()

	r.GET("/oders/")
}
