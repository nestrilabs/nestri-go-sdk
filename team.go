// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nestri

import (
	"github.com/nestrilabs/nestri-go-sdk/option"
)

// TeamService contains methods and other services that help with interacting with
// the nestri API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTeamService] method instead.
type TeamService struct {
	Options []option.RequestOption
}

// NewTeamService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTeamService(opts ...option.RequestOption) (r *TeamService) {
	r = &TeamService{}
	r.Options = opts
	return
}
