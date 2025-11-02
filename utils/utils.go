package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

var Validate = validator.New()

func ParseJSON(request *http.Request , payload any)error{
	if request.Body ==nil{
        return fmt.Errorf("missing request body")
	}

 return json.NewDecoder(request.Body).Decode(payload)
 
}

func WriteJSON(writer http.ResponseWriter,statusCode int , data any)error {
	writer.Header().Add("Content-Type","application/json")
	writer.WriteHeader(statusCode)
	return json.NewEncoder(writer).Encode(data)
}

func WriteError(writer http.ResponseWriter ,status int,err error){
	 WriteJSON(
		writer,
		status,
		map[string]string{"error":err.Error()},
	)
		
	
}


func PrintRoutes(router *mux.Router) {
    router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
        pathTemplate, _ := route.GetPathTemplate()
        methods, _ := route.GetMethods()
        fmt.Printf("Endpoint: %s \t Methods: %v\n", pathTemplate, methods)
        return nil
    })
}