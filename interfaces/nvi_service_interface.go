package interfaces

import "nvi-identity-verification/models"

type NviService interface {
	VerifyIdentity(request models.NviRequest) (bool, error)
	VerifyForeignIdentity(request models.ForeignNviRequest) (bool, error)
}
