package com.znfz.assur.znfz_android.android.test.frament_viewpage_test;


import android.graphics.Color;
import android.support.v4.app.Fragment;
import android.support.v4.view.ViewPager;
import android.view.View;
import android.widget.TextView;

import com.scwang.smartrefresh.layout.api.RefreshLayout;
import com.scwang.smartrefresh.layout.footer.ClassicsFooter;
import com.scwang.smartrefresh.layout.header.ClassicsHeader;
import com.scwang.smartrefresh.layout.listener.OnLoadMoreListener;
import com.scwang.smartrefresh.layout.listener.OnRefreshListener;
import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.app.BaseActivity;

import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Date;
import java.util.List;

/**
 * FramentViewPageActivity
 * Created by assur on 2018/4/20.
 */
public class FramentViewPageActivity extends BaseActivity {

    private TextView tv_item_one;
    private TextView tv_item_two;
    private TextView tv_item_three;
    private ViewPager myViewPager;
    private List<Fragment> list;
    private FragmentTabPagerAdapter adapter;

    @Override
    protected void initView() {
        setContentView(R.layout.frement_viewpage_activity);
        initToolBar("我的订单", true);

        tv_item_one = (TextView) findViewById(R.id.tv_item_one);
        tv_item_two = (TextView) findViewById(R.id.tv_item_two);
        tv_item_three = (TextView) findViewById(R.id.tv_item_three);
        myViewPager = (ViewPager) findViewById(R.id.myViewPager);

        // 设置菜单栏的点击事件
        tv_item_one.setOnClickListener(this);
        tv_item_two.setOnClickListener(this);
        tv_item_three.setOnClickListener(this);
        myViewPager.setOnPageChangeListener(new MyPagerChangeListener());

        //把Fragment添加到List集合里面
        list = new ArrayList<>();
        list.add(new FragmentOne());
        list.add(new FragmentTwo());
        list.add(new FragmentThree());
        adapter = new FragmentTabPagerAdapter(getSupportFragmentManager(), list);
        myViewPager.setAdapter(adapter);
        myViewPager.setCurrentItem(0);  //初始化显示第一个页面
        tv_item_one.setBackgroundColor(Color.LTGRAY);//被选中就为红色


        // 上啦刷新/下啦刷新
        RefreshLayout refreshLayout = (RefreshLayout)findViewById(R.id.refreshlayout);
        final ClassicsHeader header = (ClassicsHeader)findViewById(R.id.classicsheader);
        header.REFRESH_HEADER_REFRESHING = "刷新中...";
        header.REFRESH_HEADER_RELEASE = "松手开始刷新";
        header.REFRESH_HEADER_PULLING = "请向下拉";
        header.REFRESH_HEADER_FAILED = "刷新失败！";
        header.REFRESH_HEADER_FINISH = "刷新成功";
        header.REFRESH_HEADER_LOADING = "加载中...";
        final ClassicsFooter footer = (ClassicsFooter) findViewById(R.id.classicsFooter);
        footer.REFRESH_FOOTER_LOADING = "加载中...";
        footer.REFRESH_FOOTER_FAILED = "加载失败！";
        footer.REFRESH_FOOTER_FINISH = "加载成功";
        refreshLayout.setOnRefreshListener(new OnRefreshListener() {
            @Override
            public void onRefresh(RefreshLayout refreshlayout) {
                Date day=new Date();
                SimpleDateFormat df = new SimpleDateFormat("MM-dd HH:mm");
                header.setLastUpdateText("最后刷新时间: " + df.format(day));
                refreshlayout.finishRefresh(5000/*,false*/);//传入false表示刷新失败
            }
        });
        refreshLayout.setOnLoadMoreListener(new OnLoadMoreListener() {
            @Override
            public void onLoadMore(RefreshLayout refreshlayout) {
                refreshlayout.finishLoadMore(2000/*,false*/);//传入false表示加载失败
            }
        });
    }

    @Override
    protected void initData() {


    }

    //第一次设置点击监听事件，为菜单栏设置监听事件，监听的对象是页面的滑动
    @Override
    public void onClick(View v) {
        super.onClick(v);
        switch (v.getId()) {
            case R.id.tv_item_one:
                myViewPager.setCurrentItem(0);
                tv_item_one.setBackgroundColor(Color.LTGRAY);
                tv_item_two.setBackgroundColor(Color.WHITE);
                tv_item_three.setBackgroundColor(Color.WHITE);
                break;
            case R.id.tv_item_two:
                myViewPager.setCurrentItem(1);
                tv_item_one.setBackgroundColor(Color.WHITE);
                tv_item_two.setBackgroundColor(Color.LTGRAY);
                tv_item_three.setBackgroundColor(Color.WHITE);
                break;
            case R.id.tv_item_three:
                myViewPager.setCurrentItem(2);
                tv_item_one.setBackgroundColor(Color.WHITE);
                tv_item_two.setBackgroundColor(Color.WHITE);
                tv_item_three.setBackgroundColor(Color.LTGRAY);
                break;
        }
    }

    //第二次设置点击监听事件，为ViewPager设置监听事件，用于实现菜单栏的样式变化
    private class MyPagerChangeListener implements ViewPager.OnPageChangeListener {
        @Override
        public void onPageScrolled(int position, float positionOffset, int positionOffsetPixels) {

        }

        @Override
        public void onPageSelected(int position) {
            switch (position) {
                case 0:
                    tv_item_one.setBackgroundColor(Color.LTGRAY);
                    tv_item_two.setBackgroundColor(Color.WHITE);
                    tv_item_three.setBackgroundColor(Color.WHITE);
                    break;
                case 1:
                    tv_item_one.setBackgroundColor(Color.WHITE);
                    tv_item_two.setBackgroundColor(Color.LTGRAY);
                    tv_item_three.setBackgroundColor(Color.WHITE);
                    break;
                case 2:
                    tv_item_one.setBackgroundColor(Color.WHITE);
                    tv_item_two.setBackgroundColor(Color.WHITE);
                    tv_item_three.setBackgroundColor(Color.LTGRAY);
                    break;
            }
        }

        @Override
        public void onPageScrollStateChanged(int state) {

        }
    }

}
