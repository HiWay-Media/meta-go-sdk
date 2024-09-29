package ig

import "fmt"

const (
    BASE_URL = "https://graph.instagram.com"
    VERSION = "20.0"
)

/*
POST /{ig-user-id}/media — upload media and create media containers.

https://developers.facebook.com/docs/instagram-platform/instagram-api-with-instagram-login/content-publishing#reels-posts
*/
func apiUploadMediaUrl(igUserId string) string {
    return fmt.Sprintf("%s/%s/%s/media", BASE_URL, VERSION, igUserId)
}

// POST /{ig-user-id}/media_publish — publish uploaded media using their media containers.
func apiMediaPublishUrl(igUserId string) string {
    return fmt.Sprintf("%s/%s/%s/media_publish", BASE_URL, VERSION, igUserId)
}

// <IG_CONTAINER_ID>?fields=status_code — check media container publishing eligibility and status.
func apiCheckMediaUrl(igUserId string) string {
    return fmt.Sprintf("%s/%s/%s/fields=status_code", BASE_URL, VERSION, igUserId)
}
