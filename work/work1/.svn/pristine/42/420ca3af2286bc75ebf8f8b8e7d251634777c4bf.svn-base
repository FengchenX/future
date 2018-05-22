package com.znfz.assur.znfz_android.android.common.task;

import com.orhanobut.logger.Logger;
import com.znfz.assur.znfz_android.android.common.task.base.BaseGrpcTask;
import io.grpc.ManagedChannel;
import protocol.Api;

/**
 *  用户绑定
 * Created by assur on 2018/5/18
 */
public class BandTask extends BaseGrpcTask<protocol.Api.RespBand> {

    // 请求参数
    private  protocol.Api.ReqBand request;

    public void setRequest(Api.ReqBand request) {
        this.request = request;
    }

    // 更新UI接口
    private BandTask.UpdateUIInterface updateUIInterface;
    public void setUpdateUIInterface(BandTask.UpdateUIInterface updateUIInterface) {
        this.updateUIInterface = updateUIInterface;
    }
    public interface UpdateUIInterface {
        void onSucceed(Api.RespBand result);
        void onError();
    }


    @Override
    protected Api.RespBand doInback(ManagedChannel channel) {
        Logger.d("=====BandTask:request====\n" + request.toString());
        protocol.ApiServiceGrpc.ApiServiceBlockingStub stub = protocol.ApiServiceGrpc.newBlockingStub(channel);
        Api.RespBand response = stub.bind(request);
        Logger.d("=====BandTask:response====\n" + response.toString());
        return response;
    }

    @Override
    protected void onSucceed(Api.RespBand result) {
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
