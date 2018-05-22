package com.znfz.assur.znfz_android.android.main.publisher.bean;

/**
 * Created by dansihan on 2018/5/21.
 */

public class PublisherAddEmployeeBean {
    private String name;  //职位名称
    private String JobPersonNumber; //招聘人数
    private String CommissionRate;  //佣金比例
    private String friendname; //好友名字


    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getJobPersonNumber() {
        return JobPersonNumber;
    }

    public void setJobPersonNumber(String jobPersonNumber) {
        JobPersonNumber = jobPersonNumber;
    }

    public String getCommissionRate() {
        return CommissionRate;
    }

    public void setCommissionRate(String commissionRate) {
        CommissionRate = commissionRate;
    }

    public String getFriendname() {
        return friendname;
    }

    public void setFriendname(String friendname) {
        this.friendname = friendname;
    }
}
