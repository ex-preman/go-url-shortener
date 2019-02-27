package main

import (
	"log"
	"net/http"
	"runtime"

	"github.com/ExPreman/url-shortener-go/handler"
	storage "github.com/ExPreman/url-shortener-go/storage/mysql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// call storage
	storage := &storage.MysqlDB{}
	err := storage.Init()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/shorten", handler.EncodeHandler(storage))
	http.Handle("/", handler.RedirectHandler(storage))

	log.Print("server start at :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
