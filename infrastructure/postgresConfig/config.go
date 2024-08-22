package postgresConfig

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/oscargh945/go-Chat/utils"
	"log"
	"os"
	"time"
)

func NewConnectionConfig() string {
	// Get database url
	if url := os.Getenv("DATABASE_URL"); url != "" {
		return url
	}

	// Get data from env vars if not exist database url
	username := utils.GetEnv("POSTGRES_USER")
	password := utils.GetEnv("POSTGRES_PASSWORD")
	localHost := utils.GetEnv("POSTGRES_HOST")
	port := utils.GetEnv("POSTGRES_PORT")
	dbName := utils.GetEnv("POSTGRES_DB")

	// Create url with data
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, localHost, port, dbName)
}

func ConnectDB(ctx context.Context) *pgxpool.Pool {
	// Create config
	config, err := pgxpool.ParseConfig(NewConnectionConfig())
	if err != nil {
		panic(fmt.Sprintf("error configuring the database: %s", err))
	}

	// Create pool with config
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		panic(fmt.Sprint("Unable to connect to database:", err))
	}

	// test connection
	var now time.Time
	err = pool.QueryRow(context.Background(), "SELECT NOW()").Scan(&now)
	if err != nil {
		panic(fmt.Sprint("failed to execute query", err))
	}
	log.Println("CONNECTION WITH DATABASE ESTABLISHED!")

	return pool
}
