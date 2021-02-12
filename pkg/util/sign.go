package util

import (
	"fmt"
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
		signStr += fmt.Sprintf("%s=%v", k, param[k])
		if index != len(keys)-1 {
			signStr += ","
		}
	}
	sign, ok := param[signKey]
	return ok && sign.(string) == EncodeMD5(signStr)
}
