package com.znfz.assur.znfz_android.android.common.task;

import com.orhanobut.logger.Logger;
import com.znfz.assur.znfz_android.android.common.task.base.BaseGrpcTask;
import io.grpc.ManagedChannel;
import protocol.Api;

/**
 *  查询所有工作
 * Created by assur on 2018/4/28
 */
public class GetJobTask extends BaseGrpcTask<protocol.Api.RespGetCanApply> {

    // 请求参数
    private  protocol.Api.ReqGetCanApply request;

    public void setRequest(Api.ReqGetCanApply request) {
        this.request = request;
    }

    // 更新UI接口
    private GetJobTask.UpdateUIInterface updateUIInterface;
    public void setUpdateUIInterface(GetJobTask.UpdateUIInterface updateUIInterface) {
        this.updateUIInterface = updateUIInterface;
    }
    public interface UpdateUIInterface {
        void onSucceed(Api.RespGetCanApply result);
        void onError();
    }


    @Override
    protected Api.RespGetCanApply doInback(ManagedChannel channel) {
        Logger.d("=====GetJobTask:request====\n" + request.toString());
        protocol.ApiServiceGrpc.ApiServiceBlockingStub stub = protocol.ApiServiceGrpc.newBlockingStub(channel);
        Api.RespGetCanApply response = stub.getJob(request);
        Logger.d("=====GetJobTask:response====\n" + response.toString());
        return response;
    }

    @Override
    protected void onSucceed(Api.RespGetCanApply result) {
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
