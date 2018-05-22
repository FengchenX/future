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
import com.znfz.assur.znfz_android.android.main.publisher.adapter.PublishOrderAdapter;
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherOrderBean;
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherRoleBean;
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherSchedulingBean;

import java.util.ArrayList;
import java.util.List;
import butterknife.BindView;
import butterknife.ButterKnife;

/**
 * 订单列表界面
 * Created by assur on 2018/4/27
 */
public class OrderListActivity extends BaseActivity {


    @BindView(R.id.ll_empty_view)                       // 无数据视图
    LinearLayout llEmptyView;

    @BindView(R.id.order_list_listview)    // 订单列表
    RecyclerView orderRecyclerView;

    private List<PublisherOrderBean> orderListItems;       // 订单数据
    private PublishOrderAdapter orderAdapter;              // 订单列表适配器

    @Override
    protected void initView() {
        setContentView(R.layout.activity_order_list);
        ButterKnife.bind(this);
        initToolBar(getStringById(R.string.order_page_title), true);

        // 列表视图
        LinearLayoutManager layoutmanager = new LinearLayoutManager(OrderListActivity.this);
        orderRecyclerView.setLayoutManager(layoutmanager);
        orderRecyclerView.addItemDecoration(new RecyclerView.ItemDecoration() {
            @Override
            public void getItemOffsets(Rect outRect, View view, RecyclerView parent, RecyclerView.State state) {
                // 设置分割线
                outRect.set(10, 10, 10, 10);
            }
        });
    }


    @Override
    protected void initData() {

        // 订单列表适配器初始化
        orderAdapter = new PublishOrderAdapter(OrderListActivity.this);
        orderAdapter.setOnItemClickListener(new PublishOrderAdapter.onItemClickListener() {
            @Override
            public void onClick(int position, int viewID) {
                switch (viewID) {
                    case R.id.publisher_order_list_item_rl_more:
                        orderListItems.get(position).setOrderShowAllocation(!orderListItems.get(position).isOrderShowAllocation());
                        orderAdapter.setData(orderListItems);
                        break;
                }
            }
        });
        orderRecyclerView.setAdapter(orderAdapter);

        getOrderByScheduleRequest(UserInfo.currentHistorySchedule);
        // 初始化今日列表数据
        // initTodayListItems();
        // orderAdapter.setData(orderListItems);
    }

//
//    /**
//     * 今日订单数据
//     */
//    private void initTodayListItems() {
//        if(orderListItems != null) {
//            orderListItems.removeAll(orderListItems);
//        } else {
//            orderListItems = new ArrayList<PublisherOrderBean>();
//        }
//
//        PublisherOrderBean bean1 = new PublisherOrderBean();
//        //bean1.setOrderTime("2018-04-25 15:24");
//        bean1.setOrderPrice("660");
//        bean1.setOrderDeskNum("6号桌");
//        for (int i = 0; i < 3; i++) {
//            if(i == 0) {
//                PublisherRoleBean role = new PublisherRoleBean();
//                role.setRole("经理");
//                role.setRoleCurNum("1");
//                role.setRolePercentage("40%");
//                role.setRoleMoney("400.00");
//                bean1.getOrderRoleList().add(role);
//            }
//            if(i == 1) {
//                PublisherRoleBean role = new PublisherRoleBean();
//                role.setRole("厨师");
//                role.setRoleCurNum("3");
//                role.setRolePercentage("30%");
//                role.setRoleMoney("90.00");
//                bean1.getOrderRoleList().add(role);
//            }
//            if(i == 2) {
//                PublisherRoleBean role = new PublisherRoleBean();
//                role.setRole("服务员");
//                role.setRoleCurNum("6");
//                role.setRolePercentage("30%");
//                role.setRoleMoney("300.00");
//                bean1.getOrderRoleList().add(role);
//            }
//
//        }
//        orderListItems.add(bean1);
//
//        PublisherOrderBean bean2 = new PublisherOrderBean();
//        //bean2.setOrderTime("2018-04-22 19:24");
//        bean2.setOrderPrice("870");
//        bean2.setOrderDeskNum("2号桌");
//        for (int i = 0; i < 3; i++) {
//            if(i == 0) {
//                PublisherRoleBean role = new PublisherRoleBean();
//                role.setRole("经理");
//                role.setRoleCurNum("1");
//                role.setRolePercentage("40%");
//                role.setRoleMoney("400.00");
//                bean2.getOrderRoleList().add(role);
//            }
//            if(i == 1) {
//                PublisherRoleBean role = new PublisherRoleBean();
//                role.setRole("厨师");
//                role.setRoleCurNum("3");
//                role.setRolePercentage("30%");
//                role.setRoleMoney("90.00");
//                bean2.getOrderRoleList().add(role);
//            }
//            if(i == 2) {
//                PublisherRoleBean role = new PublisherRoleBean();
//                role.setRole("服务员");
//                role.setRoleCurNum("6");
//                role.setRolePercentage("30%");
//                role.setRoleMoney("300.00");
//                bean2.getOrderRoleList().add(role);
//            }
//        }
//        orderListItems.add(bean2);
//
//        PublisherOrderBean bean3 = new PublisherOrderBean();
//        //bean3.setOrderTime("2018-04-15 21:05");
//        bean3.setOrderPrice("870");
//        bean3.setOrderDeskNum("7号桌");
//        for (int i = 0; i < 3; i++) {
//            if(i == 0) {
//                PublisherRoleBean role = new PublisherRoleBean();
//                role.setRole("经理");
//                role.setRoleCurNum("1");
//                role.setRolePercentage("40%");
//                role.setRoleMoney("400.00");
//                bean3.getOrderRoleList().add(role);
//            }
//            if(i == 1) {
//                PublisherRoleBean role = new PublisherRoleBean();
//                role.setRole("厨师");
//                role.setRoleCurNum("5");
//                role.setRolePercentage("40%");
//                role.setRoleMoney("400.00");
//                bean3.getOrderRoleList().add(role);
//            }
//            if(i == 2) {
//                PublisherRoleBean role = new PublisherRoleBean();
//                role.setRole("服务员");
//                role.setRoleCurNum("6");
//                role.setRolePercentage("20%");
//                role.setRoleMoney("200.00");
//                bean3.getOrderRoleList().add(role);
//            }
//        }
//        orderListItems.add(bean3);
//
//    }

    /**
     * 显示订单列表数据
     */
    private void showOrderListData() {
        orderAdapter.setData(orderListItems);

        if(orderListItems == null || orderListItems.size() == 0) {
            llEmptyView.setVisibility(View.VISIBLE);
        } else {
            llEmptyView.setVisibility(View.GONE);
        }
    }


    /**
     * 获取当前排班下所有订单
     */
    private void getOrderByScheduleRequest(final PublisherSchedulingBean scheduleBean) {
        showLoading();
        GetAllOrderByScheduleTask task = new GetAllOrderByScheduleTask();
        protocol.Api.ReqGetAllOrder.Builder requestBuilder = protocol.Api.ReqGetAllOrder.newBuilder();
        requestBuilder.setUserAddress(UserInfo.USER_ADDRESS);
        requestBuilder.setJobAddress(scheduleBean.getScheduleAddress());
        task.setRequest(requestBuilder.build());
        task.setUpdateUIInterface(new GetAllOrderByScheduleTask.UpdateUIInterface() {
            @Override
            public void onSucceed(protocol.Api.RespGetAllOrder result) {
                Logger.d("GetAllOrderByScheduleTask更新UI：" + result);
                hideLoading();
                if(result.getStatusCode() == protocol.StatusOuterClass.Status.Success.ordinal()) {
                    //hintDialog(result.toString());

                    if(orderListItems != null) {
                        orderListItems.removeAll(orderListItems);
                    } else {
                        orderListItems = new ArrayList<PublisherOrderBean>();
                    }

                    List<protocol.Obj.Order> orderList = result.getOrdersList();
                    for (int i = 0; i < orderList.size(); i++) {
                        protocol.Obj.Order order = orderList.get(i);
                        PublisherOrderBean bean = new PublisherOrderBean();
                        bean.setOrderDeskNum(order.getTable() + "");
                        bean.setOrderPrice(order.getMoney() + "");
                        bean.setOrderTimeSp(order.getTimeStamp());
                        bean.setOrderRoleList(scheduleBean.getScheduleRoleList());
                        orderListItems.add(bean);
                    }

                   showOrderListData();

                } else {
                    // 请求失败
                    loadErrorDialog();
                }
            }

            @Override
            public void onError() {
                // 请求失败
                hideLoading();
                loadErrorDialog();
            }
        });
        task.execute(Variable.SERVER_IP, Variable.SERVER_PORT);
    }

}