package ig

/*
{
  "id": "17920238422030506" // Instagram Media ID
}
*/
type MediaPublishResponse struct {
    ID  string `json:"id"`
}


/*
{
  "status_code": "FINISHED",
  "id": "17889615691921648"
}
*/
type CheckMediaStatusResponse struct {
  ID          string `json:"id"`
  StatusCode  string `json:"status_code"`
}