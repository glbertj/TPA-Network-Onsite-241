package response

type CareerResponse struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Department string `json:"department"`
	Location   string `json:"location"`
}
