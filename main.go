package main

import (
	"fmt"
	"go.uber.org/dig"
	"golang-di-playground/claim"
	"golang-di-playground/company"
	"golang-di-playground/notification"
	"golang-di-playground/payment"
	"golang-di-playground/policyholder"
	"golang-di-playground/workflow"
)

func main() {
	container := dig.New()

	policyholder.Provide(container)
	company.Provide(container)
	claim.Provide(container)
	notification.Provide(container)
	payment.Provide(container)

	policyHolder := policyholder.NewPolicyHolder("John Doe")

	err := container.Invoke(func(company *company.InsuranceCompany) {
		company.PolicyHolderRepo.Save(policyHolder)
	})

	if err != nil {
		fmt.Printf("Failed to prepare data: %v\n", err)
	}

	err = container.Invoke(workflow.ClaimProcessing(policyHolder.Id))

	if err != nil {
		fmt.Printf("Failed to execute workflow: %v\n", err)
	}
}
