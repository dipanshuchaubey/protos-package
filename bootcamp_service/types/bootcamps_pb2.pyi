from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class BootcampInfo(_message.Message):
    __slots__ = ("bootcamp_id", "title", "description", "website", "email", "name_slug", "careers")
    BOOTCAMP_ID_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    WEBSITE_FIELD_NUMBER: _ClassVar[int]
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    NAME_SLUG_FIELD_NUMBER: _ClassVar[int]
    CAREERS_FIELD_NUMBER: _ClassVar[int]
    bootcamp_id: str
    title: str
    description: str
    website: str
    email: str
    name_slug: str
    careers: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, bootcamp_id: _Optional[str] = ..., title: _Optional[str] = ..., description: _Optional[str] = ..., website: _Optional[str] = ..., email: _Optional[str] = ..., name_slug: _Optional[str] = ..., careers: _Optional[_Iterable[str]] = ...) -> None: ...

class CourseInfo(_message.Message):
    __slots__ = ("course_id", "title", "description")
    COURSE_ID_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    course_id: str
    title: str
    description: str
    def __init__(self, course_id: _Optional[str] = ..., title: _Optional[str] = ..., description: _Optional[str] = ...) -> None: ...

class Review(_message.Message):
    __slots__ = ("review_id", "user_id", "title", "message", "rating")
    REVIEW_ID_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    RATING_FIELD_NUMBER: _ClassVar[int]
    review_id: str
    user_id: str
    title: str
    message: str
    rating: int
    def __init__(self, review_id: _Optional[str] = ..., user_id: _Optional[str] = ..., title: _Optional[str] = ..., message: _Optional[str] = ..., rating: _Optional[int] = ...) -> None: ...
