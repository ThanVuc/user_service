package helper

import (
	"encoding/json"
	"log"
)

type ResourceItem struct {
	Resource string   `json:"resource"`
	Actions  []string `json:"actions"`
}

var resourceList []ResourceItem = []ResourceItem{}

func AddResource(resource string, actions []string) {
	resourceList = append(resourceList, ResourceItem{
		Resource: resource,
		Actions:  actions,
	})
}

func GetResources() []byte {
	jsonData, err := json.Marshal(resourceList)
	if err != nil {
		log.Fatalf("Error marshalling resource list: %v", err)
	}

	return jsonData
}
