package policyholder

import (
	"fmt"
	"golang-di-playground/claim"
	"math/rand"
	"time"
)

type PolicyHolder struct {
	Id       string
	FullName string
}

func NewPolicyHolder(fullName string) *PolicyHolder {
	return &PolicyHolder{FullName: fullName, Id: generateRandomId()}
}

func (p *PolicyHolder) DoesAgreeToEstimate(claim *claim.Claim) bool {
	// In a real application, this would involve more complex logic
	// For this example, let's assume the policyholder agrees if the estimate is less than $2000
	fmt.Printf("Agreeing on the estimate since is less than $2000\n")
	return claim.Estimate < 2000
}

func generateRandomId() string {
	generator := rand.NewSource(time.Now().UnixNano())
	return fmt.Sprintf("%d", generator.Int63())
}
