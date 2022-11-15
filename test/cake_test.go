package test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/aldotp/cakestore/controller"
	"github.com/aldotp/cakestore/routes"
)

func TestGetAllCake(t *testing.T) {
	request, err := http.NewRequest("GET", "/cakes", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	routes.Router().ServeHTTP(response, request)
	if status := response.Code; status != http.StatusOK {
		t.Errorf("return wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `[{"id":2,"title":"Blue Cheese Cake","description":"Good taste with blue color","rating":8,"image":"https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg","created_at":"2020-02-01T10:56:31+07:00","updated_at":"2020-02-13T09:30:23+07:00"},{"id":1,"title":"Lemon cheesecake","description":"A cheesecake made of lemon","rating":7,"image":"https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg","created_at":"2020-02-01T10:56:31+07:00","updated_at":"2020-02-13T09:30:23+07:00"}]`

	if response.Body.String() != expected {
		t.Errorf("return unexpected body: got %v want %v", response.Body.String(), expected)
	}
	fmt.Println("Testing GetAllCake Success. Tidak ditemukan error, code:", response.Code)
}

func TestGetDetailCake(t *testing.T) {
	request, err := http.NewRequest("GET", "/cakes/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	routes.Router().ServeHTTP(response, request)
	if status := response.Code; status != http.StatusOK {
		t.Errorf("return wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"id":1,"title":"Lemon cheesecake","description":"A cheesecake made of lemon","rating":7,"image":"https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg","created_at":"2020-02-01T10:56:31+07:00","updated_at":"2020-02-13T09:30:23+07:00"}`

	if response.Body.String() != expected {
		t.Errorf("return unexpected body: got %v want %v", response.Body.String(), expected)
	}
	fmt.Println("Testing GetDetailCake Success. Tidak ditemukan error, code:", response.Code)
}

func TestInsertCake(t *testing.T) {
	testBody := `
	{
		"id":3,
		"title": "Cheese Cake", 
		"description": "Creamy and Salty Cake",
		"rating": 9, 
		"image":"https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
	}
	`
	request, err := http.NewRequest("POST", "/cakes", bytes.NewBufferString(testBody))
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	routes.Router().ServeHTTP(response, request)

	if status := response.Code; status != http.StatusCreated {
		t.Errorf("return wrong status code: got %v want %v", status, http.StatusCreated)
	}

	fmt.Println("Testing InsertCake Success. Tidak ditemukan error, code:", response.Code)
}

func TestUpdateCake(t *testing.T) {
	testBody := `
	{
		"id":3,
		"title": "Cheese Cake Yummy", 
		"description": "Creamy and Salty Cake Enak Banget",
		"rating": 9, 
		"image":"https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
	}
	`
	request, err := http.NewRequest("PATCH", "/cakes/3", bytes.NewBufferString(testBody))
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	routes.Router().ServeHTTP(response, request)

	if status := response.Code; status != http.StatusCreated {
		t.Errorf("return wrong status code: got %v want %v", status, http.StatusCreated)
	}

	fmt.Println("Testing Updatecake Success. Tidak ditemukan error, code:", response.Code)

}

func TestDeleteCake(t *testing.T) {
	request, err := http.NewRequest("DELETE", "/cakes/3", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	routes.Router().ServeHTTP(response, request)
	if status := response.Code; status != http.StatusOK {
		t.Errorf("return wrong status code: got %v want %v", status, http.StatusOK)
	}

	fmt.Println("Testing Deletecake Success. Tidak ditemukan error, code:", response.Code)
}
