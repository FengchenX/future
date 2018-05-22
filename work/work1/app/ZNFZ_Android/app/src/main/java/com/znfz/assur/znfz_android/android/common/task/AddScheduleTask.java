package com.znfz.assur.znfz_android.android.common.task;

import com.orhanobut.logger.Logger;
import com.znfz.assur.znfz_android.android.common.task.base.BaseGrpcTask;
import io.grpc.ManagedChannel;
import protocol.Api;

/**
 *  发布排班
 * Created by assur on 2018/4/27
 */
public class AddScheduleTask extends BaseGrpcTask<protocol.Api.RespScheduling> {

    // 请求参数
    private  protocol.Api.ReqScheduling request;

    public void setRequest(Api.ReqScheduling request) {
        this.request = request;
    }

    // 更新UI接口
    private UpdateUIInterface updateUIInterface;
    public void setUpdateUIInterface(UpdateUIInterface updateUIInterface) {
        this.updateUIInterface = updateUIInterface;
    }
    public interface UpdateUIInterface {
        void onSucceed(Api.RespScheduling result);
        void onError();
    }




    @Override
    protected Api.RespScheduling doInback(ManagedChannel channel) {
        Logger.d("=====AddScheduleTask:request====\n" + request.toString());
        protocol.ApiServiceGrpc.ApiServiceBlockingStub stub = protocol.ApiServiceGrpc.newBlockingStub(channel);
        Api.RespScheduling response = stub.setSchedule(request);
        Logger.d("=====AddScheduleTask:response====\n" + response.toString());
        return response;
    }

    @Override
    protected void onSucceed(Api.RespScheduling result) {

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
