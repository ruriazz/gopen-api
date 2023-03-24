package sliceHelper

func StringInSlice(s string, l []string) bool {
	for _, i := range l {
		if i == s {
			return true
		}
	}
	return false
}
