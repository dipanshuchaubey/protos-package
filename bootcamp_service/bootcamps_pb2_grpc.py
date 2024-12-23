# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from bootcamp_service.request import bootcamps_pb2 as bootcamp__service_dot_request_dot_bootcamps__pb2
from bootcamp_service.response import bootcamps_pb2 as bootcamp__service_dot_response_dot_bootcamps__pb2

GRPC_GENERATED_VERSION = '1.68.1'
GRPC_VERSION = grpc.__version__
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    raise RuntimeError(
        f'The grpc package installed is at version {GRPC_VERSION},'
        + f' but the generated code in bootcamp_service/bootcamps_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
    )


class BootcampServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetBootcampsDetails = channel.unary_unary(
                '/bootcamp_service.v1.BootcampService/GetBootcampsDetails',
                request_serializer=bootcamp__service_dot_request_dot_bootcamps__pb2.GetBootcampsDetailsRequest.SerializeToString,
                response_deserializer=bootcamp__service_dot_response_dot_bootcamps__pb2.GetBootcampsDetailsResponse.FromString,
                _registered_method=True)
        self.CreateBootcamp = channel.unary_unary(
                '/bootcamp_service.v1.BootcampService/CreateBootcamp',
                request_serializer=bootcamp__service_dot_request_dot_bootcamps__pb2.CreateBootcampRequest.SerializeToString,
                response_deserializer=bootcamp__service_dot_response_dot_bootcamps__pb2.CreateBootcampResponse.FromString,
                _registered_method=True)


class BootcampServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetBootcampsDetails(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CreateBootcamp(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_BootcampServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetBootcampsDetails': grpc.unary_unary_rpc_method_handler(
                    servicer.GetBootcampsDetails,
                    request_deserializer=bootcamp__service_dot_request_dot_bootcamps__pb2.GetBootcampsDetailsRequest.FromString,
                    response_serializer=bootcamp__service_dot_response_dot_bootcamps__pb2.GetBootcampsDetailsResponse.SerializeToString,
            ),
            'CreateBootcamp': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateBootcamp,
                    request_deserializer=bootcamp__service_dot_request_dot_bootcamps__pb2.CreateBootcampRequest.FromString,
                    response_serializer=bootcamp__service_dot_response_dot_bootcamps__pb2.CreateBootcampResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'bootcamp_service.v1.BootcampService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('bootcamp_service.v1.BootcampService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class BootcampService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetBootcampsDetails(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/bootcamp_service.v1.BootcampService/GetBootcampsDetails',
            bootcamp__service_dot_request_dot_bootcamps__pb2.GetBootcampsDetailsRequest.SerializeToString,
            bootcamp__service_dot_response_dot_bootcamps__pb2.GetBootcampsDetailsResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def CreateBootcamp(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/bootcamp_service.v1.BootcampService/CreateBootcamp',
            bootcamp__service_dot_request_dot_bootcamps__pb2.CreateBootcampRequest.SerializeToString,
            bootcamp__service_dot_response_dot_bootcamps__pb2.CreateBootcampResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
