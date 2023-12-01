package goutils

type Position struct {
	X int
	Y int
}

func StrReverse(s string) string {
	var ret string
	for _, v := range s {
		ret = string(v) + ret
	}
	return ret
}
