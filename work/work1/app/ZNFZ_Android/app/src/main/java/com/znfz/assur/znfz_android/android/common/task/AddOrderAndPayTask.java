package com.znfz.assur.znfz_android.android.common.task;

import com.orhanobut.logger.Logger;
import com.znfz.assur.znfz_android.android.common.task.base.BaseGrpcTask;
import io.grpc.ManagedChannel;
import protocol.Api;

/**
 *  添加订单并付款
 * Created by assur on 2018/5/5
 */
public class AddOrderAndPayTask extends BaseGrpcTask<protocol.Api.RespPay> {

    // 请求参数
    private  protocol.Api.ReqPay request;

    public void setRequest(Api.ReqPay request) {
        this.request = request;
    }

    // 更新UI接口
    private AddOrderAndPayTask.UpdateUIInterface updateUIInterface;
    public void setUpdateUIInterface(AddOrderAndPayTask.UpdateUIInterface updateUIInterface) {
        this.updateUIInterface = updateUIInterface;
    }
    public interface UpdateUIInterface {
        void onSucceed(Api.RespPay result);
        void onError();
    }


    @Override
    protected Api.RespPay doInback(ManagedChannel channel) {
        Logger.d("=====AddOrderAndPay:request====\n" + request.toString());
        protocol.ApiServiceGrpc.ApiServiceBlockingStub stub = protocol.ApiServiceGrpc.newBlockingStub(channel);
        Api.RespPay response = stub.pay(request);
        Logger.d("=====AddOrderAndPay:response====\n" + response.toString());
        return response;
    }

    @Override
    protected void onSucceed(Api.RespPay result) {
        if(updateUIInterface != null) {
            updateUIInterface.onSucceed(result);
        }
    }

    @Override
    protected void onError() {
        if(updateUIInterface != null) {
            updateUIInterface.onError();
        }
    }
}
