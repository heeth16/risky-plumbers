package api

import "errors"

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
