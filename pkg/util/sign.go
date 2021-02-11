package util

import (
	"sort"
)

const signKey = "sign"

func ValidSign(param map[string]interface{}) bool {
	// size = len(param) - 1
	keys := make([]string, 0, len(param))
	for k := range param {
		if k != signKey {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	signStr := ""
	for index, k := range keys {
		signStr += k + "=" + param[k].(string)
		if index != len(keys)-1 {
			signStr += ","
		}
	}
	sign, ok := param[signKey]
	if ok && sign.(string) == EncodeMD5(signStr) {
		return true
	}

	return false
}
