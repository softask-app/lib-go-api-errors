package apierr

import (
	"errors"
	"github.com/francoispqt/gojay"
)

type ApiError struct {
	Status    string
	Message   string
	RequestId string
}

func (a *ApiError) UnmarshalJSONObject(d *gojay.Decoder, k string) error {
	switch k {
	case JsKeyStatus:
		return d.String(&a.Status)
	case JsKeyMessage:
		return d.String(&a.Message)
	case JsKeyRequestId:
		return d.String(&a.RequestId)
	}

	return errors.New("unrecognized JSON key " + k)
}

func (a *ApiError) NKeys() int {
	return 3
}
