package orders

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"time"
	"wb/config"
)

/*func ConnectionPool(maxAttempts int, ctx context.Context) (*pgxpool.Pool, error) {
	var pool *pgxpool.Pool
	var _ error
	Connection := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", config.UserName, config.Password, config.Host, config.Port, config.Database)
	_ = Try(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		_, err := pgx.Connect(ctx, Connection)
		if err != nil {
			return err
		}
		return nil

	}, maxAttempts, 5*time.Second)
	return pool, nil
}

func Try(fn func() error, Attemtps int, delay time.Duration) (err error) {
	for Attemtps < 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			Attemtps--

			continue
		}
		return nil
		if err != nil {
			log.Fatal("Error try to connect")
		}
	}
	return err
}
*/
const (
	maxConn           = 50
	healthCheckPeriod = 3 * time.Minute
	maxConnIdleTime   = 1 * time.Minute
	maxConnLifetime   = 3 * time.Minute
	minConns          = 10
	lazyConnect       = false
)

func NewPgxConn(cfg *config.Config) (*pgx.Conn, error) {
	ctx := context.Background()
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		cfg.Host,
		cfg.Port,
		cfg.UserName,
		cfg.Database,
		cfg.Password,
	)
	poolCfg, err := pgx.ParseConfig(dataSourceName)
	//poolCfg, err := pgxpool.ParseConfig(dataSourceName)
	if err != nil {
		return nil, err
	}

	//poolCfg.MaxConns = maxConn
	//poolCfg.HealthCheckPeriod = healthCheckPeriod
	//poolCfg.MaxConnIdleTime = maxConnIdleTime
	//poolCfg.MaxConnLifetime = maxConnLifetime
	//poolCfg.MinConns = minConns
	//poolCfg.LazyConnect = lazyConnect

	//connPool, err := pgxpool.ConnConfig(ctx, poolCfg)
	connPool, err := pgx.ConnectConfig(ctx, poolCfg)
	//if err != nil {
	//return nil, errors.Wrap(err, "pgx.ConnectConfig")
	//}

	return connPool, nil
}
