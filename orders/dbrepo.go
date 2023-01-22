package orders

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"wb/config"
)

type OrderDb struct {
	database *pgxpool.Pool
}

func NewOrderDb(db *pgxpool.Pool) *OrderDb {
	return &OrderDb{database: db}
}

func (OrderDB *OrderDb) BatchCreate(ctx context.Context, Order config.Order) error {
	Batch := &pgx.Batch{}
	Batch.Queue(
		OrderRequest,
		&Order.OrderUid,
		&Order.TrackNumber,
		&Order.Entry,
		&Order.Locale,
		&Order.InternalSignature,
		&Order.CustomerId,
		&Order.DeliveryService,
		&Order.Shardkey,
		&Order.SmId,
		&Order.DateCreated,
		&Order.OofShard,
	)

	Batch.Queue(
		DeliveryRequest,
		&Order.OrderUid,
		&Order.Name,
		&Order.Phone,
		&Order.Zip,
		&Order.City,
		&Order.Address,
		&Order.Region,
		&Order.Email,
	)

	Batch.Queue(
		PaymentRequest,
		&Order.OrderUid,
		&Order.Payment,
		&Order.RequestId,
		&Order.Currency,
		&Order.Provider,
		&Order.Amount,
		&Order.PaymentDt,
		&Order.Bank,
		&Order.DeliveryCost,
		&Order.GoodTotal,
		&Order.CustomFee,
	)

	for _, o := range Order.Items {
		Batch.Queue(
			ItemRequest,
			&Order.OrderUid,
			&o.ChrtId,
			&o.TrackNumber,
			&o.Price,
			&o.Rid,
			&o.Name,
			&o.Sale,
			&o.Size,
			&o.TotalPrice,
			&o.NmId,
			&o.Brand,
			&o.Status,
		)

	}
	send := OrderDB.database.SendBatch(ctx, Batch)

	for i := 0; i < Batch.Len(); i++ {
		_, err := send.Exec()
		if err != nil {
			return err
		}
	}

	err := send.Close()
	if err != nil {
		return err
	}
	return nil
}

func (OrderDB *OrderDb) CreateOrder(ctx context.Context, Order config.Order) (*config.Order, error) {
	var (
		orderConfig config.Order
	)
	if err := OrderDB.database.QueryRow(
		ctx,
		OrderRequest,
		&Order.OrderUid,
		&Order.TrackNumber,
		&Order.Entry,
		&Order.Locale,
		&Order.InternalSignature,
		&Order.CustomerId,
		&Order.DeliveryService,
		&Order.Shardkey,
		&Order.SmId,
		&Order.DateCreated,
		&Order.OofShard,
	).Scan(
		&orderConfig.OrderUid,
		&orderConfig.TrackNumber,
		&orderConfig.Entry,
		&orderConfig.Locale,
		&orderConfig.InternalSignature,
		&orderConfig.CustomerId,
		&orderConfig.DeliveryService,
		&orderConfig.Shardkey,
		&orderConfig.SmId,
		&orderConfig.DateCreated,
		&orderConfig.OofShard,
	); err != nil {
		return nil, errors.WithMessage(err, "Error Scan Order Query")
	}
	return &orderConfig, nil
}

func (OrderDB *OrderDb) GetByUid(ctx context.Context, OrderUid string) (*config.Order, error) {
	var orderConfig config.Order

	if err := OrderDB.database.QueryRow(ctx, GetOrderByOrderUidRequest, OrderUid).Scan(
		&orderConfig.OrderUid,
		&orderConfig.TrackNumber,
		&orderConfig.Entry,
		&orderConfig.Locale,
		&orderConfig.InternalSignature,
		&orderConfig.CustomerId,
		&orderConfig.DeliveryService,
		&orderConfig.Shardkey,
		&orderConfig.SmId,
		&orderConfig.DateCreated,
		&orderConfig.OofShard); err != nil {
		return nil, errors.WithMessage(err, "Error get by uid order")
	}
	return &orderConfig, nil
}

func (OrderDB *OrderDb) GetAllByUid(ctx context.Context, OrderUid string) (*config.Order, error) {
	var orderConfig config.Order
	if err := OrderDB.database.QueryRow(ctx, GetOrderByOrderUidRequest, OrderUid).Scan(
		&orderConfig.OrderUid,
		&orderConfig.TrackNumber,
		&orderConfig.Entry,
		&orderConfig.Locale,
		&orderConfig.InternalSignature,
		&orderConfig.CustomerId,
		&orderConfig.DeliveryService,
		&orderConfig.Shardkey,
		&orderConfig.SmId,
		&orderConfig.DateCreated,
		&orderConfig.OofShard,
	); err != nil {
		errors.WithMessage(err, "Error get all by uid order")
	}
	if err := OrderDB.database.QueryRow(ctx, GetDeliveryByOrderUidRequest, OrderUid).Scan(
		&orderConfig.Name,
		&orderConfig.Phone,
		&orderConfig.Zip,
		&orderConfig.City,
		&orderConfig.Address,
		&orderConfig.Region,
		&orderConfig.Email,
	); err != nil {
		errors.WithMessage(err, "Error get all by uid delivery")
	}
	if err := OrderDB.database.QueryRow(ctx, GetPaymentByOrderUidRequest, OrderUid).Scan(
		&orderConfig.Transaction,
		&orderConfig.RequestId,
		&orderConfig.Currency,
		&orderConfig.Provider,
		&orderConfig.Amount,
		&orderConfig.PaymentDt,
		&orderConfig.Bank,
		&orderConfig.DeliveryCost,
		&orderConfig.GoodTotal,
		&orderConfig.CustomFee,
	); err != nil {
		errors.WithMessage(err, "Error get all by uid payment")
	}
	req, err := OrderDB.database.Query(ctx, GetItemsByOrderUidQueryRequest, OrderUid)
	if err != nil {
		return nil, errors.WithMessage(err, "Error request get all by uid items")
	}
	Item := config.Items{}

	for req.Next() {
		err := req.Scan(
			&Item.ChrtId,
			&Item.TrackNumber,
			&Item.Price,
			&Item.Rid,
			&Item.Name,
			&Item.Sale,
			&Item.Size,
			&Item.TotalPrice,
			&Item.NmId,
			&Item.Brand,
			&Item.Status,
		)
		if err != nil {
			errors.WithMessage(err, "Error get all by uid items")
		}
	}
	return &orderConfig, nil
}

func (OrderDB *OrderDb) GetOrderAmount(ctx context.Context) (int, error) {
	var Amount int
	if err := OrderDB.database.QueryRow(ctx, OrderAmountRequest).Scan(&Amount); err != nil {
		return 0, errors.WithMessage(err, "Error get order amount")
	}
	return Amount, nil
}

func (OrderDB *OrderDb) GetOrderUid(ctx context.Context) ([]string, error) {
	var OrderUid []string
	req, err := OrderDB.database.Query(ctx, OrderUidRequest)
	if err != nil {
		return []string{}, errors.WithMessage(err, "Error get Order uid ")
	}
	var OrderUID string
	for req.Next() {
		err := req.Scan(
			&OrderUID)
		if err != nil {
			return []string{}, errors.WithMessage(err, "Error get Order uid slice")
		}
		OrderUid = append(OrderUid, OrderUID)
	}
	return OrderUid, nil
}
