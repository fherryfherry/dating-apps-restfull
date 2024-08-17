package quotapackage

type QuotaPackageBaseResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    []QuotaPackageResponse `json:"data"`
}
type QuotaPackageResponse struct {
	Code  string `json:"code"`
	Title string `json:"title"`
	Quota int64  `json:"quota"`
}
