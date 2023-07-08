package payment

import "golang-di-playground/policyholder"

type Service interface {
	Pay(policyHolder *policyholder.PolicyHolder, amount float64) error
}
