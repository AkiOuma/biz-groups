package datasource

import (
	"github.com/AkiOuma/biz-groups/internal/domain/repository"
	"github.com/AkiOuma/biz-groups/internal/interface/datasource/sqldb"
)

type ds struct {
	sqldb *sqldb.DB
}

var _ repository.Repository = (*ds)(nil)

func NewDatasource(sqldb *sqldb.DB) repository.Repository {
	return &ds{sqldb: sqldb}
}
