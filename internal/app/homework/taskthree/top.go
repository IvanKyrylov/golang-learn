package taskthree

import (
	"sort"
	"strings"
)

func Top10(s string) []string {
	if s == "" {
		return []string{}
	}
	splitString := strings.Fields(s)

	for i := range splitString {
		splitString[i] = strings.Trim(splitString[i], ".,:;!\"")
	}

	splitMap := make(map[string]int)
	for _, v := range splitString {
		splitMap[v]++
	}

	tempSl := make([]string, 0, 10)

	for k := range splitMap {
		tempSl = append(tempSl, k)
	}

	sort.Strings(tempSl)
	sort.Slice(tempSl, func(i, j int) bool {
		return splitMap[tempSl[i]] > splitMap[tempSl[j]]
	})

	return tempSl[:10]
}
