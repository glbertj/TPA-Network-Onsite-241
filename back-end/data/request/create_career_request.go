package request

type CreateCareerRequest struct {
	Title      string `validate:"required" json:"title"`
	Department string `validate:"required" json:"department"`
	Location   string `validate:"required" json:"location"`
}
