package profile

import (
	errCommon "booking-online/commons/error"
	"booking-online/commons/fileutil"
	"booking-online/commons/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

func (r *Profile) GetMyProfileHandler(c echo.Context) error {
	loggedUser := jwt.GetClaim(c)
	if loggedUser == nil {
		return errCommon.ErrUnauthorized(c)
	}

	customer := r.customer.FindByID(loggedUser.ID)

	return c.JSON(200, ProfileResponse{
		Status:  200,
		Message: "SUCCESS",
		Data: ProfileDataResponse{
			FirstName:      customer.FirstName,
			LastName:       customer.LastName,
			Level:          customer.Level,
			Email:          customer.Email,
			Bio:            customer.Bio,
			ProfilePicture: fileutil.GetFullUrl(customer.ProfilePicture),
		},
	})
}

func (r *Profile) UpdateProfilePictureHandler(c echo.Context) error {
	customer := jwt.GetClaim(c)
	if customer == nil {
		return errCommon.ErrUnauthorized(c)
	}

	file, err := c.FormFile("image")
	if err != nil {
		return errCommon.ErrorResponseBadRequest(c, "File is broken!")
	}

	// Validate file extension
	fileExt := filepath.Ext(file.Filename)
	allowedExtensions := []string{".jpg", ".jpeg", ".png"}
	isValidExtension := false
	for _, ext := range allowedExtensions {
		if strings.EqualFold(fileExt, ext) {
			isValidExtension = true
			break
		}
	}
	if !isValidExtension {
		return errCommon.ErrorResponseBadRequest(c, "Invalid file format. Only JPG, JPEG, and PNG are allowed")
	}

	// Upload file
	imageDestinationPath := "assets/profile-picture"
	newUuid, _ := uuid.NewRandom()
	newFileName := newUuid.String() + fileExt
	if err := uploadFile(file, imageDestinationPath, newFileName); err != nil {
		return errCommon.ErrorResponseInternalError(c, err.Error())
	}
	log.Printf("File is uploaded: %s", imageDestinationPath)

	// Update to Customer data
	err = r.customer.UpdateProfilePicture(customer.ID, filepath.Join(imageDestinationPath, newFileName))
	if err != nil {
		log.Printf("Failed to update profile: %v", err.Error())
		return errCommon.ErrorResponseInternalError(c, "Something went wrong on update profile!")
	}

	return c.JSON(200, UploadResponse{
		Status:  200,
		Message: "SUCCESS",
	})
}

func uploadFile(file *multipart.FileHeader, path string, newFileName string) error {
	// Open the uploaded file.
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Create directory
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	// Create a destination file for the uploaded content.
	dst, err := os.Create(filepath.Join(path, newFileName))
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy the uploaded content to the destination file.
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}
