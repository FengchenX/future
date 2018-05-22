package com.znfz.assur.znfz_android.android.common.task;

import com.orhanobut.logger.Logger;
import com.znfz.assur.znfz_android.android.common.task.base.BaseGrpcTask;
import io.grpc.ManagedChannel;
import protocol.Api;

/**
 *  查询所有排班
 * Created by assur on 2018/4/20 0920.
 */
public class GetScheduleTask extends BaseGrpcTask<protocol.Api.RespGetSchedue> {

    // 请求参数
    private  protocol.Api.ReqGetSchedue request;

    public void setRequest(Api.ReqGetSchedue request) {
        this.request = request;
    }

    // 更新UI接口
    private GetScheduleTask.UpdateUIInterface updateUIInterface;
    public void setUpdateUIInterface(GetScheduleTask.UpdateUIInterface updateUIInterface) {
        this.updateUIInterface = updateUIInterface;
    }
    public interface UpdateUIInterface {
        void onSucceed(Api.RespGetSchedue result);
        void onError();
    }


    @Override
    protected Api.RespGetSchedue doInback(ManagedChannel channel) {
        Logger.d("=====GetScheduleTask:request====\n" + request.toString());
        protocol.ApiServiceGrpc.ApiServiceBlockingStub stub = protocol.ApiServiceGrpc.newBlockingStub(channel);
        Api.RespGetSchedue response = stub.getSchedule(request);
        Logger.d("=====GetScheduleTask:response====\n" + response.toString());
        return response;
    }

    @Override
    protected void onSucceed(Api.RespGetSchedue result) {
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
