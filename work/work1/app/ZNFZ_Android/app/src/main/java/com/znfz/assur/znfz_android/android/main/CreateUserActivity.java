package com.znfz.assur.znfz_android.android.main;

import android.Manifest;
import android.content.pm.PackageManager;
import android.support.v4.app.ActivityCompat;
import android.view.View;
import android.widget.Button;
import android.widget.TextView;
import com.orhanobut.logger.Logger;
import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.app.BaseActivity;
import com.znfz.assur.znfz_android.android.common.global.UserInfo;
import com.znfz.assur.znfz_android.android.common.global.Variable;
import com.znfz.assur.znfz_android.android.main.applicant.activity.ApplicantHomeActivity;
import org.web3j.crypto.Credentials;
import org.web3j.crypto.WalletUtils;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.Web3jFactory;
import org.web3j.protocol.http.HttpService;
import java.io.File;
import butterknife.BindView;
import butterknife.ButterKnife;

/**
 * 创建用户界面
 */
public class CreateUserActivity extends BaseActivity {



    @BindView(R.id.create_user_tv_hint)
    TextView tvHint;
    @BindView(R.id.create_user_tv_address)
    TextView tvAddress;
    @BindView(R.id.create_user_btn_create)
    Button btnCreate;

    // 按钮tag定义
    final int BUTTON_TAG_CREATE = 0;
    final int BUTTON_TAG_HOME = 1;

    private String keystoreFileName = "";


    @Override
    protected void initView() {
        setContentView(R.layout.activity_create_user);
        ButterKnife.bind(this);

        btnCreate.setTag(BUTTON_TAG_CREATE);
        btnCreate.setOnClickListener(this);
    }


    @Override
    protected void initData() {


    }


    @Override
    public void onClick(View v) {
        super.onClick(v);

        switch (v.getId()) {
            case R.id.create_user_btn_create:
                if(btnCreate.getTag().toString().endsWith(BUTTON_TAG_CREATE + "")) {
                    createUser();
                } else {
                    enterHomePage();
                }

                break;

            default:
                break;
        }
    }

    /**
     * 创建用户
     */
    private void createUser() {
        showLoading();
        String url = Variable.URL_WEB3J;
        final Web3j web3 = Web3jFactory.build(new HttpService(url));
        new Thread(new Runnable() {
            @Override
            public void run() {
                try {
                    // 创建钱包1：获取权限
                    int REQUEST_EXTERNAL_STORAGE = 1;
                    String[] PERMISSIONS_STORAGE = {
                            Manifest.permission.READ_EXTERNAL_STORAGE,
                            Manifest.permission.WRITE_EXTERNAL_STORAGE
                    };
                    int permission = ActivityCompat.checkSelfPermission(CreateUserActivity.this, Manifest.permission.WRITE_EXTERNAL_STORAGE);

                    if (permission != PackageManager.PERMISSION_GRANTED) {
                        // We don't have permission so prompt the user
                        ActivityCompat.requestPermissions(
                                CreateUserActivity.this,
                                PERMISSIONS_STORAGE,
                                REQUEST_EXTERNAL_STORAGE
                        );
                    }
                    // 创建钱包2：创建
                    final String filePath = Variable.PATH_WEB3J_USER_INFO;
                    File file = new File(filePath);
                    if (!file.exists()) {
                        file.mkdir();
                    }
                    final String password = UserInfo.USER_PASSWORD; // 二期密码默认写死
                    // String fileName = WalletUtils.generateFullNewWalletFile(password, file); // 模拟器会报内存溢出异常
                    final String fileName = WalletUtils.generateNewWalletFile(password, new File(filePath), false);// true:模拟器会报内存溢出异常
                    Logger.e("fileName(assur) :" + fileName);
                    keystoreFileName = fileName;
                    final Credentials credentials = WalletUtils.loadCredentials(password, filePath+"/"+fileName);

                    runOnUiThread(new Runnable() {
                        @Override
                        public void run() {
                            hideLoading();
                            // 创建用户成功
                            if(credentials.getAddress() != null && credentials.getAddress().length() > 0) {
                                createUserSucceed(credentials);
                            }
                        }
                    });

                } catch (Exception e1) {
                    Logger.e("e1(assur) :" + e1.toString());
                    showToast("生成分账系统身份失败:" + e1.toString());
                    hideLoading();
                }
            }
        }).start();
    }


//    /**
//     * grpc服务：绑定用户
//     */
//    private void bandRequest(final Credentials credentials) {
//        if (!isNetConnected()) {
//            // 网络没有连接
//            toast(R.string.no_net_message);
//            return;
//        }
//        showLoading();
//        final String name = "assur";
//        final String phone = "15111260025";
//        final String password = "123456";
//        final String address = credentials.getAddress();
//        final String privateKey = credentials.getEcKeyPair().getPrivateKey() + "";
//        BandTask task = new BandTask();
//        protocol.Api.ReqBand.Builder requestBuilder = protocol.Api.ReqBand.newBuilder();
//        requestBuilder.setName(name);
//        requestBuilder.setPhone(phone);
//        requestBuilder.setPassWord(password);
//        requestBuilder.setUserAddress(address);
//        requestBuilder.setAccountDescribe(privateKey);
//        task.setRequest(requestBuilder.build());
//        task.setUpdateUIInterface(new BandTask.UpdateUIInterface() {
//            @Override
//            public void onSucceed(protocol.Api.RespBand result) {
//                Logger.d("BandTask更新UI：" + result);
//                hideLoading();
//                if(result.getStatusCode() == protocol.StatusOuterClass.Status.Success.ordinal()) {
//                    UserInfo.USER_NAME = name;
//                    UserInfo.USER_PHONE = phone;
//                    UserInfo.USER_ADDRESS = address;
//                    UserInfo.USER_PASSWORD = password;
//                    UserInfo.USER_PRIVATE_KEY = privateKey;
//
//                    // 绑定成功， 更新界面
//                    createUserSucceed(credentials);
//
//                } else {
//                    // 请求失败
//                    loadErrorDialog();
//                }
//            }
//
//            @Override
//            public void onError() {
//                // 请求失败
//                hideLoading();
//                loadErrorDialog();
//            }
//        });
//        task.execute(Variable.SERVER_IP, Variable.SERVER_PORT);
//    }


    /**
     * 成功创建用户
     */
    private void createUserSucceed(final Credentials credentials) {
        Logger.e("credentials.getAddress()(assur) :" + credentials.getAddress());
        Logger.e("credentials.getEcKeyPair().getPrivateKey()(assur) :" + credentials.getEcKeyPair().getPrivateKey().toString());
        Logger.e("credentials.getEcKeyPair().getPublicKey()(assur) :" + credentials.getEcKeyPair().getPublicKey().toString());

        tvHint.setText("您的分账系统系统身份已生成");
        tvAddress.setText(credentials.getAddress());
        tvAddress.setVisibility(View.VISIBLE);
        btnCreate.setText("进入主页");
        btnCreate.setTag(BUTTON_TAG_HOME);
        spUtils.putString(Variable.SP_KEY_USER_KEYSTORE_FILENAME, keystoreFileName);
        UserInfo.USER_ADDRESS = credentials.getAddress();
        spUtils.putString(Variable.SP_KEY_USER_ADDRESS, UserInfo.USER_ADDRESS);
        Logger.e("keystoreFileName(assur) :" + spUtils.getString(Variable.SP_KEY_USER_KEYSTORE_FILENAME));
    }




    /**
     * 进入主页
     */
    public void enterHomePage() {
        // 进入主页
        startActivity(ApplicantHomeActivity.class);
    }


}
