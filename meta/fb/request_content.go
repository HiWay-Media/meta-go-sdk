package fb

// https://developers.facebook.com/docs/graph-api/guides/upload
type mediaBodyRequest struct {
    AccessToken     string  `json:"access_token"`
    Title           string  `json:"title"` 
    Description     string  `json:"description"` 
    FileUrl         string  `json:"file_url"` 
}
