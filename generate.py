import os
import os.path
import sys
import subprocess

sys.path.insert(1, os.path.join(os.getcwd(), 'py'))

import appkey_pb2
import token_pb2
import jinja2
from google.protobuf.json_format import Parse


def filter_go_encode_bytes(bval):
    return '{' + ', '.join(str(c) for c in bval) + '}'

def filter_py_repr(val):
    return repr(val)

tmpl_env = jinja2.Environment(loader=jinja2.FileSystemLoader('./'))
tmpl_env.filters['go_bytes'] = filter_go_encode_bytes
tmpl_env.filters['py_repr'] = filter_py_repr

pb_data_files = [
    ('appkey-links', appkey_pb2.Links, 'default_appkey_links_pb'),
    ('token-links', token_pb2.Links, 'default_token_links_pb')
]
for name, proto_class, ident in pb_data_files:
    src = f'{name}.json'
    pbobj = proto_class()
    Parse(open(src, 'r').read(), pbobj)
    tmpl_env.globals[ident] = pbobj.SerializeToString()

template_files = [
    ('appkey_static.go.tmpl', 'go/appkeypb/static.go'),
    ('appkey_data.py.tmpl', 'py/appkey_data.py'),
    ('token_static.go.tmpl', 'go/tokenpb/static.go'),
    ('token_data.py.tmpl', 'py/token_data.py'),
]
for tmpl_name, tmpl_dest in template_files:
    templ = tmpl_env.get_template(tmpl_name)
    dest_file = open(tmpl_dest, 'w')
    dest_file.write(templ.render())
    dest_file.close()
    if tmpl_dest.endswith('.py'):
        subprocess.check_call(['yapf', '-i', tmpl_dest])
    if tmpl_dest.endswith('.go'):
        subprocess.check_call(['go', 'fmt', tmpl_dest])
