package fb

const (
    BASE_URL = "https://graph.facebook.com"
    VERSION = "v20.0"
)

func apiUploadMediaUrl(pageId string) string {
    return fmt.Sprintf("%s/%s/%s/videos", BASE_URL, VERSION, pageId)
}
