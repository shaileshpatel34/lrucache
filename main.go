package main

import (
	"github.com/karlseguin/ccache"
	"log"
	"time"
)

func main() {
	log.Println("LRU CACHE")
	var cache = ccache.New(ccache.Configure())
	cache.Set("user:4", "Jon Snow", time.Minute*10)
	item, err := cache.Fetch("user:4", time.Minute*10, func() (interface{}, error) {
		log.Println("code to fetch the data incase of a miss. should return the data to cache and the error, if any")
		return nil, nil
	})
	log.Println(err, item.Value())
	item, err = cache.Fetch("user:1", time.Minute*10, func() (interface{}, error) {
		log.Println("code to fetch the data incase of a miss. should return the data to cache and the error, if any")
		return nil, nil
	})
	log.Println(err, item.Value())
}
