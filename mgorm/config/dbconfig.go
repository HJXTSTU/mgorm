package config

import (
	"time"
	"gopkg.in/mgo.v2"
)

const (
	MONGO_DIRECT    = false
	MONGO_TIMEOUT   = time.Second * 1
	MONGO_DATABASE  = "****"
	MONGO_POOLLIMIT = 10240
)

const(
	MGORM_MODE_MODE = mgo.Monotonic
	MGORM_MODE_REFRESH = true
)

var MONGO_IPADDRS = []string{
	"****",
}

var MGORM_COLLECTIONS = []string{
	"***",
	"***",
	"***",
	"***",
	"***",
}

const (
	MGORM_DISABLE_ID_INDEX = false
	MGORM_FORCE_ID_INDEX   = true
	MGORM_CAPPED           = false
	MGORM_MAX_BYTES        = 0
	MGORM_MAX_DOCS         = 0
)
