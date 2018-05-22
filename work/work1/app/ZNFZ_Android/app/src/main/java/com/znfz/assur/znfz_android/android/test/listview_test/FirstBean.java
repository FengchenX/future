package com.znfz.assur.znfz_android.android.test.listview_test;

import java.util.ArrayList;
import java.util.List;

public class FirstBean {

    private String name;      // 名称
    private boolean spread;   // 是否展开
    private List<SecondBean> subList;  // 二级展开数据


    public FirstBean() {
        this.name = "";
        this.spread = false;
        subList = new ArrayList<>();
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public boolean isSpread() {
        return spread;
    }

    public void setSpread(boolean spread) {
        this.spread = spread;
    }

    public List<SecondBean> getSubList() {
        return subList;
    }

    public void setSubList(List<SecondBean> subList) {
        this.subList = subList;
    }

    @Override
    public String toString() {
        return "FirstBean{" +
                "name='" + name + '\'' +
                ", spread=" + spread +
                ", subList=" + subList +
                '}';
    }
}
