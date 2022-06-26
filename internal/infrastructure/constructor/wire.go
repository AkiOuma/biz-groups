// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package constructor

import (
	"github.com/AkiOuma/biz-groups/internal/infrastructure/component"
	"github.com/AkiOuma/biz-groups/internal/infrastructure/config"
	"github.com/google/wire"
)

func InitApp(*config.Bootstrap) *component.GrpcServer {
	panic(wire.Build(
		RepoProvider,
		EntProvider,
		GrpcProvider,
		ServiceProvider,
		UsecaseProvider,
		TransportProvider,
	))
}
