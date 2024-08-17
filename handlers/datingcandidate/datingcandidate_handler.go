package datingcandidate

import (
	errCommon "booking-online/commons/error"
	"booking-online/commons/fileutil"
	"booking-online/commons/jwt"
	"github.com/labstack/echo/v4"
)

func (r *DatingCandidate) GetCandidateHandler(c echo.Context) error {
	loggedUser := jwt.GetClaim(c)
	if loggedUser == nil {
		return errCommon.ErrUnauthorized(c)
	}

	candidate := r.customer.FindNonSwipeToday(loggedUser.ID)
	if candidate == nil || candidate.FirstName == "" {
		return c.JSON(200, DatingCandidateBaseResponse{
			Status:  200,
			Message: "REACH_MAXIMUM_OR_NOT_FOUND_CANDIDATE",
		})
	}

	return c.JSON(200, DatingCandidateBaseResponse{
		Status:  200,
		Message: "SUCCESS",
		Data: &DatingCandidateResponse{
			CustomerUUID:   candidate.CustomerUUID,
			FirstName:      candidate.FirstName,
			LastName:       candidate.LastName,
			ProfilePicture: fileutil.GetFullUrl(candidate.ProfilePicture),
			Bio:            candidate.Bio,
		},
	})
}
