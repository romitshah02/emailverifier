package gateway

import (
	"gateway/internal/models"
	"gateway/internal/routes"
	"log"
	"sync"
	"time"
)

var (
	routeCache = make(map[string]*models.Route)
	mu         sync.RWMutex
)

func LoadRoutes() {
	routesList, err := routes.ListRoutes()
	if err != nil {
		log.Printf("Failed to load routes: %v", err)
		return
	}

	newCache := make(map[string]*models.Route)
	for _, r := range routesList {
		if r.Enabled {
			key := r.Method + " " + r.Path
			newCache[key] = &r
		}
	}

	mu.Lock()
	routeCache = newCache
	mu.Unlock()

	log.Printf("Loaded %d active routes", len(newCache))
}

func GetRoute(method, path string) *models.Route {
	mu.RLock()
	defer mu.RUnlock()
	key := method + " " + path
	return routeCache[key]
}

func StartRouteWatcher() {
	LoadRoutes() // Initial load
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		for range ticker.C {
			LoadRoutes()
		}
	}()
}
