package com.znfz.assur.znfz_android.android.common.task;

import com.orhanobut.logger.Logger;
import com.znfz.assur.znfz_android.android.common.task.base.BaseGrpcTask;
import java.util.HashMap;
import io.grpc.ManagedChannel;
import protocol.Api;

/**
 *  查询排班下所有订单
 * Created by assur on 2018/5/5
 */
public class GetAllOrderByScheduleTask extends BaseGrpcTask<protocol.Api.RespGetAllOrder> {

    // 请求参数
    private  protocol.Api.ReqGetAllOrder request;

    public void setRequest(Api.ReqGetAllOrder request) {
        this.request = request;
    }

    // 更新UI接口
    private GetAllOrderByScheduleTask.UpdateUIInterface updateUIInterface;
    public void setUpdateUIInterface(GetAllOrderByScheduleTask.UpdateUIInterface updateUIInterface) {
        this.updateUIInterface = updateUIInterface;
    }
    public interface UpdateUIInterface {
        void onSucceed(Api.RespGetAllOrder result);
        void onError();
    }


    @Override
    protected Api.RespGetAllOrder doInback(ManagedChannel channel) {
        Logger.d("=====GetAllOrderByScheduleTask:request====\n" + request.toString());
        protocol.ApiServiceGrpc.ApiServiceBlockingStub stub = protocol.ApiServiceGrpc.newBlockingStub(channel);
        Api.RespGetAllOrder response = stub.getAllOrderBySchedule(request);
        Logger.d("=====GetAllOrderByScheduleTask:response====\n" + response.toString());
        return response;
    }

    @Override
    protected void onSucceed(Api.RespGetAllOrder result) {
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
