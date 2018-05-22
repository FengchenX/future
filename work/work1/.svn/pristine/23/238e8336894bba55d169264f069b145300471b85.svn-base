package com.znfz.assur.znfz_android.android.common.task;

import com.znfz.assur.znfz_android.android.common.task.base.BaseGrpcTask;
import io.grpc.ManagedChannel;
import protocol.Api;

/**
 *  注册
 * Created by assur on 2018/4/20 0920.
 */
public class RegisterTask extends BaseGrpcTask<protocol.Api.RespRegister> {

    // 更新UI接口
    private UpdateUIInterface updateUIInterface;
    public void setUpdateUIInterface(UpdateUIInterface updateUIInterface) {
        this.updateUIInterface = updateUIInterface;
    }
    public interface UpdateUIInterface {
        void onSucceed(Api.RespRegister result);
        void onError();
    }




    @Override
    protected Api.RespRegister doInback(ManagedChannel channel) {

        protocol.Api.ReqRegister.Builder requestBuilder = protocol.Api.ReqRegister.newBuilder();
//        requestBuilder.setPassWord(map.get(Key_PassWord));
        protocol.Api.ReqRegister request = requestBuilder.build();

        protocol.ApiServiceGrpc.ApiServiceBlockingStub stub = protocol.ApiServiceGrpc.newBlockingStub(channel);
        Api.RespRegister response = stub.register(request);
        return response;
    }

    @Override
    protected void onSucceed(Api.RespRegister result) {

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
