package services

import "back-end/data/response"

type SearchService interface {
	Search(keyword string) ([]response.SearchResultResponse, error)
}
