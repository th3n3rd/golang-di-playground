package company

import "go.uber.org/dig"

func Provide(container *dig.Container) {
	_ = container.Provide(NewAdjuster)
	_ = container.Provide(NewInsuranceCompany)
}
