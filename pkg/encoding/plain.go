package encoding

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/encoding"
	"strconv"
	"unsafe"
)

const (
	NameForPlain = "plain"
	NameForAny   = "*"
)

// 针对 text/plain 和 */*两种情况均走自定义的解析器
func init() {
	encoding.RegisterCodec(codec{})
	encoding.RegisterCodec(codecAny{})
}

type codecAny struct {
	codec
}

func (c codecAny) Name() string {
	return NameForAny
}

type codec struct{}

func (c codec) Marshal(v interface{}) ([]byte, error) {
	switch m := v.(type) {
	case string:
		return []byte(m), nil
	case int:
		return []byte(strconv.FormatInt(int64(m), 10)), nil
	case float64:
		return []byte(strconv.FormatFloat(m, 'f', 10, 64)), nil
	case bool:
		return []byte(strconv.FormatBool(m)), nil
	default:
		str := fmt.Sprintf("%v", m)
		return []byte(str), nil
	}
}

func (c codec) Unmarshal(data []byte, v interface{}) error {
	v = *(*string)(unsafe.Pointer(&data))
	return nil
}

func (c codec) Name() string {
	return NameForPlain
}
