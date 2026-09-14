package main
import "testing"
func TestPortDefault(t *testing.T){if port()!="8080"{t.Fatal("端口默认值错误")}}
