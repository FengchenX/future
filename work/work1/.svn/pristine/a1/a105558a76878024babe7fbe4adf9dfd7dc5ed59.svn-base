package com.znfz.assur.znfz_android.android.main.publisher.bean;

import java.util.ArrayList;
import java.util.List;

/**
 * 订单
 * Created by assur on 2018/4/20
 */
public class PublisherOrderBean {

    private long orderTimeSp;      // 订单时间
    private String orderPrice;     // 订单价格
    private String orderDeskNum;   // 桌号
    private List<PublisherRoleBean> orderRoleList;  // 角色分配列表
    private boolean orderShowAllocation;  // 是否显示佣分配详情


    public PublisherOrderBean() {
        this.orderTimeSp = new Long(0);
        this.orderPrice = "";
        this.orderDeskNum = "";
        this.orderRoleList = new ArrayList<>();
        this.orderShowAllocation = false;
    }

    public long getOrderTimeSp() {
        return orderTimeSp;
    }

    public void setOrderTimeSp(long orderTimeSp) {
        this.orderTimeSp = orderTimeSp;
    }

    public String getOrderPrice() {
        return orderPrice;
    }

    public void setOrderPrice(String orderPrice) {
        this.orderPrice = orderPrice;
    }

    public String getOrderDeskNum() {
        return orderDeskNum;
    }

    public void setOrderDeskNum(String orderDeskNum) {
        this.orderDeskNum = orderDeskNum;
    }

    public List<PublisherRoleBean> getOrderRoleList() {
        return orderRoleList;
    }

    public void setOrderRoleList(List<PublisherRoleBean> orderRoleList) {
        this.orderRoleList = orderRoleList;
    }

    public boolean isOrderShowAllocation() {
        return orderShowAllocation;
    }

    public void setOrderShowAllocation(boolean orderShowAllocation) {
        this.orderShowAllocation = orderShowAllocation;
    }

    @Override
    public String toString() {
        return "PublisherOrderBean{" +
                "orderTimeSp='" + orderTimeSp + '\'' +
                ", orderPrice='" + orderPrice + '\'' +
                ", orderDeskNum='" + orderDeskNum + '\'' +
                ", orderRoleList=" + orderRoleList +
                ", orderShowAllocation=" + orderShowAllocation +
                '}';
    }
}
