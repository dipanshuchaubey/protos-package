from bootcamp_service.types import bootcamps_pb2 as _bootcamps_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GetBootcampsDetailsResponse(_message.Message):
    __slots__ = ("data",)
    class Data(_message.Message):
        __slots__ = ("bootcamp", "course", "reviews")
        BOOTCAMP_FIELD_NUMBER: _ClassVar[int]
        COURSE_FIELD_NUMBER: _ClassVar[int]
        REVIEWS_FIELD_NUMBER: _ClassVar[int]
        bootcamp: _bootcamps_pb2.BootcampInfo
        course: _bootcamps_pb2.CourseInfo
        reviews: _containers.RepeatedCompositeFieldContainer[_bootcamps_pb2.Review]
        def __init__(self, bootcamp: _Optional[_Union[_bootcamps_pb2.BootcampInfo, _Mapping]] = ..., course: _Optional[_Union[_bootcamps_pb2.CourseInfo, _Mapping]] = ..., reviews: _Optional[_Iterable[_Union[_bootcamps_pb2.Review, _Mapping]]] = ...) -> None: ...
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: _containers.RepeatedCompositeFieldContainer[GetBootcampsDetailsResponse.Data]
    def __init__(self, data: _Optional[_Iterable[_Union[GetBootcampsDetailsResponse.Data, _Mapping]]] = ...) -> None: ...

class CreateBootcampResponse(_message.Message):
    __slots__ = ("success", "data")
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    success: bool
    data: _bootcamps_pb2.BootcampInfo
    def __init__(self, success: bool = ..., data: _Optional[_Union[_bootcamps_pb2.BootcampInfo, _Mapping]] = ...) -> None: ...
