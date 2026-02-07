package version

import (
	"strings"
	"unicode"
)

func LessThan(v1, v2 string) bool {
	return Compare(v1, v2) == -1
}

func MoreThan(v1, v2 string) bool {
	return Compare(v1, v2) == 1
}

func Compare(v1, v2 string) int {
	if v1 == v2 {
		return 0
	}
	if v1 == "" {
		return -1
	}
	if v2 == "" {
		return 1
	}

	epoch1, ver1, rel1 := parseEvr(v1)
	epoch2, ver2, rel2 := parseEvr(v2)

	ret := rpmvercmp(epoch1, epoch2)
	if ret == 0 {
		ret = rpmvercmp(ver1, ver2)
		if ret == 0 && rel1 != "" && rel2 != "" {
			ret = rpmvercmp(rel1, rel2)
		}
	}

	return ret
}

func parseEvr(evr string) (string, string, string) {
	var epoch string
	var version string
	var release string

	index := 0

	if len(evr) == 0 {
		return "0", "", ""
	}

	for index < len(evr) && unicode.IsDigit(rune(evr[index])) {
		index++
	}

	v := strings.LastIndex(evr, "-")

	if index < len(evr) && evr[index] == ':' {
		epoch = evr[0:index]
		index++
		if epoch == "" {
			epoch = "0"
		}
	} else {
		epoch = "0"
		index = 0
	}

	if v == -1 {
		release = ""
		version = evr[index:]
	} else {
		release = evr[v+1:]
		version = evr[index:v]
	}

	return epoch, version, release
}

func rpmvercmp(a string, b string) int {
	var ret int
	var isNum bool
	one := 0
	two := 0
	var ptr1, ptr2 int
	var x string
	var y string

	if a == b {
		return 0
	}
	for one < len(a) && two < len(b) {
		for one < len(a) && !(unicode.IsDigit(rune(a[one])) || unicode.IsLetter(rune(a[one]))) {
			one++
		}
		for two < len(b) && !(unicode.IsDigit(rune(b[two])) || unicode.IsLetter(rune(b[two]))) {
			two++
		}

		if one > len(a) || two > len(b) {
			break
		}

		if one-ptr1 != two-ptr2 {
			if one-ptr1 < two-ptr2 {
				ret = -1
				return ret
			} else {
				ret = 1
				return ret
			}
		}

		ptr1 = one
		ptr2 = two

		if unicode.IsDigit(rune(a[ptr1])) {
			for ptr1 < len(a) && unicode.IsDigit(rune(a[ptr1])) {
				ptr1++
			}
			for ptr2 < len(b) && unicode.IsDigit(rune(b[ptr2])) {
				ptr2++
			}
			isNum = true
		} else {
			for ptr1 < len(a) && unicode.IsLetter(rune(a[ptr1])) {
				ptr1++
			}
			for ptr2 < len(b) && unicode.IsLetter(rune(b[ptr2])) {
				ptr2++
			}
			isNum = false
		}

		x = a[one:ptr1]
		y = b[two:ptr2]

		if one == ptr1 {
			ret = -1
			return ret
		}

		if two == ptr2 {
			if isNum {
				ret = 1
				return ret
			} else {
				ret = -1
				return ret
			}
		}

		if isNum {
			i := 0
			j := 0
			for i < len(x) && x[i] == '0' {
				i++
			}
			for j < len(y) && y[j] == '0' {
				j++
			}
			x = x[i:]
			y = y[j:]

			if len(x) > len(y) {
				ret = 1
				return ret
			}
			if len(x) < len(y) {
				ret = -1
				return ret
			}
		}

		if !isNum {
			ret = strings.Compare(x, y)
			if ret != 0 {
				return ret
			}
		}

		if len(x) == len(y) && isNum {
			ret = strings.Compare(x, y)
			if ret != 0 {
				return ret
			}
		}

		one = ptr1
		two = ptr2

	}
	if one > len(a) && two > len(b) {
		ret = 0
	}
	if one >= len(a) {
		if two < len(b) && unicode.IsLetter(rune(b[two])) {
			return 1
		} else {
			if two >= len(b) {
				return 0
			} else {
				return -1
			}
		}
	}
	if two >= len(b) {
		if one < len(a) && unicode.IsLetter(rune(a[one])) {
			return -1 // Alpha < Empty
		}
		return 1 // Numeric > Empty
	}

	return ret
}
