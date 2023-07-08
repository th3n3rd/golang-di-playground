package policyholder

import "go.uber.org/dig"

func Provide(container *dig.Container) {
	_ = container.Provide(NewInMemoryRepository)
}
