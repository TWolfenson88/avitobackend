package server

import (
	"avitocalls/internal/pkg/router"
	//blogDelivery "counity/internal/pkg/blog/delivery"
	//"counity/internal/pkg/router"
	"avitocalls/internal/pkg/settings"
	"github.com/dimfeld/httptreemux"
	//"sync"
	//
	//postDelivery "counity/internal/pkg/post/delivery"
	//userDelivery "counity/internal/pkg/user/delivery"
	"sync"
)

var routesMap = map[string][]settings.MapHandler{
	// USERS
	"/calls/make": {{
		Type:         "POST",
		// Handler:      userDelivery.FeedUsers,
		// CORS:         false,
		// AuthRequired: true,
		// CSRF:         false,
		// TokenRequired:true,
	}},
}

// Env variables which must to be set before running server
//var Secrets = []string{
//	"DB_NAME",
//	"DB_PASSWORD",
//	"DB_USER",
//	//"JWT_KEY",
//	//"AWS_TOKEN",
//}

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
				"http://localhost":           {},
				"http://localhost:8080":      {},
				"http://localhost:5000":      {},
				"http://127.0.0.1":           {},
				"http://127.0.0.1:8080":      {},
				"http://127.0.0.1:5000":      {},
			},
		}
		conf.InitSecure(&settings.SecureSettings)
		router.InitRouter(&conf, httptreemux.New())
	})
	return &conf
}
