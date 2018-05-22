package com.znfz.assur.znfz_android.android.common.task;

import com.orhanobut.logger.Logger;
import com.znfz.assur.znfz_android.android.common.task.base.BaseGrpcTask;
import io.grpc.ManagedChannel;
import protocol.Api;

/**
 *  查询用户历史收入详情
 * Created by assur on 2018/5/5
 */
public class GetAllIncomeTask extends BaseGrpcTask<protocol.Api.RespGetAllIncome> {

    // 请求参数
    private  protocol.Api.ReqGetAllIncome request;

    public void setRequest(Api.ReqGetAllIncome request) {
        this.request = request;
    }

    // 更新UI接口
    private GetAllIncomeTask.UpdateUIInterface updateUIInterface;
    public void setUpdateUIInterface(GetAllIncomeTask.UpdateUIInterface updateUIInterface) {
        this.updateUIInterface = updateUIInterface;
    }
    public interface UpdateUIInterface {
        void onSucceed(Api.RespGetAllIncome result);
        void onError();
    }


    @Override
    protected Api.RespGetAllIncome doInback(ManagedChannel channel) {
        Logger.d("=====GetAllIncomeTask:request====\n" + request.toString());
        protocol.ApiServiceGrpc.ApiServiceBlockingStub stub = protocol.ApiServiceGrpc.newBlockingStub(channel);
        Api.RespGetAllIncome response = stub.getAllIncome(request);
        Logger.d("=====GetAllIncomeTask:response====\n" + response.toString());
        return response;
    }

    @Override
    protected void onSucceed(Api.RespGetAllIncome result) {
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
