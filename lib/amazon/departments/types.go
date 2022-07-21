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
)

func (tp *Type) AddSubType(stp *SubType, depID, catID, subID *primitive.ObjectID) error {
	ctx := context.Background()
	collection := variables.Local.Database("Amazon").Collection("Departments")
	if stp.Exists(depID, catID, subID, tp.ID) {
		return errors.New(fmt.Sprintf(variables.CategoryExists, tp.Title))
	}
	tp.NewID()
	filter := bson.M{"_id": bson.M{"$eq": depID}, "categories._id": bson.M{"$eq": catID}, "categories.subcategories._id": bson.M{"$eq": subID}, "categories.subcategories.types._id": bson.M{"$eq": tp.ID}}
	update := bson.M{"$addToSet": bson.M{"categories.$.subcategories.$.types.$.subtypes": tp}}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New(variables.InternalServerError)
	}
	return nil
}


func (tp *Type) Exists(depID, catID, subID *primitive.ObjectID) bool {
	ctx := context.Background()
	collection := variables.Local.Database("Amazon").Collection("Departments")
	var dep Department
	err := collection.FindOne(ctx, bson.M{"_id": depID, "categories._id": catID, "categories.subcategories._id": subID, "categories.subcategories.types.title": tp.Title}).Decode(&dep)
	return err == nil
}

func (tp *Type) NewID() {
	ID := primitive.NewObjectID()
	var dep Department
	ctx := context.Background()
	collection := variables.Local.Database("Amazon").Collection("Departments")
	err := collection.FindOne(ctx, bson.M{"categories.subcategories.types._id": ID}).Decode(&dep)
	if err == nil {
		tp.NewID()
	}
	tp.ID = &ID
}


func GetTypes(ID primitive.ObjectID) ([]byte, int) {
	ctx := context.Background()
	collection := variables.Local.Database("Amazon").Collection("Departments")
	Response := variables.Response{Action: variables.GetTypes}
	var dep Department

	err := collection.FindOne(ctx, bson.M{"categories.subcategories._id": ID}).Decode(&dep)
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
	loop:
	for cindex := range dep.Categories {
		for sindex := range dep.Categories[cindex].SubCategories {
			if dep.Categories[cindex].SubCategories[sindex].ID.Hex() == ID.Hex() {
				Response.Data = dep.Categories[cindex].SubCategories[sindex].Types
				break loop
			}
		}
	}
	return variables.JsonMarshal(Response), http.StatusOK
}