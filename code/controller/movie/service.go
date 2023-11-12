package movie

import (
	"strconv"
)

func getDecadeBounds(decade string) (string, string) {
	decadeStart, _ := strconv.Atoi(decade)
	decadeEnd := decadeStart + 9

	return strconv.Itoa(decadeStart-1), strconv.Itoa(decadeEnd+1)
}
