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
    comments = "Source: api.proto")
public final class ApiServiceGrpc {

  private ApiServiceGrpc() {}

  public static final String SERVICE_NAME = "protocol.ApiService";

  // Static method descriptors that strictly reflect the proto.
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<protocol.Api.Req,
      protocol.Api.Resp> METHOD_SAY_HELLO =
      io.grpc.MethodDescriptor.<protocol.Api.Req, protocol.Api.Resp>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "protocol.ApiService", "SayHello"))
          .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.Req.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.Resp.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<protocol.Api.ReqRegister,
      protocol.Api.RespRegister> METHOD_REGISTER =
      io.grpc.MethodDescriptor.<protocol.Api.ReqRegister, protocol.Api.RespRegister>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "protocol.ApiService", "Register"))
          .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.ReqRegister.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.RespRegister.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<protocol.Api.ReqBand,
      protocol.Api.RespBand> METHOD_BIND =
      io.grpc.MethodDescriptor.<protocol.Api.ReqBand, protocol.Api.RespBand>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "protocol.ApiService", "Bind"))
          .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.ReqBand.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.RespBand.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<protocol.Api.ReqLogin,
      protocol.Api.RespLogin> METHOD_GET_BIND =
      io.grpc.MethodDescriptor.<protocol.Api.ReqLogin, protocol.Api.RespLogin>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "protocol.ApiService", "GetBind"))
          .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.ReqLogin.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.RespLogin.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<protocol.Api.ReqCheckAccount,
      protocol.Api.RespCheckAccount> METHOD_GET_ACCOUNT =
      io.grpc.MethodDescriptor.<protocol.Api.ReqCheckAccount, protocol.Api.RespCheckAccount>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "protocol.ApiService", "GetAccount"))
          .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.ReqCheckAccount.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.RespCheckAccount.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<protocol.Api.ReqGetEthBalance,
      protocol.Api.RespGetEthBalance> METHOD_GET_ETH_BALANCE =
      io.grpc.MethodDescriptor.<protocol.Api.ReqGetEthBalance, protocol.Api.RespGetEthBalance>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "protocol.ApiService", "GetEthBalance"))
          .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.ReqGetEthBalance.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.RespGetEthBalance.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<protocol.Api.ReqScheduling,
      protocol.Api.RespScheduling> METHOD_SET_SCHEDULE =
      io.grpc.MethodDescriptor.<protocol.Api.ReqScheduling, protocol.Api.RespScheduling>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "protocol.ApiService", "SetSchedule"))
          .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.ReqScheduling.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.RespScheduling.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<protocol.Api.ReqGetSchedue,
      protocol.Api.RespGetSchedue> METHOD_GET_SCHEDULE =
      io.grpc.MethodDescriptor.<protocol.Api.ReqGetSchedue, protocol.Api.RespGetSchedue>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "protocol.ApiService", "GetSchedule"))
          .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.ReqGetSchedue.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.RespGetSchedue.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<protocol.Api.ReqGetSchedue,
      protocol.Api.RespGetSchedue> METHOD_GET_CAN_JOIN =
      io.grpc.MethodDescriptor.<protocol.Api.ReqGetSchedue, protocol.Api.RespGetSchedue>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "protocol.ApiService", "GetCanJoin"))
          .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.ReqGetSchedue.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.RespGetSchedue.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<protocol.Api.ReqGetStaff,
      protocol.Api.RespGetStaff> METHOD_GET_APPLY =
      io.grpc.MethodDescriptor.<protocol.Api.ReqGetStaff, protocol.Api.RespGetStaff>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "protocol.ApiService", "GetApply"))
          .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.ReqGetStaff.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.RespGetStaff.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<protocol.Api.ReqGetCanApply,
      protocol.Api.RespGetCanApply> METHOD_GET_JOB =
      io.grpc.MethodDescriptor.<protocol.Api.ReqGetCanApply, protocol.Api.RespGetCanApply>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "protocol.ApiService", "GetJob"))
          .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.ReqGetCanApply.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.RespGetCanApply.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<protocol.Api.ReqFindJob,
      protocol.Api.RespFindJob> METHOD_APPLY_JOB =
      io.grpc.MethodDescriptor.<protocol.Api.ReqFindJob, protocol.Api.RespFindJob>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "protocol.ApiService", "ApplyJob"))
          .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.ReqFindJob.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.RespFindJob.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<protocol.Api.ReqHistoryJoin,
      protocol.Api.RespHistoryJoin> METHOD_HISTORY_JOIN =
      io.grpc.MethodDescriptor.<protocol.Api.ReqHistoryJoin, protocol.Api.RespHistoryJoin>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "protocol.ApiService", "HistoryJoin"))
          .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.ReqHistoryJoin.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.RespHistoryJoin.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<protocol.Api.ReqCheckIsOkApplication,
      protocol.Api.RespCheckIsOkApplication> METHOD_CHECK_IS_OK_APPLICATION =
      io.grpc.MethodDescriptor.<protocol.Api.ReqCheckIsOkApplication, protocol.Api.RespCheckIsOkApplication>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "protocol.ApiService", "CheckIsOkApplication"))
          .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.ReqCheckIsOkApplication.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.RespCheckIsOkApplication.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<protocol.Api.ReqPay,
      protocol.Api.RespPay> METHOD_PAY =
      io.grpc.MethodDescriptor.<protocol.Api.ReqPay, protocol.Api.RespPay>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "protocol.ApiService", "Pay"))
          .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.ReqPay.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.RespPay.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<protocol.Api.ReqGetAllOrder,
      protocol.Api.RespGetAllOrder> METHOD_GET_ALL_ORDER_BY_SCHEDULE =
      io.grpc.MethodDescriptor.<protocol.Api.ReqGetAllOrder, protocol.Api.RespGetAllOrder>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "protocol.ApiService", "GetAllOrderBySchedule"))
          .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.ReqGetAllOrder.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.RespGetAllOrder.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<protocol.Api.ReqGetContent,
      protocol.Api.RespGetContent> METHOD_GET_CONTENT =
      io.grpc.MethodDescriptor.<protocol.Api.ReqGetContent, protocol.Api.RespGetContent>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "protocol.ApiService", "GetContent"))
          .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.ReqGetContent.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.RespGetContent.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<protocol.Api.ReqGetAllIncome,
      protocol.Api.RespGetAllIncome> METHOD_GET_ALL_INCOME =
      io.grpc.MethodDescriptor.<protocol.Api.ReqGetAllIncome, protocol.Api.RespGetAllIncome>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "protocol.ApiService", "GetAllIncome"))
          .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.ReqGetAllIncome.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.RespGetAllIncome.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<protocol.Api.ReqGetIncomeBySchedule,
      protocol.Api.RespGetIncomeBySchedule> METHOD_GET_INCOME_BY_SCHEDULE =
      io.grpc.MethodDescriptor.<protocol.Api.ReqGetIncomeBySchedule, protocol.Api.RespGetIncomeBySchedule>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "protocol.ApiService", "GetIncomeBySchedule"))
          .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.ReqGetIncomeBySchedule.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.RespGetIncomeBySchedule.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<protocol.Api.ReqGetBalance,
      protocol.Api.RespGetBalance> METHOD_GET_BALANCE =
      io.grpc.MethodDescriptor.<protocol.Api.ReqGetBalance, protocol.Api.RespGetBalance>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "protocol.ApiService", "GetBalance"))
          .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.ReqGetBalance.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.RespGetBalance.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<protocol.Api.ReqGetAllMoney,
      protocol.Api.RespGetAllMoney> METHOD_GET_ALL_MONEY =
      io.grpc.MethodDescriptor.<protocol.Api.ReqGetAllMoney, protocol.Api.RespGetAllMoney>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "protocol.ApiService", "GetAllMoney"))
          .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.ReqGetAllMoney.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.RespGetAllMoney.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<protocol.Api.ReqGetNowJobAddr,
      protocol.Api.RespGetNowJobAddr> METHOD_GET_NOW_JOB_ADDRESS =
      io.grpc.MethodDescriptor.<protocol.Api.ReqGetNowJobAddr, protocol.Api.RespGetNowJobAddr>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "protocol.ApiService", "GetNowJobAddress"))
          .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.ReqGetNowJobAddr.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.RespGetNowJobAddr.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<protocol.Api.ReqAddWhite,
      protocol.Api.RespAddWhite> METHOD_ADD_WHITE_LIST =
      io.grpc.MethodDescriptor.<protocol.Api.ReqAddWhite, protocol.Api.RespAddWhite>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "protocol.ApiService", "AddWhiteList"))
          .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.ReqAddWhite.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.RespAddWhite.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<protocol.Api.ReqGetWhite,
      protocol.Api.RespGetWhite> METHOD_GET_WHITE_LIST =
      io.grpc.MethodDescriptor.<protocol.Api.ReqGetWhite, protocol.Api.RespGetWhite>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "protocol.ApiService", "GetWhiteList"))
          .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.ReqGetWhite.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.RespGetWhite.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<protocol.Api.ReqDelWhite,
      protocol.Api.RespDelWhite> METHOD_DEL_WHITE_LIST =
      io.grpc.MethodDescriptor.<protocol.Api.ReqDelWhite, protocol.Api.RespDelWhite>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "protocol.ApiService", "DelWhiteList"))
          .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.ReqDelWhite.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.RespDelWhite.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<protocol.Api.ReqThreeSetOrder,
      protocol.Api.RespThreeSetOrder> METHOD_THREE_SET_ORDER =
      io.grpc.MethodDescriptor.<protocol.Api.ReqThreeSetOrder, protocol.Api.RespThreeSetOrder>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "protocol.ApiService", "ThreeSetOrder"))
          .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.ReqThreeSetOrder.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.RespThreeSetOrder.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<protocol.Api.ReqThreeSetBill,
      protocol.Api.RespThreeSetBill> METHOD_THREE_SET_BILL =
      io.grpc.MethodDescriptor.<protocol.Api.ReqThreeSetBill, protocol.Api.RespThreeSetBill>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "protocol.ApiService", "ThreeSetBill"))
          .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.ReqThreeSetBill.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.RespThreeSetBill.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<protocol.Api.ReqManageContract,
      protocol.Api.RespManageContract> METHOD_MANAGE_CONTRACT =
      io.grpc.MethodDescriptor.<protocol.Api.ReqManageContract, protocol.Api.RespManageContract>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "protocol.ApiService", "ManageContract"))
          .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.ReqManageContract.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.RespManageContract.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<protocol.Api.ReqCheckContract,
      protocol.Api.RespCheckContract> METHOD_CHECK_CONTRACT =
      io.grpc.MethodDescriptor.<protocol.Api.ReqCheckContract, protocol.Api.RespCheckContract>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "protocol.ApiService", "CheckContract"))
          .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.ReqCheckContract.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.RespCheckContract.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<protocol.Api.CReloadConfig,
      protocol.Api.SReloadConfig> METHOD_RELOAD_CONFIG =
      io.grpc.MethodDescriptor.<protocol.Api.CReloadConfig, protocol.Api.SReloadConfig>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "protocol.ApiService", "ReloadConfig"))
          .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.CReloadConfig.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.SReloadConfig.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<protocol.Api.CReloadDeploy,
      protocol.Api.SReloadDeploy> METHOD_RELOAD_DEPLOY =
      io.grpc.MethodDescriptor.<protocol.Api.CReloadDeploy, protocol.Api.SReloadDeploy>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "protocol.ApiService", "ReloadDeploy"))
          .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.CReloadDeploy.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
              protocol.Api.SReloadDeploy.getDefaultInstance()))
          .build();

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static ApiServiceStub newStub(io.grpc.Channel channel) {
    return new ApiServiceStub(channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static ApiServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    return new ApiServiceBlockingStub(channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static ApiServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    return new ApiServiceFutureStub(channel);
  }

  /**
   * <pre>
   * api rpc 服务
   * </pre>
   */
  public static abstract class ApiServiceImplBase implements io.grpc.BindableService {

    /**
     */
    public void sayHello(protocol.Api.Req request,
        io.grpc.stub.StreamObserver<protocol.Api.Resp> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_SAY_HELLO, responseObserver);
    }

    /**
     * <pre>
     * 注册 绑定 登录
     * </pre>
     */
    public void register(protocol.Api.ReqRegister request,
        io.grpc.stub.StreamObserver<protocol.Api.RespRegister> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_REGISTER, responseObserver);
    }

    /**
     */
    public void bind(protocol.Api.ReqBand request,
        io.grpc.stub.StreamObserver<protocol.Api.RespBand> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_BIND, responseObserver);
    }

    /**
     */
    public void getBind(protocol.Api.ReqLogin request,
        io.grpc.stub.StreamObserver<protocol.Api.RespLogin> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_GET_BIND, responseObserver);
    }

    /**
     */
    public void getAccount(protocol.Api.ReqCheckAccount request,
        io.grpc.stub.StreamObserver<protocol.Api.RespCheckAccount> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_GET_ACCOUNT, responseObserver);
    }

    /**
     */
    public void getEthBalance(protocol.Api.ReqGetEthBalance request,
        io.grpc.stub.StreamObserver<protocol.Api.RespGetEthBalance> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_GET_ETH_BALANCE, responseObserver);
    }

    /**
     * <pre>
     * 排班 申请工作
     * </pre>
     */
    public void setSchedule(protocol.Api.ReqScheduling request,
        io.grpc.stub.StreamObserver<protocol.Api.RespScheduling> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_SET_SCHEDULE, responseObserver);
    }

    /**
     */
    public void getSchedule(protocol.Api.ReqGetSchedue request,
        io.grpc.stub.StreamObserver<protocol.Api.RespGetSchedue> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_GET_SCHEDULE, responseObserver);
    }

    /**
     */
    public void getCanJoin(protocol.Api.ReqGetSchedue request,
        io.grpc.stub.StreamObserver<protocol.Api.RespGetSchedue> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_GET_CAN_JOIN, responseObserver);
    }

    /**
     */
    public void getApply(protocol.Api.ReqGetStaff request,
        io.grpc.stub.StreamObserver<protocol.Api.RespGetStaff> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_GET_APPLY, responseObserver);
    }

    /**
     */
    public void getJob(protocol.Api.ReqGetCanApply request,
        io.grpc.stub.StreamObserver<protocol.Api.RespGetCanApply> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_GET_JOB, responseObserver);
    }

    /**
     */
    public void applyJob(protocol.Api.ReqFindJob request,
        io.grpc.stub.StreamObserver<protocol.Api.RespFindJob> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_APPLY_JOB, responseObserver);
    }

    /**
     */
    public void historyJoin(protocol.Api.ReqHistoryJoin request,
        io.grpc.stub.StreamObserver<protocol.Api.RespHistoryJoin> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_HISTORY_JOIN, responseObserver);
    }

    /**
     */
    public void checkIsOkApplication(protocol.Api.ReqCheckIsOkApplication request,
        io.grpc.stub.StreamObserver<protocol.Api.RespCheckIsOkApplication> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_CHECK_IS_OK_APPLICATION, responseObserver);
    }

    /**
     * <pre>
     * 收入
     * </pre>
     */
    public void pay(protocol.Api.ReqPay request,
        io.grpc.stub.StreamObserver<protocol.Api.RespPay> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_PAY, responseObserver);
    }

    /**
     */
    public void getAllOrderBySchedule(protocol.Api.ReqGetAllOrder request,
        io.grpc.stub.StreamObserver<protocol.Api.RespGetAllOrder> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_GET_ALL_ORDER_BY_SCHEDULE, responseObserver);
    }

    /**
     */
    public void getContent(protocol.Api.ReqGetContent request,
        io.grpc.stub.StreamObserver<protocol.Api.RespGetContent> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_GET_CONTENT, responseObserver);
    }

    /**
     */
    public void getAllIncome(protocol.Api.ReqGetAllIncome request,
        io.grpc.stub.StreamObserver<protocol.Api.RespGetAllIncome> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_GET_ALL_INCOME, responseObserver);
    }

    /**
     * <pre>
     * 3.5.根据排班信息获取用户收入详情
     * </pre>
     */
    public void getIncomeBySchedule(protocol.Api.ReqGetIncomeBySchedule request,
        io.grpc.stub.StreamObserver<protocol.Api.RespGetIncomeBySchedule> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_GET_INCOME_BY_SCHEDULE, responseObserver);
    }

    /**
     * <pre>
     * 查询
     * </pre>
     */
    public void getBalance(protocol.Api.ReqGetBalance request,
        io.grpc.stub.StreamObserver<protocol.Api.RespGetBalance> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_GET_BALANCE, responseObserver);
    }

    /**
     */
    public void getAllMoney(protocol.Api.ReqGetAllMoney request,
        io.grpc.stub.StreamObserver<protocol.Api.RespGetAllMoney> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_GET_ALL_MONEY, responseObserver);
    }

    /**
     */
    public void getNowJobAddress(protocol.Api.ReqGetNowJobAddr request,
        io.grpc.stub.StreamObserver<protocol.Api.RespGetNowJobAddr> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_GET_NOW_JOB_ADDRESS, responseObserver);
    }

    /**
     * <pre>
     * 用人白名单
     * </pre>
     */
    public void addWhiteList(protocol.Api.ReqAddWhite request,
        io.grpc.stub.StreamObserver<protocol.Api.RespAddWhite> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_ADD_WHITE_LIST, responseObserver);
    }

    /**
     */
    public void getWhiteList(protocol.Api.ReqGetWhite request,
        io.grpc.stub.StreamObserver<protocol.Api.RespGetWhite> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_GET_WHITE_LIST, responseObserver);
    }

    /**
     */
    public void delWhiteList(protocol.Api.ReqDelWhite request,
        io.grpc.stub.StreamObserver<protocol.Api.RespDelWhite> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_DEL_WHITE_LIST, responseObserver);
    }

    /**
     * <pre>
     * 三方
     * </pre>
     */
    public void threeSetOrder(protocol.Api.ReqThreeSetOrder request,
        io.grpc.stub.StreamObserver<protocol.Api.RespThreeSetOrder> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_THREE_SET_ORDER, responseObserver);
    }

    /**
     */
    public void threeSetBill(protocol.Api.ReqThreeSetBill request,
        io.grpc.stub.StreamObserver<protocol.Api.RespThreeSetBill> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_THREE_SET_BILL, responseObserver);
    }

    /**
     * <pre>
     * 无用
     * </pre>
     */
    public void manageContract(protocol.Api.ReqManageContract request,
        io.grpc.stub.StreamObserver<protocol.Api.RespManageContract> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_MANAGE_CONTRACT, responseObserver);
    }

    /**
     */
    public void checkContract(protocol.Api.ReqCheckContract request,
        io.grpc.stub.StreamObserver<protocol.Api.RespCheckContract> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_CHECK_CONTRACT, responseObserver);
    }

    /**
     * <pre>
     * 配置服务
     * </pre>
     */
    public void reloadConfig(protocol.Api.CReloadConfig request,
        io.grpc.stub.StreamObserver<protocol.Api.SReloadConfig> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_RELOAD_CONFIG, responseObserver);
    }

    /**
     */
    public void reloadDeploy(protocol.Api.CReloadDeploy request,
        io.grpc.stub.StreamObserver<protocol.Api.SReloadDeploy> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_RELOAD_DEPLOY, responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            METHOD_SAY_HELLO,
            asyncUnaryCall(
              new MethodHandlers<
                protocol.Api.Req,
                protocol.Api.Resp>(
                  this, METHODID_SAY_HELLO)))
          .addMethod(
            METHOD_REGISTER,
            asyncUnaryCall(
              new MethodHandlers<
                protocol.Api.ReqRegister,
                protocol.Api.RespRegister>(
                  this, METHODID_REGISTER)))
          .addMethod(
            METHOD_BIND,
            asyncUnaryCall(
              new MethodHandlers<
                protocol.Api.ReqBand,
                protocol.Api.RespBand>(
                  this, METHODID_BIND)))
          .addMethod(
            METHOD_GET_BIND,
            asyncUnaryCall(
              new MethodHandlers<
                protocol.Api.ReqLogin,
                protocol.Api.RespLogin>(
                  this, METHODID_GET_BIND)))
          .addMethod(
            METHOD_GET_ACCOUNT,
            asyncUnaryCall(
              new MethodHandlers<
                protocol.Api.ReqCheckAccount,
                protocol.Api.RespCheckAccount>(
                  this, METHODID_GET_ACCOUNT)))
          .addMethod(
            METHOD_GET_ETH_BALANCE,
            asyncUnaryCall(
              new MethodHandlers<
                protocol.Api.ReqGetEthBalance,
                protocol.Api.RespGetEthBalance>(
                  this, METHODID_GET_ETH_BALANCE)))
          .addMethod(
            METHOD_SET_SCHEDULE,
            asyncUnaryCall(
              new MethodHandlers<
                protocol.Api.ReqScheduling,
                protocol.Api.RespScheduling>(
                  this, METHODID_SET_SCHEDULE)))
          .addMethod(
            METHOD_GET_SCHEDULE,
            asyncUnaryCall(
              new MethodHandlers<
                protocol.Api.ReqGetSchedue,
                protocol.Api.RespGetSchedue>(
                  this, METHODID_GET_SCHEDULE)))
          .addMethod(
            METHOD_GET_CAN_JOIN,
            asyncUnaryCall(
              new MethodHandlers<
                protocol.Api.ReqGetSchedue,
                protocol.Api.RespGetSchedue>(
                  this, METHODID_GET_CAN_JOIN)))
          .addMethod(
            METHOD_GET_APPLY,
            asyncUnaryCall(
              new MethodHandlers<
                protocol.Api.ReqGetStaff,
                protocol.Api.RespGetStaff>(
                  this, METHODID_GET_APPLY)))
          .addMethod(
            METHOD_GET_JOB,
            asyncUnaryCall(
              new MethodHandlers<
                protocol.Api.ReqGetCanApply,
                protocol.Api.RespGetCanApply>(
                  this, METHODID_GET_JOB)))
          .addMethod(
            METHOD_APPLY_JOB,
            asyncUnaryCall(
              new MethodHandlers<
                protocol.Api.ReqFindJob,
                protocol.Api.RespFindJob>(
                  this, METHODID_APPLY_JOB)))
          .addMethod(
            METHOD_HISTORY_JOIN,
            asyncUnaryCall(
              new MethodHandlers<
                protocol.Api.ReqHistoryJoin,
                protocol.Api.RespHistoryJoin>(
                  this, METHODID_HISTORY_JOIN)))
          .addMethod(
            METHOD_CHECK_IS_OK_APPLICATION,
            asyncUnaryCall(
              new MethodHandlers<
                protocol.Api.ReqCheckIsOkApplication,
                protocol.Api.RespCheckIsOkApplication>(
                  this, METHODID_CHECK_IS_OK_APPLICATION)))
          .addMethod(
            METHOD_PAY,
            asyncUnaryCall(
              new MethodHandlers<
                protocol.Api.ReqPay,
                protocol.Api.RespPay>(
                  this, METHODID_PAY)))
          .addMethod(
            METHOD_GET_ALL_ORDER_BY_SCHEDULE,
            asyncUnaryCall(
              new MethodHandlers<
                protocol.Api.ReqGetAllOrder,
                protocol.Api.RespGetAllOrder>(
                  this, METHODID_GET_ALL_ORDER_BY_SCHEDULE)))
          .addMethod(
            METHOD_GET_CONTENT,
            asyncUnaryCall(
              new MethodHandlers<
                protocol.Api.ReqGetContent,
                protocol.Api.RespGetContent>(
                  this, METHODID_GET_CONTENT)))
          .addMethod(
            METHOD_GET_ALL_INCOME,
            asyncUnaryCall(
              new MethodHandlers<
                protocol.Api.ReqGetAllIncome,
                protocol.Api.RespGetAllIncome>(
                  this, METHODID_GET_ALL_INCOME)))
          .addMethod(
            METHOD_GET_INCOME_BY_SCHEDULE,
            asyncUnaryCall(
              new MethodHandlers<
                protocol.Api.ReqGetIncomeBySchedule,
                protocol.Api.RespGetIncomeBySchedule>(
                  this, METHODID_GET_INCOME_BY_SCHEDULE)))
          .addMethod(
            METHOD_GET_BALANCE,
            asyncUnaryCall(
              new MethodHandlers<
                protocol.Api.ReqGetBalance,
                protocol.Api.RespGetBalance>(
                  this, METHODID_GET_BALANCE)))
          .addMethod(
            METHOD_GET_ALL_MONEY,
            asyncUnaryCall(
              new MethodHandlers<
                protocol.Api.ReqGetAllMoney,
                protocol.Api.RespGetAllMoney>(
                  this, METHODID_GET_ALL_MONEY)))
          .addMethod(
            METHOD_GET_NOW_JOB_ADDRESS,
            asyncUnaryCall(
              new MethodHandlers<
                protocol.Api.ReqGetNowJobAddr,
                protocol.Api.RespGetNowJobAddr>(
                  this, METHODID_GET_NOW_JOB_ADDRESS)))
          .addMethod(
            METHOD_ADD_WHITE_LIST,
            asyncUnaryCall(
              new MethodHandlers<
                protocol.Api.ReqAddWhite,
                protocol.Api.RespAddWhite>(
                  this, METHODID_ADD_WHITE_LIST)))
          .addMethod(
            METHOD_GET_WHITE_LIST,
            asyncUnaryCall(
              new MethodHandlers<
                protocol.Api.ReqGetWhite,
                protocol.Api.RespGetWhite>(
                  this, METHODID_GET_WHITE_LIST)))
          .addMethod(
            METHOD_DEL_WHITE_LIST,
            asyncUnaryCall(
              new MethodHandlers<
                protocol.Api.ReqDelWhite,
                protocol.Api.RespDelWhite>(
                  this, METHODID_DEL_WHITE_LIST)))
          .addMethod(
            METHOD_THREE_SET_ORDER,
            asyncUnaryCall(
              new MethodHandlers<
                protocol.Api.ReqThreeSetOrder,
                protocol.Api.RespThreeSetOrder>(
                  this, METHODID_THREE_SET_ORDER)))
          .addMethod(
            METHOD_THREE_SET_BILL,
            asyncUnaryCall(
              new MethodHandlers<
                protocol.Api.ReqThreeSetBill,
                protocol.Api.RespThreeSetBill>(
                  this, METHODID_THREE_SET_BILL)))
          .addMethod(
            METHOD_MANAGE_CONTRACT,
            asyncUnaryCall(
              new MethodHandlers<
                protocol.Api.ReqManageContract,
                protocol.Api.RespManageContract>(
                  this, METHODID_MANAGE_CONTRACT)))
          .addMethod(
            METHOD_CHECK_CONTRACT,
            asyncUnaryCall(
              new MethodHandlers<
                protocol.Api.ReqCheckContract,
                protocol.Api.RespCheckContract>(
                  this, METHODID_CHECK_CONTRACT)))
          .addMethod(
            METHOD_RELOAD_CONFIG,
            asyncUnaryCall(
              new MethodHandlers<
                protocol.Api.CReloadConfig,
                protocol.Api.SReloadConfig>(
                  this, METHODID_RELOAD_CONFIG)))
          .addMethod(
            METHOD_RELOAD_DEPLOY,
            asyncUnaryCall(
              new MethodHandlers<
                protocol.Api.CReloadDeploy,
                protocol.Api.SReloadDeploy>(
                  this, METHODID_RELOAD_DEPLOY)))
          .build();
    }
  }

  /**
   * <pre>
   * api rpc 服务
   * </pre>
   */
  public static final class ApiServiceStub extends io.grpc.stub.AbstractStub<ApiServiceStub> {
    private ApiServiceStub(io.grpc.Channel channel) {
      super(channel);
    }

    private ApiServiceStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ApiServiceStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new ApiServiceStub(channel, callOptions);
    }

    /**
     */
    public void sayHello(protocol.Api.Req request,
        io.grpc.stub.StreamObserver<protocol.Api.Resp> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_SAY_HELLO, getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 注册 绑定 登录
     * </pre>
     */
    public void register(protocol.Api.ReqRegister request,
        io.grpc.stub.StreamObserver<protocol.Api.RespRegister> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_REGISTER, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void bind(protocol.Api.ReqBand request,
        io.grpc.stub.StreamObserver<protocol.Api.RespBand> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_BIND, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getBind(protocol.Api.ReqLogin request,
        io.grpc.stub.StreamObserver<protocol.Api.RespLogin> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_GET_BIND, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getAccount(protocol.Api.ReqCheckAccount request,
        io.grpc.stub.StreamObserver<protocol.Api.RespCheckAccount> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_GET_ACCOUNT, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getEthBalance(protocol.Api.ReqGetEthBalance request,
        io.grpc.stub.StreamObserver<protocol.Api.RespGetEthBalance> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_GET_ETH_BALANCE, getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 排班 申请工作
     * </pre>
     */
    public void setSchedule(protocol.Api.ReqScheduling request,
        io.grpc.stub.StreamObserver<protocol.Api.RespScheduling> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_SET_SCHEDULE, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getSchedule(protocol.Api.ReqGetSchedue request,
        io.grpc.stub.StreamObserver<protocol.Api.RespGetSchedue> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_GET_SCHEDULE, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getCanJoin(protocol.Api.ReqGetSchedue request,
        io.grpc.stub.StreamObserver<protocol.Api.RespGetSchedue> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_GET_CAN_JOIN, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getApply(protocol.Api.ReqGetStaff request,
        io.grpc.stub.StreamObserver<protocol.Api.RespGetStaff> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_GET_APPLY, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getJob(protocol.Api.ReqGetCanApply request,
        io.grpc.stub.StreamObserver<protocol.Api.RespGetCanApply> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_GET_JOB, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void applyJob(protocol.Api.ReqFindJob request,
        io.grpc.stub.StreamObserver<protocol.Api.RespFindJob> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_APPLY_JOB, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void historyJoin(protocol.Api.ReqHistoryJoin request,
        io.grpc.stub.StreamObserver<protocol.Api.RespHistoryJoin> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_HISTORY_JOIN, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void checkIsOkApplication(protocol.Api.ReqCheckIsOkApplication request,
        io.grpc.stub.StreamObserver<protocol.Api.RespCheckIsOkApplication> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_CHECK_IS_OK_APPLICATION, getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 收入
     * </pre>
     */
    public void pay(protocol.Api.ReqPay request,
        io.grpc.stub.StreamObserver<protocol.Api.RespPay> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_PAY, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getAllOrderBySchedule(protocol.Api.ReqGetAllOrder request,
        io.grpc.stub.StreamObserver<protocol.Api.RespGetAllOrder> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_GET_ALL_ORDER_BY_SCHEDULE, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getContent(protocol.Api.ReqGetContent request,
        io.grpc.stub.StreamObserver<protocol.Api.RespGetContent> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_GET_CONTENT, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getAllIncome(protocol.Api.ReqGetAllIncome request,
        io.grpc.stub.StreamObserver<protocol.Api.RespGetAllIncome> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_GET_ALL_INCOME, getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 3.5.根据排班信息获取用户收入详情
     * </pre>
     */
    public void getIncomeBySchedule(protocol.Api.ReqGetIncomeBySchedule request,
        io.grpc.stub.StreamObserver<protocol.Api.RespGetIncomeBySchedule> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_GET_INCOME_BY_SCHEDULE, getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 查询
     * </pre>
     */
    public void getBalance(protocol.Api.ReqGetBalance request,
        io.grpc.stub.StreamObserver<protocol.Api.RespGetBalance> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_GET_BALANCE, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getAllMoney(protocol.Api.ReqGetAllMoney request,
        io.grpc.stub.StreamObserver<protocol.Api.RespGetAllMoney> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_GET_ALL_MONEY, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getNowJobAddress(protocol.Api.ReqGetNowJobAddr request,
        io.grpc.stub.StreamObserver<protocol.Api.RespGetNowJobAddr> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_GET_NOW_JOB_ADDRESS, getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 用人白名单
     * </pre>
     */
    public void addWhiteList(protocol.Api.ReqAddWhite request,
        io.grpc.stub.StreamObserver<protocol.Api.RespAddWhite> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_ADD_WHITE_LIST, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getWhiteList(protocol.Api.ReqGetWhite request,
        io.grpc.stub.StreamObserver<protocol.Api.RespGetWhite> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_GET_WHITE_LIST, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void delWhiteList(protocol.Api.ReqDelWhite request,
        io.grpc.stub.StreamObserver<protocol.Api.RespDelWhite> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_DEL_WHITE_LIST, getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 三方
     * </pre>
     */
    public void threeSetOrder(protocol.Api.ReqThreeSetOrder request,
        io.grpc.stub.StreamObserver<protocol.Api.RespThreeSetOrder> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_THREE_SET_ORDER, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void threeSetBill(protocol.Api.ReqThreeSetBill request,
        io.grpc.stub.StreamObserver<protocol.Api.RespThreeSetBill> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_THREE_SET_BILL, getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 无用
     * </pre>
     */
    public void manageContract(protocol.Api.ReqManageContract request,
        io.grpc.stub.StreamObserver<protocol.Api.RespManageContract> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_MANAGE_CONTRACT, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void checkContract(protocol.Api.ReqCheckContract request,
        io.grpc.stub.StreamObserver<protocol.Api.RespCheckContract> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_CHECK_CONTRACT, getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 配置服务
     * </pre>
     */
    public void reloadConfig(protocol.Api.CReloadConfig request,
        io.grpc.stub.StreamObserver<protocol.Api.SReloadConfig> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_RELOAD_CONFIG, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void reloadDeploy(protocol.Api.CReloadDeploy request,
        io.grpc.stub.StreamObserver<protocol.Api.SReloadDeploy> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_RELOAD_DEPLOY, getCallOptions()), request, responseObserver);
    }
  }

  /**
   * <pre>
   * api rpc 服务
   * </pre>
   */
  public static final class ApiServiceBlockingStub extends io.grpc.stub.AbstractStub<ApiServiceBlockingStub> {
    private ApiServiceBlockingStub(io.grpc.Channel channel) {
      super(channel);
    }

    private ApiServiceBlockingStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ApiServiceBlockingStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new ApiServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public protocol.Api.Resp sayHello(protocol.Api.Req request) {
      return blockingUnaryCall(
          getChannel(), METHOD_SAY_HELLO, getCallOptions(), request);
    }

    /**
     * <pre>
     * 注册 绑定 登录
     * </pre>
     */
    public protocol.Api.RespRegister register(protocol.Api.ReqRegister request) {
      return blockingUnaryCall(
          getChannel(), METHOD_REGISTER, getCallOptions(), request);
    }

    /**
     */
    public protocol.Api.RespBand bind(protocol.Api.ReqBand request) {
      return blockingUnaryCall(
          getChannel(), METHOD_BIND, getCallOptions(), request);
    }

    /**
     */
    public protocol.Api.RespLogin getBind(protocol.Api.ReqLogin request) {
      return blockingUnaryCall(
          getChannel(), METHOD_GET_BIND, getCallOptions(), request);
    }

    /**
     */
    public protocol.Api.RespCheckAccount getAccount(protocol.Api.ReqCheckAccount request) {
      return blockingUnaryCall(
          getChannel(), METHOD_GET_ACCOUNT, getCallOptions(), request);
    }

    /**
     */
    public protocol.Api.RespGetEthBalance getEthBalance(protocol.Api.ReqGetEthBalance request) {
      return blockingUnaryCall(
          getChannel(), METHOD_GET_ETH_BALANCE, getCallOptions(), request);
    }

    /**
     * <pre>
     * 排班 申请工作
     * </pre>
     */
    public protocol.Api.RespScheduling setSchedule(protocol.Api.ReqScheduling request) {
      return blockingUnaryCall(
          getChannel(), METHOD_SET_SCHEDULE, getCallOptions(), request);
    }

    /**
     */
    public protocol.Api.RespGetSchedue getSchedule(protocol.Api.ReqGetSchedue request) {
      return blockingUnaryCall(
          getChannel(), METHOD_GET_SCHEDULE, getCallOptions(), request);
    }

    /**
     */
    public protocol.Api.RespGetSchedue getCanJoin(protocol.Api.ReqGetSchedue request) {
      return blockingUnaryCall(
          getChannel(), METHOD_GET_CAN_JOIN, getCallOptions(), request);
    }

    /**
     */
    public protocol.Api.RespGetStaff getApply(protocol.Api.ReqGetStaff request) {
      return blockingUnaryCall(
          getChannel(), METHOD_GET_APPLY, getCallOptions(), request);
    }

    /**
     */
    public protocol.Api.RespGetCanApply getJob(protocol.Api.ReqGetCanApply request) {
      return blockingUnaryCall(
          getChannel(), METHOD_GET_JOB, getCallOptions(), request);
    }

    /**
     */
    public protocol.Api.RespFindJob applyJob(protocol.Api.ReqFindJob request) {
      return blockingUnaryCall(
          getChannel(), METHOD_APPLY_JOB, getCallOptions(), request);
    }

    /**
     */
    public protocol.Api.RespHistoryJoin historyJoin(protocol.Api.ReqHistoryJoin request) {
      return blockingUnaryCall(
          getChannel(), METHOD_HISTORY_JOIN, getCallOptions(), request);
    }

    /**
     */
    public protocol.Api.RespCheckIsOkApplication checkIsOkApplication(protocol.Api.ReqCheckIsOkApplication request) {
      return blockingUnaryCall(
          getChannel(), METHOD_CHECK_IS_OK_APPLICATION, getCallOptions(), request);
    }

    /**
     * <pre>
     * 收入
     * </pre>
     */
    public protocol.Api.RespPay pay(protocol.Api.ReqPay request) {
      return blockingUnaryCall(
          getChannel(), METHOD_PAY, getCallOptions(), request);
    }

    /**
     */
    public protocol.Api.RespGetAllOrder getAllOrderBySchedule(protocol.Api.ReqGetAllOrder request) {
      return blockingUnaryCall(
          getChannel(), METHOD_GET_ALL_ORDER_BY_SCHEDULE, getCallOptions(), request);
    }

    /**
     */
    public protocol.Api.RespGetContent getContent(protocol.Api.ReqGetContent request) {
      return blockingUnaryCall(
          getChannel(), METHOD_GET_CONTENT, getCallOptions(), request);
    }

    /**
     */
    public protocol.Api.RespGetAllIncome getAllIncome(protocol.Api.ReqGetAllIncome request) {
      return blockingUnaryCall(
          getChannel(), METHOD_GET_ALL_INCOME, getCallOptions(), request);
    }

    /**
     * <pre>
     * 3.5.根据排班信息获取用户收入详情
     * </pre>
     */
    public protocol.Api.RespGetIncomeBySchedule getIncomeBySchedule(protocol.Api.ReqGetIncomeBySchedule request) {
      return blockingUnaryCall(
          getChannel(), METHOD_GET_INCOME_BY_SCHEDULE, getCallOptions(), request);
    }

    /**
     * <pre>
     * 查询
     * </pre>
     */
    public protocol.Api.RespGetBalance getBalance(protocol.Api.ReqGetBalance request) {
      return blockingUnaryCall(
          getChannel(), METHOD_GET_BALANCE, getCallOptions(), request);
    }

    /**
     */
    public protocol.Api.RespGetAllMoney getAllMoney(protocol.Api.ReqGetAllMoney request) {
      return blockingUnaryCall(
          getChannel(), METHOD_GET_ALL_MONEY, getCallOptions(), request);
    }

    /**
     */
    public protocol.Api.RespGetNowJobAddr getNowJobAddress(protocol.Api.ReqGetNowJobAddr request) {
      return blockingUnaryCall(
          getChannel(), METHOD_GET_NOW_JOB_ADDRESS, getCallOptions(), request);
    }

    /**
     * <pre>
     * 用人白名单
     * </pre>
     */
    public protocol.Api.RespAddWhite addWhiteList(protocol.Api.ReqAddWhite request) {
      return blockingUnaryCall(
          getChannel(), METHOD_ADD_WHITE_LIST, getCallOptions(), request);
    }

    /**
     */
    public protocol.Api.RespGetWhite getWhiteList(protocol.Api.ReqGetWhite request) {
      return blockingUnaryCall(
          getChannel(), METHOD_GET_WHITE_LIST, getCallOptions(), request);
    }

    /**
     */
    public protocol.Api.RespDelWhite delWhiteList(protocol.Api.ReqDelWhite request) {
      return blockingUnaryCall(
          getChannel(), METHOD_DEL_WHITE_LIST, getCallOptions(), request);
    }

    /**
     * <pre>
     * 三方
     * </pre>
     */
    public protocol.Api.RespThreeSetOrder threeSetOrder(protocol.Api.ReqThreeSetOrder request) {
      return blockingUnaryCall(
          getChannel(), METHOD_THREE_SET_ORDER, getCallOptions(), request);
    }

    /**
     */
    public protocol.Api.RespThreeSetBill threeSetBill(protocol.Api.ReqThreeSetBill request) {
      return blockingUnaryCall(
          getChannel(), METHOD_THREE_SET_BILL, getCallOptions(), request);
    }

    /**
     * <pre>
     * 无用
     * </pre>
     */
    public protocol.Api.RespManageContract manageContract(protocol.Api.ReqManageContract request) {
      return blockingUnaryCall(
          getChannel(), METHOD_MANAGE_CONTRACT, getCallOptions(), request);
    }

    /**
     */
    public protocol.Api.RespCheckContract checkContract(protocol.Api.ReqCheckContract request) {
      return blockingUnaryCall(
          getChannel(), METHOD_CHECK_CONTRACT, getCallOptions(), request);
    }

    /**
     * <pre>
     * 配置服务
     * </pre>
     */
    public protocol.Api.SReloadConfig reloadConfig(protocol.Api.CReloadConfig request) {
      return blockingUnaryCall(
          getChannel(), METHOD_RELOAD_CONFIG, getCallOptions(), request);
    }

    /**
     */
    public protocol.Api.SReloadDeploy reloadDeploy(protocol.Api.CReloadDeploy request) {
      return blockingUnaryCall(
          getChannel(), METHOD_RELOAD_DEPLOY, getCallOptions(), request);
    }
  }

  /**
   * <pre>
   * api rpc 服务
   * </pre>
   */
  public static final class ApiServiceFutureStub extends io.grpc.stub.AbstractStub<ApiServiceFutureStub> {
    private ApiServiceFutureStub(io.grpc.Channel channel) {
      super(channel);
    }

    private ApiServiceFutureStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ApiServiceFutureStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new ApiServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<protocol.Api.Resp> sayHello(
        protocol.Api.Req request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_SAY_HELLO, getCallOptions()), request);
    }

    /**
     * <pre>
     * 注册 绑定 登录
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<protocol.Api.RespRegister> register(
        protocol.Api.ReqRegister request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_REGISTER, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<protocol.Api.RespBand> bind(
        protocol.Api.ReqBand request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_BIND, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<protocol.Api.RespLogin> getBind(
        protocol.Api.ReqLogin request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_GET_BIND, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<protocol.Api.RespCheckAccount> getAccount(
        protocol.Api.ReqCheckAccount request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_GET_ACCOUNT, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<protocol.Api.RespGetEthBalance> getEthBalance(
        protocol.Api.ReqGetEthBalance request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_GET_ETH_BALANCE, getCallOptions()), request);
    }

    /**
     * <pre>
     * 排班 申请工作
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<protocol.Api.RespScheduling> setSchedule(
        protocol.Api.ReqScheduling request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_SET_SCHEDULE, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<protocol.Api.RespGetSchedue> getSchedule(
        protocol.Api.ReqGetSchedue request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_GET_SCHEDULE, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<protocol.Api.RespGetSchedue> getCanJoin(
        protocol.Api.ReqGetSchedue request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_GET_CAN_JOIN, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<protocol.Api.RespGetStaff> getApply(
        protocol.Api.ReqGetStaff request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_GET_APPLY, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<protocol.Api.RespGetCanApply> getJob(
        protocol.Api.ReqGetCanApply request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_GET_JOB, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<protocol.Api.RespFindJob> applyJob(
        protocol.Api.ReqFindJob request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_APPLY_JOB, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<protocol.Api.RespHistoryJoin> historyJoin(
        protocol.Api.ReqHistoryJoin request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_HISTORY_JOIN, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<protocol.Api.RespCheckIsOkApplication> checkIsOkApplication(
        protocol.Api.ReqCheckIsOkApplication request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_CHECK_IS_OK_APPLICATION, getCallOptions()), request);
    }

    /**
     * <pre>
     * 收入
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<protocol.Api.RespPay> pay(
        protocol.Api.ReqPay request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_PAY, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<protocol.Api.RespGetAllOrder> getAllOrderBySchedule(
        protocol.Api.ReqGetAllOrder request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_GET_ALL_ORDER_BY_SCHEDULE, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<protocol.Api.RespGetContent> getContent(
        protocol.Api.ReqGetContent request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_GET_CONTENT, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<protocol.Api.RespGetAllIncome> getAllIncome(
        protocol.Api.ReqGetAllIncome request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_GET_ALL_INCOME, getCallOptions()), request);
    }

    /**
     * <pre>
     * 3.5.根据排班信息获取用户收入详情
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<protocol.Api.RespGetIncomeBySchedule> getIncomeBySchedule(
        protocol.Api.ReqGetIncomeBySchedule request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_GET_INCOME_BY_SCHEDULE, getCallOptions()), request);
    }

    /**
     * <pre>
     * 查询
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<protocol.Api.RespGetBalance> getBalance(
        protocol.Api.ReqGetBalance request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_GET_BALANCE, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<protocol.Api.RespGetAllMoney> getAllMoney(
        protocol.Api.ReqGetAllMoney request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_GET_ALL_MONEY, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<protocol.Api.RespGetNowJobAddr> getNowJobAddress(
        protocol.Api.ReqGetNowJobAddr request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_GET_NOW_JOB_ADDRESS, getCallOptions()), request);
    }

    /**
     * <pre>
     * 用人白名单
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<protocol.Api.RespAddWhite> addWhiteList(
        protocol.Api.ReqAddWhite request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_ADD_WHITE_LIST, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<protocol.Api.RespGetWhite> getWhiteList(
        protocol.Api.ReqGetWhite request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_GET_WHITE_LIST, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<protocol.Api.RespDelWhite> delWhiteList(
        protocol.Api.ReqDelWhite request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_DEL_WHITE_LIST, getCallOptions()), request);
    }

    /**
     * <pre>
     * 三方
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<protocol.Api.RespThreeSetOrder> threeSetOrder(
        protocol.Api.ReqThreeSetOrder request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_THREE_SET_ORDER, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<protocol.Api.RespThreeSetBill> threeSetBill(
        protocol.Api.ReqThreeSetBill request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_THREE_SET_BILL, getCallOptions()), request);
    }

    /**
     * <pre>
     * 无用
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<protocol.Api.RespManageContract> manageContract(
        protocol.Api.ReqManageContract request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_MANAGE_CONTRACT, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<protocol.Api.RespCheckContract> checkContract(
        protocol.Api.ReqCheckContract request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_CHECK_CONTRACT, getCallOptions()), request);
    }

    /**
     * <pre>
     * 配置服务
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<protocol.Api.SReloadConfig> reloadConfig(
        protocol.Api.CReloadConfig request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_RELOAD_CONFIG, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<protocol.Api.SReloadDeploy> reloadDeploy(
        protocol.Api.CReloadDeploy request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_RELOAD_DEPLOY, getCallOptions()), request);
    }
  }

  private static final int METHODID_SAY_HELLO = 0;
  private static final int METHODID_REGISTER = 1;
  private static final int METHODID_BIND = 2;
  private static final int METHODID_GET_BIND = 3;
  private static final int METHODID_GET_ACCOUNT = 4;
  private static final int METHODID_GET_ETH_BALANCE = 5;
  private static final int METHODID_SET_SCHEDULE = 6;
  private static final int METHODID_GET_SCHEDULE = 7;
  private static final int METHODID_GET_CAN_JOIN = 8;
  private static final int METHODID_GET_APPLY = 9;
  private static final int METHODID_GET_JOB = 10;
  private static final int METHODID_APPLY_JOB = 11;
  private static final int METHODID_HISTORY_JOIN = 12;
  private static final int METHODID_CHECK_IS_OK_APPLICATION = 13;
  private static final int METHODID_PAY = 14;
  private static final int METHODID_GET_ALL_ORDER_BY_SCHEDULE = 15;
  private static final int METHODID_GET_CONTENT = 16;
  private static final int METHODID_GET_ALL_INCOME = 17;
  private static final int METHODID_GET_INCOME_BY_SCHEDULE = 18;
  private static final int METHODID_GET_BALANCE = 19;
  private static final int METHODID_GET_ALL_MONEY = 20;
  private static final int METHODID_GET_NOW_JOB_ADDRESS = 21;
  private static final int METHODID_ADD_WHITE_LIST = 22;
  private static final int METHODID_GET_WHITE_LIST = 23;
  private static final int METHODID_DEL_WHITE_LIST = 24;
  private static final int METHODID_THREE_SET_ORDER = 25;
  private static final int METHODID_THREE_SET_BILL = 26;
  private static final int METHODID_MANAGE_CONTRACT = 27;
  private static final int METHODID_CHECK_CONTRACT = 28;
  private static final int METHODID_RELOAD_CONFIG = 29;
  private static final int METHODID_RELOAD_DEPLOY = 30;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final ApiServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(ApiServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_SAY_HELLO:
          serviceImpl.sayHello((protocol.Api.Req) request,
              (io.grpc.stub.StreamObserver<protocol.Api.Resp>) responseObserver);
          break;
        case METHODID_REGISTER:
          serviceImpl.register((protocol.Api.ReqRegister) request,
              (io.grpc.stub.StreamObserver<protocol.Api.RespRegister>) responseObserver);
          break;
        case METHODID_BIND:
          serviceImpl.bind((protocol.Api.ReqBand) request,
              (io.grpc.stub.StreamObserver<protocol.Api.RespBand>) responseObserver);
          break;
        case METHODID_GET_BIND:
          serviceImpl.getBind((protocol.Api.ReqLogin) request,
              (io.grpc.stub.StreamObserver<protocol.Api.RespLogin>) responseObserver);
          break;
        case METHODID_GET_ACCOUNT:
          serviceImpl.getAccount((protocol.Api.ReqCheckAccount) request,
              (io.grpc.stub.StreamObserver<protocol.Api.RespCheckAccount>) responseObserver);
          break;
        case METHODID_GET_ETH_BALANCE:
          serviceImpl.getEthBalance((protocol.Api.ReqGetEthBalance) request,
              (io.grpc.stub.StreamObserver<protocol.Api.RespGetEthBalance>) responseObserver);
          break;
        case METHODID_SET_SCHEDULE:
          serviceImpl.setSchedule((protocol.Api.ReqScheduling) request,
              (io.grpc.stub.StreamObserver<protocol.Api.RespScheduling>) responseObserver);
          break;
        case METHODID_GET_SCHEDULE:
          serviceImpl.getSchedule((protocol.Api.ReqGetSchedue) request,
              (io.grpc.stub.StreamObserver<protocol.Api.RespGetSchedue>) responseObserver);
          break;
        case METHODID_GET_CAN_JOIN:
          serviceImpl.getCanJoin((protocol.Api.ReqGetSchedue) request,
              (io.grpc.stub.StreamObserver<protocol.Api.RespGetSchedue>) responseObserver);
          break;
        case METHODID_GET_APPLY:
          serviceImpl.getApply((protocol.Api.ReqGetStaff) request,
              (io.grpc.stub.StreamObserver<protocol.Api.RespGetStaff>) responseObserver);
          break;
        case METHODID_GET_JOB:
          serviceImpl.getJob((protocol.Api.ReqGetCanApply) request,
              (io.grpc.stub.StreamObserver<protocol.Api.RespGetCanApply>) responseObserver);
          break;
        case METHODID_APPLY_JOB:
          serviceImpl.applyJob((protocol.Api.ReqFindJob) request,
              (io.grpc.stub.StreamObserver<protocol.Api.RespFindJob>) responseObserver);
          break;
        case METHODID_HISTORY_JOIN:
          serviceImpl.historyJoin((protocol.Api.ReqHistoryJoin) request,
              (io.grpc.stub.StreamObserver<protocol.Api.RespHistoryJoin>) responseObserver);
          break;
        case METHODID_CHECK_IS_OK_APPLICATION:
          serviceImpl.checkIsOkApplication((protocol.Api.ReqCheckIsOkApplication) request,
              (io.grpc.stub.StreamObserver<protocol.Api.RespCheckIsOkApplication>) responseObserver);
          break;
        case METHODID_PAY:
          serviceImpl.pay((protocol.Api.ReqPay) request,
              (io.grpc.stub.StreamObserver<protocol.Api.RespPay>) responseObserver);
          break;
        case METHODID_GET_ALL_ORDER_BY_SCHEDULE:
          serviceImpl.getAllOrderBySchedule((protocol.Api.ReqGetAllOrder) request,
              (io.grpc.stub.StreamObserver<protocol.Api.RespGetAllOrder>) responseObserver);
          break;
        case METHODID_GET_CONTENT:
          serviceImpl.getContent((protocol.Api.ReqGetContent) request,
              (io.grpc.stub.StreamObserver<protocol.Api.RespGetContent>) responseObserver);
          break;
        case METHODID_GET_ALL_INCOME:
          serviceImpl.getAllIncome((protocol.Api.ReqGetAllIncome) request,
              (io.grpc.stub.StreamObserver<protocol.Api.RespGetAllIncome>) responseObserver);
          break;
        case METHODID_GET_INCOME_BY_SCHEDULE:
          serviceImpl.getIncomeBySchedule((protocol.Api.ReqGetIncomeBySchedule) request,
              (io.grpc.stub.StreamObserver<protocol.Api.RespGetIncomeBySchedule>) responseObserver);
          break;
        case METHODID_GET_BALANCE:
          serviceImpl.getBalance((protocol.Api.ReqGetBalance) request,
              (io.grpc.stub.StreamObserver<protocol.Api.RespGetBalance>) responseObserver);
          break;
        case METHODID_GET_ALL_MONEY:
          serviceImpl.getAllMoney((protocol.Api.ReqGetAllMoney) request,
              (io.grpc.stub.StreamObserver<protocol.Api.RespGetAllMoney>) responseObserver);
          break;
        case METHODID_GET_NOW_JOB_ADDRESS:
          serviceImpl.getNowJobAddress((protocol.Api.ReqGetNowJobAddr) request,
              (io.grpc.stub.StreamObserver<protocol.Api.RespGetNowJobAddr>) responseObserver);
          break;
        case METHODID_ADD_WHITE_LIST:
          serviceImpl.addWhiteList((protocol.Api.ReqAddWhite) request,
              (io.grpc.stub.StreamObserver<protocol.Api.RespAddWhite>) responseObserver);
          break;
        case METHODID_GET_WHITE_LIST:
          serviceImpl.getWhiteList((protocol.Api.ReqGetWhite) request,
              (io.grpc.stub.StreamObserver<protocol.Api.RespGetWhite>) responseObserver);
          break;
        case METHODID_DEL_WHITE_LIST:
          serviceImpl.delWhiteList((protocol.Api.ReqDelWhite) request,
              (io.grpc.stub.StreamObserver<protocol.Api.RespDelWhite>) responseObserver);
          break;
        case METHODID_THREE_SET_ORDER:
          serviceImpl.threeSetOrder((protocol.Api.ReqThreeSetOrder) request,
              (io.grpc.stub.StreamObserver<protocol.Api.RespThreeSetOrder>) responseObserver);
          break;
        case METHODID_THREE_SET_BILL:
          serviceImpl.threeSetBill((protocol.Api.ReqThreeSetBill) request,
              (io.grpc.stub.StreamObserver<protocol.Api.RespThreeSetBill>) responseObserver);
          break;
        case METHODID_MANAGE_CONTRACT:
          serviceImpl.manageContract((protocol.Api.ReqManageContract) request,
              (io.grpc.stub.StreamObserver<protocol.Api.RespManageContract>) responseObserver);
          break;
        case METHODID_CHECK_CONTRACT:
          serviceImpl.checkContract((protocol.Api.ReqCheckContract) request,
              (io.grpc.stub.StreamObserver<protocol.Api.RespCheckContract>) responseObserver);
          break;
        case METHODID_RELOAD_CONFIG:
          serviceImpl.reloadConfig((protocol.Api.CReloadConfig) request,
              (io.grpc.stub.StreamObserver<protocol.Api.SReloadConfig>) responseObserver);
          break;
        case METHODID_RELOAD_DEPLOY:
          serviceImpl.reloadDeploy((protocol.Api.CReloadDeploy) request,
              (io.grpc.stub.StreamObserver<protocol.Api.SReloadDeploy>) responseObserver);
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
      synchronized (ApiServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .addMethod(METHOD_SAY_HELLO)
              .addMethod(METHOD_REGISTER)
              .addMethod(METHOD_BIND)
              .addMethod(METHOD_GET_BIND)
              .addMethod(METHOD_GET_ACCOUNT)
              .addMethod(METHOD_GET_ETH_BALANCE)
              .addMethod(METHOD_SET_SCHEDULE)
              .addMethod(METHOD_GET_SCHEDULE)
              .addMethod(METHOD_GET_CAN_JOIN)
              .addMethod(METHOD_GET_APPLY)
              .addMethod(METHOD_GET_JOB)
              .addMethod(METHOD_APPLY_JOB)
              .addMethod(METHOD_HISTORY_JOIN)
              .addMethod(METHOD_CHECK_IS_OK_APPLICATION)
              .addMethod(METHOD_PAY)
              .addMethod(METHOD_GET_ALL_ORDER_BY_SCHEDULE)
              .addMethod(METHOD_GET_CONTENT)
              .addMethod(METHOD_GET_ALL_INCOME)
              .addMethod(METHOD_GET_INCOME_BY_SCHEDULE)
              .addMethod(METHOD_GET_BALANCE)
              .addMethod(METHOD_GET_ALL_MONEY)
              .addMethod(METHOD_GET_NOW_JOB_ADDRESS)
              .addMethod(METHOD_ADD_WHITE_LIST)
              .addMethod(METHOD_GET_WHITE_LIST)
              .addMethod(METHOD_DEL_WHITE_LIST)
              .addMethod(METHOD_THREE_SET_ORDER)
              .addMethod(METHOD_THREE_SET_BILL)
              .addMethod(METHOD_MANAGE_CONTRACT)
              .addMethod(METHOD_CHECK_CONTRACT)
              .addMethod(METHOD_RELOAD_CONFIG)
              .addMethod(METHOD_RELOAD_DEPLOY)
              .build();
        }
      }
    }
    return result;
  }
}
