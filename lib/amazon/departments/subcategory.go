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


func (sub *SubCategory) AddType(tp *Type, depID, catID *primitive.ObjectID) error {
	ctx := context.Background()
	collection := variables.Local.Database("Amazon").Collection("Departments")
	if tp.Exists(depID, catID, sub.ID) {
		return errors.New(fmt.Sprintf(variables.CategoryExists, tp.Title))
	}
	tp.NewID()
	filter := bson.M{"_id": bson.M{"$eq": depID}, "categories._id": bson.M{"$eq": catID}, "categories.subcategories._id": bson.M{"$eq": sub.ID}}
	update := bson.M{"$addToSet": bson.M{"categories.$.subcategories.types": tp}}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New(variables.InternalServerError)
	}
	return nil
}


func (sub *SubCategory) Exists(depID, catID *primitive.ObjectID) bool {
	ctx := context.Background()
	collection := variables.Local.Database("Amazon").Collection("Departments")
	var dep Department
	err := collection.FindOne(ctx, bson.M{"_id": depID, "categories._id": catID, "categories.subcategories.title": sub.Title}).Decode(&dep)
	return err == nil
}


func (sub *SubCategory) NewID() {
	ID := primitive.NewObjectID()
	var dep Department
	ctx := context.Background()
	collection := variables.Local.Database("Amazon").Collection("Departments")
	err := collection.FindOne(ctx, bson.M{"categories.subcategories._id": ID}).Decode(&dep)
	if err == nil {
		sub.NewID()
	}
	sub.ID = &ID
}


// get subcategories
func GetSubCategories(ID primitive.ObjectID) ([]byte, int) {
	ctx := context.Background()
	collection := variables.Local.Database("Amazon").Collection("Departments")
	Response := variables.Response{Action: variables.GetSubCategories}
	var dep Department
	opts := options.FindOne().SetProjection(bson.M{"categories": bson.M{"$elemMatch": bson.M{"_id": ID}}})

	err := collection.FindOne(ctx, bson.M{"categories._id": ID}, opts).Decode(&dep)
	if err != nil {
		Response.Failed = true
		if err == mongo.ErrNilDocument {
			Response.Error = fmt.Sprintf(variables.SubCategoryNotFound, ID.Hex())
			return variables.JsonMarshal(Response), http.StatusNotFound
		}
		fmt.Println(err.Error())
		Response.Error = variables.InternalServerError
		return variables.JsonMarshal(Response), http.StatusInternalServerError
	}
	Response.Success = true
	if len(dep.Categories) > 0 {
		for index := range dep.Categories[0].SubCategories {
			dep.Categories[0].SubCategories[index].Types = []Type{}
		}
		Response.Data = dep.Categories[0].SubCategories
	}
	return variables.JsonMarshal(Response), http.StatusOK
}