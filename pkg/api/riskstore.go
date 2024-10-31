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
	w.Write([]byte("Calling PostRisks\n"))

	var newRisk Risk
	if err := json.NewDecoder(r.Body).Decode(&newRisk); err != nil {
		raiseRiskStoreError(w, http.StatusBadRequest, "Invalid format for Risk")
		return
	}
}

func (rs *RiskStore) GetRisksId(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	w.Write([]byte("Calling GetRisks\n"))
}
