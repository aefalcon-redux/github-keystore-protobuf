PROTOC ?= protoc
PYTHON ?= python3

.PHONY: all clean

all: go/appkeypb/appkey.pb.go py/appkey_pb2.py go/appkeypb/static.go go/tokenpb/token.pb.go py/token_pb2.py

go/appkeypb/appkey.pb.go py/appkey_pb2.py: appkey.proto go/appkeypb py
	$(PROTOC) --python_out=py --go_out=go/appkeypb $<

go/tokenpb/token.pb.go py/token_pb2.py: token.proto go/tokenpb py
	$(PROTOC) --python_out=py --go_out=go/tokenpb $<

py go go/appkeypb go/tokenpb:
	mkdir -p $@

links.pbf go/appkeypb/static.go py/appkey_data.py: links.json appkey.proto generate.py static.go.tmpl appkey_data.py.tmpl
	$(PYTHON) generate.py

clean:
	rm -fr go py *.pbf
