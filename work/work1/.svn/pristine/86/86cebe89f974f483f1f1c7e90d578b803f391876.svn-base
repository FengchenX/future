package com.znfz.assur.znfz_android.android.main;

import android.widget.Button;
import com.orhanobut.logger.Logger;
import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.app.BaseActivity;
import com.znfz.assur.znfz_android.android.common.global.Variable;
import com.znfz.assur.znfz_android.android.common.task.RegisterTask;
import com.znfz.assur.znfz_android.android.main.publisher.activity.PublisherHomeActivity;
import java.util.HashMap;
import butterknife.BindView;
import butterknife.ButterKnife;
import butterknife.OnClick;
import protocol.Api;

/**
 * 注册界面
 * Created by assur on 2018/4/20.
 */
public class RegisterActivity extends BaseActivity {


    @BindView(R.id.register_btn_register)  // 注册按钮
    Button btnRegister;



    @Override
    protected void initView() {
        setContentView(R.layout.activity_register);
        ButterKnife.bind(this);
        initToolBar(getResources().getString(R.string.register_page_title), false);

    }

    @Override
    protected void initData() {

    }


    @OnClick(R.id.register_btn_register)
    public void btnClick() {
        //register();
        startActivity(PublisherHomeActivity.class);
    }




    /**
     * grpc服务：注册
     */
    private void register() {
        if(!isNetConnected()) {
            toast(R.string.no_net_message);
            return;
        }
        showLoading();
        HashMap<String, String> map = new HashMap<String, String>();
//        map.put(RegisterTask.Key_Name, "assur");
//        map.put(RegisterTask.Key_Role, "234567890345678");
//        map.put(RegisterTask.Key_PassWord, "99999999");
//        map.put(RegisterTask.Key_Phone, "15111260025");
        RegisterTask task = new RegisterTask();
//        task.setMap(map);
        task.setUpdateUIInterface(new RegisterTask.UpdateUIInterface() {
            @Override
            public void onSucceed(Api.RespRegister result) {
                Logger.d("register更新UI：" + result);
                hideLoading();
            }

            @Override
            public void onError() {
                hideLoading();
                toast(R.string.load_fail);
            }
        });
        task.execute(Variable.SERVER_IP, Variable.SERVER_PORT);

    }

}
