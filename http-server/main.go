package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("handler started")
	defer log.Println("handler ended")
	select {
	case <-time.After(5 * time.Second):
		log.Println("request processed")
		w.Write([]byte("Hello from http"))
	case <-ctx.Done():
		log.Println(ctx.Err())
		http.Error(w, ctx.Err().Error(), http.StatusInternalServerError)
	}

}
