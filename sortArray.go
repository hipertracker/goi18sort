package main

import (
	"github.com/tidwall/collate"
	"sort"
)

func sortArray(arr []string, lang string) {
	less := collate.IndexString(lang)
	sort.SliceStable(arr, func(i, j int) bool {
		return less(arr[i], arr[j])
	})
}
