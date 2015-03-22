package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func index(d Data) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(d)
		if err != nil {
			panic(err)
		}*/
		page := make(map[string]Data)
		page[d.GetName()+"s"] = d
		err := json.NewEncoder(w).Encode(page)
		if err != nil {
			panic(err)
		}
	}
}
func show(d Data) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		vars := mux.Vars(r)
		recordID, _ := strconv.Atoi(vars["recordID"])
		i := d.GetItem(recordID)
		page := make(map[string]Item)
		page[d.GetName()] = i
		err := json.NewEncoder(w).Encode(page)
		if err != nil {
			panic(err)
		}
	}
}

func create(d Data) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		i, err := d.Decode(r.Body)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(422)
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
		}
		d.Append(i)
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(i)
		if err != nil {
			panic(err)
		}
		d.Save()
	}
}
func update(d Data) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		vars := mux.Vars(r)
		recordID, _ := strconv.Atoi(vars["recordID"])
		i, err := d.Decode(r.Body)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		err = r.Body.Close()
		d.Edit(recordID, i)
		d.Save()
	}
}
func delete(d Data) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		vars := mux.Vars(r)
		recordID, _ := strconv.Atoi(vars["recordID"])
		d.Remove(recordID)
		d.Save()
	}
}
