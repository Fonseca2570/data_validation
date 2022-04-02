package validator

import (
	"data_validation/services/data_loader"
	"fmt"
)

type validator struct {
	Data data_loader.Loader
}

func NewValidator(loader data_loader.Loader) validator{
	return validator{
		Data: loader,
	}
}

func (v *validator) First() error {
	fmt.Println("first")
	return nil
}
