package com.znfz.assur.znfz_android.android.common.global;

import android.os.Environment;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Variable {

    public static final String SERVER_IP = "39.108.80.66";   // grpc服务IP（线上）
    public static final String SERVER_PORT  = "8899";        // grpc服务端口（线上）

//    public static final String SERVER_IP = "192.168.83.200";   // grpc服务IP（梁思昊本地）
//    public static final String SERVER_PORT  = "8899";          // grpc服务端口（梁思昊本地）

//    public static final String SERVER_IP = "192.168.83.56";   // grpc服务IP（辛冰本地）
//    public static final String SERVER_PORT  = "65535";        // grpc服务端口（辛冰本地）


    public static final String URL_WEB3J = "http://39.108.80.66:8546/your api-key";  // 以太坊节点地址
    public static final String PATH_WEB3J_USER_INFO = Environment.getExternalStorageDirectory().toString() + "/weh3j_user";




    // SharePreference key
    public static final String SP_KEY_USER_KEYSTORE_FILENAME = "SP_KEY_USER_KEYSTORE_FILENAME"; // 本地保存已经创建的用户的keystore全路径+文件名
    public static final String SP_KEY_USER_ADDRESS = "SP_KEY_USER_ADDRESS";   // 本地保存用户地址的key
    public static final String SP_KEY_USER_TYPE = "SP_KEY_USER_TYPE";         // 本地保存用户类型("publisher"/"applicant")的key
    public static final String SP_VALUE_USER_TYPE_PUBLISHER = "SP_VALUE_USER_TYPE_PUBLISHER";  // 本地保存用户类型的值（发布者）
    public static final String SP_VALUE_USER_TYPE_APPLICANT = "SP_VALUE_USER_TYPE_PUBLISHER";  // 本地保存用户类型的值（应聘者者）

    // SharePreference value




    //bundle
    public static final String KEY_SCHEDULING_ADDRESS = "KEY_SCHEDULING_ADDRESS";       // 排班地址key



    // 角色类型
    public static final String USER_ROLE_NONE = "0";                // 无角色
    public static final String USER_ROLE_PUBLISHER_MANAGER = "1";   // 发布者（经理）
    public static final String USER_ROLE_APPLICANT_COOK = "2";      // 申请者（厨师）
    public static final String USER_ROLE_APPLICANT_WAITER = "3";    // 申请者（服务员）





    // 第一期，写死公司名称/地址，只有一个公司
    public static final String CONPANY_NAME = "正稻"; // 公司账户名称
    public static final String CONPANY_ACCOUNT = "0x373356A20d4992277b610502102875506008D838"; // 公司账户地址（管理费的收款地址）
    public static final String CONPANY_STORES_NUMBER = "正稻001";  // 门店编号
    public static final int CONPANY_RATIO  = 12;  // 公司管理费的固定比例


    // 第一期，测试账户
    public static final String managerAddr = "0x3a32e7d31b1418b0468505d2c9b1892053159bc4";
    public static final String cookerAddr = "0x07f4370Fb847bc72Fcb9a131EE9daC2956E169F4";
    public static final String huangAddr = "0x759D4c2E15587Fae036f183202F36CA3C667ccbD";


    // ===========临时，换设备都用这个账户==============
    public static final String keyStoreFileFullPath = Environment.getExternalStorageDirectory().toString()
            + "/weh3j_user_test"
            + "/UTC--2018-05-19T11-24-09.082--569ef95c3c40d7bfadf53d02edda3afe0c9bb17a.json";
    public static final String storeFileContent = "{\"address\":\"569ef95c3c40d7bfadf53d02edda3afe0c9bb17a\",\"id\":\"6f996145-6b4b-455d-bf99-16089b7f6509\",\"version\":3,\"crypto\":{\"cipher\":\"aes-128-ctr\",\"cipherparams\":{\"iv\":\"21d3b415546d67309b02799fb488f98c\"},\"ciphertext\":\"fc10d4bc320a91742e99c85fc041a211db9bf40dd407b6611d7f57ff156b545b\",\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":4096,\"p\":6,\"r\":8,\"salt\":\"4cb53365095e37064ed94c788fcb5ffceb78e84dad92e64d3e1b29cf1b0f2adb\"},\"mac\":\"5a7bfbe500a30cd5315cc34f14b1d6460018f1ff5eab92ce07b9598839b17fa3\"}}";
    public static final String storeAddress = "569ef95c3c40d7bfadf53d02edda3afe0c9bb17a";

    // ===========临时，换设备都用这个账户==============

}
