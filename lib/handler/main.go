package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alagha-go/go-amazon/crawler"
	"github.com/alagha-go/go-amazon/lib/amazon/departments"
	"github.com/alagha-go/go-amazon/lib/variables"
	gosocketio "github.com/ambelovsky/gosf-socketio"
	"github.com/ambelovsky/gosf-socketio/transport"
	"github.com/gorilla/mux"
)

var (
	Router = mux.NewRouter()
	ServeMux = http.NewServeMux()
	Server = gosocketio.NewServer(transport.GetDefaultWebsocketTransport())
)


func init() {
	fmt.Println("starting server...")
	go StartCollection()
	ServeMux.Handle("/", Router)
	Router.HandleFunc("/departments", GetDepartments)
	Router.HandleFunc("/departments/{id}", GetCategories)
	Router.HandleFunc("/departments/categories/{id}", GetSubCategories)
	Router.HandleFunc("/departments/categories/subcategories/{id}", GetTypes)
	Router.HandleFunc("/departments/categories/subcategories/types/{id}", GetSubTypes)
}


func StartCollection() {
	var Deps []departments.Department
	deps := crawler.GetDepartments()
	json.Unmarshal(variables.JsonMarshal(deps), &Deps)
	for index := range Deps {
		Deps[index].Create()
		for i := range Deps[index].Categories {
			Deps[index].AddCategory(&Deps[index].Categories[i])
		}
	}
	for dindex := range deps {
		depID := Deps[dindex].ID
		for cindex := range deps[dindex].Categories {
			catID := Deps[dindex].Categories[cindex].ID
			deps[dindex].Categories[cindex].SetSubCategories()
			for sindex := range deps[dindex].Categories[cindex].SubCategories {
				var sub departments.SubCategory
				json.Unmarshal(variables.JsonMarshal(deps[dindex].Categories[cindex].SubCategories[sindex]), &sub)
				Deps[dindex].Categories[cindex].AddSubCategory(&sub, depID)
				deps[dindex].Categories[cindex].SubCategories[sindex].SetTypes()
				for tindex := range deps[dindex].Categories[cindex].SubCategories[sindex].Types {
					var tp departments.Type
					json.Unmarshal(variables.JsonMarshal(deps[dindex].Categories[cindex].SubCategories[sindex].Types[tindex]), &tp)
					Deps[dindex].Categories[cindex].SubCategories[sindex].AddType(&tp, depID, catID)
					deps[dindex].Categories[cindex].SubCategories[sindex].Types[tindex].SetSubTypes()
					for s2index := range deps[dindex].Categories[cindex].SubCategories[sindex].Types[tindex].SubTypes {
						var stp departments.SubType
						json.Unmarshal(variables.JsonMarshal(deps[dindex].Categories[cindex].SubCategories[sindex].Types[tindex].SubTypes[s2index]), &stp)
						Deps[dindex].Categories[cindex].SubCategories[sindex].Types[tindex].AddSubType(&stp, depID, catID, sub.ID)
					}
				}
			}
		}
		fmt.Println(dindex+1, len(deps))
	}
	fmt.Println("done")
}