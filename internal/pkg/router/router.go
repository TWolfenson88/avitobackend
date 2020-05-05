package router

import (
	"avitocalls/internal/pkg/middleware"

	"avitocalls/internal/pkg/settings"
	"github.com/dimfeld/httptreemux"
	"log"
)

// Parse route map and return configured Router
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
			//if pack.AuthRequired {
			//	handler = middleware.Authenticate(handler)
			//}
			//handler = middleware.CheckToken(handler)
			handler = middleware.SetAllowOrigin(handler)
			handler = middleware.DecodeBody(handler)


			//if pack.CORS {
			//	s.Secure.CORSMap[pack.Type] = struct{}{}
			//	handler = middleware.CORS(handler)
			//}
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
	//// generate "GET, POST, OPTIONS, HEAD, PUT" string
	//for key, _ := range s.Secure.CORSMap {
	//	s.Secure.CORSMethods += key + ", "
	//}
	s.Router = router
}
