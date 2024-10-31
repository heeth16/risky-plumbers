package api

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type RiskStore struct {
	sync.Mutex
	Risks map[uuid.UUID]Risk
}

// Ensure RiskStore always implements the ServerInterface
var _ ServerInterface = (*RiskStore)(nil)

func NewRiskStore() *RiskStore {
	return &RiskStore{
		Risks: make(map[uuid.UUID]Risk),
	}
}

// GetRisks retrieves a list of Risks.
//
// It responds with a JSON body containing all stored Risks and a 200 OK status code.
func (rs *RiskStore) GetRisks(w http.ResponseWriter, r *http.Request) {
	rs.Lock()
	defer rs.Unlock()

	risks := make([]*Risk, 0, len(rs.Risks))
	for _, val := range rs.Risks {
		risks = append(risks, &val)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(risks)
}

// PostRisks creates a new Risk based on the provided request body.
//
// It responds with a JSON body containing the auto-generated UUID of the created Risk and a 201 Created status code.
// If the request body is not valid JSON, it responds with a 400 Bad Request status code.
// If the Risk in the request body is invalid, it responds with a 400 Bad Request status code.
func (rs *RiskStore) PostRisks(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		returnRiskStoreResponse(w, http.StatusBadRequest, "Invalid Content-Type format. Supported is application/json")
		return
	}

	var newRisk RiskRequest
	if err := json.NewDecoder(r.Body).Decode(&newRisk); err != nil {
		returnRiskStoreResponse(w, http.StatusBadRequest, "Invalid format for Risk")
		return
	}

	if err := newRisk.validate(); err != nil {
		returnRiskStoreResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id := uuid.New()

	rs.Lock()
	defer rs.Unlock()

	rs.Risks[id] = Risk{
		Id:          id,
		Title:       newRisk.Title,
		Description: newRisk.Description,
		State:       RiskState(newRisk.State),
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(struct { // Unnamed struct
		ID uuid.UUID
	}{
		ID: id,
	})
}

// GetRisksId retrieves an individual Risk by its UUID.
//
// It responds with a JSON body containing the Risk and a 200 OK status code.
// If the Risk ID is not found, it responds with a 400 Bad Request status code.
func (rs *RiskStore) GetRisksId(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	rs.Lock()
	defer rs.Unlock()

	if _, ok := rs.Risks[id]; !ok {
		returnRiskStoreResponse(w, http.StatusBadRequest, "Risk ID not found")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rs.Risks[id])
}
