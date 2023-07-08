package claim

import (
	"fmt"
)

type Claim struct {
	Number         string
	Details        string
	PolicyHolderId string
	Estimate       float64
	EstimateAgreed bool
	Paid           bool
}

func (c *Claim) Resolve() {
	fmt.Printf("Claim %s resolved\n", c.Number)
}

func (c *Claim) MarkAsPaid() {
	c.Paid = true
}

func (c *Claim) EstimateDamages(value float64) {
	c.Estimate = value
}

func (c *Claim) AgreeToEstimate() {
	c.EstimateAgreed = true
}
