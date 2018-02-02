package global

import (
	"gopkg.in/mgo.v2"
	"os"
	"projects/server/data/mgorm/config"
	"log"
)

var session *mgo.Session
var collections = make(map[string]*mgo.Collection)

type DataStore struct{
	ds_session *mgo.Session
}

func (ds DataStore)Close(){
	ds.ds_session.Close()
}

func (ds DataStore)DB()*mgo.Database{
	return ds.ds_session.DB("")
}

func NewDataStore()DataStore{
	ds :=DataStore{
		ds_session:session.Copy(),
	}
	return ds
}

func Mgorm_Init(exit_signal chan os.Signal) {
	var err error
	// TODO::Init Session
	// TODO::Init Db
	// TODO::Init Collections
	dialInfo := mgo.DialInfo{
		Addrs:     config.MONGO_IPADDRS,
		Direct:    config.MONGO_DIRECT,
		Timeout:   config.MONGO_TIMEOUT,
		Database:  config.MONGO_DATABASE,
		PoolLimit: config.MONGO_POOLLIMIT,
	}
	log.Println("[INFO] mgorm:connect session:")
	session, err = mgo.DialWithInfo(&dialInfo)
	if err != nil {
		panic(err)
	}

	session.SetMode(config.MGORM_MODE_MODE, config.MGORM_MODE_REFRESH)
	log.Println("[INFO] mgorm:connect session done")
	//	Use default
	//	If database not define
	//	It will create a new one
	db := session.DB("")
	log.Println("[INFO] mgorm:get database name:" + db.Name + " done")
	info := mgo.CollectionInfo{
		DisableIdIndex: config.MGORM_DISABLE_ID_INDEX,
		ForceIdIndex:   config.MGORM_FORCE_ID_INDEX,
		Capped:         config.MGORM_CAPPED,
		MaxBytes:       config.MGORM_MAX_BYTES,
		MaxDocs:        config.MGORM_MAX_DOCS,
	}
	log.Println("[INFO] mgorm:create collections for:" + config.MONGO_DATABASE)
	for _, v := range config.MGORM_COLLECTIONS {
		db.C(v).Create(&info)
		collections[v] = db.C(v)
	}
	log.Println("[INFO] mgorm:create collections for:" + config.MONGO_DATABASE + " done")
	log.Println("[INFO] mgorm:init done")
	go func(exit_signal chan os.Signal) {
		<-exit_signal
		log.Println("[INFO] mgorm:close connection")
		session.Close()
		os.Exit(0)
	}(exit_signal)
}
