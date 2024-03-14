package parser

import (
	"path/filepath"
)

type FileParser interface {
	Parse() ([]Player, error)
}

type Player struct {
	Name  string `json:"name"`
	Score int    `json:"high_score"`
}

func GetPath(fileName string) string {
	return filepath.Join("./examples", fileName)
}

func FindMaxAndMin(players []Player) (max, min Player) {
	max = players[0]
	min = players[0]

	for i := 0; i < len(players); i++ {
		if players[i].Score > max.Score {
			max = players[i]
		}

		if players[i].Score < min.Score {
			min = players[i]
		}
	}
	return max, min
}
