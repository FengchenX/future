package com.znfz.assur.znfz_android.android.main.applicant.activity;

import android.content.Context;
import android.os.Bundle;
import android.support.v7.widget.RecyclerView;
import android.view.View;
import android.widget.ImageView;
import android.widget.LinearLayout;
import android.widget.RelativeLayout;
import android.widget.TextView;

import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.app.BaseActivity;
import com.znfz.assur.znfz_android.android.main.applicant.adapter.ApplicantMenuDetailAdapter;
import com.znfz.assur.znfz_android.android.main.applicant.adapter.ApplicantOrderCommissionDetailAdapter;
import com.znfz.assur.znfz_android.android.main.applicant.bean.ApplicantMenuBean;
import com.znfz.assur.znfz_android.android.main.applicant.bean.ApplicantherRoleBean;

import java.util.ArrayList;
import java.util.List;

import butterknife.BindView;
import butterknife.ButterKnife;

/**
 * Created by dansihan on 2018/5/18.
 */

public class OrderDetailActivity extends BaseActivity {
    @BindView(R.id.ll_image_back)
    LinearLayout llImageBack;
    @BindView(R.id.tv_title)
    TextView tvTitle;
    @BindView(R.id.iv_right)
    ImageView ivRight;
    @BindView(R.id.tv_right)
    TextView tvRight;
    @BindView(R.id.rl_toolbar)
    RelativeLayout rlToolbar;
    @BindView(R.id.tv_time)
    TextView tvTime;
    @BindView(R.id.tv_order_time_show)
    TextView tvOrderTimeShow;
    @BindView(R.id.tv_order_table)
    TextView tvOrderTable;
    @BindView(R.id.tv_money)
    TextView tvMoney;
    @BindView(R.id.tv_order_money_show)
    TextView tvOrderMoneyShow;
    @BindView(R.id.tv_company_commission)
    TextView tvCompanyCommission;
    @BindView(R.id.tv_company_commission_show)
    TextView tvCompanyCommissionShow;
    @BindView(R.id.tv_company_commission_percent)
    TextView tvCompanyCommissionPercent;
    @BindView(R.id.order_detail_personnel)
    RecyclerView orderDetailPersonnel;
    @BindView(R.id.order_detail_menu)
    RecyclerView orderDetailMenu;

    private Context mContext;

    private List<ApplicantherRoleBean> orderListItems;       // 订单详情佣金分配数据
    private ApplicantOrderCommissionDetailAdapter orderAdapter;// 订单详情佣金分配数据适配器

    private List<ApplicantMenuBean> menuBeanList; //订单详情菜单数据
    private ApplicantMenuDetailAdapter menuDetailAdapter; //订单详情菜单适配器

    @Override
    protected void initView() {
        setContentView(R.layout.activity_order_detail);
        ButterKnife.bind(this);
        mContext = this;
        tvTitle.setText(getResources().getString(R.string.order_detail));
    }

    @Override
    protected void initData() {
        orderListItems = new ArrayList<>();
        orderAdapter = new ApplicantOrderCommissionDetailAdapter(mContext);
        orderDetailPersonnel.setAdapter(orderAdapter);

        menuBeanList = new ArrayList<>();
        menuDetailAdapter = new ApplicantMenuDetailAdapter(mContext);
    }

    @Override
    public void onClick(View v) {
        super.onClick(v);

        switch (v.getId()){
            case R.id.ll_image_back:
                finish();
                break;
        }
    }
}
