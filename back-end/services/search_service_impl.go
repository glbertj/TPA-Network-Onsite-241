package services

import (
	"back-end/data/response"
	"back-end/model"
	"back-end/repository"
	"back-end/utils"
	"fmt"
	"sort"
)

type SearchServiceImpl struct {
	SongRepository   repository.SongRepository
	ArtistRepository repository.ArtistRepository
	AlbumRepository  repository.AlbumRepository
	FollowRepository repository.FollowRepository
}

func NewSearchService(songRepository repository.SongRepository, artistRepository repository.ArtistRepository, albumRepository repository.AlbumRepository, followRepository repository.FollowRepository) SearchService {
	return &SearchServiceImpl{
		SongRepository:   songRepository,
		ArtistRepository: artistRepository,
		AlbumRepository:  albumRepository,
		FollowRepository: followRepository,
	}
}

func (s SearchServiceImpl) Search(keyword string) ([]response.SearchResultResponse, error) {
	var res []response.SearchResponse

	fmt.Println("keyword: ", keyword)

	songs, err := s.SongRepository.FindSongByTitle(keyword)
	if err != nil {
		return []response.SearchResultResponse{}, err
	}

	albums, err := s.AlbumRepository.GetAlbumsByTitle(keyword)
	if err != nil {
		return []response.SearchResultResponse{}, err
	}
	//
	users, err := s.ArtistRepository.GetArtistByName(keyword)
	if err != nil {
		return []response.SearchResultResponse{}, err
	}

	for _, user := range users {
		res = append(res, response.SearchResponse{
			Payload: user,
			Type:    "artist",
			Count:   user.FollowCount,
			Title:   user.Username,
		})
	}

	for _, song := range songs {
		res = append(res, response.SearchResponse{
			Payload: song,
			Type:    "song",
			Count:   song.PlayCount,
			Title:   song.Title,
		})
	}

	for _, album := range albums {
		res = append(res, response.SearchResponse{
			Payload: album,
			Type:    "album",
			Count:   album.PlayCount,
			Title:   album.Title,
		})
	}

	sort.Slice(res, func(i, j int) bool {
		if res[i].Count == res[j].Count {
			return utils.GetDistance(keyword, res[i].Title) > utils.GetDistance(keyword, res[j].Title)
		}
		return res[i].Count > res[j].Count
	})

	var result []response.SearchResultResponse

	if len(res) > 0 {
		if res[0].Type == "song" {
			count := 0

			for _, r := range res {
				if r.Type != "song" {
					continue
				}
				if count == 5 {
					break
				}
				song, err := s.SongRepository.GetSongById(r.Payload.(response.SongSearch).SongId)
				if err != nil {
					return []response.SearchResultResponse{}, err
				}

				songResponse := response.SongResponse{
					SongId:      song.SongId,
					Title:       song.Title,
					ArtistId:    song.ArtistId,
					AlbumId:     song.AlbumId,
					ReleaseDate: song.ReleaseDate,
					Duration:    song.Duration,
					File:        song.File,
					Album:       song.Album,
					Play:        song.Play,
					Artist:      song.Artist,
				}

				if r.Type == "song" {
					result = append(result, response.SearchResultResponse{
						Song:  songResponse,
						Type:  "",
						Count: 0,
						Title: "",
					})
					count++
				}
			}
		} else if res[0].Type == "album" {
			top5track, err := s.SongRepository.GetTop5TrackFromAlbum(res[0].Payload.(response.AlbumSearch).AlbumId)
			if err != nil {
				return []response.SearchResultResponse{}, err
			}

			for _, song := range top5track {
				songResponse := response.SongResponse{
					SongId:      song.SongId,
					Title:       song.Title,
					ArtistId:    song.ArtistId,
					AlbumId:     song.AlbumId,
					ReleaseDate: song.ReleaseDate,
					Duration:    song.Duration,
					File:        song.File,
					Album:       song.Album,
					Play:        song.Play,
					Artist:      song.Artist,
				}
				result = append(result, response.SearchResultResponse{
					Song:  songResponse,
					Type:  "album",
					Title: song.Title,
				})
			}
		} else if res[0].Type == "artist" {
			top5track, err := s.SongRepository.GetTop5TrackFromArtist(res[0].Payload.(response.ArtistSearch).ArtistId)
			if err != nil {
				return []response.SearchResultResponse{}, err
			}

			for _, song := range top5track {
				songResponse := response.SongResponse{
					SongId:      song.SongId,
					Title:       song.Title,
					ArtistId:    song.ArtistId,
					AlbumId:     song.AlbumId,
					ReleaseDate: song.ReleaseDate,
					Duration:    song.Duration,
					File:        song.File,
					Album:       song.Album,
					Play:        song.Play,
					Artist:      song.Artist,
				}
				result = append(result, response.SearchResultResponse{
					Song:  songResponse,
					Type:  "artist",
					Title: song.Title,
				})
			}
		}
	}

	var songRes []response.SearchResultResponse
	if len(result) <= 0 {
		allSong, err := s.SongRepository.GetAllSong()
		if err != nil {
			return result, nil
		}

		var filteredSongs []model.Song
		for _, song := range allSong {
			if utils.GetDistance(keyword, song.Title) <= 20 {
				filteredSongs = append(filteredSongs, song)
			}
		}

		sort.Slice(filteredSongs, func(i, j int) bool {
			return utils.GetDistance(keyword, filteredSongs[i].Title) > utils.GetDistance(keyword, filteredSongs[j].Title)
		})

		count := 0
		for _, song := range filteredSongs {
			if count == 5 {
				break
			}

			songResponse := response.SongResponse{
				SongId:      song.SongId,
				Title:       song.Title,
				ArtistId:    song.ArtistId,
				AlbumId:     song.AlbumId,
				ReleaseDate: song.ReleaseDate,
				Duration:    song.Duration,
				File:        song.File,
				Album:       song.Album,
				Play:        song.Play,
				Artist:      song.Artist,
			}

			songRes = append(songRes, response.SearchResultResponse{
				Song:  songResponse,
				Type:  "song",
				Title: song.Title,
			})
			count++
		}
		return songRes, nil
	}

	//if len(result) < 6 {
	//	count := 5 - len(result)
	//	for _, song := range songRes {
	//		if count >= 5 {
	//			break
	//		}
	//		result = append(result, song)
	//	}
	//	return result, nil
	//}

	return result, nil
}
