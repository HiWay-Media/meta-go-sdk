package ig

const (
    BASE_URL = "https://graph.instagram.com/"
    VERSION = "20.0"
)


//POST /{ig-user-id}/media — upload media and create media containers.
func UploadMediaUrl(igUserId string) string {
    return fmt.Sprintf("%s/%s/%s/media", BASE_URL, VERSION, igUserId)
}