package sqldb

import "github.com/AkiOuma/biz-groups/internal/interface/datasource/sqldb/ent"

type DB struct {
	client *ent.Client
}

func NewDB(client *ent.Client) *DB {
	return &DB{
		client: client,
	}
}
