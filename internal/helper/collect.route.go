package helper

/*
	@Author: Sinh
	@Date: 2025/6/1
	@Description: Collect route information for documentation, backend usage, and backup.
*/
import (
	"fmt"
	"os"
	"user_service/global"
)

// path format: method::http://domain/v1/api/path

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

// RegisterRoute registers a new route with the specified method and path.
// To view details, please refer to the routers in the internal/routers directory.
func RegisterRoute(method, path string) {
	routeList = append(routeList, NewRoute(method, path))
}

func GetRoutes() []string {
	var DOMAIN = fmt.Sprintf("http://%s:%d", global.Config.Server.Host, global.Config.Server.Port)
	var routes []string = []string{}
	for _, route := range routeList {
		routes = append(routes, fmt.Sprintf("%s::%s/%s", route.Method, DOMAIN, route.Path))
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
		routePath := fmt.Sprintf(route)
		_, err = file.WriteString(routePath + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}
