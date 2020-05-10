package router

import (
	"avitocalls/internal/pkg/middleware"

	"avitocalls/internal/pkg/settings"
	"github.com/dimfeld/httptreemux"
	"log"
)

func InitRouter(s *settings.ServerSettings, router *httptreemux.TreeMux) {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal("Error was occurred", r)
		}
	}()

	var optionsHandler settings.HandlerFunc = nil
	for key, list := range s.Routes {
		for _, pack := range list {
			handler := pack.Handler

			handler = middleware.SetAllowOrigin(handler)
			handler = middleware.DecodeBody(handler)

			switch pack.Type {
			case "GET":
				(*router).GET(key, httptreemux.HandlerFunc(handler))
			case "PUT":
				(*router).PUT(key, httptreemux.HandlerFunc(handler))
			case "POST":
				(*router).POST(key, httptreemux.HandlerFunc(handler))
			case "DELETE":
				(*router).DELETE(key, httptreemux.HandlerFunc(handler))
			case "OPTIONS":
				optionsHandler = handler
			}

		}
	}

	if optionsHandler != nil {
		for key, _ := range s.Routes {
			(*router).OPTIONS(key, httptreemux.HandlerFunc(optionsHandler))
		}
	}
	s.Router = router
}
