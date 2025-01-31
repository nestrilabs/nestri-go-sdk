// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

type UnionString string

func (UnionString) ImplementsSessionGetResponseDataStartedAtUnion()        {}
func (UnionString) ImplementsSessionGetResponseDataEndedAtUnion()          {}
func (UnionString) ImplementsSessionActiveListResponseDataStartedAtUnion() {}
func (UnionString) ImplementsSessionActiveListResponseDataEndedAtUnion()   {}
func (UnionString) ImplementsUserGetResponseDataCreatedAtUnion()           {}
func (UnionString) ImplementsUserGetResponseDataDiscriminatorUnion()       {}
func (UnionString) ImplementsUserGetResponseDataUpdatedAtUnion()           {}
func (UnionString) ImplementsSubscriptionListResponseDataCanceledAtUnion() {}

type UnionFloat float64

func (UnionFloat) ImplementsSessionGetResponseDataStartedAtUnion()        {}
func (UnionFloat) ImplementsSessionGetResponseDataEndedAtUnion()          {}
func (UnionFloat) ImplementsSessionActiveListResponseDataStartedAtUnion() {}
func (UnionFloat) ImplementsSessionActiveListResponseDataEndedAtUnion()   {}
func (UnionFloat) ImplementsUserGetResponseDataCreatedAtUnion()           {}
func (UnionFloat) ImplementsUserGetResponseDataDiscriminatorUnion()       {}
func (UnionFloat) ImplementsUserGetResponseDataUpdatedAtUnion()           {}
func (UnionFloat) ImplementsSubscriptionListResponseDataCanceledAtUnion() {}
