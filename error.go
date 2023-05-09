package error

import (
	"errors"
	"fmt"
	"github.com/Islooper/error/tagerror"
	"sync"
	"time"
)

type terror struct {
	error *tagerror.TagError
	p     *terror
	n     *terror
}

type LinkError struct {
	h     *terror
	t     *terror
	size  int64
	mutex *sync.Mutex
}

func newLinkError() *LinkError {
	return &LinkError{
		h:     nil,
		t:     nil,
		size:  0,
		mutex: &sync.Mutex{},
	}
}

func (l *LinkError) Error(tag string, err string, format string, a ...any) *LinkError {
	if l == nil {
		l = newLinkError()
	}

	if tag == "" {
		tag = time.Now().String()
	}

	l.mutex.Lock()
	defer l.mutex.Unlock()
	node := &terror{
		error: tagerror.New(tag, errors.New(err), fmt.Sprintf(format, a...)),
	}

	if l.t == nil {
		l.h = node
		l.t = node
	} else {
		l.t.n = node
		node.p = l.t
		l.t = node
	}

	l.size += 1

	return l
}

// GetTagError tag 模式下只支持获取自定义tag
func (l *LinkError) GetTagError(tag string) []*tagerror.TagError {
	errs := make([]*tagerror.TagError, 0)

	if l.size == 0 {
		return errs
	}

	l.mutex.Lock()
	defer l.mutex.Unlock()

	if tag == "" {
		return errs
	}

	curr := l.h

	for curr != nil {
		if curr.error.Tag() == tag {
			errs = append(errs, curr.error)
		}
		curr = curr.n
	}

	return errs
}

func (l *LinkError) GetIndex(index int64) *tagerror.TagError {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.size == 0 || index >= l.size {
		return nil
	}

	if index == 0 {
		return l.h.error
	}

	node := l.h
	for i := 0; i < int(index); i++ {
		node = node.n
	}

	return node.error
}

func (l *LinkError) GetSize() int64 {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	return l.size
}

// 输出所有string
func (l *LinkError) String() string {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.h == nil {
		return ""
	}

	curr := l.h
	var result string
	for curr != nil {
		result += fmt.Sprintf("tag:%s => err:%s => extra:%s  \n", curr.error.Tag(),
			curr.error.Error(), curr.error.Extra())
		curr = curr.n
	}

	return result
}
