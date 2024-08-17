package customers

type CreateCustomerPayload struct {
	UUID           string
	FirstName      string
	LastName       string
	Bio            string
	Email          string
	Password       string
	Level          string
	SwipeQuota     int64
	ProfilePicture string
}
