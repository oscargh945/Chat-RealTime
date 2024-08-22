package postgresConfig

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

type Postgres struct {
	Pool *pgxpool.Pool
	ctx  context.Context
}

func NewPostgres(ctx context.Context) Postgres {
	return Postgres{
		Pool: ConnectDB(ctx),
		ctx:  ctx,
	}
}

func (p Postgres) createExtension() error {
	row, _ := os.ReadFile("server/infrastructure/postgresConfig/scripts/create_extension.sql")
	if _, err := p.Pool.Exec(p.ctx, string(row)); err != nil {
		panic(err)
	}
	return nil
}

func (p Postgres) createTables() error {
	tables, _ := os.ReadFile("server/infrastructure/postgresConfig/scripts/create_table.sql")
	if _, err := p.Pool.Exec(p.ctx, string(tables)); err != nil {
		panic(err)
	}
	return nil
}

func (p Postgres) InitPostgresDB() {
	log.Println("Reading script")
	_ = p.createExtension()
	_ = p.createTables()
	log.Println("Scripts successfully")
}
