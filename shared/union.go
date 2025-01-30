// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

type UnionString string

func (UnionString) ImplementsUserGetResponseDataCreatedAtUnion()           {}
func (UnionString) ImplementsUserGetResponseDataDiscriminatorUnion()       {}
func (UnionString) ImplementsUserGetResponseDataUpdatedAtUnion()           {}
func (UnionString) ImplementsSubscriptionListResponseDataCanceledAtUnion() {}

type UnionFloat float64

func (UnionFloat) ImplementsUserGetResponseDataCreatedAtUnion()           {}
func (UnionFloat) ImplementsUserGetResponseDataDiscriminatorUnion()       {}
func (UnionFloat) ImplementsUserGetResponseDataUpdatedAtUnion()           {}
func (UnionFloat) ImplementsSubscriptionListResponseDataCanceledAtUnion() {}
