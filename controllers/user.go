package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/louismomo66/mongo-golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct{
	session *mongo.Client
}

func NewUserController(s *mongo.Client) *UserController{
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request,p httprouter.Params){
	id := p.ByName("id")

	if !primitive.IsValidObjectID(id){
		w.WriteHeader(http.StatusNotFound)
	}
oid := bson.M{"_id": id}

u := models.User{}

if err := uc.session.Database("sample_mflix").Collection("movies").FindOne(context.Background(),oid).Decode(&u); err != nil{
	w.WriteHeader(404)
	return
}

uj, err := json.Marshal(u)
if err != nil{
	fmt.Println(err)
}

w.Header().Set("Content-Type", "application/json")
w.WriteHeader(200)
fmt.Fprintf(w, "%s\n",uj)

}

func (uc UserController) CreateUser (w http.ResponseWriter, r *http.Request, _ httprouter.Params){
u := models.User{}
json.NewDecoder(r.Body).Decode(&u)
u.Id = primitive.NewObjectID()

uc.session.Database("mongo-goland").Collection("users").InsertOne(context.Background(),u)

uj, err := json.Marshal(u)
fmt.Println(string(uj))
if err != nil{
	fmt.Println(err)
}
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusCreated) //201
fmt.Fprintf(w,"%s\n" ,uj)


}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
 id := p.ByName("id")

 if !primitive.IsValidObjectID(id){
w.WriteHeader(404)
return
 }

 oid := bson.M{"_id": id}

if _,err := uc.session.Database("mongo-goland").Collection("users").DeleteOne(context.Background(),oid); err != nil {
	w.WriteHeader(404)
}

w.WriteHeader(200)
fmt.Fprint(w, "Deleted user", oid, "\n")

}