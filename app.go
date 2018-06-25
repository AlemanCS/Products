package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var config = Config{}
var dao = ProductsDAO{}

var connection, err = InitRecommendationCache()

func AllProductsEndPoint(w http.ResponseWriter, r *http.Request) {
	products, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, products)
}

func FindProductEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	product, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}
	connection.ForceCache(product.Description, product.Name)
	fmt.Println(connection.getCache(product.Description))
	respondWithJson(w, http.StatusOK, product)

}

func CreateProductEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	product.ID = bson.NewObjectId()
	if err := dao.Insert(product); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, product)
}

func UpdateProductEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(product); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteProductEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(product); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func main() {

	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	err := dao.Connect()
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/products", AllProductsEndPoint).Methods("GET")
	r.HandleFunc("/products", CreateProductEndPoint).Methods("POST")
	r.HandleFunc("/products", UpdateProductEndPoint).Methods("PUT")
	r.HandleFunc("/products", DeleteProductEndPoint).Methods("DELETE")
	r.HandleFunc("/products/{id}", FindProductEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
