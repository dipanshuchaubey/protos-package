from google.protobuf import descriptor_pb2 as _descriptor_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class UserTypes(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    USER_TYPE_UNSPECIFIED: _ClassVar[UserTypes]
    USER_TYPE_ACTIVE: _ClassVar[UserTypes]
    USER_TYPE_INACTIVE: _ClassVar[UserTypes]
USER_TYPE_UNSPECIFIED: UserTypes
USER_TYPE_ACTIVE: UserTypes
USER_TYPE_INACTIVE: UserTypes
LABEL_FIELD_NUMBER: _ClassVar[int]
label: _descriptor.FieldDescriptor

class GetUserRequest(_message.Message):
    __slots__ = ("user_id", "tenant_id")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    TENANT_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: str
    tenant_id: str
    def __init__(self, user_id: _Optional[str] = ..., tenant_id: _Optional[str] = ...) -> None: ...

class GetUsersRequest(_message.Message):
    __slots__ = ("user_ids", "tenant_id")
    USER_IDS_FIELD_NUMBER: _ClassVar[int]
    TENANT_ID_FIELD_NUMBER: _ClassVar[int]
    user_ids: _containers.RepeatedScalarFieldContainer[str]
    tenant_id: str
    def __init__(self, user_ids: _Optional[_Iterable[str]] = ..., tenant_id: _Optional[str] = ...) -> None: ...

class UserInfo(_message.Message):
    __slots__ = ("user_id", "tenant_id", "full_name", "email", "user_type")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    TENANT_ID_FIELD_NUMBER: _ClassVar[int]
    FULL_NAME_FIELD_NUMBER: _ClassVar[int]
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    USER_TYPE_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    tenant_id: str
    full_name: str
    email: str
    user_type: UserTypes
    def __init__(self, user_id: _Optional[int] = ..., tenant_id: _Optional[str] = ..., full_name: _Optional[str] = ..., email: _Optional[str] = ..., user_type: _Optional[_Union[UserTypes, str]] = ...) -> None: ...

class GetUserResponse(_message.Message):
    __slots__ = ("success", "data")
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    success: bool
    data: UserInfo
    def __init__(self, success: bool = ..., data: _Optional[_Union[UserInfo, _Mapping]] = ...) -> None: ...

class GetUsersResponse(_message.Message):
    __slots__ = ("success", "data")
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    success: bool
    data: _containers.RepeatedCompositeFieldContainer[UserInfo]
    def __init__(self, success: bool = ..., data: _Optional[_Iterable[_Union[UserInfo, _Mapping]]] = ...) -> None: ...
