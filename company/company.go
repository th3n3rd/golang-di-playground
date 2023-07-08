package company

import (
	"golang-di-playground/claim"
	"golang-di-playground/notification"
	"golang-di-playground/payment"
	"golang-di-playground/policyholder"
)

type InsuranceCompany struct {
	Adjuster         Adjuster
	ClaimRepo        claim.Repository
	PolicyHolderRepo policyholder.Repository
	Payment          payment.Service
	Notifier         notification.Service
}

func NewInsuranceCompany(
	adjuster *Adjuster,
	claimRepo claim.Repository,
	policyHolderRepo policyholder.Repository,
	payment payment.Service,
	notifier notification.Service,
) *InsuranceCompany {
	return &InsuranceCompany{
		Adjuster:         *adjuster,
		ClaimRepo:        claimRepo,
		PolicyHolderRepo: policyHolderRepo,
		Payment:          payment,
		Notifier:         notifier,
	}
}

func (c *InsuranceCompany) FindPolicyHolder(policyHolderId string) *policyholder.PolicyHolder {
	return c.PolicyHolderRepo.FindById(policyHolderId)
}

func (c *InsuranceCompany) ReportIncident(policyHolder *policyholder.PolicyHolder, details string) *claim.Claim {
	newClaim := &claim.Claim{
		Number:         "123456",
		Details:        details,
		PolicyHolderId: policyHolder.Id,
	}
	c.ClaimRepo.Save(newClaim)
	c.Notifier.Notify(policyHolder, "Your claim has been created")
	return newClaim
}

func (c *InsuranceCompany) AssessDamage(newClaim *claim.Claim) {
	policyHolder := c.PolicyHolderRepo.FindById(newClaim.PolicyHolderId)
	estimate := c.Adjuster.VerifyAndEstimate(newClaim.Details)
	newClaim.EstimateDamages(estimate)
	c.ClaimRepo.Save(newClaim)
	c.Notifier.Notify(policyHolder, "Your claim has been estimated")
}

func (c *InsuranceCompany) EstimateAgreement(estimatedClaim *claim.Claim) {
	policyHolder := c.PolicyHolderRepo.FindById(estimatedClaim.PolicyHolderId)
	if policyHolder.DoesAgreeToEstimate(estimatedClaim) {
		estimatedClaim.AgreeToEstimate()
		c.ClaimRepo.Save(estimatedClaim)
		c.Notifier.Notify(policyHolder, "Thank you for reviewing and confirming the estimate")
	}
}

func (c *InsuranceCompany) PayClaim(agreedClaim *claim.Claim) {
	policyHolder := c.PolicyHolderRepo.FindById(agreedClaim.PolicyHolderId)
	_ = c.Payment.Pay(policyHolder, agreedClaim.Estimate)
	agreedClaim.MarkAsPaid()
	c.ClaimRepo.Save(agreedClaim)
	c.Notifier.Notify(policyHolder, "Your claim has been paid")
}

func (c *InsuranceCompany) ResolveClaim(processedClaim *claim.Claim) {
	processedClaim.Resolve()
	c.ClaimRepo.Save(processedClaim)
	policyHolder := c.PolicyHolderRepo.FindById(processedClaim.PolicyHolderId)
	c.Notifier.Notify(policyHolder, "Your claim has been resolved")
}

func (c *InsuranceCompany) FindClaim(claimNumber string) *claim.Claim {
	return c.ClaimRepo.FindByNumber(claimNumber)
}
