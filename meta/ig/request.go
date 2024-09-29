package ig

type MediaRequest struct {
    Caption     string  `json:"caption"`
    CoverUrl    string `json:"cover_url"`
    MediaType   string `json:"media_type"`
}

type mediaPublishRequest struct {
    CreationId string `json:"creation_id"`
}