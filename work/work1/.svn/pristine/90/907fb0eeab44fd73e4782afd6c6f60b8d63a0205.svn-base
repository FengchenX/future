package com.znfz.assur.znfz_android.android.main.applicant.activity;

import android.content.Context;
import android.content.Intent;
import android.graphics.Rect;
import android.os.Bundle;
import android.support.v7.widget.LinearLayoutManager;
import android.support.v7.widget.RecyclerView;
import android.util.Log;
import android.view.View;
import android.widget.Button;
import android.widget.ImageView;
import android.widget.LinearLayout;
import android.widget.RelativeLayout;
import android.widget.TextView;

import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.app.BaseActivity;
import com.znfz.assur.znfz_android.android.main.applicant.adapter.ApplicantMyFriendAdapter;
import com.znfz.assur.znfz_android.android.main.applicant.bean.ApplicantMyFriendBean;
import com.znfz.assur.znfz_android.android.main.publisher.activity.OrderListActivity;

import java.io.Serializable;
import java.util.ArrayList;
import java.util.List;

import butterknife.BindView;
import butterknife.ButterKnife;

/**
 * Created by dansihan on 2018/5/18.
 */

public class ApplicantMyFriendActivity extends BaseActivity {
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
    @BindView(R.id.applicant_my_friend)
    RecyclerView applicantMyFriend;
    @BindView(R.id.img_no_friend)
    ImageView imgNoFriend;
    @BindView(R.id.btn_job_detail_apply_immediately)
    Button btnJobDetailApplyImmediately;


    private ApplicantMyFriendAdapter myFriendAdapter;
    private List<ApplicantMyFriendBean> myFriendBeanList;
    private Context mContext;
    private boolean ifshowcheck = false;

    private List<ApplicantMyFriendBean> friendBeanListchecked;

    @Override
    protected void initView() {
        setContentView(R.layout.activity_applicant_my_frend);
        ButterKnife.bind(this);
        mContext = this;
        tvTitle.setText(getResources().getString(R.string.applicant_my_friend));
        llImageBack.setOnClickListener(this);
        btnJobDetailApplyImmediately.setOnClickListener(this);
        // 列表视图
        LinearLayoutManager layoutmanager = new LinearLayoutManager(ApplicantMyFriendActivity.this);
        applicantMyFriend.setLayoutManager(layoutmanager);

        applicantMyFriend.addItemDecoration(new RecyclerView.ItemDecoration() {
            @Override
            public void getItemOffsets(Rect outRect, View view, RecyclerView parent, RecyclerView.State state) {
                // 设置分割线
                outRect.set(10, 10, 10, 10);
            }
        });
    }

    @Override
    protected void initData() {
        myFriendAdapter = new ApplicantMyFriendAdapter(mContext);
        myFriendBeanList = new ArrayList<>();
        ifshowcheck = getIntent().getBooleanExtra("ifShowSelectFriend", false);

        if (ifshowcheck) {
            btnJobDetailApplyImmediately.setVisibility(View.GONE);
            tvRight.setVisibility(View.VISIBLE);
            tvRight.setText(getResources().getString(R.string.dialog_ok));
            tvRight.setOnClickListener(this);
        } else {
            btnJobDetailApplyImmediately.setVisibility(View.VISIBLE);
        }
        for (int i = 0; i < 5; i++) {
            ApplicantMyFriendBean friendBean = new ApplicantMyFriendBean();
            friendBean.setIfshowcheck(ifshowcheck);
            friendBean.setName("小明" + i + 1);
            friendBean.setPhonenumber("1305213695" + i);
            friendBean.setUrl("http://img4.duitang.com/uploads/item/201411/06/20141106211621_KzYnZ.jpeg");
            myFriendBeanList.add(friendBean);
        }
        applicantMyFriend.setAdapter(myFriendAdapter);
        myFriendAdapter.setData(myFriendBeanList);
        myFriendAdapter.setOnItemClickListener(new ApplicantMyFriendAdapter.onItemClickListener() {
            @Override
            public void onClick(int position, int viewID) {
                switch (viewID) {
                    case R.id.Rll_friend_item:
                        if (myFriendBeanList.get(position).isChecked()) {
                            myFriendBeanList.get(position).setChecked(false);

                        } else {
                            myFriendBeanList.get(position).setChecked(true);
                        }
                        break;

                    case R.id.img_check_friend:
                        if (myFriendBeanList.get(position).isChecked()) {
                            myFriendBeanList.get(position).setChecked(false);

                        } else {
                            myFriendBeanList.get(position).setChecked(true);
                        }
                        break;
                }
                myFriendAdapter.notifyDataSetChanged();
            }
        });
    }

    @Override
    public void onClick(View v) {
        super.onClick(v);
        switch (v.getId()) {

            case R.id.ll_image_back:
                finish();
                break;

            case R.id.btn_job_detail_apply_immediately:
                startActivityForResult(ApplicantAddNewFrinedActivity.class, 1);
                break;

            case R.id.tv_right:
                friendBeanListchecked = new ArrayList<>();
                for (ApplicantMyFriendBean bean : myFriendBeanList) {
                    if (bean.isChecked()) {
                        friendBeanListchecked.add(bean);
                    }
                }
                Intent intent = new Intent();
                intent.putExtra("friendname", (Serializable) friendBeanListchecked);
                setResult(10, intent);
                finish();
                break;
        }
    }

    @Override
    protected void onActivityResult(int requestCode, int resultCode, Intent data) {
        super.onActivityResult(requestCode, resultCode, data);
    }
}
