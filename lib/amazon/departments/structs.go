package departments

import "go.mongodb.org/mongo-driver/bson/primitive"

type Departments []Department


type Department struct {
	ID								*primitive.ObjectID								`json:"_id,omitempty" bson:"_id,omitempty"`
	Title							string											`json:"title,omitempty" bson:"title,omitempty"`
	Url								string											`json:"url,omitempty" bson:"url,omitempty"`
	Departments						[]Department									`json:"departments,omitempty" bson:"departments,omitempty"`
}