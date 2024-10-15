package adapter

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/edusantanaw/desafio_backend_with_golang/pkg/utils"
)

type IController[T interface{}] func(data T) utils.HttpResponse

func AdapterWithBody[T comparable](controller IController[T], schema T, route string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := json.NewDecoder(r.Body).Decode(&schema)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response := controller(schema)
		params := buildParams(route, r.RequestURI)
		body, err := json.Marshal(params)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
		w.WriteHeader(response.Code)
	}
}

func AdapterWithQuery(controller IController[map[string]string], route string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		queryMap := make(map[string]string)
		params := buildParams(r.RequestURI, route)
		for key, values := range query {
			if len(values) > 0 {
				queryMap[key] = values[0]
			}
		}
		response := controller(queryMap)
		body, err := json.Marshal(params)

		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
		w.WriteHeader(response.Code)
	}
}

func buildParams(route string, routerRoute string) map[string]string {
	routePath := strings.Split(strings.Split(route, "?")[0], "/")
	routerPath := strings.Split(routerRoute, "/")
	params := filterAndBind(routerPath, routePath, strings.Contains, ":")
	return params
}

func filterAndBind(slice []string, bind []string, filterHanlder func(v string, e string) bool, equals string) map[string]string {
	res := make(map[string]string)
	for key, value := range slice {
		println(1, value)
		if filterHanlder(value, equals) {
			println(key, value, bind[key])
			res[strings.Replace(value, ":", "", 1)] = bind[key]
		}
	}
	return res
}
