# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: debugtalk.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0f\x64\x65\x62ugtalk.proto\x12\x05proto\"\x07\n\x05\x45mpty\"!\n\x10GetNamesResponse\x12\r\n\x05names\x18\x01 \x03(\t\")\n\x0b\x43\x61llRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0c\n\x04\x61rgs\x18\x02 \x01(\x0c\"\x1d\n\x0c\x43\x61llResponse\x12\r\n\x05value\x18\x01 \x01(\x0c\x32o\n\tDebugTalk\x12\x31\n\x08GetNames\x12\x0c.proto.Empty\x1a\x17.proto.GetNamesResponse\x12/\n\x04\x43\x61ll\x12\x12.proto.CallRequest\x1a\x13.proto.CallResponseB\x11Z\x0fplugin/go/protob\x06proto3')



_EMPTY = DESCRIPTOR.message_types_by_name['Empty']
_GETNAMESRESPONSE = DESCRIPTOR.message_types_by_name['GetNamesResponse']
_CALLREQUEST = DESCRIPTOR.message_types_by_name['CallRequest']
_CALLRESPONSE = DESCRIPTOR.message_types_by_name['CallResponse']
Empty = _reflection.GeneratedProtocolMessageType('Empty', (_message.Message,), {
  'DESCRIPTOR' : _EMPTY,
  '__module__' : 'debugtalk_pb2'
  # @@protoc_insertion_point(class_scope:proto.Empty)
  })
_sym_db.RegisterMessage(Empty)

GetNamesResponse = _reflection.GeneratedProtocolMessageType('GetNamesResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETNAMESRESPONSE,
  '__module__' : 'debugtalk_pb2'
  # @@protoc_insertion_point(class_scope:proto.GetNamesResponse)
  })
_sym_db.RegisterMessage(GetNamesResponse)

CallRequest = _reflection.GeneratedProtocolMessageType('CallRequest', (_message.Message,), {
  'DESCRIPTOR' : _CALLREQUEST,
  '__module__' : 'debugtalk_pb2'
  # @@protoc_insertion_point(class_scope:proto.CallRequest)
  })
_sym_db.RegisterMessage(CallRequest)

CallResponse = _reflection.GeneratedProtocolMessageType('CallResponse', (_message.Message,), {
  'DESCRIPTOR' : _CALLRESPONSE,
  '__module__' : 'debugtalk_pb2'
  # @@protoc_insertion_point(class_scope:proto.CallResponse)
  })
_sym_db.RegisterMessage(CallResponse)

_DEBUGTALK = DESCRIPTOR.services_by_name['DebugTalk']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\017plugin/go/proto'
  _EMPTY._serialized_start=26
  _EMPTY._serialized_end=33
  _GETNAMESRESPONSE._serialized_start=35
  _GETNAMESRESPONSE._serialized_end=68
  _CALLREQUEST._serialized_start=70
  _CALLREQUEST._serialized_end=111
  _CALLRESPONSE._serialized_start=113
  _CALLRESPONSE._serialized_end=142
  _DEBUGTALK._serialized_start=144
  _DEBUGTALK._serialized_end=255
# @@protoc_insertion_point(module_scope)
