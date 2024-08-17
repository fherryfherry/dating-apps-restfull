package datingaction

import (
	errCommon "booking-online/commons/error"
	"booking-online/commons/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func (r *DatingAction) SwipeHandler(c echo.Context) error {
	loggedUser := jwt.GetClaim(c)
	if loggedUser == nil {
		return errCommon.ErrUnauthorized(c)
	}
	var request DatingActionRequest

	if err := c.Bind(&request); err != nil {
		return err
	}

	if err := c.Validate(request); err != nil {
		return errCommon.ErrorResponseBadRequest(c, err.Error())
	}

	// Validate swipe type
	if request.SwipeType != "LIKE" && request.SwipeType != "PASS" {
		log.Printf("Swipe: Customer %v invalid swiping type %v", loggedUser.ID, request.SwipeType)
		return errCommon.ErrorResponseBadRequest(c, "Swipe type is invalid!")
	}

	// get logged customer
	customer := r.customer.FindByID(loggedUser.ID)

	// Get customer by uuid to get real id
	swipeCustomer := r.customer.FindByUuid(request.SwipeCustomerUuid)
	if swipeCustomer == nil {
		log.Printf("Swipe: Customer %v invalid swipe customer uuid %v", customer.ID, request.SwipeCustomerUuid)
		return errCommon.ErrorResponseBadRequest(c, "Something went wrong, invalid customer!")
	}

	log.Printf("Swipe: Customer %v | Swipe customer ID %v", loggedUser.ID, swipeCustomer.ID)

	// Validate already swipe
	if check := r.swipe.CheckSwipe(loggedUser.ID, swipeCustomer.ID); check {
		log.Printf("Swipe: Customer %v has already swipe the customer %v for today", customer.ID, swipeCustomer.ID)
		return errCommon.ErrorResponseBadRequest(c, "The candidate has already swipe!")
	}

	// Validate if user already reach the maximum
	swipeTotalToday := r.swipe.GetSwipeToday(customer.ID)
	if customer.SwipeQuota != -1 && swipeTotalToday >= customer.SwipeQuota {
		log.Printf("Swipe: Customer %v has reached maximum quota %v", customer.ID, customer.SwipeQuota)
		return errCommon.ErrorResponseBadRequest(c, "The customer has reached maximum quota!")
	}

	// Save the swipe
	err := r.swipe.AddSwipe(loggedUser.ID, swipeCustomer.ID, request.SwipeType)
	if err != nil {
		log.Printf("Swipe: Error on swipe loggedUser %v | swipeCustomer %v | swipeType %v | Error: %v", loggedUser.ID, customer.ID, request.SwipeType, err.Error())
		return errCommon.ErrorResponseInternalError(c, "Something went wrong while add swipe!")
	}

	return c.JSON(200, DatingActionResponse{
		Status:  200,
		Message: "SUCCESS",
	})
}
