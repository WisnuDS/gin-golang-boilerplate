package configs

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(InitEnv),
	fx.Provide(InitDatabase),
	fx.Provide(InitLogger),
)
