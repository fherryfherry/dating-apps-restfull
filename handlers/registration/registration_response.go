package registration

type RegistrationResponsePayload struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    RegistrationResponseData `json:"data"`
}
type RegistrationResponseData struct {
	ID string `json:"id"`
}

type UploadResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    UploadData `json:"data"`
}
type UploadData struct {
	ImagePath string `json:"image_path"`
}
