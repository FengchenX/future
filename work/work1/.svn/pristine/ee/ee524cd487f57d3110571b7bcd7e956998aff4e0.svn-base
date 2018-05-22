package com.znfz.assur.znfz_android.android.main.applicant.activity;

import android.support.v4.app.FragmentTransaction;

import com.ashokvarma.bottomnavigation.BottomNavigationBar;
import com.ashokvarma.bottomnavigation.BottomNavigationItem;
import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.app.BaseActivity;
import com.znfz.assur.znfz_android.android.main.applicant.frament.ApplicantFragmentIncome;
import com.znfz.assur.znfz_android.android.main.applicant.frament.ApplicantFragmentMy;
import com.znfz.assur.znfz_android.android.main.applicant.frament.ApplicantFragmentOrder;
import com.znfz.assur.znfz_android.android.main.applicant.frament.ApplicantFragmentScheduling;
import com.znfz.assur.znfz_android.android.main.publisher.frament.PublisherFragmentEmployees;
import com.znfz.assur.znfz_android.android.main.publisher.frament.PublisherFragmentIncome;
import com.znfz.assur.znfz_android.android.main.publisher.frament.PublisherFragmentOrder;
import com.znfz.assur.znfz_android.android.main.publisher.frament.PublisherFragmentScheduling;

import butterknife.ButterKnife;

/**
 * 申请者主界面
 * Created by assur on 2018/4/20
 */
public class ApplicantHomeActivity extends BaseActivity implements BottomNavigationBar.OnTabSelectedListener{


    private BottomNavigationBar mBottomNavigationBar;

    private ApplicantFragmentScheduling fragmentScheduling;
    private ApplicantFragmentOrder fragmentOrder;
    private ApplicantFragmentMy fragmentMy;


    @Override
    protected void initView() {
        setContentView(R.layout.activity_home_applicant);
        ButterKnife.bind(this);

        // 底部菜单初始化
        mBottomNavigationBar = (BottomNavigationBar)findViewById(R.id.bottom_navigation_bar);
        mBottomNavigationBar.setMode(BottomNavigationBar.MODE_FIXED);
        mBottomNavigationBar.setBackgroundStyle(BottomNavigationBar.BACKGROUND_STYLE_STATIC);
        mBottomNavigationBar.setBarBackgroundColor(R.color.white);
        mBottomNavigationBar.setInActiveColor(R.color.white);
        mBottomNavigationBar.setBackgroundColor(getResources().getColor(R.color.white));
        mBottomNavigationBar .addItem(new BottomNavigationItem(R.drawable.ic_tab_scheduling, R.string.applicant_home_tab_scheduling).setActiveColorResource(R.color.status_blue).setInActiveColorResource(R.color.alpha_gray))
                .addItem(new BottomNavigationItem(R.drawable.ic_tab_order, R.string.applicant_home_tab_order).setActiveColorResource(R.color.status_blue).setInActiveColorResource(R.color.alpha_gray))
                .addItem(new BottomNavigationItem(R.drawable.ic_tab_my, R.string.applicant_home_tab_my).setActiveColorResource(R.color.status_blue).setInActiveColorResource(R.color.alpha_gray))
                .setFirstSelectedPosition(0)
                .initialise();
        mBottomNavigationBar.setTabSelectedListener(this);
        onTabSelected(0); // 默认显示订单
    }

    @Override
    protected void initData() {

    }

    /**
     * 底部菜单选中回调
     * @param position
     */
    @Override
    public void onTabSelected(int position) {

        FragmentTransaction transaction = getSupportFragmentManager().beginTransaction();
        switch (position) {
            case 0:
                if (fragmentScheduling == null) {
                    fragmentScheduling = new ApplicantFragmentScheduling();
                }
                transaction.replace(R.id.ll_content, fragmentScheduling);
                break;

            case 1:
                if (fragmentOrder == null) {
                    fragmentOrder = new ApplicantFragmentOrder();
                }
                transaction.replace(R.id.ll_content, fragmentOrder);
                break;

            case 2:
                if (fragmentMy == null) {
                    fragmentMy = new ApplicantFragmentMy();
                }
                transaction.replace(R.id.ll_content, fragmentMy);
                break;

            default:
                if (fragmentMy == null) {
                    fragmentMy = new ApplicantFragmentMy();
                }
                transaction.replace(R.id.ll_content, fragmentMy);
                break;
        }
        transaction.commit();

    }

    /**
     * 底部菜单取消选中回调
     * @param position
     */
    @Override
    public void onTabUnselected(int position) {


    }

    /**
     * 底部菜单重复选中回调
     * @param position
     */
    @Override
    public void onTabReselected(int position) {

    }


}
