package logger

import (
	"log"
	"gopkg.in/mgo.v2"
)

const(
	logger_error ="[ERROR]%s,%s\n"
	logger_info = "[INFO]%s\n"
	logger_info_find = "[INFO]%s,%v\n"
	logger_info_upsert = "[INFO]%s,%v\n"
)

func Logger_Error(msg,error string){
	log.Printf(logger_error,msg,error)
}

func Logger_Info(msg string){
	log.Printf(logger_info,msg)
}

func Logger_Info_Find(msg,v interface{}){
	log.Printf(logger_info_find,msg,v)
}

func Logger_Info_Upsert(msg string,info *mgo.ChangeInfo){
	log.Printf(logger_info_upsert,msg,*info)
}
