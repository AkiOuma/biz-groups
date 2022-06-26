package component

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/AkiOuma/biz-groups/internal/infrastructure/config"
	"github.com/AkiOuma/biz-groups/internal/interface/datasource/sqldb/ent"
	"github.com/AkiOuma/biz-groups/internal/interface/datasource/sqldb/ent/migrate"
	_ "github.com/go-sql-driver/mysql"
)

func NewEntClient(c *config.Bootstrap) *ent.Client {
	client, err := ent.Open("mysql", buildDSN(c))
	if err != nil {
		log.Fatal("ent client init failed, reason: " + err.Error())
	}
	err = client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatal("ent schema init failed, reason: " + err.Error())
	}
	if os.Getenv("ENV") == "DEV" {
		client = client.Debug()
	}
	return client
}

func buildDSN(c *config.Bootstrap) string {
	var sb strings.Builder
	sb.WriteString(c.GetData().GetDb().GetUser())
	sb.WriteRune(':')
	sb.WriteString(c.GetData().GetDb().GetPassword())
	sb.WriteString("@tcp(")
	sb.WriteString(c.GetData().GetDb().GetHost())
	sb.WriteRune(':')
	sb.WriteString(c.GetData().GetDb().GetPort())
	sb.WriteString(")/")
	sb.WriteString(c.GetData().GetDb().GetDatabase())
	sb.WriteString("?parseTime=True")
	return sb.String()
}
