PROTOC ?= protoc
PYTHON ?= python3

PB_TARGETS := APPKEY TOKEN
pb_src_APPKEY := appkey.proto
pb_dest_dir_py_APPKEY := py
pb_dest_dir_go_APPKEY := go/appkeypb
pb_dest_dir_APPKEY := $(pb_dest_dir_py_APPKEY) $(pb_dest_dir_go_APPKEY)
pb_gen_APPKEY := go/appkeypb/appkey.pb.go py/appkey_pb2.py
pb_src_TOKEN := token.proto
pb_dest_dir_py_TOKEN := py
pb_dest_dir_go_TOKEN := go/tokenpb
pb_dest_dir_TOKEN := $(pb_dest_dir_py_TOKEN) $(pb_dest_dir_go_TOKEN)
pb_gen_TOKEN := go/tokenpb/token.pb.go py/token_pb2.py
PB_GEN := $(pb_gen_APPKEY) $(pb_gen_TOKEN)
GENERATED_DEST_FILES := go/appkeypb/static.go py/appkey_data.py
GENERATED_SRC_FILES := appkey_data.py.tmpl appkey_static.go.tmpl appkey-links.json token-links.json token_data.py.tmpl token_static.go.tmpl
GENERATED_FILES := $(PB_GEN_APPKEY) $(PB_GEN_TOKEN) $(GENERATED_DEST_FILES)

.PHONY: all clean

all: $(GENERATED_FILES)

define PB_template =
$$(pb_gen_$(1)): $$(pb_src_$(1))
	mkdir -p $(pb_dest_dir_$(1))
	$(PROTOC) --python_out=$(pb_dest_dir_py_$(1)) --go_out=$(pb_dest_dir_go_$(1)) $(pb_src_$(1))
endef

$(foreach PB_TARGET,$(PB_TARGETS),$(eval $(call PB_template,$(PB_TARGET)))): 

$(GENERATED_DEST_FILES): $(GENERATED_SRC_FILES) $(PB_GEN) generate.py
	$(PYTHON) generate.py

clean:
	rm -fr go py
