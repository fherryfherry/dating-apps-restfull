package datingcandidate

type DatingCandidateBaseResponse struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    *DatingCandidateResponse `json:"data"`
}
type DatingCandidateResponse struct {
	CustomerUUID   string `json:"customer_uuid"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	ProfilePicture string `json:"profile_picture"`
	Bio            string `json:"bio"`
}
