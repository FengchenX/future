package com.znfz.assur.znfz_android.android.common.task;

import com.orhanobut.logger.Logger;
import com.znfz.assur.znfz_android.android.common.task.base.BaseGrpcTask;
import io.grpc.ManagedChannel;
import protocol.Api;

/**
 *  获取用户是否绑定
 * Created by assur on 2018/5/19
 */
public class GetBandTask extends BaseGrpcTask<protocol.Api.RespLogin> {

    // 请求参数
    private  protocol.Api.ReqLogin request;

    public void setRequest(Api.ReqLogin request) {
        this.request = request;
    }

    // 更新UI接口
    private GetBandTask.UpdateUIInterface updateUIInterface;
    public void setUpdateUIInterface(GetBandTask.UpdateUIInterface updateUIInterface) {
        this.updateUIInterface = updateUIInterface;
    }
    public interface UpdateUIInterface {
        void onSucceed(Api.RespLogin result);
        void onError();
    }


    @Override
    protected Api.RespLogin doInback(ManagedChannel channel) {
        Logger.d("=====GetBandTask:request====\n" + request.toString());
        protocol.ApiServiceGrpc.ApiServiceBlockingStub stub = protocol.ApiServiceGrpc.newBlockingStub(channel);
        Api.RespLogin response = stub.getBind(request);
        Logger.d("=====GetBandTask:response====\n" + response.toString());
        return response;
    }

    @Override
    protected void onSucceed(Api.RespLogin result) {
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
