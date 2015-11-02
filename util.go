package gw2api

import (
	"bytes"
	"strconv"
)

func stringSlice(ids []int) []string {
	newIds := make([]string, len(ids))
	for i, id := range ids {
		newIds[i] = strconv.Itoa(id)
	}
	return newIds
}

func commaList(ids []string) string {
	var appendix bytes.Buffer
	for i, id := range ids {
		if i > 0 {
			appendix.WriteString(",")
		}
		appendix.WriteString(id)
	}
	return appendix.String()
}

func flagGet(n, pos uint) bool {
	return n&1<<pos == 0
}

func flagSet(n, pos uint) {
	n |= 1 << pos
}
