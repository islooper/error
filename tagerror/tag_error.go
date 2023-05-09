package tagerror

import "errors"

type TagError struct {
	tag   string
	err   error
	extra string
}

func (t *TagError) Error() string {
	return t.err.Error()
}

func New(tag string, err error, extra string) *TagError {
	return &TagError{
		tag:   tag,
		err:   err,
		extra: extra,
	}
}

func (t *TagError) Extra() string {
	return t.extra
}

func (t *TagError) Tag() string {
	return t.tag
}

func (t *TagError) Is(target error) bool {
	return errors.Is(t.err, target)
}

func (t *TagError) As(target interface{}) bool {
	return errors.As(t.err, &target)
}
