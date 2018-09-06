PROTOC ?= protoc

.PHONY: all clean

all: go/appkeypb/appkey.pb.go py/appkey_pb2.py

go/appkeypb/appkey.pb.go py/appkey_pb2.py: appkey.proto go/appkeypb py
	$(PROTOC) --python_out=py --go_out=go/appkeypb $<

py go go/appkeypb:
	mkdir -p $@

clean:
	rm -fr go py
