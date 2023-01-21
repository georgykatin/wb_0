package httpserver

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"os/signal"
	"syscall"
	"wb/Using"
	"wb/cache"
	"wb/config"
	"wb/httpmeth"
	"wb/orders"
	"wb/orders/dilivery"
	"wb/orders/interfaces"
)

type Server struct {
	cfg            config.Config
	pgxpool        *pgxpool.Pool
	cache          *cache.Cache
	gin            *gin.Engine
	natsConnection stan.Conn
	Cacheinterf    interfaces.CacheRepo
	Dbinterf       interfaces.DbRepo
	Sub            dilivery.OrderSub
}

type PublishServer struct {
	cfg      *config.Config
	natsConn stan.Conn
	pool     *pgxpool.Pool
	gin      *gin.Engine
	cache    *cache.Cache
}

func (s PublishServer) Run() any {
	return nil
}

func NewServer(
	cfg *config.Config,
	natsConnection stan.Conn,
	pool *pgxpool.Pool,
	cache *cache.Cache,
) *PublishServer {
	gin.SetMode(gin.ReleaseMode)
	return &PublishServer{cfg: cfg, natsConn: natsConnection, pool: pool, gin: gin.New(), cache: cache}
}

func (serv *Server) InitServer() error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	orderDbRepository := orders.NewOrderDb(serv.pgxpool)
	orderCacheRepository := orders.OrderCacheRepo(serv.cache)
	orderUseCase := Using.NewOrderUseCase(orderDbRepository, orderCacheRepository)

	var validate *validator.Validate
	validate = validator.New()

	go func() {
		OrderSubscriber := serv.Sub.NewOrderSub(serv.natsConnection, orderUseCase, validate)
		OrderSubscriber.Run(ctx)
	}()
	go func() {
		log.Printf("Server listen and serve on port:8080")
		serv.RunServer()
	}()
	serv.gin.Use(cors.Default())
	httpv1 := serv.gin.Group("/api/v1", cors.Default())
	orderHandlers := httpmeth.OrderHandlers(httpv1.Group("/order"), orderUseCase)
	orderHandlers.MapRoutes()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		log.Fatalf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		log.Fatalf("ctx.Done: %v", done)
	}

	//log.Println("Server Exited Property")

	return nil
}
