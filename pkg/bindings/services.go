package bindings

import (
	"backend/pkg/bindings/profile"

	"go.uber.org/fx"
)

var ServicesOptions = fx.Options(
	profile.Module,
)
