package error

func GetError() *LinkError {
	return newLinkError()
}

func Destroy(l *LinkError) {
	l = nil
}
