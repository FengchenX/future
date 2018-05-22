package com.znfz.assur.znfz_android.android.main.applicant.frament;
import android.graphics.Rect;
import android.os.Bundle;
import android.support.v7.widget.LinearLayoutManager;
import android.support.v7.widget.RecyclerView;
import android.view.View;
import android.widget.Button;
import android.widget.LinearLayout;
import android.widget.RelativeLayout;
import android.widget.TextView;

import com.orhanobut.logger.Logger;
import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.app.BaseFragment;
import com.znfz.assur.znfz_android.android.common.global.UserInfo;
import com.znfz.assur.znfz_android.android.common.global.Variable;
import com.znfz.assur.znfz_android.android.common.task.GetAllIncomeTask;
import com.znfz.assur.znfz_android.android.common.utils.RoleUtils;
import com.znfz.assur.znfz_android.android.main.applicant.activity.ApplicantHomeActivity;
import com.znfz.assur.znfz_android.android.main.applicant.activity.ApplicantMyFriendActivity;
import com.znfz.assur.znfz_android.android.main.applicant.activity.ApplicantMyInfoActivity;
import com.znfz.assur.znfz_android.android.main.publisher.activity.PublisherHomeActivity;
import com.znfz.assur.znfz_android.android.main.publisher.adapter.PublishIncomeAdapter;
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherIncomeBean;
import java.util.ArrayList;
import java.util.List;
import butterknife.BindView;
import butterknife.ButterKnife;

/**
 * 我的界面
 * Created by assur on 2018/5/18.
 */

public class ApplicantFragmentMy extends BaseFragment {

    @BindView(R.id.applicant_my_tv_username)           // 用户名
            TextView tvUserName;

    @BindView(R.id.applicant_my_btn_changerole)        // 角色切换
            Button btnRoleChange;

    @BindView(R.id.applicant_my_wallet_rl)   // 钱包
            RelativeLayout rlWallet;

    @BindView(R.id.applicant_mypublish_rl)   // 我的发布
            RelativeLayout rlMyPublish;

    @BindView(R.id.applicant_income_rl)      // 积分记录
            RelativeLayout rlIncome;

    @BindView(R.id.applicant_myinfo_rl)      // 我的资料
            RelativeLayout rlMyInfo;

    @BindView(R.id.applicant_myfriend_rl)    // 我的好友
            RelativeLayout rlMyFriend;

    @BindView(R.id.applicant_mysystemid_rl)  // 我的系统身份
            RelativeLayout rlMySystemId;



    @Override
    protected int setView() {
        return R.layout.frement_applicant_my;
    }

    @Override
    protected void init(View view) {
        ButterKnife.bind(this, view);

        btnRoleChange.setOnClickListener(this);
        rlWallet.setOnClickListener(this);
        rlMyPublish.setOnClickListener(this);
        rlIncome.setOnClickListener(this);
        rlMyInfo.setOnClickListener(this);
        rlMyFriend.setOnClickListener(this);
        rlMySystemId.setOnClickListener(this);

        updateMyPublishViewShow();
    }

    @Override
    protected void initData(Bundle savedInstanceState) {

        //tvUserName.setText(UserInfo.USER_NAME);

    }

    @Override
    public void onClick(View v) {
        super.onClick(v);
        switch(v.getId()) {

            case R.id.applicant_my_btn_changerole:
                // 切换角色
                if(UserInfo.roleState == UserInfo.ROLE_STATE_APPLACATION) {
                    UserInfo.roleState = UserInfo.ROLE_STATE_PUBLISH;
                    startActivity(PublisherHomeActivity.class);
                } else {
                    UserInfo.roleState = UserInfo.ROLE_STATE_APPLACATION;
                    startActivity(ApplicantHomeActivity.class);
                }


                break;

            case R.id.applicant_my_wallet_rl:
                // 我的钱包

                break;

            case R.id.applicant_mypublish_rl:
                // 我的发布

                break;

            case R.id.applicant_income_rl:
                // 积分记录

                break;

            case R.id.applicant_myinfo_rl:
                // 我的资料
                startActivity(ApplicantMyInfoActivity.class);
                break;

            case R.id.applicant_myfriend_rl:
                // 我的好友
                Bundle bundle = new Bundle();
                bundle.putBoolean("ifShowSelectFriend",false);
                startActivity(ApplicantMyFriendActivity.class,bundle);
                break;

            case R.id.applicant_mysystemid_rl:
                // 我的系统身份

                break;

        }
    }

    /**
     * 更新"我的发布"菜单的显示
     */
    private void updateMyPublishViewShow() {
        if(UserInfo.roleState == UserInfo.ROLE_STATE_PUBLISH) {
            rlMyPublish.setVisibility(View.VISIBLE);
            btnRoleChange.setText("切换申请");
        } else {
            rlMyPublish.setVisibility(View.GONE);
            btnRoleChange.setText("切换发布");
        }

    }
}
