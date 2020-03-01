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

var resourceRoles = "/v1/roles"

func TestGetRoles(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, resourceRoles, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	var handler = http.Handler(http.HandlerFunc(roleController.GetRoles))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("failed to request, return code: %v, payload: %v", rr.Code, rr.Body)
	}

	if rr.Header().Get("Content-Type") != "application/json; charset=UTF-8" {
		t.Errorf("Content-Type returned: %s, expected Application/json; charset=UTF-8", rr.Header().Get("Content-Type"))
	}

	var roles = make(map[string]interface{})
	err = json.NewDecoder(rr.Body).Decode(&roles)
	if err != nil {
		t.Fatal(err)
	}

	if roles["current"] == nil {
		t.Errorf("atribute current not found %v", roles)
	}

	if roles["total_page"] == nil {
		t.Errorf("atribute total_page not found %v", roles)
	}

	if roles["per_page"] == nil {
		t.Errorf("atribute per_page not found %v", roles)
	}

	if roles["total_records"] == nil {
		t.Errorf("atribute total_records not found %v", roles)
	}
}

func TestPostRole(t *testing.T) {
	data := make(map[string]interface{})
	data["name"] = fmt.Sprintf("test_%s", randomVal)

	payload, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, resourceRoles, bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	var handler = http.Handler(http.HandlerFunc(roleController.PostRole))
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

	if userCreated["name"] != data["name"] {
		t.Errorf("invalid atribute name, returned %v, extected %v", userCreated["name"], data["name"])
	}
}

func TestGetRole(t *testing.T) {
	var orderBy, orderDir, limit, page, filters = "name", "asc", 1, 1, make(map[string]interface{})
	paginateData, err := roleService.Paginate(&filters, &orderBy, &orderDir, &limit, &page)
	if paginateData == nil {
		t.Fatal("there are no records")
	}

	firstRow := paginateData.Data.([]*models.Role)[0]
	if firstRow == nil {
		t.Fatal("failed to return first row for load user")
	}

	req, err := http.NewRequest(http.MethodGet, resourceRoles, nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{
		"uuid": firstRow.ID,
	})

	rr := httptest.NewRecorder()
	var handler = http.Handler(http.HandlerFunc(roleController.GetRole))
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

	if userLoad["name"] == nil {
		t.Errorf("atribute name not found: %v", userLoad)
	}
}

func TestPutRole(t *testing.T) {
	var orderBy, orderDir, limit, page, filters = "name", "desc", 1, 1, make(map[string]interface{})
	paginateData, err := roleService.Paginate(&filters, &orderBy, &orderDir, &limit, &page)
	if paginateData == nil {
		t.Fatal("there are no records")
	}

	firstRow := paginateData.Data.([]*models.Role)[0]
	if firstRow == nil {
		t.Fatal("failed to return first row for load user")
	}

	data := make(map[string]interface{})
	data["name"] = fmt.Sprintf("test_%s", randomVal)

	payload, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPut, resourceRoles, bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{
		"uuid": firstRow.ID,
	})

	rr := httptest.NewRecorder()
	var handler = http.Handler(http.HandlerFunc(roleController.PutRole))
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

	if userCreated["name"] != data["name"] {
		t.Errorf("invalid atribute name, returned %v, extected %v", userCreated["activated"], true)
	}
}

func TestDeleteRole(t *testing.T) {
	var orderBy, orderDir, limit, page, filters = "name", "desc", 1, 1, make(map[string]interface{})
	paginateData, err := roleService.Paginate(&filters, &orderBy, &orderDir, &limit, &page)
	if paginateData == nil {
		t.Fatal("there are no records")
	}

	firstRow := paginateData.Data.([]*models.Role)[0]
	if firstRow == nil {
		t.Fatal("failed to return first row for load user")
	}

	req, err := http.NewRequest(http.MethodDelete, resourceRoles, nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{
		"uuid": firstRow.ID,
	})

	rr := httptest.NewRecorder()
	var handler = http.Handler(http.HandlerFunc(roleController.DeleteRole))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("failed to delete resource, return code: %v, payload: %v", rr.Code, rr.Body)
	}
}

func TestGetRoleNotFound(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, resourceRoles, nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{
		"uuid": "not_found",
	})

	rr := httptest.NewRecorder()
	var handler = http.Handler(http.HandlerFunc(roleController.DeleteRole))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("failed to load resource, return code: %v, payload: %v", rr.Code, rr.Body)
	}
}

func TestDeleteRoleNotFound(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, resourceRoles, nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{
		"uuid": "not_found",
	})

	rr := httptest.NewRecorder()
	var handler = http.Handler(http.HandlerFunc(roleController.DeleteRole))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("failed to delete resource, return code: %v, payload: %v", rr.Code, rr.Body)
	}
}

func TestPostRoleBadRequest(t *testing.T) {
	data := make(map[string]interface{})

	payload, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, resourceRoles, bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	var handler = http.Handler(http.HandlerFunc(roleController.PostRole))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("failed to create resource, return code: %v, payload: %v", rr.Code, rr.Body)
	}
}

func TestPutRoleBadRequest(t *testing.T) {
	var orderBy, orderDir, limit, page, filters = "name", "desc", 1, 1, make(map[string]interface{})
	paginateData, err := roleService.Paginate(&filters, &orderBy, &orderDir, &limit, &page)
	if paginateData == nil {
		t.Fatal("there are no records")
	}

	firstRow := paginateData.Data.([]*models.Role)[0]
	if firstRow == nil {
		t.Fatal("failed to return first row for load role")
	}

	data := make(map[string]interface{})

	payload, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPut, resourceRoles, bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{
		"uuid": firstRow.ID,
	})

	rr := httptest.NewRecorder()
	var handler = http.Handler(http.HandlerFunc(roleController.PutRole))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("failed to update resource, return code: %v, payload: %v", rr.Code, rr.Body)
	}
}