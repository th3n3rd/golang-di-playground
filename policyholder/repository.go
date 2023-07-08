package policyholder

type Repository interface {
	FindById(id string) *PolicyHolder
	Save(policyHolder *PolicyHolder)
}
