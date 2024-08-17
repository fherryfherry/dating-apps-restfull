package orders

type CreateOrderPayload struct {
	CustomerID    uint
	CustomerName  string
	CustomerEmail string
	PackageID     uint
	PackageTitle  string
	PackageQuota  int64
	PackagePrice  float32
}
