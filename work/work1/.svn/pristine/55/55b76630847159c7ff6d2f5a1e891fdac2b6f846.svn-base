package com.znfz.assur.znfz_android.android.main.publisher.activity;

import android.graphics.Rect;
import android.support.v7.widget.LinearLayoutManager;
import android.support.v7.widget.RecyclerView;
import android.view.View;
import android.widget.LinearLayout;

import com.orhanobut.logger.Logger;
import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.app.BaseActivity;
import com.znfz.assur.znfz_android.android.common.global.UserInfo;
import com.znfz.assur.znfz_android.android.common.global.Variable;
import com.znfz.assur.znfz_android.android.common.task.GetAllOrderByScheduleTask;
import com.znfz.assur.znfz_android.android.main.publisher.adapter.PublishSchedulingAdapter;
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherOrderBean;
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherSchedulingBean;

import java.util.ArrayList;
import java.util.List;
import butterknife.BindView;
import butterknife.ButterKnife;

/**
 * 历史排版列表界面
 * Created by assur on 2018/5/21
 */
public class ScheduleListActivity extends BaseActivity {


    @BindView(R.id.ll_empty_view)                         // 无数据视图
    LinearLayout llEmptyView;
    @BindView(R.id.publisher_listview_scheduling)         // 列表视图
    RecyclerView recyclerView;

    private List<PublisherSchedulingBean> listItems;          // 排班数据
    private PublishSchedulingAdapter adapter;                 // 排班列表适配器


    @Override
    protected void initView() {
        setContentView(R.layout.activity_publisher_scheduling);
        ButterKnife.bind(this);
        initToolBar(getStringById(R.string.publisher_home_scheduling_list), true);

        // 列表视图
        LinearLayoutManager layoutmanager = new LinearLayoutManager(ScheduleListActivity.this);
        recyclerView.setLayoutManager(layoutmanager);
        recyclerView.addItemDecoration(new RecyclerView.ItemDecoration() {
            @Override
            public void getItemOffsets(Rect outRect, View view, RecyclerView parent, RecyclerView.State state) {
                // 设置分割线
                outRect.set(10, 10, 10, 10);
            }
        });
    }


    @Override
    protected void initData() {

        listItems = new ArrayList<>(UserInfo.scheduleistItemsOfOneCompany);

        // 当前排班列表适配器初始化
        adapter = new PublishSchedulingAdapter(this);
        adapter.setOnItemClickListener(new PublishSchedulingAdapter.onItemClickListener() {
            @Override
            public void onClick(int position, int viewID) {
                switch (viewID) {
                    case R.id.publisher_scheduling_list_item_rl_btn_detail:
                        listItems.get(position).setShowRoleList(!listItems.get(position).isShowRoleList());
                        adapter.setData(listItems);
                        break;
                }
            }
        });

        // 如果没有数据，显示空白视图
        if(listItems == null || listItems.size() == 0) {
            llEmptyView.setVisibility(View.VISIBLE);
        } else {
            llEmptyView.setVisibility(View.GONE);
            recyclerView.setAdapter(adapter);
            adapter.setData(listItems);
        }

    }

}