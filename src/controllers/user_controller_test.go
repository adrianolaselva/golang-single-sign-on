package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"oauth2/src/models"
	"testing"
)

var resourceUsers = "/v1/users"

func TestGetUsers(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, resourceUsers, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	var handler = http.Handler(http.HandlerFunc(userController.GetUsers))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("failed to request, return code: %v, payload: %v", rr.Code, rr.Body)
	}

	if rr.Header().Get("Content-Type") != "application/json; charset=UTF-8" {
		t.Errorf("Content-Type returned: %s, expected Application/json; charset=UTF-8", rr.Header().Get("Content-Type"))
	}

	var users = make(map[string]interface{})
	err = json.NewDecoder(rr.Body).Decode(&users)
	if err != nil {
		t.Fatal(err)
	}

	if users["current"] == nil {
		t.Errorf("atribute current not found %v", users)
	}

	if users["total_page"] == nil {
		t.Errorf("atribute total_page not found %v", users)
	}

	if users["per_page"] == nil {
		t.Errorf("atribute per_page not found %v", users)
	}

	if users["total_records"] == nil {
		t.Errorf("atribute total_records not found %v", users)
	}
}

func TestPostUser(t *testing.T) {
	password, err := hash.BCryptGenerate(fmt.Sprintf("test_%s", randomVal))
	if err != nil {
		t.Fatal(err)
	}

	data := make(map[string]interface{})
	data["name"] = fmt.Sprintf("test_%s", randomVal)
	data["last_name"] = fmt.Sprintf("test_%s", randomVal)
	data["username"] = fmt.Sprintf("test_%s", randomVal)
	data["email"] = fmt.Sprintf("test_%s@gmail.com", randomVal)
	data["password"] = password
	data["birthday"] = dateCommon.ConvertFromDateStr("1987-02-11").Format("2006-01-02")
	data["activated"] = true

	payload, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, resourceUsers, bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	var handler = http.Handler(http.HandlerFunc(userController.PostUser))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("failed to create resource, return code: %v, payload: %v", rr.Code, rr.Body)
	}

	var userCreated = make(map[string]interface{})
	err = json.NewDecoder(rr.Body).Decode(&userCreated)
	if err != nil {
		t.Fatal(err)
	}

	if userCreated["id"] == nil {
		t.Errorf("atribute not returned")
	}

	if userCreated["activated"] != data["activated"] {
		t.Errorf("invalid atribute activated, returned %v, extected %v", userCreated["activated"], true)
	}

	if userCreated["name"] != data["name"] {
		t.Errorf("invalid atribute name, returned %v, extected %v", userCreated["activated"], true)
	}

	if userCreated["last_name"] != data["last_name"] {
		t.Errorf("invalid atribute last_name, returned %v, extected %v", userCreated["activated"], true)
	}

	if userCreated["username"] != data["username"] {
		t.Errorf("invalid atribute username, returned %v, extected %v", userCreated["activated"], true)
	}

	if userCreated["email"] != data["email"] {
		t.Errorf("invalid atribute email, returned %v, extected %v", userCreated["activated"], true)
	}

	if userCreated["birthday"] != data["birthday"] {
		t.Errorf("invalid atribute birthday, returned %v, extected %v", userCreated["activated"], true)
	}

	if userCreated["password"] != nil {
		t.Error("password must not be returned")
	}
}

func TestGetUser(t *testing.T) {
	var orderBy, orderDir, limit, page, filters = "created_at", "asc", 1, 1, make(map[string]interface{})
	paginateData, err := userService.Paginate(&filters, &orderBy, &orderDir, &limit, &page)
	if paginateData == nil {
		t.Fatal("there are no records")
	}

	firstRow := paginateData.Data.([]*models.User)[0]
	if firstRow == nil {
		t.Fatal("failed to return first row for load user")
	}

	req, err := http.NewRequest(http.MethodGet, resourceUsers, nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{
		"uuid": firstRow.ID,
	})

	rr := httptest.NewRecorder()
	var handler = http.Handler(http.HandlerFunc(userController.GetUser))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("failed to load resource, return code: %v, payload: %v", rr.Code, rr.Body)
	}

	var userLoad = make(map[string]interface{})
	err = json.NewDecoder(rr.Body).Decode(&userLoad)
	if err != nil {
		t.Fatal(err)
	}

	if userLoad["id"] == nil {
		t.Errorf("atribute id not found: %v", userLoad)
	}

	if userLoad["activated"] == nil {
		t.Errorf("atribute activated not found: %v", userLoad)
	}

	if userLoad["name"] == nil {
		t.Errorf("atribute name not found: %v", userLoad)
	}

	if userLoad["last_name"] == nil {
		t.Errorf("atribute last_name not found: %v", userLoad)
	}

	if userLoad["username"] == nil {
		t.Errorf("atribute username not found: %v", userLoad)
	}

	if userLoad["email"] == nil {
		t.Errorf("atribute email not found: %v", userLoad)
	}

	if userLoad["birthday"] == nil {
		t.Errorf("atribute birthday not found: %v", userLoad)
	}

	if userLoad["password"] != nil {
		t.Errorf("atribute password found: %v", userLoad)
	}
}

func TestPutUser(t *testing.T) {
	var orderBy, orderDir, limit, page, filters = "created_at", "desc", 1, 1, make(map[string]interface{})
	paginateData, err := userService.Paginate(&filters, &orderBy, &orderDir, &limit, &page)
	if paginateData == nil {
		t.Fatal("there are no records")
	}

	firstRow := paginateData.Data.([]*models.User)[0]
	if firstRow == nil {
		t.Fatal("failed to return first row for load user")
	}

	password, err := hash.BCryptGenerate(fmt.Sprintf("test_%s", randomVal))
	if err != nil {
		t.Fatal(err)
	}

	data := make(map[string]interface{})
	data["name"] = fmt.Sprintf("test_%s", randomVal)
	data["last_name"] = fmt.Sprintf("test_%s", randomVal)
	data["username"] = fmt.Sprintf("test_%s", randomVal)
	data["email"] = fmt.Sprintf("test_%s@gmail.com", randomVal)
	data["password"] = password
	data["birthday"] = dateCommon.ConvertFromDateStr("1987-02-11").Format("2006-01-02")
	data["activated"] = true

	payload, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPut, resourceUsers, bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{
		"uuid": firstRow.ID,
	})

	rr := httptest.NewRecorder()
	var handler = http.Handler(http.HandlerFunc(userController.PutUser))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("failed to update resource, return code: %v, payload: %v", rr.Code, rr.Body)
	}

	var userCreated = make(map[string]interface{})
	err = json.NewDecoder(rr.Body).Decode(&userCreated)
	if err != nil {
		t.Fatal(err)
	}

	if userCreated["id"] == nil {
		t.Errorf("atribute not returned")
	}

	if userCreated["activated"] != data["activated"] {
		t.Errorf("invalid atribute activated, returned %v, extected %v", userCreated["activated"], true)
	}

	if userCreated["name"] != data["name"] {
		t.Errorf("invalid atribute name, returned %v, extected %v", userCreated["activated"], true)
	}

	if userCreated["last_name"] != data["last_name"] {
		t.Errorf("invalid atribute last_name, returned %v, extected %v", userCreated["activated"], true)
	}

	if userCreated["username"] != data["username"] {
		t.Errorf("invalid atribute username, returned %v, extected %v", userCreated["activated"], true)
	}

	if userCreated["email"] != data["email"] {
		t.Errorf("invalid atribute email, returned %v, extected %v", userCreated["activated"], true)
	}

	if userCreated["birthday"] != data["birthday"] {
		t.Errorf("invalid atribute birthday, returned %v, extected %v", userCreated["activated"], true)
	}

	if userCreated["password"] != nil {
		t.Error("password must not be returned")
	}
}

func TestDeleteUser(t *testing.T) {
	var orderBy, orderDir, limit, page, filters = "created_at", "desc", 1, 1, make(map[string]interface{})
	paginateData, err := userService.Paginate(&filters, &orderBy, &orderDir, &limit, &page)
	if paginateData == nil {
		t.Fatal("there are no records")
	}

	firstRow := paginateData.Data.([]*models.User)[0]
	if firstRow == nil {
		t.Fatal("failed to return first row for load user")
	}

	req, err := http.NewRequest(http.MethodDelete, resourceUsers, nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{
		"uuid": firstRow.ID,
	})

	rr := httptest.NewRecorder()
	var handler = http.Handler(http.HandlerFunc(userController.DeleteUser))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("failed to delete resource, return code: %v, payload: %v", rr.Code, rr.Body)
	}
}

func TestGetUserNotFound(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, resourceUsers, nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{
		"uuid": "not_found",
	})

	rr := httptest.NewRecorder()
	var handler = http.Handler(http.HandlerFunc(userController.DeleteUser))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("failed to load resource, return code: %v, payload: %v", rr.Code, rr.Body)
	}
}

func TestDeleteUserNotFound(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, resourceUsers, nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{
		"uuid": "not_found",
	})

	rr := httptest.NewRecorder()
	var handler = http.Handler(http.HandlerFunc(userController.DeleteUser))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("failed to delete resource, return code: %v, payload: %v", rr.Code, rr.Body)
	}
}

func TestPostUserBadRequest(t *testing.T) {
	data := make(map[string]interface{})
	data["name"] = fmt.Sprintf("test_%s", randomVal)

	payload, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, resourceUsers, bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	var handler = http.Handler(http.HandlerFunc(userController.PostUser))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("failed to create resource, return code: %v, payload: %v", rr.Code, rr.Body)
	}
}

func TestPutUserBadRequest(t *testing.T) {
	data := make(map[string]interface{})
	data["name"] = fmt.Sprintf("test_%s", randomVal)

	payload, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPut, resourceUsers, bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{
		"uuid": "bad",
	})

	rr := httptest.NewRecorder()
	var handler = http.Handler(http.HandlerFunc(userController.PutUser))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("failed to update resource, return code: %v, payload: %v", rr.Code, rr.Body)
	}

	var userCreated = make(map[string]interface{})
	err = json.NewDecoder(rr.Body).Decode(&userCreated)
	if err != nil {
		t.Fatal(err)
	}
}