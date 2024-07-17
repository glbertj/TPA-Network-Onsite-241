package request

type AlbumRequest struct {
	Title    string `json:"title"`
	Type     string `json:"type"`
	Banner   string `json:"banner"`
	ArtistId string `json:"artistId"`
}
