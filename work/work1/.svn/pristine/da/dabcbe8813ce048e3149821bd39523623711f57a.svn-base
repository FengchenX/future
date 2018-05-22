package com.znfz.assur.znfz_android.android.main.publisher.activity;

import android.support.v4.app.FragmentTransaction;

import com.ashokvarma.bottomnavigation.BottomNavigationBar;
import com.ashokvarma.bottomnavigation.BottomNavigationItem;
import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.app.BaseActivity;
import com.znfz.assur.znfz_android.android.main.applicant.frament.ApplicantFragmentMy;
import com.znfz.assur.znfz_android.android.main.publisher.frament.PublisherFragmentAddEmployee;
import com.znfz.assur.znfz_android.android.main.publisher.frament.PublisherFragmentEmployees;
import com.znfz.assur.znfz_android.android.main.publisher.frament.PublisherFragmentIncome;
import com.znfz.assur.znfz_android.android.main.publisher.frament.PublisherFragmentOrder;
import com.znfz.assur.znfz_android.android.main.publisher.frament.PublisherFragmentScheduling;
import butterknife.ButterKnife;

/**
 * 发布者主界面
 * Created by assur on 2018/4/20
 */
public class PublisherHomeActivity extends BaseActivity implements BottomNavigationBar.OnTabSelectedListener {

    private BottomNavigationBar mBottomNavigationBar;
    private PublisherFragmentScheduling publisherFragmentScheduling;
    private PublisherFragmentOrder publisherFragmentOrder;
    private PublisherFragmentIncome publisherFragmentIncome;
    private PublisherFragmentAddEmployee publisherFragmentAddEmployee;
    private ApplicantFragmentMy fragmentMy;


    @Override
    protected void initView() {
        setContentView(R.layout.activity_home_publisher);
        ButterKnife.bind(this);


        // 底部菜单初始化
        mBottomNavigationBar = (BottomNavigationBar)findViewById(R.id.bottom_navigation_bar);
        mBottomNavigationBar.setMode(BottomNavigationBar.MODE_FIXED);
        mBottomNavigationBar.setBackgroundStyle(BottomNavigationBar.BACKGROUND_STYLE_STATIC);
        mBottomNavigationBar.setBarBackgroundColor(R.color.white);
        mBottomNavigationBar.setInActiveColor(R.color.white);
        mBottomNavigationBar.setBackgroundColor(getResources().getColor(R.color.white));
        mBottomNavigationBar.addItem(new BottomNavigationItem(R.drawable.ic_tab_scheduling, R.string.publisher_home_tab_scheduling).setActiveColorResource(R.color.status_blue).setInActiveColorResource(R.color.alpha_gray))
                .addItem(new BottomNavigationItem(R.drawable.ic_tab_order, R.string.publisher_home_tab_order).setActiveColorResource(R.color.status_blue).setInActiveColorResource(R.color.alpha_gray))
                .addItem(new BottomNavigationItem(R.drawable.ic_tab_income, R.string.publisher_home_tab_addschedule).setActiveColorResource(R.color.status_blue).setInActiveColorResource(R.color.alpha_gray))
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
                if (publisherFragmentScheduling == null) {
                    publisherFragmentScheduling = new PublisherFragmentScheduling();
                }
                transaction.replace(R.id.ll_content, publisherFragmentScheduling);

                break;
            case 1:
                if (publisherFragmentOrder == null) {
                    publisherFragmentOrder = new PublisherFragmentOrder();
                }
                transaction.replace(R.id.ll_content, publisherFragmentOrder);
                break;
            case 2:
                if (publisherFragmentAddEmployee == null) {
                    publisherFragmentAddEmployee = new PublisherFragmentAddEmployee();
                }
                transaction.replace(R.id.ll_content, publisherFragmentAddEmployee);
                break;
            case 3:
                if (fragmentMy == null) {
                    fragmentMy = new ApplicantFragmentMy();
                }
                transaction.replace(R.id.ll_content, fragmentMy);
                break;
            default:
                if (publisherFragmentOrder == null) {
                    publisherFragmentOrder = new PublisherFragmentOrder();
                }
                transaction.replace(R.id.ll_content, publisherFragmentOrder);
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
