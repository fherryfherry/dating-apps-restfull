package order

import (
	errCommon "booking-online/commons/error"
	"booking-online/commons/jwt"
	"booking-online/domains/orders"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func (r *Order) CheckoutOrder(c echo.Context) error {
	// Get jwt account
	customer := jwt.GetClaim(c)
	if customer == nil {
		return errCommon.ErrUnauthorized(c)
	}

	reqPayload := OrderRequestPayload{}
	if err := c.Bind(&reqPayload); err != nil {
		return errCommon.ErrorResponseBadRequest(c, err.Error())
	}

	if err := c.Validate(reqPayload); err != nil {
		return errCommon.ErrorResponseBadRequest(c, err.Error())
	}

	// get package detail
	packages := r.packages.FindByCode(reqPayload.PackageCode)
	if packages == nil {
		return errCommon.ErrorResponseBadRequest(c, "Package is not found!")
	}

	payload := orders.CreateOrderPayload{
		CustomerID:    customer.ID,
		CustomerName:  customer.Name,
		CustomerEmail: customer.Email,
		PackageID:     packages.ID,
		PackageTitle:  packages.Title,
		PackagePrice:  packages.Price,
		PackageQuota:  packages.Quota,
	}
	orderNo, err := r.orderSrv.CreateOrder(payload)
	if err != nil {
		log.Printf("Error on create order CustomerID %v | PackageID %v | Error : %v", customer.ID, packages.ID, err.Error())
		return errCommon.ErrorResponseInternalError(c, "Failed on create order!")
	}

	// Update Customer to new level
	err = r.customer.UpdatePackageLevel(customer.ID, packages.Code, packages.Quota)
	if err != nil {
		log.Printf("Error on update package level CustomerID %v | Error = %v", customer.ID, err.Error())
		return errCommon.ErrorResponseInternalError(c, "Failed to update package level")
	}

	return c.JSON(200, OrderResponsePayload{
		Status:  200,
		Message: "SUCCESS",
		Data:    OrderResponseData{orderNo},
	})
}
