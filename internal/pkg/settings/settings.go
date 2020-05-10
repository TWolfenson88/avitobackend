package settings

import (
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request, map[string]string)

type MapHandler struct {
	Type          string
	Handler       HandlerFunc
	//TokenRequired bool
	//AuthRequired  bool
	//CORS          bool
	//CSRF          bool
}

type GlobalSecure struct {
	//CORSMethods  string
	//CORSMap      map[string]struct{}
	//AllowedHosts map[string]struct{}
}

var SecureSettings GlobalSecure

type ServerSettings struct {
	Port   int
	Ip     string
	Routes map[string][]MapHandler
	Router http.Handler
	Secure *GlobalSecure
}

func (s *ServerSettings) GetSettings() ServerSettings {
	return *s
}


func (s *ServerSettings) SetRoute(reqType, url string, handler HandlerFunc) {
	s.Routes[url] = append(s.Routes[url], MapHandler{Type: reqType, Handler: handler})
}

func (s *ServerSettings) SetRouter(handler http.Handler) {
	s.Router = handler
}

type RouterInterface interface {
	http.Handler
	POST(path string, handler HandlerFunc)
	GET(path string, handler HandlerFunc)
	PUT(path string, handler HandlerFunc)
	DELETE(path string, handler HandlerFunc)
	OPTIONS(path string, handler HandlerFunc)
}

func (s *ServerSettings) GetRouter() http.Handler {
	return s.Router
}
