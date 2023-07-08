package workflow

import (
	"golang-di-playground/claim"
	"golang-di-playground/company"
	"golang-di-playground/policyholder"
	"testing"
)

func TestClaimProcessingWorkflow(t *testing.T) {
	policyHoldersRepo := policyholder.NewInMemoryRepository()
	johnDoe := policyholder.NewPolicyHolder("John Doe")
	policyHoldersRepo.Save(johnDoe)

	paymentService := &FakePaymentService{}
	notificationService := &FakeNotificationService{}

	claimsRepo := claim.NewInMemoryRepository()
	insuranceCompany := company.NewInsuranceCompany(
		company.NewAdjuster(),
		claimsRepo,
		policyHoldersRepo,
		paymentService,
		notificationService,
	)

	ClaimProcessing(johnDoe.Id)(insuranceCompany)

	expectNotificationIsSent(t, notificationService, "Your claim has been created")
	expectNotificationIsSent(t, notificationService, "Your claim has been estimated")
	expectNotificationIsSent(t, notificationService, "Thank you for reviewing and confirming the estimate")
	expectNotificationIsSent(t, notificationService, "Your claim has been paid")

	expectPaymentIsMade(t, paymentService, 1000)

	expectNotificationIsSent(t, notificationService, "Your claim has been resolved")

	claims := claimsRepo.FindAllByPolicyHolderId(johnDoe.Id)
	expectClaimHasBeenProcessed(t, claims[0])
}

func expectClaimHasBeenProcessed(t *testing.T, actualClaim *claim.Claim) {
	if actualClaim.Estimate == 0 {
		t.Errorf("Expected claim estimate to be greater than 0")
	}
	if !actualClaim.EstimateAgreed {
		t.Errorf("Expected claim estimate to be agreed")
	}
	if !actualClaim.Paid {
		t.Errorf("Expected claim to be paid")
	}
}

func expectPaymentIsMade(t *testing.T, paymentService *FakePaymentService, expected float64) {
	if amount, ok := paymentService.NextPayment(); !ok || amount != expected {
		t.Errorf("Expected payment %.2f, but got %.2f", expected, amount)
	}
}

func expectNotificationIsSent(t *testing.T, notificationService *FakeNotificationService, expected string) {
	if message, ok := notificationService.NextMessage(); !ok || message != expected {
		t.Errorf("Expected notification %q, but got %q", expected, message)
	}
}
