package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var categories = []Category{
	{ID: 1, Name: "Buah", Description: "Buah-buahan"},
	{ID: 2, Name: "Elektronik", Description: "Elektronik"},
}

func getCategory(id int) (Category, error) {
	for _, c := range categories {
		if c.ID == id {
			return c, nil
		}
	}
	return Category{}, fmt.Errorf("category with id %d not found", id)
}

func addCategory(c Category) {
	var newId int
	if len(categories) > 0 {
		newId = categories[len(categories)-1].ID + 1
	}
	c.ID = newId
	categories = append(categories, c)
}

func updateCategory(id int, updatedC Category) {
	for i := range categories {
		if categories[i].ID == id {
			categories[i].Name = updatedC.Name
			categories[i].Description = updatedC.Description
		}
	}
}

func deleteCategory(id int) {
	for i := range categories {
		if categories[i].ID == id {
			categories = append(categories[:i], categories[i+1:]...)
			return
		}
	}
}

func getCategories() []Category {
	return categories
}

func main() {

	http.HandleFunc("/categories/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			http.Error(w, "Invalid Category ID", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case "GET":
			category, err := getCategory(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			json.NewEncoder(w).Encode(category)
		case "PUT":
			var updateC Category
			err := json.NewDecoder(r.Body).Decode(&updateC)
			if err != nil {
				http.Error(w, "request salah", http.StatusBadRequest)
				return
			}
			updateCategory(id, updateC)
			category, err := getCategory(id)
			json.NewEncoder(w).Encode(category)

		case "DELETE":
			deleteCategory(id)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Sukses delete",
			})
		}
	})

	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {
			// GET Function
			json.NewEncoder(w).Encode(getCategories())
			return
		}

		if r.Method == "POST" {
			var newCat Category
			err := json.NewDecoder(r.Body).Decode(&newCat)
			if err != nil {
				http.Error(w, "request salah", http.StatusBadRequest)
				return
			}
			addCategory(newCat)
			json.NewEncoder(w).Encode(getCategories())
			return
		}
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode((map[string]string{
			"status":  "Ok",
			"message": "API Running",
		}))
	})

	fmt.Println("Server running di localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Gagal running server")
	}
}
