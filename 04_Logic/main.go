package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupByAnagram(listOfStrings []string) *[][]string {
	list := make(map[string][]string)
	var result [][]string

	for _, word := range listOfStrings {
		key := sortStr(word)
		list[key] = append(list[key], word)
	}

	for _, words := range list {
		result = append(result, words)
	}
	return &result
}

func sortStr(k string) string {
	s := strings.Split(k, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	list := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	res := groupByAnagram(list)
	fmt.Println(*res)
}
