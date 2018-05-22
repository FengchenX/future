package com.znfz.assur.znfz_android.android.common.task;

import com.orhanobut.logger.Logger;
import com.znfz.assur.znfz_android.android.common.task.base.BaseGrpcTask;
import io.grpc.ManagedChannel;
import protocol.Api;

/**
 *  申请工作
 * Created by assur on 2018/4/28
 */
public class ApplyJobTask extends BaseGrpcTask<protocol.Api.RespFindJob> {

    // 请求参数
    private  protocol.Api.ReqFindJob request;

    public void setRequest(Api.ReqFindJob request) {
        this.request = request;
    }

    // 更新UI接口
    private ApplyJobTask.UpdateUIInterface updateUIInterface;
    public void setUpdateUIInterface(ApplyJobTask.UpdateUIInterface updateUIInterface) {
        this.updateUIInterface = updateUIInterface;
    }
    public interface UpdateUIInterface {
        void onSucceed(Api.RespFindJob result);
        void onError();
    }


    @Override
    protected Api.RespFindJob doInback(ManagedChannel channel) {
        Logger.d("=====ApplyJobTask:request====\n" + request.toString());
        protocol.ApiServiceGrpc.ApiServiceBlockingStub stub = protocol.ApiServiceGrpc.newBlockingStub(channel);
        Api.RespFindJob response = stub.applyJob(request);
        Logger.d("=====ApplyJobTask:response====\n" + response.toString());
        return response;
    }

    @Override
    protected void onSucceed(Api.RespFindJob result) {
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
