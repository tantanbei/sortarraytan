package sortarraytan

import (
	"errors"
)

func IntArrayReverse(origin []int) []int {
	originLen := len(origin)
	for i := 0; i < originLen/2; i++ {
		origin[i], origin[originLen-1-i] = origin[originLen-1-i], origin[i]
	}
	return origin
}

func ByteArrayReverse(origin []byte) []byte {
	originLen := len(origin)
	for i := 0; i < originLen/2; i++ {
		origin[i], origin[originLen-1-i] = origin[originLen-1-i], origin[i]
	}
	return origin
}

func StringArrayReverse(origin []string) []string {
	originLen := len(origin)
	for i := 0; i < originLen/2; i++ {
		origin[i], origin[originLen-1-i] = origin[originLen-1-i], origin[i]
	}
	return origin
}

func InterfaceArrayReverse(origin interface{}) (interface{}, error) {
	switch origin.(type) {
	case []byte:
		originB := origin.([]byte)
		originLen := len(originB)
		for i := 0; i < originLen/2; i++ {
			originB[i], originB[originLen-1-i] = originB[originLen-1-i], originB[i]
		}
		return originB, nil
	case []int:
		originI := origin.([]int)
		originLen := len(originI)
		for i := 0; i < originLen/2; i++ {
			originI[i], originI[originLen-1-i] = originI[originLen-1-i], originI[i]
		}
		return originI, nil
	case []string:
		originS := origin.([]string)
		originLen := len(originS)
		for i := 0; i < originLen/2; i++ {
			originS[i], originS[originLen-1-i] = originS[originLen-1-i], originS[i]
		}
		return originS, nil
	default:
		return nil, errors.New("The input is not []byte / []int / []string")
	}
}
