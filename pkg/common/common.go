package common

import (
	"fmt"
	"strconv"
	"strings"
)

// DaysInMonth := map[int]int{
// 	1:1,

// }

func DateStringToInt(date string) int {
	split := strings.Split(date, "-")
	year, _ := strconv.Atoi(split[0])
	month, _ := strconv.Atoi(split[1])
	day, _ := strconv.Atoi(split[2])
	result := year*10000 + month*100 + day
	return result
}

func DateIntToString(date int) string {
	day := strconv.Itoa(date % 100)
	date = (date - (date % 100)) / 100
	month := strconv.Itoa(date % 100)
	date = (date - (date % 100)) / 100
	year := strconv.Itoa(date)

	return fmt.Sprintf("%s-%s-%s", year, month, day)
}
