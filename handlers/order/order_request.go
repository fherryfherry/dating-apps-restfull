package order

type OrderRequestPayload struct {
	PackageCode string `json:"package_code" validate:"required"`
}
