package helper

import (
	"fmt"
	"os"
)

// path format: method::http://domain/v1/api/path

var DOMAIN = fmt.Sprintf("http://%s:%d", "localhost", 8080)

type Route struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

func NewRoute(method, path string) *Route {
	return &Route{
		Method: method,
		Path:   path,
	}
}

var (
	routeList = []*Route{}
)

func RegisterRoute(method, path string) {
	routeList = append(routeList, NewRoute(method, path))
}

func GetRoutes() []string {
	var routes []string = []string{}
	for _, route := range routeList {
		routes = append(routes, fmt.Sprintf("%s::%s%s", route.Method, DOMAIN, route.Path))
	}
	return routes
}

func WriteRouteToFile() {
	const PATH = "./backup"
	if _, err := os.Stat(PATH); os.IsNotExist(err) {
		err := os.MkdirAll(PATH, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
	}
	const FILE_PATH = PATH + "/route_collections.txt"
	file, err := os.OpenFile(FILE_PATH, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	routes := GetRoutes()
	for _, route := range routes {
		routePath := fmt.Sprintf("%s::%s", route, DOMAIN)
		_, err = file.WriteString(routePath + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}
