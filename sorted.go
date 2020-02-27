package main

import (
	"github.com/tidwall/collate"
	"sort"
)

type i18nStruct struct {
	lang string
	arr  []string
}

func (i18n i18nStruct) Less(i, j int) bool {
	less := collate.IndexString(i18n.lang)
	return less(i18n.arr[i], i18n.arr[j])
}

func (i18n i18nStruct) Len() int {
	return len(i18n.arr)
}

func (i18n i18nStruct) Swap(i, j int) {
	i18n.arr[i], i18n.arr[j] = i18n.arr[j], i18n.arr[i]
}

func sorted(arr []string, lang string) []string {
	i18n := i18nStruct{lang, arr}
	sort.Sort(i18n)
	return i18n.arr
}
