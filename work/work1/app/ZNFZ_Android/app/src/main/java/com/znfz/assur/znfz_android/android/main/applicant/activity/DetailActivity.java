package com.znfz.assur.znfz_android.android.main.applicant.activity;

import android.os.Bundle;
import android.view.View;
import android.widget.Button;
import android.widget.ImageView;
import android.widget.LinearLayout;
import android.widget.RelativeLayout;
import android.widget.TextView;

import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.app.BaseActivity;

import butterknife.BindView;
import butterknife.ButterKnife;

/**
 * Created by dansihan on 2018/5/18.
 */

public class DetailActivity extends BaseActivity {

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
    @BindView(R.id.tv_chef)
    TextView tvChef;
    @BindView(R.id.tv_jod_detail_trick)  //厨师人数
            TextView tvJodDetailTrick;
    @BindView(R.id.tv_jod_detail_commission)  //佣金百分比
            TextView tvJodDetailCommission;
    @BindView(R.id.jod_detail_address)  //工作地址大区：南山区
            TextView jodDetailAddress;
    @BindView(R.id.jod_detail_time)  //工作时间
            TextView jodDetailTime;
    @BindView(R.id.tv_job_detail)  //工作描述
            TextView tvJobDetail;
    @BindView(R.id.tv_job_detail_address) //工作详情地址
            TextView tvJobDetailAddress;
    @BindView(R.id.tv_job_detail_company)  //工作公司名字
            TextView tvJobDetailCompany;
    @BindView(R.id.tv_icon)
    TextView tvIcon;
    @BindView(R.id.btn_job_detail_apply_immediately) //立即申请按钮
    Button btnJobDetailApplyImmediately;

    @Override
    protected void initView() {
        setContentView(R.layout.activity_applicant_detail);
        ButterKnife.bind(this);
        llImageBack.setOnClickListener(this);
        tvTitle.setText(getResources().getString(R.string.details));
        btnJobDetailApplyImmediately.setOnClickListener(this);
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

            case R.id.btn_job_detail_apply_immediately:


                break;
        }
    }

}
