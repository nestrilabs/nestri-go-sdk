// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

type UnionString string

func (UnionString) ImplementsMachineGetResponseDataCreatedAtUnion()  {}
func (UnionString) ImplementsMachineListResponseDataCreatedAtUnion() {}
func (UnionString) ImplementsSessionGetResponseDataStartedAtUnion()  {}
func (UnionString) ImplementsSessionGetResponseDataEndedAtUnion()    {}
func (UnionString) ImplementsSessionListResponseDataStartedAtUnion() {}
func (UnionString) ImplementsSessionListResponseDataEndedAtUnion()   {}

type UnionFloat float64

func (UnionFloat) ImplementsMachineGetResponseDataCreatedAtUnion()  {}
func (UnionFloat) ImplementsMachineListResponseDataCreatedAtUnion() {}
func (UnionFloat) ImplementsSessionGetResponseDataStartedAtUnion()  {}
func (UnionFloat) ImplementsSessionGetResponseDataEndedAtUnion()    {}
func (UnionFloat) ImplementsSessionListResponseDataStartedAtUnion() {}
func (UnionFloat) ImplementsSessionListResponseDataEndedAtUnion()   {}
