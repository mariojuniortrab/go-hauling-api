package validation

import infra_errors "github.com/mariojuniortrab/hauling-api/internal/infra/errors"

type Validator interface {
	Validate(s interface{}) *infra_errors.CustomError
}
