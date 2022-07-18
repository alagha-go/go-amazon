package departments

import (
	"context"
	"fmt"
	"net/http"

	"github.com/alagha-go/go-amazon/lib/variables"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// create this department if it does not exist
func (dep *Department) Create() ([]byte, int) {
	ctx := context.Background()
	collection := variables.Local.Database("Amazon").Collection("Departments")
	Response := variables.Response{Action: variables.CreateDepartment}
	if dep.Exists() {
		Response.Failed = true
		Response.Error = fmt.Sprintf(variables.DepartMentExists, dep.Title)
		return variables.JsonMarshal(Response), http.StatusConflict
	}
	dep.NewID()
	_, err := collection.InsertOne(ctx, dep)
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InternalServerError
		return variables.JsonMarshal(Response), http.StatusInternalServerError
	}
	Response.Success = true
	Response.Data = dep
	return variables.JsonMarshal(Response), http.StatusOK
}

// check if a department with this name already exists
func (dep *Department) Exists() bool {
	var department Department
	ctx := context.Background()
	collection := variables.Local.Database("Amazon").Collection("Departments")
	err := collection.FindOne(ctx, bson.M{"title": dep.Title}).Decode(&department)
	return err == nil
}

// give the department a unique ID
func (dep *Department) NewID() {
	ID := primitive.NewObjectID()
	var department Department
	ctx := context.Background()
	collection := variables.Local.Database("Amazon").Collection("Departments")
	err := collection.FindOne(ctx, bson.M{"_id": ID}).Decode(&department)
	if err == nil {
		dep.NewID()
	}
	dep.ID = &ID
}