package workflow

import (
	"fmt"
	"golang-di-playground/company"
)

// ClaimProcessing implements the following steps:
// 1. Policy-holder reports an incident to the insurance company, either online or via phone.
// 2. Policy-holder provides necessary details about the incident including when, where, and how it happened.
// 3. Policy-holder receives a claim number to track the process.
// 4. Insurance company dispatches an adjuster to verify the incident and estimate the cost.
// 5. Policy-holder receives an estimate and agrees to it.
// 6. Policy-holder receives payment or services for the claim.
// 7. Policy-holder receives notification about the claim resolution.
func ClaimProcessing(policyHolderId string) func(company *company.InsuranceCompany) {
	return func(company *company.InsuranceCompany) {
		policyHolder := company.FindPolicyHolder(policyHolderId)
		claim := company.ReportIncident(policyHolder, "Car accident")
		company.AssessDamage(claim)
		company.EstimateAgreement(claim)
		company.PayClaim(claim)
		company.ResolveClaim(claim)
		savedClaim := company.FindClaim(claim.Number)
		fmt.Printf("Retrieved claim: %+v\n", savedClaim)
	}
}
