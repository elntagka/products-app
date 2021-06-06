package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProductsHandler_AllProducts(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(ProductsHandler))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Errorf("failed to perform get request: %s", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Didn't get 200 got %d", res.StatusCode)
	}
	b, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Errorf("unable to read response: %s", err)
	}
	p := []product{}

	err = json.Unmarshal(b, &p)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	if len(p) != 40 {
		t.Errorf("unexpected num of products, got %d", len(p))
	}

}

func TestProductsHandler_GetPage(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(ProductsHandler))
	defer ts.Close()

	res, err := http.Get(ts.URL + "?page=5")
	if err != nil {
		t.Errorf("failed to perform get request: %s", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Didn't get 200 got %d", res.StatusCode)
	}
	b, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Errorf("unable to read response: %s", err)
	}
	p := []product{}

	err = json.Unmarshal(b, &p)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	if len(p) != 5 {
		t.Errorf("unexpected num of products, got %d", len(p))
	}

}

func TestProductsHandler_MalformedQuery(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(ProductsHandler))
	defer ts.Close()

	res, err := http.Get(ts.URL + "?page=five")
	if err != nil {
		t.Errorf("failed to perform get request: %s", err)
	}
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Didn't get 400 got %d", res.StatusCode)
	}

	b, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Errorf("unable to read response: %s", err)
	}
	if string(b) != `{"message": "page should be integer"}` {
		t.Errorf("Expected bad request message but got: %s", b)
	}

}
