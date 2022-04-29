package biz

import "github.com/google/wire"

type Placeholder struct {
}

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewBizPlaceholder)

func NewBizPlaceholder() *Placeholder {
	return &Placeholder{}
}
