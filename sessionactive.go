// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nestri

import (
	"github.com/nestrilabs/nestri-go-sdk/option"
)

// SessionActiveService contains methods and other services that help with
// interacting with the nestri API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSessionActiveService] method instead.
type SessionActiveService struct {
	Options []option.RequestOption
	Public  *SessionActivePublicService
}

// NewSessionActiveService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSessionActiveService(opts ...option.RequestOption) (r *SessionActiveService) {
	r = &SessionActiveService{}
	r.Options = opts
	r.Public = NewSessionActivePublicService(opts...)
	return
}
