package policyholder

type InMemoryRepository struct {
	policyHolders map[string]*PolicyHolder
}

func NewInMemoryRepository() Repository {
	return &InMemoryRepository{policyHolders: make(map[string]*PolicyHolder)}
}

func (r *InMemoryRepository) Save(policyHolder *PolicyHolder) {
	r.policyHolders[policyHolder.Id] = policyHolder
}

func (r *InMemoryRepository) FindById(id string) *PolicyHolder {
	return r.policyHolders[id]
}
