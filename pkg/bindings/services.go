package bindings

import (
	"backend/pkg/bindings/profile"
	"backend/pkg/bindings/tasks"

	"go.uber.org/fx"
)

var ServicesOptions = fx.Options(
	tasks.Module,
	profile.Module,
)
