package com.znfz.assur.znfz_android.android.common.global;


import android.os.Environment;

import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherSchedulingBean;

import java.util.List;

public class UserInfo {


    // 当前用户信息
    public static String USER_KEYSTIRE_FILE_CONTENT;  // 读取的用户keystore文件内容
    public static String USER_NAME;
    public static String USER_ROLE;
    public static String USER_PHONE;
    public static String USER_PASSWORD = "123456";    // 注册用户二期写死
    public static String USER_ADDRESS;
    public static String USER_PRIVATE_KEY;


    // 当前是发布者/申请者切换标志(默认0是申请者，其他值是发布者)
    public static int ROLE_STATE_APPLACATION = 0;
    public static int ROLE_STATE_PUBLISH = 1;
    public static int roleState = ROLE_STATE_APPLACATION;


    // 当前排班信息（最后发布的排班作为当前排班，暂时存在这里）
    public static PublisherSchedulingBean currentSchedule;

    // 正在查看历史订单所在的排班信息（暂时存在这里）
    public static PublisherSchedulingBean currentHistorySchedule;

    // 某公司下的所有排版（暂时保存在这里）
    public static List<PublisherSchedulingBean> scheduleistItemsOfOneCompany;

}
