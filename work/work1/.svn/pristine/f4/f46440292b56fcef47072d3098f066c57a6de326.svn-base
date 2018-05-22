package com.znfz.assur.znfz_android.android.test.grpc_interface_test;

import android.view.View;
import android.widget.Button;

import com.orhanobut.logger.Logger;
import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.app.BaseActivity;
import com.znfz.assur.znfz_android.android.common.global.UserInfo;
import com.znfz.assur.znfz_android.android.common.global.Variable;
import com.znfz.assur.znfz_android.android.common.task.AddOrderAndPayTask;
import com.znfz.assur.znfz_android.android.common.task.AddScheduleTask;
import com.znfz.assur.znfz_android.android.common.task.ApplyJobTask;
import com.znfz.assur.znfz_android.android.common.task.CheckAccountTask;
import com.znfz.assur.znfz_android.android.common.task.GetAllOrderByScheduleTask;
import com.znfz.assur.znfz_android.android.common.task.GetScheduleTask;
import com.znfz.assur.znfz_android.android.common.task.LoginTask;
import com.znfz.assur.znfz_android.android.common.task.RegisterTask;
import com.znfz.assur.znfz_android.android.common.utils.DialogUtils;
import com.znfz.assur.znfz_android.android.common.utils.TimeUtils;
import com.znfz.assur.znfz_android.android.common.utils.ToastUtils;
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherOrderBean;
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherRoleBean;
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherSchedulingBean;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import butterknife.BindView;
import butterknife.ButterKnife;
import butterknife.OnClick;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import protocol.Api;


/**
 * 接口测试界面
 * Created by assur on 2018/4/20.
 *
 *
 * manager 86111111
 * client 86222222
 *
以下测试用户：
 var Phone       = "55555"
 var managerAddr = "0x87b67fCf385CE4919884Fb982DD63be36ED3Ef38"
 var managerPass = "9999"
 var managerDesp = "{\"address\":\"87b67fcf385ce4919884fb982dd63be36ed3ef38\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"598ef34d426fed459fdfe47354ff105e2052496f0853306c45dac2cb37173be9\",\"cipherparams\":{\"iv\":\"bc6d86dec9eee329f07208e279c6dce8\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"02dbc0761bb3220b27a3a14de8d21f87e210312156e3efdcd43a33930800baf4\"},\"mac\":\"b121af4e14390fdb2b29ff823f747c7a7589456a4daaaacef87e5dad9ee2e11a\"},\"id\":\"9aec95aa-35e4-4ff6-b0b8-d39acfa89c18\",\"version\":3}"



 // 经理
 manager{
 name:"assur_manager",
 role:"1",
 password:"8888",
 phone:"11111",
 address:"0x8F392805FC157AcfCbc9585E61Ab97e8d383a97b",
 }

 // 厨师
 cooker{
 name:"assur_cooker",
 role:"2",
 password:"8888",
 phone:"22222",
 address:"0x5B413FBE6Fb09f9255453A94c3ea8E7120E2734A",
 }

 // 服务员
 waiter{
 name:"assur_waiter",
 role:"3",
 password:"8888",
 phone:"33333",
 address:"0x3d83F0128862BEb8cd249152bF1fC8D419928Ae4",
 }
 *
 *
 *
 *
 */
public class TestInterfaceActivity extends BaseActivity {


    @BindView(R.id.btn1)
    Button btn1;
    @BindView(R.id.btn2)
    Button btn2;
    @BindView(R.id.btn3)
    Button btn3;
    @BindView(R.id.btn4)
    Button btn4;
    @BindView(R.id.btn5)
    Button btn5;
    @BindView(R.id.btn6)
    Button btn6;
    @BindView(R.id.btn7)
    Button btn7;
    @BindView(R.id.btn8)
    Button btn8;
    @BindView(R.id.btn9)
    Button btn9;
    @BindView(R.id.btn10)
    Button btn10;
    @BindView(R.id.btn11)
    Button btn11;
    @BindView(R.id.btn12)
    Button btn12;
    @BindView(R.id.btn13)
    Button btn13;
    @BindView(R.id.btn14)
    Button btn14;

    @Override
    protected void initView() {
        setContentView(R.layout.activity_test_interface);
        ButterKnife.bind(this);
    }

    @Override
    protected void initData() {

    }


    ManagedChannel mTestChannel;
    @OnClick(R.id.btn1)
    public void btn1Click() {
        sayHello();
    }

    @OnClick(R.id.btn2)
    public void btn2Click() {
        register();
    }

    @OnClick(R.id.btn3)
    public void btn3Click() {
        manageContract();
    }

    @OnClick(R.id.btn4)
    public void btn4Click() {
        checkContract();
    }

    @OnClick(R.id.btn5)
    public void btn5Click() {
        setSchedule();
    }

    @OnClick(R.id.btn6)
    public void btn6Click() {
        getSchedule();
    }

    @OnClick(R.id.btn7)
    public void btn7Click() {
        getApply();
    }

    @OnClick(R.id.btn8)
    public void btn8Click() {
        applyJob();
    }

    @OnClick(R.id.btn9)
    public void btn9Click() {
        getOrderByScheduleRequest();
    }

    @OnClick(R.id.btn10)
    public void btn10Click() {
        checkAccount();
    }

    @OnClick(R.id.btn11)
    public void btn11Click() {
        setContent();
    }

    @OnClick(R.id.btn12)
    public void btn12Click() {
        getContent();
    }

    @OnClick(R.id.btn13)
    public void btn13Click() {
        getBalance();
    }

    @OnClick(R.id.btn14)
    public void btn14Click() {
        checkIsOkApplication();
    }

    @OnClick(R.id.btn15)
    public void btn15Click() {
        login();
    }





    /**
     * grpc服务：sayHello测试
     */
    private  void sayHello() {
        showLoading();
        mTestChannel = ManagedChannelBuilder.forAddress(Variable.SERVER_IP, Integer.parseInt(Variable.SERVER_PORT))
                .usePlaintext(true)
                .build();

        Api.Req.Builder requestBuilder = Api.Req.newBuilder();
        requestBuilder.setName("assur");
        Api.Req request = requestBuilder.build();

        protocol.ApiServiceGrpc.ApiServiceBlockingStub stub = protocol.ApiServiceGrpc.newBlockingStub(mTestChannel);
        Api.Resp response = stub.sayHello(request);

        DialogUtils.hintDialog(TestInterfaceActivity.this, "请求结果:\n" + response.getMessage(), "确定");
        hideLoading();
    }


    /**
     * grpc服务：注册
     */
    private void register() {
        if (!isNetConnected()) {
            ToastUtils.showToast(TestInterfaceActivity.this, getResources().getString(R.string.no_net_message));
            return;
        }
        showLoading();
//        HashMap<String, String> map = new HashMap<String, String>();
//        map.put(RegisterTask.Key_Name, "assur");
//        map.put(RegisterTask.Key_Role, "0");
//        map.put(RegisterTask.Key_PassWord, "99999999");
//        map.put(RegisterTask.Key_Phone, "15111260025");
        RegisterTask task = new RegisterTask();
//        task.setMap(map);
        task.setUpdateUIInterface(new RegisterTask.UpdateUIInterface() {
            @Override
            public void onSucceed(Api.RespRegister result) {
                Logger.d("register更新UI：" + result);
                hideLoading();
                DialogUtils.hintDialog(TestInterfaceActivity.this, "请求结果:\n" + result.getUserAddress(), "确定");
            }

            @Override
            public void onError() {
                hideLoading();
                loadErrorDialog();
            }
        });
        task.execute(Variable.SERVER_IP, Variable.SERVER_PORT);

    }

    /**
     * grpc服务：登陆
     */
    private void login() {
        if (!isNetConnected()) {
            ToastUtils.showToast(TestInterfaceActivity.this, getResources().getString(R.string.no_net_message));
            return;
        }
        showLoading();
//        HashMap<String, String> map = new HashMap<String, String>();
//        map.put(RegisterTask.Key_Phone, "11111");
        LoginTask task = new LoginTask();
//        task.setMap(map);
        task.setUpdateUIInterface(new LoginTask.UpdateUIInterface() {
            @Override
            public void onSucceed(Api.RespLogin result) {
                Logger.d("register更新UI：" + result);
                hideLoading();
//                DialogUtils.hintDialog(TestInterfaceActivity.this, "请求结果:\n"
//                        + "状态吗:" + result.getStatusCode() + "\n"
//                        + "用户名:" + result.getName()+ "\n"
//                        + "角色:" + result.getRole()+ "\n"
//                        + "密码:" + result.getPassWord()+ "\n"
//                        + "手机号:" + result.getPhone()+ "\n"
//                        + "地址:" + result.getAddress()+ "\n", "确定");
//
//                UserInfo.USER_NAME = result.getName();
//                UserInfo.USER_ROLE = result.getRole();
//                UserInfo.USER_PHONE = result.getPhone();
//                UserInfo.USER_ADDRESS = result.getAddress();
            }

            @Override
            public void onError() {
                hideLoading();
                loadErrorDialog();
            }
        });
        task.execute(Variable.SERVER_IP, Variable.SERVER_PORT);

    }



    /**
     * grpc服务：管理合约
     */
    private void manageContract() {

    }

    /**
     * grpc服务：查询合约
     */
    private void checkContract() {
//        if (!isNetConnected()) {
//            ToastUtils.showToast(TestInterfaceActivity.this, getResources().getString(R.string.no_net_message));
//            return;
//        }
//        showLoading();
////        HashMap<String, String> map = new HashMap<String, String>();
////        map.put(CheckAccountTask.Key_Address, UserInfo.USER_ADDRESS);
//        CheckContractTask task = new CheckContractTask();
////        task.setMap(map);
//        task.setUpdateUIInterface(new CheckContractTask.UpdateUIInterface() {
//            @Override
//            public void onSucceed(Api.RespCheckContract result) {
//                Logger.d("checkContract更新UI：" + result);
//                hideLoading();
//                DialogUtils.hintDialog(TestInterfaceActivity.this, "请求结果:\n" + result.getHashCode(), "确定");
//            }
//
//            @Override
//            public void onError() {
//                hideLoading();
//                loadErrorDialog();
//            }
//        });
//        task.execute(Variable.SERVER_IP, Variable.SERVER_PORT);
    }


    /**
     * grpc服务：发布排班
     */
    private void setSchedule() {
        showLoading();
        AddScheduleTask task = new AddScheduleTask();
        protocol.Api.ReqScheduling.Builder requestBuilder = protocol.Api.ReqScheduling.newBuilder();
        requestBuilder.setUserAddress(UserInfo.USER_ADDRESS);
        requestBuilder.setPassWord(UserInfo.USER_PASSWORD);
        requestBuilder.setAccountDescribe(UserInfo.USER_KEYSTIRE_FILE_CONTENT);
        requestBuilder.setCompany(Variable.CONPANY_NAME);
        requestBuilder.setPayAccount(Variable.huangAddr);
        requestBuilder.setManagerPayee(Variable.CONPANY_ACCOUNT);
        requestBuilder.setManagerRatio(Variable.CONPANY_RATIO);
        requestBuilder.setStoresNumber(Variable.CONPANY_STORES_NUMBER);
        requestBuilder.setPostscript("备注信息");
        requestBuilder.setMyRatio(100 - Variable.CONPANY_RATIO);
        requestBuilder.setTimeStamp(TimeUtils.getTimeSp("2018-05-01 09:00:00"));
        for (int i = 0; i<3; i++) {

            // 职位人数 与 白名单人数对应，如：3个人，白名单数组中就有3个，可以重复
            int peopleCount = 3;
            final List<String> whiteList = new ArrayList<String>();
            for (int j = 0; j < peopleCount; j++) {
                whiteList.add(Variable.cookerAddr);
            }

            protocol.Obj.Job job =  protocol.Obj.Job.newBuilder()
                    .setRole("职位" + i)
                    .setCount(3)
                    .setRadio(10)
                    .setCompany("公司" + i)
                    .addAllWhiteList(whiteList)
                    .setTimeStamp(TimeUtils.getTimeSp("2018-01-01 00:00:00"))
                    .build();
            requestBuilder.addJobs(job);
        }
        task.setRequest(requestBuilder.build());
        task.setUpdateUIInterface(new AddScheduleTask.UpdateUIInterface() {
            @Override
            public void onSucceed(protocol.Api.RespScheduling result) {
                Logger.d("AddScheduleTask更新UI：" + result);
                hideLoading();
                if(result.getStatusCode() == protocol.StatusOuterClass.Status.Success.ordinal()) {
                    if(result.getJobAddress() == null || result.getJobAddress().equals("")) {
                        hintDialog("添加排班失败，排班地址为空！");
                    } else {
                        showToast("添加排班成功！");
                        finish();
                    }
                } else {
                    // 请求失败
                    loadErrorDialog();
                }
            }

            @Override
            public void onError() {
                // 请求失败
                hideLoading();
                loadErrorDialog();
            }
        });
        task.execute(Variable.SERVER_IP, Variable.SERVER_PORT);

    }

    /**
     * grpc服务：查询排班
     */
    private void getSchedule() {
        showLoading();
        GetScheduleTask task = new GetScheduleTask();
        protocol.Api.ReqGetSchedue.Builder requestBuilder = protocol.Api.ReqGetSchedue.newBuilder();
        requestBuilder.setCompanyName(Variable.CONPANY_NAME);
        requestBuilder.setUserAddress(UserInfo.USER_ADDRESS);
        task.setRequest(requestBuilder.build());
        task.setUpdateUIInterface(new GetScheduleTask.UpdateUIInterface() {
            @Override
            public void onSucceed(protocol.Api.RespGetSchedue result) {
                Logger.d("GetScheduleTask更新UI：" + result);
                hideLoading();
                if(result.getStatusCode() == protocol.StatusOuterClass.Status.Success.ordinal()) {
                    hintDialog(result.toString());
                    //showToast("查询排班成功！");

                } else {
                    // 请求失败
                    loadErrorDialog();
                }
            }

            @Override
            public void onError() {
                // 请求失败
                hideLoading();
                loadErrorDialog();
            }
        });
        task.execute(Variable.SERVER_IP, Variable.SERVER_PORT);
    }

    /**
     * grpc服务：查询申请
     */
    private void getApply() {

    }

    /**
     * grpc服务：申请工作
     */
    private void applyJob() {
        showLoading();
        ApplyJobTask task = new ApplyJobTask();
        protocol.Api.ReqFindJob.Builder requestBuilder = protocol.Api.ReqFindJob.newBuilder();
        requestBuilder.setUserAddress(UserInfo.USER_ADDRESS);
        requestBuilder.setPassWord(UserInfo.USER_PASSWORD);
        requestBuilder.setAccountDescribe(UserInfo.USER_KEYSTIRE_FILE_CONTENT);
        requestBuilder.setMyJob(protocol.Obj.Job.newBuilder()
                .setJobAddress("0x4356258907a9c7C98009CA142B5De630B7A26102")
                .setJobId(1)
                .setRadio(10)
                .setFatherAddr("0x569ef95c3c40d7bfadf53d02edda3afe0c9bb17a")
                .build());
        task.setRequest(requestBuilder.build());
        task.setUpdateUIInterface(new ApplyJobTask.UpdateUIInterface() {
            @Override
            public void onSucceed(protocol.Api.RespFindJob result) {
                Logger.d("ApplyJobTask更新UI：" + result);
                hideLoading();
                if(result.getStatusCode() == protocol.StatusOuterClass.Status.Success.ordinal()) {
                    hintDialog(result.toString());
                    showToast("申请排班成功！");
                } else {
                    // 请求失败
                    loadErrorDialog();
                }
            }

            @Override
            public void onError() {
                // 请求失败
                hideLoading();
                loadErrorDialog();
            }
        });
        task.execute(Variable.SERVER_IP, Variable.SERVER_PORT);

    }

    /**
     * grpc服务：查询某排班下的订单(该排班需要你申请过或者你发布的才有返回)
     */
    private void getOrderByScheduleRequest() {
        showLoading();
        GetAllOrderByScheduleTask task = new GetAllOrderByScheduleTask();
        protocol.Api.ReqGetAllOrder.Builder requestBuilder = protocol.Api.ReqGetAllOrder.newBuilder();
        requestBuilder.setUserAddress(UserInfo.USER_ADDRESS);
        requestBuilder.setJobAddress("0x4B740ee6D118a1F55E9b0080E3858348FE1c85Fe");
        task.setRequest(requestBuilder.build());
        task.setUpdateUIInterface(new GetAllOrderByScheduleTask.UpdateUIInterface() {
            @Override
            public void onSucceed(protocol.Api.RespGetAllOrder result) {
                Logger.d("GetAllOrderByScheduleTask更新UI：" + result);
                hideLoading();
                if(result.getStatusCode() == protocol.StatusOuterClass.Status.Success.ordinal()) {
                    hintDialog(result.getOrdersList().toString());
                } else {
                    // 请求失败
                    loadErrorDialog();
                }
            }

            @Override
            public void onError() {
                // 请求失败
                hideLoading();
                loadErrorDialog();
            }
        });
        task.execute(Variable.SERVER_IP, Variable.SERVER_PORT);
    }

    /**
     * grpc服务：申请付款（接口已取消）
     */
    private void addOrderAndPay() {
        showLoading();
        AddOrderAndPayTask task = new AddOrderAndPayTask();
        protocol.Api.ReqPay.Builder requestBuilder = protocol.Api.ReqPay.newBuilder();
        protocol.Obj.Order order = protocol.Obj.Order.newBuilder()
                .setTable(1)
                .setMoney(100)
                .setTimeStamp(TimeUtils.getTimeSp("2018-01-01 00:00:00"))
                .build();
        requestBuilder.setContent(order);
        requestBuilder.setJobAddress("0x4356258907a9c7C98009CA142B5De630B7A26102");
        requestBuilder.setMoney(100);
        requestBuilder.setPassWord(UserInfo.USER_PASSWORD);
        requestBuilder.setAccountDescribe(UserInfo.USER_KEYSTIRE_FILE_CONTENT);
        requestBuilder.setUserAddress(UserInfo.USER_ADDRESS);

        task.setRequest(requestBuilder.build());
        task.setUpdateUIInterface(new AddOrderAndPayTask.UpdateUIInterface() {
            @Override
            public void onSucceed(protocol.Api.RespPay result) {
                Logger.d("GetScheduleTask更新UI：" + result);
                hideLoading();
                if(result.getStatusCode() == protocol.StatusOuterClass.Status.Success.ordinal()) {
                    showToast("添加订单成功");
                } else {
                    // 请求失败
                    loadErrorDialog();
                }
            }

            @Override
            public void onError() {
                // 请求失败
                hideLoading();
                loadErrorDialog();
            }
        });
        task.execute(Variable.SERVER_IP, Variable.SERVER_PORT);
    }

    /**
     * grpc服务：查询账户
     */
    private void checkAccount() {
        if (!isNetConnected()) {
            ToastUtils.showToast(TestInterfaceActivity.this, getResources().getString(R.string.no_net_message));
            return;
        }
        showLoading();
        CheckAccountTask task = new CheckAccountTask();
        Api.ReqCheckAccount.Builder requestBuilder = Api.ReqCheckAccount.newBuilder();
        requestBuilder.setUserAddress(UserInfo.USER_ADDRESS);
        task.setRequest(requestBuilder.build());
        task.setUpdateUIInterface(new CheckAccountTask.UpdateUIInterface() {
            @Override
            public void onSucceed(Api.RespCheckAccount result) {
                Logger.d("checkAccount更新UI：" + result);
                hideLoading();
                DialogUtils.hintDialog(TestInterfaceActivity.this, "请求结果:\n"
                        + result.getName() + "\n"
                        + result.getPassWord() + "\n"
                        + result.getRole() + "\n"
                        + result.getAccountDescribe() + "\n"
                        + result.getPhone(), "确定");
            }

            @Override
            public void onError() {
                hideLoading();
                loadErrorDialog();
            }
        });
        task.execute(Variable.SERVER_IP, Variable.SERVER_PORT);
    }



    /**
     * grpc服务：提交订单
     */
    private void setContent() {

    }

    /**
     * grpc服务：查询订单
     */
    private void getContent() {

    }

    /**
     * grpc服务：查询某一班的收入
     */
    private void getBalance() {

    }

    /**
     * grpc服务：检查是否申请工作成功
     */
    private void checkIsOkApplication() {

    }



}
