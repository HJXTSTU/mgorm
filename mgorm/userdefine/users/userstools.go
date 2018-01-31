package users

import (
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"projects/mgorm/logger"
	"projects/mgorm/global"
)

const collection_name="users"

type User struct{
	Id bson.ObjectId `bson:"_id" json:"_id"`
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
	Gradename string `bson:"gradename" json:"gradename"`
}



func NewUsers(username,password,gradename string){
	user := User{
		Username:username,
		Password:password,
		Gradename:gradename,
	}
	db := global.DB()
	logmsg:="Insert Users:"+username
	err :=db.C(collection_name).Insert(&user)
	if err == nil{
		logger.Logger_Error(logmsg,err.Error())
	}else{
		logger.Logger_Info(logmsg)
	}
}

func FindUserByUsername(username string)(User,string){
	user := User{}
	db := global.DB()
	logmsg := "Find User:"+username
	err := db.C(collection_name).Find(bson.M{"username":username}).One(&user)
	if err != nil{
		logger.Logger_Error(logmsg,err.Error())
	}else{
		logger.Logger_Info_Find(logmsg,user)
	}
	js,_ :=json.Marshal(&user)
	return user,string(js)
}