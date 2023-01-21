package orders

const (
	OrderRequest = `INSERT INTO order (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)" `
	DeliveryRequest = `INSERT INTO delivery (order_uid, name, phone, zip, city, address, region, email)
	VALUES ($1, $2, $3, $4, $5, $6,$7,$8)`
	PaymentRequest = `INSERT INTO payment (order_uid, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
	VALUES ($1, $2, $3, $4, $5, $6,$7,$8, $9, $10, $11)`
	ItemRequest = `INSERT INTO items (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status)
	VALUES ($1, $2, $3, $4, $5, $6,$7,$8, $9, $10, $11, $12)`
	GetOrderByOrderUidRequest      = `SELECT order_uid, track_number, entry, locale, internal_signature, customer_id,delivery_service,shardkey, sm_id, date_created, oof_shard FROM order WHERE order_uid = $1`
	GetItemsByOrderUidQueryRequest = `SELECT chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status FROM items WHERE order_uid = $1`
	GetPaymentByOrderUidRequest    = `SELECT transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee FROM payment WHERE order_uid = $1`
	GetDeliveryByOrderUidRequest   = `SELECT name, phone, zip, city, address, region, email FROM delivery where order_uid = $1`
	OrderAmountRequest             = `SELECT COUNT(*) FROM order`
	OrderUidRequest                = `SELECT order_uid from order`
)
