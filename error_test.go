package error

import (
	"testing"
)

func TestLinkError_Normal(t *testing.T) {

	err := GetError().Error("TestLinkError_Success", "TestLinkError_Success", "").
		Error("this is a test tag", "this is a test error", "").
		Error("", "test empty tag", "")

	defer func() {
		Destroy(err)
	}()

	t.Logf(err.String())
}

func TestLinkError_NoInit(t *testing.T) {
	var err *LinkError
	err = err.Error("TestLinkError_NoInit", "TestLinkError_NoInit", "")

	defer func() {
		Destroy(err)
	}()

	t.Logf(err.String())
}

func TestLinkError_GetTag(t *testing.T) {
	var err *LinkError
	err = err.Error("TestLinkError_GetTag", "TestLinkError_GetTag errors 1", "").
		Error("TestLinkError_GetTag", "TestLinkError_GetTag errors 2", "").
		Error("", "test error", "")

	defer func() {
		Destroy(err)
	}()

	t.Logf(err.String())

	errs := err.GetTagError("TestLinkError_GetTag")

	for _, e := range errs {
		t.Logf("tag: %s, error: %s", e.Tag(), e.Error())
	}

}
