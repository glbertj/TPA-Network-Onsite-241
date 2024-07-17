package utils

func GetDistance(query string, content string) int {
	lenQuery := len(query)
	lenContent := len(content)

	distance := make([][]int, lenQuery+1)
	for i := range distance {
		distance[i] = make([]int, lenContent+1)
	}

	for i := 0; i <= lenQuery; i++ {
		distance[i][0] = i
	}

	for j := 0; j <= lenContent; j++ {
		distance[0][j] = j
	}

	for i := 1; i <= lenQuery; i++ {
		for j := 1; j <= lenContent; j++ {
			cost := 0
			if query[i-1] != content[j-1] {
				cost = 1
			}
			distance[i][j] = min(distance[i-1][j]+1, distance[i][j-1]+1, distance[i-1][j-1]+cost)
		}
	}

	return distance[lenQuery][lenContent]

}
