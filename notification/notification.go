package notification

import (
	"golang-di-playground/policyholder"
)

type Service interface {
	Notify(policyholder *policyholder.PolicyHolder, message string)
}
