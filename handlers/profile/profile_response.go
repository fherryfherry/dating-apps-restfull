package profile

type ProfileResponse struct {
	Status  int                 `json:"status"`
	Message string              `json:"message"`
	Data    ProfileDataResponse `json:"data"`
}
type ProfileDataResponse struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Level          string `json:"level"`
	Email          string `json:"email"`
	Bio            string `json:"json"`
	ProfilePicture string `json:"profile_picture"`
}

type UploadResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
