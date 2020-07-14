package arrutil

import (
	"fmt"
	"strings"
)

func GroupByMultiKey(arr []map[string]string, keyArr []string, splitKey string) map[string][]map[string]string {
	groupedHash := map[string][]map[string]string{}
	for _, v := range arr {
		var hashKeyArr []string
		var hashKey string
		hashKey = ""
		for _, k := range keyArr {
			hashKeyArr = append(hashKeyArr, v[k])
		}
		hashKey = strings.Join(hashKeyArr, splitKey)
		groupedHash[hashKey] = append(groupedHash[hashKey], v)
	}
	return groupedHash
}

func InArray(val string, array []string) (isInArray bool) {
	for _, v := range array {
		if val == v {
			return true
		}
	}
	return false
}

func MakeArrToHash(listKey string, arr []map[string]string) map[string][]map[string]string {
	hashArr := map[string][]map[string]string{}
	for _, v := range arr {
		hashArr[v[listKey]] = append(hashArr[v[listKey]], v)
	}
	return hashArr
}

func setMap() {
	mapdata := map[string]int{
		"月": 1,
		"火": 2,
		"水": 3,
	}
	fmt.Println(mapdata["火"])
}
