package core

func IsSpace(b byte) bool {
	return b == 32
}

func IsAlpha(b byte) bool {
	return (b >= 97 && b <= 122) || (b >= 65 && b <= 90)
}

func IsDigit(b byte) bool {
	return b >= 48 && b <= 57
}

func IsAlNum(b byte) bool {
	return IsAlpha(b) || IsDigit(b)
}

func ISOC(b byte) bool {
	return b == 40 || b == 41 || b == 123 || b == 125
}
