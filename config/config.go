package config

import (
	"github.com/tkanos/gonfig"
	"path"
	"path/filepath"
	"runtime"
	"time"
)

type PlacesCacheSettings struct {
	Ttl time.Duration
}

type CacheSettings struct {
	Places PlacesCacheSettings
}

type PlacesAviasalesRequest struct {
	Url string
	Timeout time.Duration
}

type AviasalesRequests struct {
	Places PlacesAviasalesRequest
}

type ServerConfig struct {
	Host string
	Port string
}

type RedisConfig struct {
	Host string
	Port string
	Password string
}

type Requests struct {
	Aviasales AviasalesRequests
}

type StorageConfig struct {
	Redis RedisConfig
}

type Config struct {
	Server ServerConfig
	Storage StorageConfig
	Requests Requests
	Cache CacheSettings
}

var Configuration = Config{}

var _, dirname, _, _ = runtime.Caller(0)
var filePath = path.Join(filepath.Dir(dirname), "default.json")

var _ = gonfig.GetConf(filePath, &Configuration)