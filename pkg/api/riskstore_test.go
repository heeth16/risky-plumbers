package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/stretchr/testify/assert"
)

func TestGetRisks(t *testing.T) {
	riskStore := NewRiskStore()
	id := uuid.New()

	riskStore.Risks[id] = Risk{
		Id:          id,
		Title:       "Test Risk",
		Description: "This is a test risk",
		State:       "open",
	}

	req, err := http.NewRequest("GET", "/risks", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(riskStore.GetRisks)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var risks []*Risk
	err = json.NewDecoder(rr.Body).Decode(&risks)
	assert.NoError(t, err)
	assert.NotEmpty(t, risks)
}

func TestPostRisks(t *testing.T) {
	riskStore := NewRiskStore()
	riskRequest := RiskRequest{
		Title:       "Test Risk",
		Description: "This is a test risk",
		State:       "open",
	}

	body, _ := json.Marshal(riskRequest)
	req, err := http.NewRequest("POST", "/risks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(riskStore.PostRisks)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var response struct {
		ID uuid.UUID `json:"ID"`
	}
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, response.ID)
}

func TestGetRisksId(t *testing.T) {
	riskStore := NewRiskStore()
	id := uuid.New()
	riskStore.Risks[id] = Risk{
		Id:          id,
		Title:       "Test Risk",
		Description: "This is a test risk",
		State:       "open",
	}

	req, err := http.NewRequest("GET", "/risks/"+id.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		riskStore.GetRisksId(w, r, openapi_types.UUID(id))
	})

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var risk Risk
	err = json.NewDecoder(rr.Body).Decode(&risk)
	assert.NoError(t, err)
	assert.Equal(t, id, risk.Id)
}

func TestGetRisksId_NotFound(t *testing.T) {
	riskStore := NewRiskStore()
	id := uuid.New() // Use an ID that doesn't exist

	req, err := http.NewRequest("GET", "/risks/"+id.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		riskStore.GetRisksId(w, r, openapi_types.UUID(id))
	})

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
