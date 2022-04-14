package dependency

import (
	"context"

	"boilerplate.com/v1/cmd/server"
	"boilerplate.com/v1/configs"
	"go.uber.org/fx"
)

/** Module exported for initializing application */
var Module = fx.Options(
	server.Module,
	configs.Module,
	fx.Invoke(injection),
)

func injection(
	lifecycle fx.Lifecycle,
	env configs.Env,
	engine server.Engine,
	logger configs.Logger,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			PORT := ":3002"
			if env.ServerPort != "" {
				PORT = ":" + env.ServerPort
			}

			go func() {
				engine.Gin.Run(PORT)
			}()

			logger.Zap.Info("✅ Server is up and running on port:", env.ServerPort)
			return nil
		},
	})
}
