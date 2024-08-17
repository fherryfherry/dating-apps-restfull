package datingaction

type DatingActionRequest struct {
	SwipeType         string `json:"swipe_type" validate:"required"`
	SwipeCustomerUuid string `json:"swipe_customer_uuid" validate:"required"`
}
