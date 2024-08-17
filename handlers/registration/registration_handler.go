package registration

import (
	errCommon "booking-online/commons/error"
	"booking-online/domains/customers"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

func (r *Registration) RegisterHandler(c echo.Context) error {

	reqPayload := RegistrationRequestPayload{}
	if err := c.Bind(&reqPayload); err != nil {
		return errCommon.ErrorResponseBadRequest(c, err.Error())
	}

	if err := c.Validate(reqPayload); err != nil {
		return errCommon.ErrorResponseBadRequest(c, err.Error())
	}

	// Validate exists of account by email
	if r.customerSrv.CheckExistByEmail(reqPayload.Email) {
		return errCommon.ErrorResponseBadRequest(c, "Account already exists!")
	}

	// Get default package for free
	defaultLevel := viper.GetString("default_level")
	packages := r.packages.FindByCode(defaultLevel)
	if packages == nil {
		return errCommon.ErrorResponseBadRequest(c, "Package is not found!")
	}

	// Generate new uuid for customer
	newUuid, err := uuid.NewRandom()
	if err != nil {
		log.Printf("Generate uuid failed: %v", err.Error())
		return errCommon.ErrorResponseInternalError(c, "Something went wrong generate uuid")
	}

	// Mapping to create customer payload
	payload := customers.CreateCustomerPayload{
		UUID:       newUuid.String(),
		FirstName:  reqPayload.FirstName,
		LastName:   reqPayload.LastName,
		Bio:        reqPayload.Bio,
		Email:      reqPayload.Email,
		Password:   reqPayload.Password,
		Level:      packages.Code,
		SwipeQuota: packages.Quota,
	}

	// Insert process
	customerModel, err := r.customerSrv.CreateCustomer(payload)
	if err != nil {
		log.Printf("Registration Payload %v | failed error = %v", payload, err.Error())
		return errCommon.ErrorResponseInternalError(c, "Registration failed!")
	}

	return c.JSON(200, RegistrationResponsePayload{
		Status:  200,
		Message: "SUCCESS",
		Data:    RegistrationResponseData{customerModel.CustomerUUID},
	})
}
