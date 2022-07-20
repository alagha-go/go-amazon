package departments

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/alagha-go/go-amazon/lib/variables"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// add subcategory to the database
func (cat *Category) AddSubCategory(sub *SubCategory, depID *primitive.ObjectID) error {
	ctx := context.Background()
	collection := variables.Local.Database("Amazon").Collection("Departments")
	if sub.Exists(depID, cat.ID) {
		return errors.New(fmt.Sprintf(variables.CategoryExists, cat.Title))
	}
	sub.NewID()
	filter := bson.M{"_id": bson.M{"$eq": depID}, "categories._id": bson.M{"$eq": cat.ID}}
	update := bson.M{"$addToSet": bson.M{"categories.$.subcategories": sub}}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New(variables.InternalServerError)
	}
	return nil
}

// check if category with the title exists
func (cat *Category) Exists(depID *primitive.ObjectID) bool {
	ctx := context.Background()
	collection := variables.Local.Database("Amazon").Collection("Departments")
	var dep Department
	err := collection.FindOne(ctx, bson.M{"_id": depID, "categories.title": cat.Title}).Decode(&dep)
	return err == nil
}

// give the category a unique ID
func (cat *Category) NewID() {
	ID := primitive.NewObjectID()
	var dep Department
	ctx := context.Background()
	collection := variables.Local.Database("Amazon").Collection("Departments")
	err := collection.FindOne(ctx, bson.M{"categories._id": ID}).Decode(&dep)
	if err == nil {
		cat.NewID()
	}
	cat.ID = &ID
}

func GetCategories(ID primitive.ObjectID) ([]byte, int) {
	ctx := context.Background()
	collection := variables.Local.Database("Amazon").Collection("Departments")
	Response := variables.Response{Action: variables.GetCategories}
	var dep Department
	opts := options.FindOne().SetProjection(bson.M{"categories.title": 1, "categories.url": 1, "categories._id": 1})

	err := collection.FindOne(ctx, bson.M{"_id": ID}, opts).Decode(&dep)
	if err != nil {
		Response.Failed = true
		if err == mongo.ErrNilDocument {
			Response.Error = fmt.Sprintf(variables.CategoryNotFound, ID.Hex())
			return variables.JsonMarshal(Response), http.StatusNotFound
		}
		fmt.Println(err.Error())
		Response.Error = variables.InternalServerError
		return variables.JsonMarshal(Response), http.StatusInternalServerError
	}
	Response.Success = true
	Response.Data = dep.Categories
	return variables.JsonMarshal(Response), http.StatusOK
}