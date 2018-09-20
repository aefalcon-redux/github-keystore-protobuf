# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: location.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='location.proto',
  package='mobettersoftware.protobuf.location',
  syntax='proto3',
  serialized_options=_b('Z:github.com/aefalcon/github-keystore-protobuf/go/locationpb'),
  serialized_pb=_b('\n\x0elocation.proto\x12\"mobettersoftware.protobuf.location\"^\n\x08Location\x12\x37\n\x02s3\x18\x01 \x01(\x0b\x32).mobettersoftware.protobuf.location.S3RefH\x00\x12\r\n\x03url\x18\x02 \x01(\tH\x00\x42\n\n\x08location\"4\n\x05S3Ref\x12\x0e\n\x06\x62ucket\x18\x01 \x01(\t\x12\x0b\n\x03key\x18\x02 \x01(\t\x12\x0e\n\x06region\x18\x03 \x01(\tB<Z:github.com/aefalcon/github-keystore-protobuf/go/locationpbb\x06proto3')
)




_LOCATION = _descriptor.Descriptor(
  name='Location',
  full_name='mobettersoftware.protobuf.location.Location',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='s3', full_name='mobettersoftware.protobuf.location.Location.s3', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='url', full_name='mobettersoftware.protobuf.location.Location.url', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='location', full_name='mobettersoftware.protobuf.location.Location.location',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=54,
  serialized_end=148,
)


_S3REF = _descriptor.Descriptor(
  name='S3Ref',
  full_name='mobettersoftware.protobuf.location.S3Ref',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='bucket', full_name='mobettersoftware.protobuf.location.S3Ref.bucket', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='key', full_name='mobettersoftware.protobuf.location.S3Ref.key', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='region', full_name='mobettersoftware.protobuf.location.S3Ref.region', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=150,
  serialized_end=202,
)

_LOCATION.fields_by_name['s3'].message_type = _S3REF
_LOCATION.oneofs_by_name['location'].fields.append(
  _LOCATION.fields_by_name['s3'])
_LOCATION.fields_by_name['s3'].containing_oneof = _LOCATION.oneofs_by_name['location']
_LOCATION.oneofs_by_name['location'].fields.append(
  _LOCATION.fields_by_name['url'])
_LOCATION.fields_by_name['url'].containing_oneof = _LOCATION.oneofs_by_name['location']
DESCRIPTOR.message_types_by_name['Location'] = _LOCATION
DESCRIPTOR.message_types_by_name['S3Ref'] = _S3REF
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Location = _reflection.GeneratedProtocolMessageType('Location', (_message.Message,), dict(
  DESCRIPTOR = _LOCATION,
  __module__ = 'location_pb2'
  # @@protoc_insertion_point(class_scope:mobettersoftware.protobuf.location.Location)
  ))
_sym_db.RegisterMessage(Location)

S3Ref = _reflection.GeneratedProtocolMessageType('S3Ref', (_message.Message,), dict(
  DESCRIPTOR = _S3REF,
  __module__ = 'location_pb2'
  # @@protoc_insertion_point(class_scope:mobettersoftware.protobuf.location.S3Ref)
  ))
_sym_db.RegisterMessage(S3Ref)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
