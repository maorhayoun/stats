package helpers

import (
	//"fmt"
	"regexp"
	"strconv"
)

// time expression to support days and hours, e.g. 2.5d, 7h, 1d2h, 3 (hours)
var timeSpanExp = regexp.MustCompile(`((?P<days>[-]?\d+)?d)?((?P<hours>[-]?\d+)h?)?`)

func ParseHours(s string) int {
	match := timeSpanExp.FindStringSubmatch(s)
	result := make(map[string]string)
	for i, name := range timeSpanExp.SubexpNames() {
		result[name] = match[i]
	}

	days, _ := strconv.Atoi(result["days"])
	hours, _ := strconv.Atoi(result["hours"])

	//fmt.Printf("days: %d\n", days)
	//fmt.Printf("hours: %d\n", hours)

	return days*24 + hours
}
