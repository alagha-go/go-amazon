package handler

import (
	"encoding/json"
	"net/http"

	"github.com/alagha-go/go-amazon/lib/amazon/departments"
	"github.com/alagha-go/go-amazon/lib/variables"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



func GetDepartments(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	data, status := departments.GetDepartments()
	res.WriteHeader(status)
	res.Write(data)
}


func GetCategories(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	Response := variables.Response{Action: variables.GetCategories}
	var dep departments.Department
	ID, err := primitive.ObjectIDFromHex(mux.Vars(req)["id"])
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InvalidID
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(Response)
		return
	}
	data, status := departments.GetCategories(ID)
	Response.Success = true
	Response.Data = dep.Categories
	res.WriteHeader(status)
	res.Write(data)
}


func GetSubCategories(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	Response := variables.Response{Action: variables.GetSubCategories}
	var dep departments.Department
	ID, err := primitive.ObjectIDFromHex(mux.Vars(req)["id"])
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InvalidID
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(Response)
		return
	}
	data, status := departments.GetSubCategories(ID)
	Response.Success = true
	Response.Data = dep.Categories
	res.WriteHeader(status)
	res.Write(data)
}


func GetTypes(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	Response := variables.Response{Action: variables.GetTypes}
	var dep departments.Department
	ID, err := primitive.ObjectIDFromHex(mux.Vars(req)["id"])
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InvalidID
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(Response)
		return
	}
	data, status := departments.GetTypes(ID)
	Response.Success = true
	Response.Data = dep.Categories
	res.WriteHeader(status)
	res.Write(data)
}


func GetSubTypes(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	Response := variables.Response{Action: variables.GetTypes}
	var dep departments.Department
	ID, err := primitive.ObjectIDFromHex(mux.Vars(req)["id"])
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InvalidID
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(Response)
		return
	}
	data, status := departments.GetSubTypes(ID)
	Response.Success = true
	Response.Data = dep.Categories
	res.WriteHeader(status)
	res.Write(data)
}