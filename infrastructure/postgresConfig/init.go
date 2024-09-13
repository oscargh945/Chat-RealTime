package postgresConfig

import (
	"context"
	"fmt"
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
	row, err := os.ReadFile("server/infrastructure/postgresConfig/scripts/create_extension.sql")
	if err != nil {
		return fmt.Errorf("error reading create_extension.sql: %w", err)
	}
	if _, err := p.Pool.Exec(p.ctx, string(row)); err != nil {
		return fmt.Errorf("error executing create_extension.sql: %w", err)
	}
	return nil
}

func (p Postgres) createTables() error {
	tables, err := os.ReadFile("server/infrastructure/postgresConfig/scripts/create_table.sql")
	if err != nil {
		return fmt.Errorf("error reading create_table.sql: %w", err)
	}
	if _, err := p.Pool.Exec(p.ctx, string(tables)); err != nil {
		return fmt.Errorf("error executing create_table.sql: %w", err)
	}
	return nil
}

func (p Postgres) InitPostgresDB() {
	log.Println("Reading script")
	_ = p.createExtension()
	_ = p.createTables()
	log.Println("Scripts successfully")
}
