package core_http_request

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func DecodeAndValidate(r *http.Request, req any) error {
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return fmt.Errorf("decode request: %w", err)
	}

	if err := validate.Struct(req); err != nil {
		return fmt.Errorf("validate request: %w", err)
	}

	return nil
}
