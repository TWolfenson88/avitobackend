package server

import (
	callDelivery "avitocalls/internal/pkg/call/delivery"
	"avitocalls/internal/pkg/router"
	"avitocalls/internal/pkg/settings"
	userDelivery "avitocalls/internal/pkg/user/delivery"
	"github.com/dimfeld/httptreemux"
	"sync"
)

var routesMap = map[string][]settings.MapHandler{
	// CALLS
	"/calls/start": {{
		Type:    "POST",
		Handler: callDelivery.StartCall,
	}},
	"/calls/stop": {{
		Type:    "POST",
		Handler: callDelivery.EndCall,
	}},
	"/calls/history": {{
		Type:    "GET",
		Handler: callDelivery.GetHistory,
	}},
	// PROFILES
	"/users/reg": {{
		Type:    "POST",
		Handler: userDelivery.RegisterUser,
	}},
	"/users/all": {{
		Type:    "GET",
		Handler: userDelivery.FeedUsers,
	}},
	"/users/login": {{
		Type:    "POST",
		Handler: userDelivery.LoginUser,
	}},
	//"/ws": {{
	//	Type:
	//}}
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
		// settings.SecureSettings =
		settings.SecureSettings = settings.GlobalSecure{
			CORSMethods: "",
			CORSMap:     map[string]struct{}{},
			AllowedHosts: map[string]struct{}{
				"http://0.0.0.0":      {},
			},
		}
		conf.InitSecure(&settings.SecureSettings)
		router.InitRouter(&conf, httptreemux.New())
	})
	return &conf
}
