package protocol;

import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.asyncServerStreamingCall;
import static io.grpc.stub.ClientCalls.asyncClientStreamingCall;
import static io.grpc.stub.ClientCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ClientCalls.blockingUnaryCall;
import static io.grpc.stub.ClientCalls.blockingServerStreamingCall;
import static io.grpc.stub.ClientCalls.futureUnaryCall;
import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.asyncServerStreamingCall;
import static io.grpc.stub.ServerCalls.asyncClientStreamingCall;
import static io.grpc.stub.ServerCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedStreamingCall;

/**
 * <pre>
 * api rpc 服务
 * </pre>
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.4.0)",
    comments = "Source: config.proto")
public final class ConfServerGrpc {

  private ConfServerGrpc() {}

  public static final String SERVICE_NAME = "protocol.ConfServer";

  // Static method descriptors that strictly reflect the proto.
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<protocol.Config.ReqGetConfig,
      protocol.Config.RespGetConfig> METHOD_GET_CONFIG =
      io.grpc.MethodDescriptor.<protocol.Config.ReqGetConfig, protocol.Config.RespGetConfig>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "protocol.ConfServer", "GetConfig"))
          .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Config.ReqGetConfig.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Config.RespGetConfig.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<protocol.Config.ReqSetConfig,
      protocol.Config.RespSetConfig> METHOD_SET_CONFIG =
      io.grpc.MethodDescriptor.<protocol.Config.ReqSetConfig, protocol.Config.RespSetConfig>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "protocol.ConfServer", "SetConfig"))
          .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Config.ReqSetConfig.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Config.RespSetConfig.getDefaultInstance()))
          .build();

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static ConfServerStub newStub(io.grpc.Channel channel) {
    return new ConfServerStub(channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static ConfServerBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    return new ConfServerBlockingStub(channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static ConfServerFutureStub newFutureStub(
      io.grpc.Channel channel) {
    return new ConfServerFutureStub(channel);
  }

  /**
   * <pre>
   * api rpc 服务
   * </pre>
   */
  public static abstract class ConfServerImplBase implements io.grpc.BindableService {

    /**
     */
    public void getConfig(protocol.Config.ReqGetConfig request,
        io.grpc.stub.StreamObserver<protocol.Config.RespGetConfig> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_GET_CONFIG, responseObserver);
    }

    /**
     */
    public void setConfig(protocol.Config.ReqSetConfig request,
        io.grpc.stub.StreamObserver<protocol.Config.RespSetConfig> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_SET_CONFIG, responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            METHOD_GET_CONFIG,
            asyncUnaryCall(
              new MethodHandlers<
                protocol.Config.ReqGetConfig,
                protocol.Config.RespGetConfig>(
                  this, METHODID_GET_CONFIG)))
          .addMethod(
            METHOD_SET_CONFIG,
            asyncUnaryCall(
              new MethodHandlers<
                protocol.Config.ReqSetConfig,
                protocol.Config.RespSetConfig>(
                  this, METHODID_SET_CONFIG)))
          .build();
    }
  }

  /**
   * <pre>
   * api rpc 服务
   * </pre>
   */
  public static final class ConfServerStub extends io.grpc.stub.AbstractStub<ConfServerStub> {
    private ConfServerStub(io.grpc.Channel channel) {
      super(channel);
    }

    private ConfServerStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ConfServerStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new ConfServerStub(channel, callOptions);
    }

    /**
     */
    public void getConfig(protocol.Config.ReqGetConfig request,
        io.grpc.stub.StreamObserver<protocol.Config.RespGetConfig> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_GET_CONFIG, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void setConfig(protocol.Config.ReqSetConfig request,
        io.grpc.stub.StreamObserver<protocol.Config.RespSetConfig> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_SET_CONFIG, getCallOptions()), request, responseObserver);
    }
  }

  /**
   * <pre>
   * api rpc 服务
   * </pre>
   */
  public static final class ConfServerBlockingStub extends io.grpc.stub.AbstractStub<ConfServerBlockingStub> {
    private ConfServerBlockingStub(io.grpc.Channel channel) {
      super(channel);
    }

    private ConfServerBlockingStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ConfServerBlockingStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new ConfServerBlockingStub(channel, callOptions);
    }

    /**
     */
    public protocol.Config.RespGetConfig getConfig(protocol.Config.ReqGetConfig request) {
      return blockingUnaryCall(
          getChannel(), METHOD_GET_CONFIG, getCallOptions(), request);
    }

    /**
     */
    public protocol.Config.RespSetConfig setConfig(protocol.Config.ReqSetConfig request) {
      return blockingUnaryCall(
          getChannel(), METHOD_SET_CONFIG, getCallOptions(), request);
    }
  }

  /**
   * <pre>
   * api rpc 服务
   * </pre>
   */
  public static final class ConfServerFutureStub extends io.grpc.stub.AbstractStub<ConfServerFutureStub> {
    private ConfServerFutureStub(io.grpc.Channel channel) {
      super(channel);
    }

    private ConfServerFutureStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ConfServerFutureStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new ConfServerFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<protocol.Config.RespGetConfig> getConfig(
        protocol.Config.ReqGetConfig request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_GET_CONFIG, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<protocol.Config.RespSetConfig> setConfig(
        protocol.Config.ReqSetConfig request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_SET_CONFIG, getCallOptions()), request);
    }
  }

  private static final int METHODID_GET_CONFIG = 0;
  private static final int METHODID_SET_CONFIG = 1;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final ConfServerImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(ConfServerImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_GET_CONFIG:
          serviceImpl.getConfig((protocol.Config.ReqGetConfig) request,
              (io.grpc.stub.StreamObserver<protocol.Config.RespGetConfig>) responseObserver);
          break;
        case METHODID_SET_CONFIG:
          serviceImpl.setConfig((protocol.Config.ReqSetConfig) request,
              (io.grpc.stub.StreamObserver<protocol.Config.RespSetConfig>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (ConfServerGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .addMethod(METHOD_GET_CONFIG)
              .addMethod(METHOD_SET_CONFIG)
              .build();
        }
      }
    }
    return result;
  }
}
