package semerr

import "errors"

var (
	ErrProviderCategoryNotFound = errors.New("the given provider category wasnt found")
)
