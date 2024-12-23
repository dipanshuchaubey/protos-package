from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class GetBootcampsDetailsRequest(_message.Message):
    __slots__ = ("bootcamp_ids",)
    BOOTCAMP_IDS_FIELD_NUMBER: _ClassVar[int]
    bootcamp_ids: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, bootcamp_ids: _Optional[_Iterable[str]] = ...) -> None: ...

class CreateBootcampRequest(_message.Message):
    __slots__ = ("title", "description", "website", "phone", "email", "careers", "address")
    TITLE_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    WEBSITE_FIELD_NUMBER: _ClassVar[int]
    PHONE_FIELD_NUMBER: _ClassVar[int]
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    CAREERS_FIELD_NUMBER: _ClassVar[int]
    ADDRESS_FIELD_NUMBER: _ClassVar[int]
    title: str
    description: str
    website: str
    phone: str
    email: str
    careers: _containers.RepeatedScalarFieldContainer[str]
    address: str
    def __init__(self, title: _Optional[str] = ..., description: _Optional[str] = ..., website: _Optional[str] = ..., phone: _Optional[str] = ..., email: _Optional[str] = ..., careers: _Optional[_Iterable[str]] = ..., address: _Optional[str] = ...) -> None: ...
