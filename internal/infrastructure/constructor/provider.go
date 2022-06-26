package constructor

import (
	"github.com/AkiOuma/biz-groups/internal/domain/service"
	"github.com/AkiOuma/biz-groups/internal/infrastructure/component"
	"github.com/AkiOuma/biz-groups/internal/interface/datasource"
	"github.com/AkiOuma/biz-groups/internal/interface/datasource/sqldb"
	"github.com/AkiOuma/biz-groups/internal/interface/transport"
	"github.com/AkiOuma/biz-groups/internal/usecase"
	"github.com/google/wire"
)

var EntProvider = wire.NewSet(component.NewEntClient)
var GrpcProvider = wire.NewSet(component.NewGrpcServer)
var RepoProvider = wire.NewSet(datasource.NewDatasource, sqldb.NewDB)

var ServiceProvider = wire.NewSet(service.NewService)
var UsecaseProvider = wire.NewSet(usecase.NewUsecase)
var TransportProvider = wire.NewSet(transport.NewTransport)
