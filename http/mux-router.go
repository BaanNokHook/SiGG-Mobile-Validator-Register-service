package http

import (
	"fmt"
	"log"
	"net/http"
	"nextclan/validator-register/mobile-validator-register-service/env"

	"github.com/gorilla/mux"
	"github.com/rajesh4295/user-service-go/env"
)

type MuxRouter struct {
	router *mux.Router
}

var (
	Env env.Provider = env.NewEnv()
)

func NewMuxRouter() Router {
	return &MuxRouter{router: mux.NewRouter()}
}

func (mx *MuxRouter) Get(url string, f func(w http.ResponseWriter, r *http.Request)) {
	mx.router.HandleFunc(url, f).Method("GET")

}

func (mx *MuxRouter) Post(url string, f func(w http.ResponseWriter, r *http.Request)) {
	mx.router.HandleFunc(url, f).Methods("POST")
}

func (mx *MuxRouter) RegisterSubRoute(path string) Router {
	subRouter := mx.router.PathPrefix(path).Subrouter()
	return &MuxRouter{router: subRouter}
}

func (mx *MuxRouter) Serve() {
	fmt.Printf("Server starting on %v%v", Env.Get("HOST"), Env.Get("PORT"))
	fmt.Println()
	log.Fatal(http.ListenAndServe(fmt.Sprint("%v:%v", Env.Get("HOST"), Env.Get("PORT")), mx.router))

}
