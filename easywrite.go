package easy

import (
	"fmt"
	"bytes"
	"io"
	"reflect"
	"github.com/AlasdairF/Conv"
)


func Write(w io.Writer, a ...interface{}) {
	for _, p := range a {
		switch reflect.TypeOf(p).Kind() {
			case reflect.String:
				w.Write([]byte(reflect.ValueOf(p).String()))
			case reflect.Slice:
				w.Write(reflect.ValueOf(p).Bytes())
			case reflect.Int: case reflect.Int8: case reflect.Int16: case reflect.Int32: case reflect.Int64:
				w.Write(conv.Bytes(int(reflect.ValueOf(p).Int())))
			case reflect.Uint: case reflect.Uint8: case reflect.Uint16: case reflect.Uint32: case reflect.Uint64:
				w.Write(conv.Bytes(int(reflect.ValueOf(p).Uint())))
		}
	}
}
