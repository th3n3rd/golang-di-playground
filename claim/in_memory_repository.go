package claim

type InMemoryRepository struct {
	claims map[string]*Claim
}

func NewInMemoryRepository() Repository {
	return &InMemoryRepository{claims: make(map[string]*Claim)}
}

func (r *InMemoryRepository) Save(claim *Claim) {
	r.claims[claim.Number] = claim
}

func (r *InMemoryRepository) FindByNumber(number string) *Claim {
	return r.claims[number]
}

func (r *InMemoryRepository) FindAllByPolicyHolderId(policyHolderId string) []*Claim {
	var matchingClaims []*Claim
	for _, claim := range r.claims {
		if claim.PolicyHolderId == policyHolderId {
			matchingClaims = append(matchingClaims, claim)
		}
	}
	return matchingClaims
}
