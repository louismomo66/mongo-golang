package main

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/louismomo66/mongo-golang/controllers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main(){
r := httprouter.New()
uc := controllers.NewUserController(getSession())
r.GET("/user/:id",uc.GetUser)
r.POST("/user",uc.CreateUser)
r.DELETE("/user/:id",uc.DeleteUser)
http.ListenAndServe("localhost:9000",r)
}

func getSession() *mongo.Client {
	// s, err := mgo.Dial("mongodb+srv://louis:09147625@cluster0.t9yxjh9.mongodb.net/")
	s, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://louis:databasekey@cluster0.t9yxjh9.mongodb.net/"))
	if err != nil {
		panic(err)
	}
	return s
}
