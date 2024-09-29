package ig

import (
	"encoding/json"
	"fmt"
)

/*
curl -X POST "https://graph.instagram.com/v20.0/17841400008460056/media_publish"
    -H "Content-Type: application/json" 
    -d '{
           "creation_id=17889455560051444"
         }'
*/
func (s *igService) MediaPublish(igUserId, creationId string) (*MediaPublishResponse, error) {
    body := mediaPublishRequest{CreationId: creationId}

	resp, err := o.restyPost(apiMediaPublishUrl(igUserId), body)
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
	o.debugPrint(obj)
	return &obj, nil
}