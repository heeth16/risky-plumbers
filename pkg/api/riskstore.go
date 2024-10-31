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

func (rs *RiskStore) GetRisks(w http.ResponseWriter, r *http.Request) {

}

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

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct { // Unnamed struct
		ID uuid.UUID
	}{
		ID: id,
	})
}

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
