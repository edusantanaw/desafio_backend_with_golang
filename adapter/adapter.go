package adapter

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/edusantanaw/desafio_backend_with_golang/pkg/utils"
)

type AdapterContext[T comparable] struct {
	Body   T
	Params map[string]string
	Query  map[string]string
}

type IAdapterWithBodyController[T comparable] func(ctx *AdapterContext[T]) utils.HttpResponse

func AdapterWithBody[T comparable](controller IAdapterWithBodyController[T], schema T, route string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := json.NewDecoder(r.Body).Decode(&schema)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		params := buildParams(route, r.RequestURI)
		ctx := &AdapterContext[T]{Params: params, Body: schema}
		response := controller(ctx)
		body, err := json.Marshal(response.Body)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.Code)
		w.Write(body)
	}
}

type GetAdapterContext struct {
	Params map[string]string
	Query  map[string]string
}

type IAdapterWithQueryController[T interface{}] func(ctx *GetAdapterContext) utils.HttpResponse

func AdapterWithQuery(controller IAdapterWithQueryController[map[string]string], route string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		queryMap := make(map[string]string)
		params := buildParams(r.RequestURI, route)
		for key, values := range query {
			if len(values) > 0 {
				queryMap[key] = values[0]
			}
		}
		ctx := &GetAdapterContext{Params: params, Query: queryMap}
		response := controller(ctx)
		body, err := json.Marshal(response.Body)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.Code)
		w.Write(body)
	}
}

func buildParams(route string, routerRoute string) map[string]string {
	routePath := strings.Split(strings.Split(route, "?")[0], "/")
	routerPath := strings.Split(routerRoute, "/")
	params := filterAndBind(routerPath, routePath, strings.Contains, ":")
	return params
}

type IFilter func(v string, e string) bool

func filterAndBind(slice []string, bind []string, filterHanlder IFilter, equals string) map[string]string {
	res := make(map[string]string)
	for key, value := range slice {
		if filterHanlder(value, equals) {
			println(key, value, bind[key])
			res[strings.Replace(value, ":", "", 1)] = bind[key]
		}
	}
	return res
}
