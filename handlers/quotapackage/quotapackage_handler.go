package quotapackage

import (
	errCommon "booking-online/commons/error"
	"booking-online/commons/jwt"
	"github.com/labstack/echo/v4"
)

func (r *QuotaPackage) GetPackagesHandler(c echo.Context) error {
	loggedUser := jwt.GetClaim(c)
	if loggedUser == nil {
		return errCommon.ErrUnauthorized(c)
	}

	// Get package list
	queryData := r.packages.GetList()

	// Mapping fields
	result := []QuotaPackageResponse{}
	for _, item := range queryData {
		result = append(result, QuotaPackageResponse{
			Code:  item.Code,
			Title: item.Title,
			Quota: item.Quota,
		})
	}

	return c.JSON(200, QuotaPackageBaseResponse{
		Status:  200,
		Message: "SUCCESS",
		Data:    result,
	})
}
