# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: bootcamp_service/bootcamps.proto
# Protobuf Python Version: 5.28.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    28,
    1,
    '',
    'bootcamp_service/bootcamps.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from bootcamp_service.request import bootcamps_pb2 as bootcamp__service_dot_request_dot_bootcamps__pb2
from bootcamp_service.response import bootcamps_pb2 as bootcamp__service_dot_response_dot_bootcamps__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n bootcamp_service/bootcamps.proto\x12\x13\x62ootcamp_service.v1\x1a(bootcamp_service/request/bootcamps.proto\x1a)bootcamp_service/response/bootcamps.proto2\x8d\x02\n\x0f\x42ootcampService\x12\x83\x01\n\x13GetBootcampsDetails\x12\x34.bootcamp_service.request.GetBootcampsDetailsRequest\x1a\x36.bootcamp_service.response.GetBootcampsDetailsResponse\x12t\n\x0e\x43reateBootcamp\x12/.bootcamp_service.request.CreateBootcampRequest\x1a\x31.bootcamp_service.response.CreateBootcampResponseB?Z=github.com/dipanshuchaubey/protos-package/bootcamp_service/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'bootcamp_service.bootcamps_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z=github.com/dipanshuchaubey/protos-package/bootcamp_service/v1'
  _globals['_BOOTCAMPSERVICE']._serialized_start=143
  _globals['_BOOTCAMPSERVICE']._serialized_end=412
# @@protoc_insertion_point(module_scope)
