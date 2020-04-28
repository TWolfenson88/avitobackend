package server

import (
	callDelivery "avitocalls/internal/pkg/call/delivery"
	"avitocalls/internal/pkg/router"
	"avitocalls/internal/pkg/settings"
	"github.com/dimfeld/httptreemux"
	"sync"
)

var routesMap = map[string][]settings.MapHandler{
	// CALLS
	"/calls/make": {{  // toDo /calls/make/id
		Type:    			"POST",
		Handler: 			callDelivery.CallUser,
		// CORS:         false,
		// AuthRequired: true,
		// CSRF:         false,
		// TokenRequired:true,
	}},
	"/calls/wait": {{
		Type:    			"GET",
		Handler: 			callDelivery.WaitForCall,
	}},
}

var doOnce sync.Once
var conf settings.ServerSettings

func GetConfig() *settings.ServerSettings {
	doOnce.Do(func() {
		conf = settings.ServerSettings{
			Port:   5000,
			Ip:     "0.0.0.0",
			Routes: routesMap,
		}
		settings.SecureSettings = settings.GlobalSecure{
			CORSMethods: "",
			CORSMap:     map[string]struct{}{},
			AllowedHosts: map[string]struct{}{
				"http://localhost":      {},
				"http://localhost:8080": {},
				"http://localhost:5000": {},
				"http://127.0.0.1":      {},
				"http://127.0.0.1:8080": {},
				"http://127.0.0.1:5000": {},
			},
		}
		conf.InitSecure(&settings.SecureSettings)
		router.InitRouter(&conf, httptreemux.New())
	})
	return &conf
}
