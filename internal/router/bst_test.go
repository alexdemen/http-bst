package router

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/alexdemen/http-bst/internal/core"
	"github.com/alexdemen/http-bst/internal/dto"
	"github.com/alexdemen/http-bst/internal/service"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchHandle(t *testing.T) {
	treeService := service.NewTreeService(nil)
	err := treeService.Add(100)
	if err != nil {
		t.Fatalf("failed to add key")
	}

	router := NewBSTRouter(*treeService)

	handler := router.searchHandle()

	req := httptest.NewRequest("GET", "http://localhost/search?val=100", nil)
	resp := httptest.NewRecorder()
	handler(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatal("invalid status code")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("invalid body")
	}

	respBody := dto.TreeResponse{}
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		t.Fatalf("failed to parse json: %s", err.Error())
	}

	if respBody.Val != 100 {
		t.Fatalf("invalid node key")
	}
}

func TestSearchHandle_NotFound(t *testing.T) {
	treeService := service.NewTreeService(nil)

	router := NewBSTRouter(*treeService)

	handler := router.searchHandle()

	req := httptest.NewRequest("GET", "http://localhost/search?val=100", nil)
	resp := httptest.NewRecorder()
	handler(resp, req)

	if resp.Code != http.StatusNotFound {
		t.Fatal("invalid status code")
	}
}

func TestInsertHandle(t *testing.T) {
	treeService := service.NewTreeService(nil)

	router := NewBSTRouter(*treeService)

	handler := router.insertHandler()

	body := dto.TreeRequest{Val: 100}
	byteBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("failed parse request body")
	}

	req := httptest.NewRequest("POST", "http://localhost/insert", bytes.NewReader(byteBody))
	resp := httptest.NewRecorder()

	handler(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatal("invalid status code")
	}

	node, err := treeService.Find(100)
	if err != nil {
		t.Fatalf("failed to find key 100")
	}

	if node != body.Val {
		t.Fatalf("invalid founded key")
	}
}

func TestInsertHandle_BadData(t *testing.T) {
	treeService := service.NewTreeService(nil)

	router := NewBSTRouter(*treeService)

	handler := router.insertHandler()

	body := []byte("bad data")

	req := httptest.NewRequest("POST", "http://localhost/insert", bytes.NewReader(body))
	resp := httptest.NewRecorder()

	handler(resp, req)

	if resp.Code != http.StatusBadRequest {
		t.Fatal("invalid status code")
	}

	_, err := treeService.Find(100)
	if err != nil && !errors.Is(err, core.NewKeyNotFoundError(100)) {
		t.Fatalf("failed to find key 100")
	}
}

func TestDeleteHandle(t *testing.T) {
	treeService := service.NewTreeService(nil)
	err := treeService.Add(100)
	if err != nil {
		t.Fatalf("failed to add key")
	}

	router := NewBSTRouter(*treeService)

	handler := router.deleteHandler()

	req := httptest.NewRequest("DELETE", "http://localhost/delete?val=100", nil)
	resp := httptest.NewRecorder()

	handler(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatal("invalid status code")
	}

	_, err = treeService.Find(100)
	if err != nil && !errors.Is(err, core.NewKeyNotFoundError(100)) {
		t.Fatalf("failed to find not existing key")
	}
}
