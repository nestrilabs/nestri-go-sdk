// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nestri

import (
	"github.com/nestrilabs/nestri-go-sdk/option"
)

// SessionActivePublicService contains methods and other services that help with
// interacting with the nestri API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSessionActivePublicService] method instead.
type SessionActivePublicService struct {
	Options []option.RequestOption
}

// NewSessionActivePublicService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSessionActivePublicService(opts ...option.RequestOption) (r *SessionActivePublicService) {
	r = &SessionActivePublicService{}
	r.Options = opts
	return
}
