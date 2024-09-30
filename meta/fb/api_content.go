package fb

import (
	"encoding/json"
	"fmt"
)

/*
Sintassi della richiesta
Formattato per una maggiore leggibilità.

curl -i -X POST "https://graph.facebook.com/v20.0/<APP_ID>/uploads

	?file_name=<FILE_NAME>
	&file_length=<FILE_LENGTH>
	&file_type=<FILE_TYPE>
	&access_token=<USER_ACCESS_TOKEN>"

In caso di azione eseguita correttamente, l'app riceve una risposta JSON con l'ID della sessione di caricamento.

	{
	  "id": "upload:<UPLOAD_SESSION_ID>"
	}
*/
func (s *fbService) UploadMedia(fileUrl, pageId, title, description string) (*UploadMediaResponse, error) {
	params := &mediaBodyRequest{
        AccessToken:        s.accessToken,
        Title:              title,
        Description:        description,
        FileUrl:            fileUrl,
    }
	//
	resp, err := s.restyPostWithQueryParams(apiUploadMediaUrl(pageId), nil, params)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("CheckStatus error %s", resp.String())
	}
	var obj UploadMediaResponse
	if err := json.Unmarshal(resp.Body(), &obj); err != nil {
		return nil, err
	}
	s.debugPrint(obj)
	return &obj, nil
}
