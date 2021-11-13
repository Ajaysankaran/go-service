package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const servicePath = "services"

func handleServices(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		serviceList, err := GetServices()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		j, err := json.Marshal(serviceList)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPost:
		var service Service
		err := json.NewDecoder(r.Body).Decode(&service)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		serviceID, err := CreateService(service)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf(`{"productId":%d}`, serviceID)))
	case http.MethodOptions:
		return
	}
}

// func handleService(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		return
// 	}
// }

func SetupRoutes(apiBasePath string) {
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, servicePath), http.HandlerFunc(handleServices))
}
