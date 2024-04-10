package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Handletest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "testando servidor")
}

type countHandler int

func (ch countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "numero de acesso %d", ch)
}
func main() {
	//http.HandleFunc("/", Handletest)
	c := countHandler(0)
	s := &http.Server{
		Addr:           ":8080",
		Handler:        c,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())

}
