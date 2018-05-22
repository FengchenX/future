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

public class ApplicantAddNewFrinedActivity extends BaseActivity {
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
    @BindView(R.id.tv_nick_name)
    TextView tvNickName;
    @BindView(R.id.ed_my_friend_nick_name)
    EditText edMyFriendNickName;
    @BindView(R.id.tv_my_friend_phone_number)
    EditText tvMyFriendPhoneNumber;

    @Override
    protected void initView() {
        setContentView(R.layout.activity_add_new_friend);
        ButterKnife.bind(this);

        llImageBack.setOnClickListener(this);
        tvTitle.setText(getResources().getString(R.string.applicant_my_friend_info));
        tvRight.setVisibility(View.VISIBLE);
        tvRight.setText(getResources().getString(R.string.applicant_add_friend_save));
        tvRight.setOnClickListener(this);
        tvRight.setTextColor(getResources().getColor(R.color.text_title_color));
    }

    @Override
    protected void initData() {

    }

    @Override
    public void onClick(View v) {
        super.onClick(v);

        switch (v.getId()){
            case R.id.ll_image_back:
                setResult(RESULT_OK);
                finish();
                break;

            case R.id.tv_right:

                break;
        }
    }
}
