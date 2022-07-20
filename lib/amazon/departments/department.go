package departments

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/alagha-go/go-amazon/lib/variables"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func GetDepartments() ([]byte, int) {
	Response := variables.Response{Action: variables.GetDepartments}
	ctx := context.Background()
	collection := variables.Local.Database("Amazon").Collection("Departments")
	var departments []Department
	opts := options.Find().SetProjection(bson.M{"title": 1})
	cursor, err := collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		fmt.Println(err.Error())
		Response.Failed = true
		Response.Error = variables.InternalServerError
		return variables.JsonMarshal(Response), http.StatusInternalServerError
	}
	err = cursor.All(ctx, &departments)
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InternalServerError
		return variables.JsonMarshal(Response), http.StatusInternalServerError
	}
	Response.Success = true
	Response.Data = departments
	return variables.JsonMarshal(Response), http.StatusOK
}

func (dep *Department) AddCategory(cat *Category) error {
	ctx := context.Background()
	collection := variables.Local.Database("Amazon").Collection("Departments")
	if cat.Exists(dep.ID) {
		return nil
	}
	cat.NewID()
	filter := bson.M{"_id": bson.M{"$eq": dep.ID}}
	update := bson.M{"$addToSet": bson.M{"categories": cat}}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New(variables.InternalServerError)
	}
	return nil
}