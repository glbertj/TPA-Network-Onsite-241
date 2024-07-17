package request

type PlayRequest struct {
	SongId string `json:"songId"`
	UserId string `json:"userId"`
}
