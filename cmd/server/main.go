package main

import (
	"gopher-translator/pkg/translation"
	"gopher-translator/pkg/storage"
	"gopher-translator/pkg/handlers"
	"os"
	"log"
	"github.com/bouk/httprouter"
	"net/http"
)

func main()  {
	port := "8080"
	if len(os.Args) >= 3 && os.Args[1] == "-port" {
		port = os.Args[2]
	}

	repository := storage.CreateNewRedisStorage("localhost:6379","",0)
	translator := translation.CreateNewGopherTranslator()
	handler := handlers.CreateNewTranslatorHandler(repository, translator)
	router := httprouter.New()

	router.GET("/", handler.Welcome())
	router.POST("/word", handler.TranslateWord())
	router.POST("/sentence", handler.TranslateSentence())
	router.GET("/history", handler.TranslateHistory())

	log.Fatal(http.ListenAndServe(":" + port, router))
}