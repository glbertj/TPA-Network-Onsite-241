package request

type ArtistRequest struct {
	UserId      string  `json:"userId"`
	Description *string `json:"description"`
	Banner      *string `json:"banner"`
}
