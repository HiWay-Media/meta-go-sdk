package ig

import (
	"encoding/json"
	"fmt"
)

/*
curl -X POST "https://graph.instagram.com/v20.0/<IG_ID>/media"
     -H "Content-Type: application/json" 
     -d '{
           "image_url":"https://www.example.com/images/bronz-fonz.jpg"
         }'
*/
func (s *igService) UploadMedia(igUserId string, videoUrl, imageUrl *string, caption string ) (*MediaPublishResponse, error) {
    body := MediaRequest{
        Caption: caption,
    }
    if videoUrl != nil {
        body.VideoUrl = videoUrl
        body.MediaType = "VIDEO"
    }
    if imageUrl != nil {
        body.ImageUrl = imageUrl
    }
    resp, err := s.restyPost(apiUploadMediaUrl(igUserId), body)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("creator info error %s", resp.String())
	}
    var obj MediaPublishResponse
	if err := json.Unmarshal(resp.Body(), &obj); err != nil {
		return nil, err
	}
	s.debugPrint(obj)
	return &obj, nil
}

/*
curl -X POST "https://graph.instagram.com/v20.0/17841400008460056/media_publish"
    -H "Content-Type: application/json" 
    -d '{
           "creation_id=17889455560051444"
         }'
*/
func (s *igService) MediaPublish(igUserId, creationId string) (*MediaPublishResponse, error) {
    body := mediaPublishRequest{CreationId: creationId}

	resp, err := s.restyPost(apiMediaPublishUrl(igUserId), body)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("creator info error %s", resp.String())
	}
    var obj MediaPublishResponse
	if err := json.Unmarshal(resp.Body(), &obj); err != nil {
		return nil, err
	}
	s.debugPrint(obj)
	return &obj, nil
}