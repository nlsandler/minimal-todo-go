// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package minimaltodo

import (
	"github.com/nlsandler/minimal-todo-go/internal/apierror"
	"github.com/nlsandler/minimal-todo-go/packages/param"
	"github.com/nlsandler/minimal-todo-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// This is an alias to an internal type.
type Tag = shared.Tag

// This is an alias to an internal type.
type TagParam = shared.TagParam
