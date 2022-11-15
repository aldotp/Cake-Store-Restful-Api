package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aldotp/cakestore/routes"
	"github.com/stretchr/testify/assert"
)

func TestGetAllCake(t *testing.T) {
	request, _ := http.NewRequest("GET", "/cakes", nil)
	response := httptest.NewRecorder()
	routes.Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}
func TestGetDetailCake(t *testing.T) {
	request, _ := http.NewRequest("GET", "/cakes/1", nil)
	response := httptest.NewRecorder()
	routes.Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestInsertCake(t *testing.T) {
	testBody := `
	{
		"title": "b", 
		"description": "B cheesecake made of lemon",
		"rating": 5, 
		"image":"https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
	}
	`
	request, _ := http.NewRequest("POST", "/cakes", bytes.NewBufferString(testBody))
	response := httptest.NewRecorder()
	routes.Router().ServeHTTP(response, request)
	assert.Equal(t, 201, response.Code, "OK response is expected")
}

func TestUpdateCake(t *testing.T) {
	testBody := `
	{
		"title": "b", 
		"description": 
		"B cheesecake made of lemon",
		"rating": 6, 
		"image":"https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
	}
	`
	request, _ := http.NewRequest("PATCH", "/cakes/1", bytes.NewBufferString(testBody))
	response := httptest.NewRecorder()
	routes.Router().ServeHTTP(response, request)
	assert.Equal(t, 201, response.Code, "OK response is expected")
}

func TestDeleteCake(t *testing.T) {
	request, _ := http.NewRequest("DELETE", "/cakes/1", nil)
	response := httptest.NewRecorder()
	routes.Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}
