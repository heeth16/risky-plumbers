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
	w.Write([]byte("Calling GetRisks\n"))
}

func (rs *RiskStore) PostRisks(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		raiseRiskStoreError(w, http.StatusBadRequest, "Invalid Content-Type format. Supported is application/json")
		return
	}

	var newRisk RiskRequest
	if err := json.NewDecoder(r.Body).Decode(&newRisk); err != nil {
		raiseRiskStoreError(w, http.StatusBadRequest, "Invalid format for Risk")
		return
	}

	if err := newRisk.validate(); err != nil {
		raiseRiskStoreError(w, http.StatusBadRequest, err.Error())
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
	w.Write([]byte("Calling GetRisks\n"))
}
