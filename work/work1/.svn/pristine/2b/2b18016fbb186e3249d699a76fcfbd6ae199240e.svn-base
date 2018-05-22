package com.znfz.assur.znfz_android.android.main;

import android.widget.EditText;
import com.orhanobut.logger.Logger;
import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.app.BaseActivity;
import com.znfz.assur.znfz_android.android.common.global.UserInfo;
import com.znfz.assur.znfz_android.android.common.global.Variable;
import com.znfz.assur.znfz_android.android.common.task.LoginTask;
import com.znfz.assur.znfz_android.android.main.applicant.activity.ApplicantHomeActivity;
import com.znfz.assur.znfz_android.android.main.publisher.activity.PublisherHomeActivity;
import butterknife.BindView;
import butterknife.ButterKnife;
import butterknife.OnClick;
import protocol.StatusOuterClass;


/**
 * 登陆界面
 * Created by assur on 2018/4/23.
 */
public class LoginActivity extends BaseActivity {


    @BindView(R.id.login_et_userphone)
    EditText etUserPhone;

    @Override
    protected void initView() {
        setContentView(R.layout.activity_login);
        ButterKnife.bind(this);
    }

    @Override
    protected void initData() {

        //86222222（厨师）
        //15219438281（经理）
        etUserPhone.setText("15219438281");
    }

    @OnClick(R.id.login_btn_login)
    public void loginBtnClick() {
        login();
    }


    /**
     * grpc服务：登陆
     */
    private void login() {
        if (!isNetConnected()) {
            // 网络没有连接
            toast(R.string.no_net_message);
            return;
        }
        if(etUserPhone.getText().toString() == null || etUserPhone.getText().toString().length() == 0) {
            // 电话号码为空
            hintDialog(R.string.login_phone_empty_hint);
            return;
        }
        showLoading();
        LoginTask task = new LoginTask();
        protocol.Api.ReqLogin.Builder requestBuilder = protocol.Api.ReqLogin.newBuilder();
        requestBuilder.setPhone(etUserPhone.getText().toString());
        task.setRequest(requestBuilder.build());
        task.setUpdateUIInterface(new LoginTask.UpdateUIInterface() {
            @Override
            public void onSucceed(protocol.Api.RespLogin result) {
                Logger.d("LoginTask更新UI：" + result);
                hideLoading();
                if(result.getStatusCode() == StatusOuterClass.Status.Success.ordinal()) {
//                    UserInfo.USER_NAME = result.getName();
//                    UserInfo.USER_ROLE = result.getRole();
//                    UserInfo.USER_PHONE = result.getPhone();
//                    UserInfo.USER_ADDRESS = result.getAddress();
//                    UserInfo.USER_PASSWORD = result.getPassWord();
//                    UserInfo.USER_PRIVATE_KEY = result.getAccountDescribe();
//                    if(UserInfo.USER_ROLE.equals(Variable.USER_ROLE_NONE)) {
//                        // 没有用户角色（不存在）
//                        hintDialog(R.string.login_no_role_hint);
//                    } else {
//                        if(UserInfo.USER_ROLE.equals(Variable.USER_ROLE_PUBLISHER_MANAGER)) {
//                            // 进入发布者主页
//                            startActivity(PublisherHomeActivity.class);
//                        } else {
//                            // 进入申请者主页
//                            startActivity(ApplicantHomeActivity.class);
//                        }
//                    }

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
}
