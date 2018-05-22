package com.znfz.assur.znfz_android.android.main.publisher.activity;

import android.content.Intent;
import android.os.Bundle;
import android.view.View;
import android.widget.AdapterView;
import android.widget.Button;
import android.widget.EditText;
import android.widget.ImageView;
import android.widget.LinearLayout;
import android.widget.RelativeLayout;
import android.widget.TextView;

import com.flyco.dialog.entity.DialogMenuItem;
import com.flyco.dialog.listener.OnOperItemClickL;
import com.flyco.dialog.widget.ActionSheetDialog;
import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.app.BaseActivity;
import com.znfz.assur.znfz_android.android.main.applicant.activity.ApplicantMyFriendActivity;
import com.znfz.assur.znfz_android.android.main.applicant.bean.ApplicantMyFriendBean;

import java.util.ArrayList;
import java.util.List;

import butterknife.BindView;
import butterknife.ButterKnife;
import jnr.ffi.annotations.In;

/**
 * Created by dansihan on 2018/5/19.
 */

public class publisherJobInformationActivity extends BaseActivity {
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
    @BindView(R.id.ed_job)
    EditText edJob;  //招聘职位
    @BindView(R.id.tv_select_commission_rate) //选择佣金比例
            TextView tvSelectCommissionRate;
    @BindView(R.id.tv_select_number) //选择招聘人数
            TextView tvSelectNumber;
    @BindView(R.id.tv_select_fiend)  //选择好友
            TextView tvSelectFiend;
    @BindView(R.id.ed_job_notes)  //职位描述
            EditText edJobNotes;
    @BindView(R.id.btn_publisher_job_add) //保存按钮
            Button btnPublisherJobAdd;
    @BindView(R.id.Rll_select_commission_rate)
    RelativeLayout RllSelectCommissionRate;
    @BindView(R.id.Rll_select_number)
    RelativeLayout RllSelectNumber;
    @BindView(R.id.Rll_select_fiend)
    RelativeLayout RllSelectFiend;

    private String JobPersonNumber = ""; //保存招聘人数
    private String CommissionRate = ""; //保存佣金比例
    private List<ApplicantMyFriendBean> friendBeanListchecked;
    private boolean Ifmodify = false;

    @Override
    protected void initView() {
        setContentView(R.layout.activity_job_information);
        ButterKnife.bind(this);

        tvTitle.setText(getResources().getString(R.string.publisher_job_information));
        tvSelectFiend.setOnClickListener(this);
        tvSelectCommissionRate.setOnClickListener(this);
        tvSelectNumber.setOnClickListener(this);
        llImageBack.setOnClickListener(this);
        RllSelectCommissionRate.setOnClickListener(this);
        RllSelectNumber.setOnClickListener(this);
        RllSelectFiend.setOnClickListener(this);
        btnPublisherJobAdd.setOnClickListener(this);

    }

    @Override
    protected void initData() {
        Intent intent = getIntent();
        Ifmodify = intent.getBooleanExtra("item", false);
        if (Ifmodify) {
            edJob.setText(intent.getStringExtra("jobname"));
            JobPersonNumber = intent.getStringExtra("JobPersonNumber");
            tvSelectNumber.setText(JobPersonNumber);
            tvSelectFiend.setText(intent.getStringExtra("FriendName"));
            CommissionRate = intent.getStringExtra("CommissionRate");
            tvSelectCommissionRate.setText(CommissionRate + "%");
        }
    }

    @Override
    public void onClick(View v) {
        super.onClick(v);

        switch (v.getId()) {
            case R.id.ll_image_back:
                finish();
                break;

            case R.id.btn_publisher_job_add:
                Intent intent = new Intent();
                intent.putExtra("jobname", edJob.getText().toString());
                intent.putExtra("JobPersonNumber", JobPersonNumber);
                intent.putExtra("CommissionRate", CommissionRate);
                intent.putExtra("FriendName", tvSelectFiend.getText().toString());
                setResult(2, intent);
                finish();
                break;

            case R.id.Rll_select_commission_rate: //选择佣金比例
                SelectCommissionRate();
                break;
            case R.id.tv_select_commission_rate: //选择佣金比例
                SelectCommissionRate();
                break;

            case R.id.tv_select_number: //选择招聘人数
                SelectNumber();
                break;
            case R.id.Rll_select_number: //选择招聘人数
                SelectNumber();
                break;

            case R.id.tv_select_fiend: //选择好友
                startToApplicantMyFriendActivity();
                break;
            case R.id.Rll_select_fiend: //选择好友
                startToApplicantMyFriendActivity();
                break;
        }
    }

    @Override
    protected void onActivityResult(int requestCode, int resultCode, Intent data) {
        super.onActivityResult(requestCode, resultCode, data);
        if (requestCode == 1 && resultCode == 10) {
            tvSelectFiend.setText("");
            friendBeanListchecked = (List<ApplicantMyFriendBean>) data.getSerializableExtra("friendname");
            StringBuffer name = new StringBuffer();
            for (ApplicantMyFriendBean bean : friendBeanListchecked) {
                name.append(bean.getName() + "  ,");
            }
            tvSelectFiend.setText(name.toString());
        }
    }

    /**
     * 选择好友
     */
    private void startToApplicantMyFriendActivity() {
        //选择好友
        Bundle bundle = new Bundle();
        bundle.putBoolean("ifShowSelectFriend", true);
        startActivityForResult(ApplicantMyFriendActivity.class, bundle, 1);
    }

    /**
     * 选择招聘人数
     */
    private void SelectNumber() {
        final ArrayList<DialogMenuItem> perItemArray = new ArrayList<>();
        for (int i = 1; i < 50; i++) {
            DialogMenuItem item = new DialogMenuItem(i + "", i);
            perItemArray.add(item);
        }
        final ActionSheetDialog perDialog = new ActionSheetDialog(context, perItemArray, null);
        perDialog.title(getStringById(R.string.publisher_add_scheduling_per_hint))//
                .titleTextSize_SP(14.5f)
                .show();
        perDialog.setOnOperItemClickL(new OnOperItemClickL() {
            @Override
            public void onOperItemClick(AdapterView<?> parent, View view, int p, long id) {
                perDialog.dismiss();
                JobPersonNumber = String.valueOf(p + 1);
                tvSelectNumber.setText(JobPersonNumber + "  ");
            }
        });
    }

    private void SelectCommissionRate() {

        final ArrayList<DialogMenuItem> commissionRate = new ArrayList<>();
        for (int i = 1; i < 100; i++) {
            DialogMenuItem item = new DialogMenuItem(i + "%", i);
            commissionRate.add(item);
        }
        final ActionSheetDialog commissionRateDialog = new ActionSheetDialog(context, commissionRate, null);
        commissionRateDialog.title(getStringById(R.string.publisher_add_scheduling_per_hint))//
                .titleTextSize_SP(14.5f)
                .show();
        commissionRateDialog.setOnOperItemClickL(new OnOperItemClickL() {
            @Override
            public void onOperItemClick(AdapterView<?> parent, View view, int p, long id) {
                commissionRateDialog.dismiss();
                CommissionRate = String.valueOf(p + 1);
                tvSelectCommissionRate.setText(CommissionRate + "%  ");
            }
        });

    }
}
