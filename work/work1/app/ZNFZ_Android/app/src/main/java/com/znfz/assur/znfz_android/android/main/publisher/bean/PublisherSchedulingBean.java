package com.znfz.assur.znfz_android.android.main.publisher.bean;

import java.util.ArrayList;
import java.util.List;

/**
 * 排班
 * Created by assur on 2018/4/25
 */
public class PublisherSchedulingBean {

    private String scheduleAddress;       // 排班地址（区块链）
    private String scheduleFatherAddress; // 排班发布者地址（区块链）
    private String scheduleCompany; // 排班公司名称
    private long scheduleCompanyPer; // 排班公司管理费百分比
    private long scheduleStorePer;   // 排班门店百分比
    private Long scheduleTimeSp;    // 排班开始时间
    private boolean isCurSchedule;  // 是否当前排班
    private boolean isShowRoleList;  // 是否显示角色详情
    private List<PublisherRoleBean> scheduleRoleList; // 排班角色列表


    public PublisherSchedulingBean() {
        scheduleAddress = "";
        scheduleFatherAddress = "";
        scheduleCompany = "";
        scheduleCompanyPer = 0;
        scheduleStorePer = 0;
        scheduleTimeSp = new Long(0);
        isCurSchedule = false;
        isShowRoleList = false;
        scheduleRoleList = new ArrayList<>();
    }

    public String getScheduleAddress() {
        return scheduleAddress;
    }

    public void setScheduleAddress(String scheduleAddress) {
        this.scheduleAddress = scheduleAddress;
    }

    public String getScheduleFatherAddress() {
        return scheduleFatherAddress;
    }

    public void setScheduleFatherAddress(String scheduleFatherAddress) {
        this.scheduleFatherAddress = scheduleFatherAddress;
    }

    public long getScheduleCompanyPer() {
        return scheduleCompanyPer;
    }

    public void setScheduleCompanyPer(long scheduleCompanyPer) {
        this.scheduleCompanyPer = scheduleCompanyPer;
    }

    public long getScheduleStorePer() {
        return scheduleStorePer;
    }

    public void setScheduleStorePer(long scheduleStorePer) {
        this.scheduleStorePer = scheduleStorePer;
    }

    public Long getScheduleTimeSp() {
        return scheduleTimeSp;
    }

    public void setScheduleTimeSp(Long scheduleTimeSp) {
        this.scheduleTimeSp = scheduleTimeSp;
    }

    public String getScheduleCompany() {
        return scheduleCompany;
    }

    public void setScheduleCompany(String scheduleCompany) {
        this.scheduleCompany = scheduleCompany;
    }

    public boolean isCurSchedule() {
        return isCurSchedule;
    }

    public void setCurSchedule(boolean curSchedule) {
        isCurSchedule = curSchedule;
    }

    public List<PublisherRoleBean> getScheduleRoleList() {
        return scheduleRoleList;
    }

    public void setScheduleRoleList(List<PublisherRoleBean> scheduleRoleList) {
        this.scheduleRoleList = scheduleRoleList;
    }

    public boolean isShowRoleList() {
        return isShowRoleList;
    }

    public void setShowRoleList(boolean showRoleList) {
        isShowRoleList = showRoleList;
    }

    @Override
    public String toString() {
        return "PublisherSchedulingBean{" +
                "scheduleAddress='" + scheduleAddress + '\'' +
                ", scheduleFatherAddress='" + scheduleFatherAddress + '\'' +
                ", scheduleCompany='" + scheduleCompany + '\'' +
                ", scheduleCompanyPer=" + scheduleCompanyPer +
                ", scheduleStorePer=" + scheduleStorePer +
                ", scheduleTimeSp=" + scheduleTimeSp +
                ", isCurSchedule=" + isCurSchedule +
                ", isShowRoleList=" + isShowRoleList +
                ", scheduleRoleList=" + scheduleRoleList +
                '}';
    }
}
