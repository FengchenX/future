package com.znfz.assur.znfz_android.android.common.task;

import com.orhanobut.logger.Logger;
import com.znfz.assur.znfz_android.android.common.task.base.BaseGrpcTask;
import io.grpc.ManagedChannel;
import protocol.Api;

/**
 *  查询申请情况
 * Created by assur on 2018/5/6
 */
public class GetApplyTask extends BaseGrpcTask<protocol.Api.RespGetStaff> {

    // 请求参数
    private  protocol.Api.ReqGetStaff request;

    public void setRequest(Api.ReqGetStaff request) {
        this.request = request;
    }

    // 更新UI接口
    private GetApplyTask.UpdateUIInterface updateUIInterface;
    public void setUpdateUIInterface(GetApplyTask.UpdateUIInterface updateUIInterface) {
        this.updateUIInterface = updateUIInterface;
    }
    public interface UpdateUIInterface {
        void onSucceed(Api.RespGetStaff result);
        void onError();
    }


    @Override
    protected Api.RespGetStaff doInback(ManagedChannel channel) {
        Logger.d("=====GetApplyTask:request====\n" + request.toString());
        protocol.ApiServiceGrpc.ApiServiceBlockingStub stub = protocol.ApiServiceGrpc.newBlockingStub(channel);
        Api.RespGetStaff response = stub.getApply(request);
        Logger.d("=====GetApplyTask:response====\n" + response.toString());
        return response;
    }

    @Override
    protected void onSucceed(Api.RespGetStaff result) {
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
