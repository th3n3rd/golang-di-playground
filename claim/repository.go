package claim

type Repository interface {
	Save(claim *Claim)
	FindByNumber(number string) *Claim
	FindAllByPolicyHolderId(policyHolderId string) []*Claim
}
