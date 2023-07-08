package workflow

import "golang-di-playground/policyholder"

type FakePaymentService struct {
	Payments []float64
}

func (p *FakePaymentService) Pay(policyHolder *policyholder.PolicyHolder, amount float64) error {
	p.Payments = append(p.Payments, amount)
	return nil
}

func (p *FakePaymentService) NextPayment() (float64, bool) {
	if len(p.Payments) == 0 {
		return 0, false
	}
	payment := p.Payments[0]
	p.Payments = p.Payments[1:]
	return payment, true
}
