package order

type OrderResponsePayload struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    OrderResponseData `json:"data"`
}
type OrderResponseData struct {
	OrderNo string `json:"order_no"`
}
