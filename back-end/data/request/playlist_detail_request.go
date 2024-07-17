package request

type PlayListDetailRequest struct {
	UserId     string `json:"userId"`
	PlaylistID string `json:"playlistId"`
	SongID     string `json:"songId"`
}
