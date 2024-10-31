package api

import "errors"

// validate checks the RiskRequest fields for validity.
//
// It ensures that the Title, Description, and State are non-empty.
// Additionally, it verifies that the State is one of the allowed values:
// [open, closed, accepted, investigating].
// Returns an error if any validation rule is violated, otherwise returns nil.
func (rs *RiskRequest) validate() error {
	if len(rs.Title) == 0 {
		return errors.New("risk ritle is required")
	}

	if len(rs.Description) == 0 {
		return errors.New("risk description is required")
	}

	if len(rs.State) == 0 {
		return errors.New("risk state is required")
	}

	if rs.State != RiskRequestStateOpen && rs.State != RiskRequestStateAccepted && rs.State != RiskRequestStateClosed && rs.State != RiskRequestStateInvestigating {
		return errors.New("allowed risk states are [open, closed, accepted, investigating]")
	}

	return nil
}
