package business

import (
	"errors"
	"nvi-identity-verification/interfaces"
	"nvi-identity-verification/models"
	"nvi-identity-verification/utils"
)

type NviBusiness struct {
	Service interfaces.NviService
}

func NewNviBusiness(service interfaces.NviService) *NviBusiness {
	return &NviBusiness{Service: service}
}

// CheckTurkishIdentity validates and processes Turkish identity verification
func (b *NviBusiness) CheckTurkishIdentity(request models.NviRequest) (bool, error) {

	if len(request.ID) != 11 || !utils.IsNumeric(request.ID) {
		return false, errors.New("invalid ID format: must be 11 numeric characters")
	}
	if len(request.Name) == 0 || len(request.Surname) == 0 {
		return false, errors.New("name and surname cannot be empty")
	}
	if !utils.IsYearValid(request.BirthYear) {
		return false, errors.New("invalid birth year format: must be a 4-digit year")
	}

	return b.Service.VerifyIdentity(request)
}

// CheckForeignIdentity validates and processes foreign identity verification
func (b *NviBusiness) CheckForeignIdentity(request models.ForeignNviRequest) (bool, error) {

	if len(request.ID) == 0 || !utils.IsNumeric(request.ID) {
		return false, errors.New("invalid ID format: must be numeric")
	}
	if len(request.Name) == 0 || len(request.Surname) == 0 {
		return false, errors.New("name and surname cannot be empty")
	}
	if request.BirthDay < 1 || request.BirthDay > 31 {
		return false, errors.New("invalid birth day: must be between 1 and 31")
	}
	if request.BirthMonth < 1 || request.BirthMonth > 12 {
		return false, errors.New("invalid birth month: must be between 1 and 12")
	}
	if request.BirthYear < 1900 || request.BirthYear > 2100 {
		return false, errors.New("invalid birth year: must be a realistic year")
	}

	return b.Service.VerifyForeignIdentity(request)
}
