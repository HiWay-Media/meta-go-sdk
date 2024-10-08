package ig

type MediaRequest struct {
    Caption             string  `json:"caption"`
    CoverUrl            string `json:"cover_url"`
    IsCarouselItem      string `json:"is_carousel_item,omitempty"`
    MediaType           string `json:"media_type,omitempty"`
    ImageUrl            string `json:"image_url,omitempty"`
    VideoUrl            string `json:"video_url,omitempty"`
}

type mediaPublishRequest struct {
    CreationId string `json:"creation_id"`
}