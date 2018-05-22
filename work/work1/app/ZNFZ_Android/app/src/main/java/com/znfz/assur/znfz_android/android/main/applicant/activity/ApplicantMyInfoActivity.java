package com.znfz.assur.znfz_android.android.main.applicant.activity;

import android.os.Bundle;
import android.view.View;
import android.widget.EditText;
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

public class ApplicantMyInfoActivity extends BaseActivity {
    @BindView(R.id.ll_image_back)
    LinearLayout ll_image_back;
    @BindView(R.id.tv_title)
    TextView tvTitle;
    @BindView(R.id.iv_right)
    ImageView ivRight;
    @BindView(R.id.tv_right)
    TextView tvRight;
    @BindView(R.id.rl_toolbar)
    RelativeLayout rlToolbar;
    @BindView(R.id.ed_my_nick_name)
    EditText edMyNickName;
    @BindView(R.id.tv_my_phone_number)
    TextView tvMyPhoneNumber;

    @Override
    protected void initView() {
        setContentView(R.layout.activity_my_info);
        ButterKnife.bind(this);
        tvTitle.setText(getResources().getString(R.string.my_info));
        ll_image_back.setOnClickListener(this);
    }

    @Override
    protected void initData() {

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
