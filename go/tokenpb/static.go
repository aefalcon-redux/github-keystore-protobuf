package tokenpb

import (
	"github.com/golang/protobuf/proto"
)

var DefaultLinks Links
var DefaultLinksPb = []byte{10, 21, 97, 112, 112, 47, 123, 65, 112, 112, 73, 100, 125, 47, 116, 111, 107, 101, 110, 46, 112, 98, 102, 18, 37, 105, 110, 115, 116, 97, 108, 108, 47, 123, 65, 112, 112, 73, 100, 125, 47, 123, 73, 110, 115, 116, 97, 108, 108, 73, 100, 125, 47, 116, 111, 107, 101, 110, 46, 112, 98, 102}

func init() {
	err := proto.Unmarshal(DefaultLinksPb, &DefaultLinks)
	if err != nil {
		panic(err.Error())
	}
}
