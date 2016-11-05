package main

import geoip2 "github.com/oschwald/geoip2-golang"
import "log"
import "github.com/howeyc/fsnotify"
import "net/http"
import "net"
import "encoding/json"
import "sync"

func openDB() geoip2.Reader {
	db, err := geoip2.Open("/data/GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	return *db
}

func main() {
	var dbMutex sync.RWMutex
	dbMutex.Lock()
	db := openDB()
	dbMutex.Unlock()

	// Responsible for database live reloads
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	watcher.Watch("/data/GeoLite2-City.mmdb")
	go func() {
		for {
			event := <-watcher.Event
			if event.IsRename() || event.IsCreate() || event.IsModify() {
				log.Printf("Starting live reload")
				dbMutex.Lock()
				db = openDB()
				dbMutex.Unlock()
				log.Printf("Finished live reload")
			}
		}
	}()

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		rawAddress := req.URL.Path[1:]
		address := net.ParseIP(rawAddress)
		if address == nil {
			res.WriteHeader(400)
			res.Write([]byte("Invalid IP address"))
			return
		}
		dbMutex.RLock()
		record, err := db.City(address)
		dbMutex.RUnlock()
		if err != nil {
			res.WriteHeader(404)
			res.Write([]byte(err.Error()))
			return
		}
		res.Header().Add("Content-Type", "application/json")
		result, _ := json.Marshal(record)
		res.Write(result)
	})
	log.Printf("Started geoip-server on port 4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
