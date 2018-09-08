import os
import os.path
import sys
import subprocess

sys.path.insert(1, os.path.join(os.getcwd(), 'py'))

import appkey_pb2
from google.protobuf.json_format import Parse

links = appkey_pb2.Links()
Parse(open('links.json', 'r').read(), links)
links_pb = links.SerializeToString()
open('links.pbf', 'wb').write(links_pb)
go_static_tmpl = open('static.go.tmpl', 'r').read()
default_links_go = '{' + ', '.join(str(c) for c in links_pb) + '}'
go_static = go_static_tmpl.format(default_links_pb=default_links_go)
open('go/appkeypb/static.go', 'w').write(go_static)
subprocess.call(['go', 'fmt', 'static.go'])
py_data_tmpl = open('appkey_data.py.tmpl', 'r').read()
py_data = py_data_tmpl.format(default_links_pb=repr(links_pb))
open('py/appkey_data.py', 'w').write(py_data)
subprocess.call(['yapf', '-i', 'py/appkey_data.py'])
