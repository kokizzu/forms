package common

import (
	"encoding/json"
	"strings"

	"github.com/webx-top/com"
)

func ParseCheckedValue(value string) []string {
	var checkedValues []string
	if value == `` {
		return checkedValues
	}
	if value[0] == '[' && value[len(value)-1] == ']' {
		var values []interface{}
		if err := json.Unmarshal(com.Str2bytes(value), &values); err == nil {
			checkedValues = make([]string, len(values))
			for i, v := range values {
				checkedValues[i] = com.String(v)
			}
		}
	} else {
		checkedValues = strings.Split(value, `,`)
	}
	return checkedValues
}
