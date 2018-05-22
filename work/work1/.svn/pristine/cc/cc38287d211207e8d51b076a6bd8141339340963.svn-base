package com.znfz.assur.znfz_android.android.common.task;

import com.orhanobut.logger.Logger;
import com.znfz.assur.znfz_android.android.common.task.base.BaseGrpcTask;
import io.grpc.ManagedChannel;
import protocol.Api;

/**
 *  查询用户信息
 * Created by assur on 2018/4/20 0920.
 */
public class CheckAccountTask extends BaseGrpcTask<Api.RespCheckAccount> {

    // 请求参数
    private  protocol.Api.ReqCheckAccount request;

    public void setRequest(Api.ReqCheckAccount request) {
        this.request = request;
    }

    // 更新UI接口
    private UpdateUIInterface updateUIInterface;
    public void setUpdateUIInterface(UpdateUIInterface updateUIInterface) {
        this.updateUIInterface = updateUIInterface;
    }
    public interface UpdateUIInterface {
        void onSucceed(Api.RespCheckAccount result);
        void onError();
    }




    @Override
    protected Api.RespCheckAccount doInback(ManagedChannel channel) {

        Logger.d("=====CheckAccountTask:request====\n" + request.toString());
        protocol.ApiServiceGrpc.ApiServiceBlockingStub stub = protocol.ApiServiceGrpc.newBlockingStub(channel);
        Api.RespCheckAccount response = stub.getAccount(request);
        Logger.d("=====CheckAccountTask:response====\n" + response.toString());
        return response;
    }

    @Override
    protected void onSucceed(Api.RespCheckAccount result) {

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
