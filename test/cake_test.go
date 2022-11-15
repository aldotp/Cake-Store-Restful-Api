package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aldotp/cakestore/controller"
	"github.com/aldotp/cakestore/routes"
	"github.com/stretchr/testify/assert"
)

func TestGetAllCake(t *testing.T) {
	request, _ := http.NewRequest("GET", "/cakes", nil)
	response := httptest.NewRecorder()
	routes.Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")

}
func TestGetCake(t *testing.T) {
	req, err := http.NewRequest("GET", "/entries", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetCakes)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `test`
	if rr.Body.String() != expected {
		t.Errorf("Handler return unexpected body : got %v want %v", rr.Body.String(), expected)
	}
}

func TestInsertCake(t *testing.T) {
	testBody := `{"title": "b", "description": "B cheesecake made of lemon","rating": 5, "image":"https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"}`
	request, _ := http.NewRequest("POST", "/cakes", bytes.NewBufferString(testBody))
	response := httptest.NewRecorder()
	routes.Router().ServeHTTP(response, request)
	assert.Equal(t, 201, response.Code, "OK response is expected")
}
