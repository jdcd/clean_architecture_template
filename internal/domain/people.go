package domain

import (
	"errors"
	"fmt"

	"github.com/jdcd9001/clean-architecture-template/pkg"
)

const (
	minAge        = 1
	maxAge        = 130
	invalidName   = "the field name is invalid, please check it"
	invalidAge    = "the field age is invalid, please check it"
	ageOutOfRange = "the age %d is out of the valid range [%d - %d]"
	emptyName     = "the field name cannot be empty"
)

type People struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func (r *People) Validate() error {
	if err := r.checkName(); err != nil {
		return err
	}

	if err := r.checkAge(); err != nil {
		return err
	}

	return nil
}

func (r *People) checkName() error {
	if r.Name == "" {
		formattedError := pkg.CreateFormatError(pkg.DataValidation, invalidName, emptyName)
		pkg.ErrorLogger().Println(emptyName)
		return errors.New(formattedError)
	}
	return nil
}

func (r *People) checkAge() error {
	if r.Age < minAge || r.Age > maxAge {
		detailErr := fmt.Sprintf(ageOutOfRange, r.Age, minAge, maxAge)
		formattedError := pkg.CreateFormatError(pkg.DataValidation, invalidAge, detailErr)
		pkg.ErrorLogger().Println(detailErr)
		return errors.New(formattedError)
	}

	return nil
}
