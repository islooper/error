
# Tag Error 🍺

灵感来源：日志代码项目侵入严重，日志平台一次错误日志上下文太长，不利于快速定位问题。总结：方便写日志。

自用error库，用tag给error分组，支持快速查找。支持链式新增，支持自定义error以及对出错的参数进行记录。

## 快速开始 🍺

### 安装

```shell

go get github.com/Islooer/error

```

### 使用

```go
 package error

import (
	"errors"
	"testing"
)

func TestLinkError_Success(t *testing.T) {

	err := GetError().Error("TestLinkError_Success", errors.New("TestLinkError_Success"), "").
		Error("this is a test tag", errors.New("this is a test error"), "").
		Error("", errors.New("test empty tag"), "")

	defer func() {
		Destroy(err)
	}()

	t.Logf(err.String())
}

func TestLinkError_NoInit(t *testing.T) {
	var err *LinkError
	err = err.Error("TestLinkError_NoInit", errors.New("TestLinkError_NoInit"), "")

	defer func() {
		Destroy(err)
	}()

	t.Logf(err.String())
}

func TestLinkError_GetTag(t *testing.T) {
	var err *LinkError
	err = err.Error("TestLinkError_GetTag", errors.New("TestLinkError_GetTag errors 1"), "").
		Error("TestLinkError_GetTag", errors.New("TestLinkError_GetTag errors 2"), "").
		Error("", errors.New("test error"), "")

	defer func() {
		Destroy(err)
	}()

	t.Logf(err.String())

	errs := err.GetTagError("TestLinkError_GetTag")

	for _, e := range errs {
		t.Logf("tag: %s, error: %s", e.Tag(), e.Error())
	}

}

```