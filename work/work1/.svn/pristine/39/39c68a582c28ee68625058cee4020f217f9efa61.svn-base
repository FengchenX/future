package com.znfz.assur.znfz_android.android.main;

import android.Manifest;
import android.content.pm.PackageManager;
import android.support.v4.app.ActivityCompat;

import com.orhanobut.logger.Logger;
import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.app.BaseActivity;
import com.znfz.assur.znfz_android.android.common.global.UserInfo;
import com.znfz.assur.znfz_android.android.common.global.Variable;
import com.znfz.assur.znfz_android.android.common.task.BandTask;
import com.znfz.assur.znfz_android.android.common.task.GetBandTask;
import com.znfz.assur.znfz_android.android.main.applicant.activity.ApplicantHomeActivity;
import com.znfz.assur.znfz_android.android.main.publisher.activity.PublisherHomeActivity;
import com.znfz.assur.znfz_android.android.test.grpc_interface_test.TestInterfaceActivity;

import org.web3j.crypto.Credentials;
import org.web3j.crypto.WalletUtils;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.util.List;


/**
 * 主界面
 * Created by assur on 2018/4/20.
 */
public class MainActivity extends BaseActivity {

     private Credentials credentials;


    @Override
    protected void initView() {

    }

    @Override
    protected void initData() {

        // 获取权限
        getPermission();

        // ==============测试阶段，换设备使用同一个账户，需要将次账户写入===============
        if(spUtils.getString(Variable.SP_KEY_USER_ADDRESS) == null || spUtils.getString(Variable.SP_KEY_USER_ADDRESS).length() == 0) {
            try{
                File keystoreFile = new File(Variable.PATH_WEB3J_USER_INFO, "UTC--2018-05-19T11-24-09.082--569ef95c3c40d7bfadf53d02edda3afe0c9bb17a.json");
                FileOutputStream outStream = new FileOutputStream(keystoreFile);
                outStream.write(Variable.storeFileContent.getBytes());
                outStream.close();
                spUtils.putString(Variable.SP_KEY_USER_KEYSTORE_FILENAME, keystoreFile.getName());
                try {
                    credentials = WalletUtils.loadCredentials(UserInfo.USER_PASSWORD, keystoreFile.getAbsolutePath());
                    Logger.e("credentials.getAddress()(assur) :" + credentials.getAddress());
                    // 这里耗时比较久，后面改成创建用户的时候就保存用户地址等信息，不用每次查询这么久
                    UserInfo.USER_ADDRESS = credentials.getAddress();
                    spUtils.putString(Variable.SP_KEY_USER_ADDRESS, UserInfo.USER_ADDRESS);

                    // 判断是否绑定
                    getBingRequest();

                } catch (Exception e1) {
                    Logger.e("e1：" + e1);
                }

            } catch (Exception e) {
                e.printStackTrace();
            }
        } else {
            startActivity(PublisherHomeActivity.class);
            finish();
        }

        // ======================================================================










        String keystoreFileFullPath  = Variable.PATH_WEB3J_USER_INFO + "/" + spUtils.getString(Variable.SP_KEY_USER_KEYSTORE_FILENAME);
        String content = readKeystoreFile(keystoreFileFullPath);
        UserInfo.USER_KEYSTIRE_FILE_CONTENT = content;
        Logger.e("UserInfo.USER_KEYSTIRE_FILE_CONTENT:" + UserInfo.USER_KEYSTIRE_FILE_CONTENT);


        UserInfo.USER_ADDRESS = spUtils.getString(Variable.SP_KEY_USER_ADDRESS);
        Logger.e("UserInfo.USER_ADDRESS:" + UserInfo.USER_ADDRESS);
        if(UserInfo.USER_ADDRESS != null &&  UserInfo.USER_ADDRESS.length() > 0) {
//            // 进入主页
//            startActivity(TestInterfaceActivity.class);
//            finish();
//            return;
        } else {
            showLoading();
            new Thread(new Runnable() {
                @Override
                public void run() {
                    checkUser();
                }
            }).start();

        }



    }

    /**
     * 检查是否存在用户
     *
     * 判断是否存在已经创建的用户keystore文件，没有创建则跳转到创建界面
     *
     */
    private void checkUser() {

        String keyStoreFileName = spUtils.getString(Variable.SP_KEY_USER_KEYSTORE_FILENAME);
        if(keyStoreFileName == null || keyStoreFileName.length() == 0) {
            runOnUiThread(new Runnable() {
                @Override
                public void run() {
                    hideLoading();
                    // 创建用户
                    startActivity(CreateUserActivity.class);
                    finish();
                }
            });
        } else {
            String keystoreFileFullPath = Variable.PATH_WEB3J_USER_INFO + "/" + spUtils.getString(Variable.SP_KEY_USER_KEYSTORE_FILENAME);
            Logger.e("keystoreFileFullPath(assur) :" + keystoreFileFullPath);
            if(keystoreFileFullPath != null && keystoreFileFullPath.length() > 0) {

                if(new File(keystoreFileFullPath) != null) {
                    Logger.e("keystoreFileName(:keystoreFileName文件经存在，直接进入主页");
                    try {
                        final Credentials credentials = WalletUtils.loadCredentials(UserInfo.USER_PASSWORD, keystoreFileFullPath);
                        Logger.e("credentials.getAddress()(assur) :" + credentials.getAddress());
                        // 这里耗时比较久，后面改成创建用户的时候就保存用户地址等信息，不用每次查询这么久
                        UserInfo.USER_ADDRESS = credentials.getAddress();
                        spUtils.putString(Variable.SP_KEY_USER_ADDRESS, UserInfo.USER_ADDRESS);

                    } catch (Exception e1) {
                        Logger.e("e1：" + e1);
                    }

                    runOnUiThread(new Runnable() {
                        @Override
                        public void run() {
                            hideLoading();
                            // 进入主页
                            startActivity(ApplicantHomeActivity.class);
                            finish();
                        }
                    });

                } else {
                    Logger.e("keystoreFileName:您的账户文件已丢失，请重新创建，或者找回账户");
                    showToast("您的账户文件已丢失，请重新创建，或者找回账户");
                    runOnUiThread(new Runnable() {
                        @Override
                        public void run() {
                            hideLoading();
                            // 创建用户
                            startActivity(CreateUserActivity.class);
                            finish();
                        }
                    });

                }

            } else {
                Logger.e("keystoreFileName :keystoreFileName文件不存在，进入创建界面");
                runOnUiThread(new Runnable() {
                    @Override
                    public void run() {
                        hideLoading();
                        // 创建用户
                        startActivity(CreateUserActivity.class);
                        finish();
                    }
                });

            }
        }


    }


    private void getPermission() {
        int REQUEST_EXTERNAL_STORAGE = 1;
        String[] PERMISSIONS_STORAGE = {
                Manifest.permission.READ_EXTERNAL_STORAGE,
                Manifest.permission.WRITE_EXTERNAL_STORAGE
        };
        int permission = ActivityCompat.checkSelfPermission(MainActivity.this, Manifest.permission.WRITE_EXTERNAL_STORAGE);

        if (permission != PackageManager.PERMISSION_GRANTED) {
            // We don't have permission so prompt the user
            ActivityCompat.requestPermissions(
                    MainActivity.this,
                    PERMISSIONS_STORAGE,
                    REQUEST_EXTERNAL_STORAGE
            );
        }
    }

    private String readKeystoreFile(String fileFullPath) {
        File file = new File(fileFullPath);
        String str = null;
        StringBuffer stringBuffer = new StringBuffer();
        try {
            InputStream is = new FileInputStream(file);
            InputStreamReader input = new InputStreamReader(is, "UTF-8");
            BufferedReader reader = new BufferedReader(input);
            while ((str = reader.readLine()) != null) {
                stringBuffer.append(str);
            }

        } catch (FileNotFoundException e) {
            // TODO Auto-generated catch block
            e.printStackTrace();
        } catch (IOException e) {
            // TODO Auto-generated catch block
            e.printStackTrace();
        }
        return stringBuffer.toString();
    }


    /**
     * grpc服务：用户绑定
     */
    private void bandRequest(final Credentials credentials) {
//        if (!isNetConnected()) {
//            // 网络没有连接
//            toast(R.string.no_net_message);
//            return;
//        }
        showLoading();
        final String name = "assur";
        final String phone = "15111260025";
        final String password = UserInfo.USER_PASSWORD;
        final String address = credentials.getAddress();
        BandTask task = new BandTask();
        protocol.Api.ReqBand.Builder requestBuilder = protocol.Api.ReqBand.newBuilder();
        requestBuilder.setName(name);
        requestBuilder.setPhone(phone);
        requestBuilder.setPassWord(password);
        requestBuilder.setUserAddress(address);
        requestBuilder.setAccountDescribe(Variable.storeFileContent);
        task.setRequest(requestBuilder.build());
        task.setUpdateUIInterface(new BandTask.UpdateUIInterface() {
            @Override
            public void onSucceed(protocol.Api.RespBand result) {
                Logger.d("BandTask更新UI：" + result);
                hideLoading();
                if(result.getStatusCode() == protocol.StatusOuterClass.Status.Success.ordinal()) {
                    UserInfo.USER_NAME = name;
                    UserInfo.USER_PHONE = phone;
                    UserInfo.USER_ADDRESS = address;
                    UserInfo.USER_PASSWORD = password;

                    startActivity(TestInterfaceActivity.class);
                    finish();

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
     * grpc服务：用户是否绑定（需要一段时间才能真正绑定成功）
     */
    private void getBingRequest() {
        showLoading();
        GetBandTask task = new GetBandTask();
        protocol.Api.ReqLogin.Builder requestBuilder = protocol.Api.ReqLogin.newBuilder();
        requestBuilder.setPhone("15111260025");
        task.setRequest(requestBuilder.build());
        task.setUpdateUIInterface(new GetBandTask.UpdateUIInterface() {
            @Override
            public void onSucceed(protocol.Api.RespLogin result) {
                Logger.d("GetBandTask更新UI：" + result);
                hideLoading();
                if(result.getStatusCode() == protocol.StatusOuterClass.Status.Success.ordinal()) {

                    showToast("已绑定");
                    // 已经绑定
                    startActivity(PublisherHomeActivity.class);
                    finish();

                } else {
                    // 请求失败
                    //loadErrorDialog();
                    showToast("未绑定，进行绑定请求");

                    // 绑定用户，注意只要一次就行
                    if(credentials != null) {
                        bandRequest(credentials);
                    }

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
