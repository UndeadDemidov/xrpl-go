package transaction

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// PathStep represents a single step in a payment path, including optional account, currency, and issuer.
type PathStep struct {
	Account  types.Address `json:"account,omitempty"`
	Currency string        `json:"currency,omitempty"`
	Issuer   types.Address `json:"issuer,omitempty"`
}

// Flatten returns a map representation of the PathStep.
func (p *PathStep) Flatten() map[string]interface{} {
	flattened := make(map[string]interface{})

	if p.Account != "" {
		flattened["account"] = p.Account.String()
	}

	if p.Currency != "" {
		flattened["currency"] = p.Currency
	}

	if p.Issuer != "" {
		flattened["issuer"] = p.Issuer.String()
	}

	return flattened

}
