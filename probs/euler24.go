package probs

func Euler24() string {
	str := "0123456789"
	return permutation("", str)
}

var count int = 0

func permutation(prefix, str string) string {
	if len(str) == 0 {
		count++
		if count == 1000000 {
			return prefix
		}
		return ""
	}
	val := ""
	for i := 0; i < len(str); i++ {
		val = permutation(prefix+string(str[i]), str[:i]+str[i+1:])
		if len(val) != 0 {
			break
		}
	}

	return val
}
