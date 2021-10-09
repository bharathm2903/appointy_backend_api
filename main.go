package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)
 func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}
 func main(){
	router := httprouter.New()
    router.GET("/", Index)
 }