package controller

import (
	"bytes"
	"catalogService/model"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	reflect "reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestCreateRestaurantHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockServiceDB := NewMockIServiceDb(ctrl)

	handler := createRestaurant(mockServiceDB)
	restaurantRequest := model.Restaurants{
		Name:     "Test Restaurant",
		Location: "Test Location",
		Menu:     []model.MenuItems{},
	}
	jsonBytes, _ := json.Marshal(restaurantRequest)
	mockServiceDB.EXPECT().
		CreateRestaurant(gomock.Eq(restaurantRequest)).
		Return(model.RestaurantResponse{
			Message: "Restaurant created successfully",
			Restaurant: model.Restaurant{
				ID:       1,
				Name:     "Test Restaurant",
				Location: "Test Location",
				Menu:     []model.MenuItem{},
			},
		}, nil)

	req, err := http.NewRequest("POST", "/restaurants", bytes.NewBuffer(jsonBytes))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()

	handler(rr, req)


	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}


	expectedResponse := model.RestaurantResponse{
		Message: "Restaurant created successfully",
		Restaurant: model.Restaurant{
			ID:       1,
			Name:     "Test Restaurant",
			Location: "Test Location",
			Menu:     []model.MenuItem{},
		},
	}
	actualResponse := model.RestaurantResponse{}
	err = json.NewDecoder(rr.Body).Decode(&actualResponse)
	if err != nil {
		t.Fatalf("Error decoding response body: %v", err)
	}

	if !reflect.DeepEqual(actualResponse, expectedResponse) {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			actualResponse, expectedResponse)
	}
}


func TestFetchAllRestaurantsHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockServiceDB := NewMockIServiceDb(ctrl)

	handler := fetchAllRestaurants(mockServiceDB)
	mockRestaurants := []model.Restaurant{
		{
			ID:       1,
			Name:     "Restaurant 1",
			Location: "Location 1",
			Menu:     []model.MenuItem{},
		},
		{
			ID:       2,
			Name:     "Restaurant 2",
			Location: "Location 2",
			Menu:     []model.MenuItem{},
		},
	}
	mockServiceDB.EXPECT().
		GetAllRestaurants().
		Return(mockRestaurants, nil)

	req, err := http.NewRequest("GET", "/restaurants", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()

	handler(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedResponse := mockRestaurants
	actualResponse := []model.Restaurant{}
	err = json.NewDecoder(rr.Body).Decode(&actualResponse)
	if err != nil {
		t.Fatalf("Error decoding response body: %v", err)
	}

	if !reflect.DeepEqual(actualResponse, expectedResponse) {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			actualResponse, expectedResponse)
	}
}

func TestFetchAllRestaurantsHandler_Error(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockServiceDB := NewMockIServiceDb(ctrl)

	handler := fetchAllRestaurants(mockServiceDB)

	mockServiceDB.EXPECT().
		GetAllRestaurants().
		Return(nil, errors.New("error fetching restaurants"))

	req, err := http.NewRequest("GET", "/restaurants", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()

	handler(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

}