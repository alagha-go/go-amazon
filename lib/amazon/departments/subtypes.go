package departments

import (
	"context"
	"fmt"
	"net/http"

	"github.com/alagha-go/go-amazon/lib/variables"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func (stp *SubType) Exists(depID, catID, subID, tpID *primitive.ObjectID) bool {
	ctx := context.Background()
	collection := variables.Local.Database("Amazon").Collection("Departments")
	var dep Department
	err := collection.FindOne(ctx, bson.M{"_id": depID, "categories._id": catID, "categories.subcategories._id": subID, "categories.subcategories.types._id": tpID, "categories.subcategories.types.subtypes.title": stp.Title}).Decode(&dep)
	return err == nil
}

func (stp *SubType) NewID() {
	ID := primitive.NewObjectID()
	var dep Department
	ctx := context.Background()
	collection := variables.Local.Database("Amazon").Collection("Departments")
	err := collection.FindOne(ctx, bson.M{"categories.subcategories.types.subtypes._id": ID}).Decode(&dep)
	if err == nil {
		stp.NewID()
	}
	stp.ID = &ID
}

func GetSubTypes(ID primitive.ObjectID) ([]byte, int) {
	ctx := context.Background()
	collection := variables.Local.Database("Amazon").Collection("Departments")
	Response := variables.Response{Action: variables.GetSubTypes}
	var dep Department
	opts := options.FindOne().SetProjection(bson.M{"categories.subcategories.types": bson.M{"$elemMatch": bson.M{"_id": ID}}})

	err := collection.FindOne(ctx, bson.M{"categories.subcategories.types._id": ID}, opts).Decode(&dep)
	if err != nil {
		Response.Failed = true
		if err == mongo.ErrNilDocument {
			Response.Error = fmt.Sprintf(variables.TypeNotFound, ID.Hex())
			return variables.JsonMarshal(Response), http.StatusNotFound
		}
		fmt.Println(err.Error())
		Response.Error = variables.InternalServerError
		return variables.JsonMarshal(Response), http.StatusInternalServerError
	}
	Response.Success = true
	if len(dep.Categories) > 0 {
		if len(dep.Categories[0].SubCategories) > 0 {
			if len(dep.Categories[0].SubCategories[0].Types) > 0 {
				Response.Data = dep.Categories[0].SubCategories[0].Types[0].SubTypes
			}
		}
	}
	return variables.JsonMarshal(Response), http.StatusOK
}