package com.znfz.assur.znfz_android.android.main.publisher.activity;

import android.os.Bundle;
import android.view.View;
import android.widget.EditText;
import android.widget.ImageView;
import android.widget.LinearLayout;
import android.widget.ListView;
import android.widget.RelativeLayout;
import android.widget.TextView;

import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.app.BaseActivity;

import butterknife.BindView;
import butterknife.ButterKnife;

/**
 * Created by dansihan on 2018/5/21.
 */

public class PublisherAddEmployeeActivity extends BaseActivity {
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
    @BindView(R.id.ed_publisher_company_name) //输入公司名称
            EditText edPublisherCompanyName;
    @BindView(R.id.ed_publisher_company_account) //输入公司账户
            EditText edPublisherCompanyAccount;
    @BindView(R.id.ed_publisher_store_name)  //输入门店名称
            EditText edPublisherStoreName;
    @BindView(R.id.ed_publisher_store_number) //输入门店编号
            EditText edPublisherStoreNumber;
    @BindView(R.id.tv_publisher_company_proportion) //选择公司比例
            TextView tvPublisherCompanyProportion;
    @BindView(R.id.tv_publisher_store_proportion)  //选择门店比例
            TextView tvPublisherStoreProportion;
    @BindView(R.id.tv_publisher_time_start)  //开始工作时间
            TextView tvPublisherTimeStart;
    @BindView(R.id.tv_publisher_time_end)  //结束工作时间
            TextView tvPublisherTimeEnd;
    @BindView(R.id.ed_publisher_address)  //工作大概区域
            EditText edPublisherAddress;
    @BindView(R.id.ed_publisher_address_detail)  //工作详细地址
            EditText edPublisherAddressDetail;
    @BindView(R.id.lv_add_employee_list)  //招聘工作人员列表
            ListView lvAddEmployeeList;
    @BindView(R.id.tv_add_employee)  //添加招聘工作人员
            TextView tvAddEmployee;

    @Override
    protected void initView() {
        setContentView(R.layout.activity_add_empioyee);
        ButterKnife.bind(this);

        tvTitle.setText(getResources().getString(R.string.publisher_Release_schedule));
    }

    @Override
    protected void initData() {

    }

    @Override
    public void onClick(View v) {
        super.onClick(v);
        switch (v.getId()) {
            case R.id.ll_image_back:
                finish();
                break;

            case R.id.tv_add_employee:

                startActivityForResult(publisherJobInformationActivity.class,10);
                break;
        }
    }
}
