package com.znfz.assur.znfz_android.android.main.applicant.frament;

import android.graphics.Rect;
import android.os.Bundle;
import android.support.v7.widget.LinearLayoutManager;
import android.support.v7.widget.RecyclerView;
import android.view.View;
import android.widget.Button;
import android.widget.LinearLayout;

import com.orhanobut.logger.Logger;
import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.app.BaseFragment;
import com.znfz.assur.znfz_android.android.common.global.UserInfo;
import com.znfz.assur.znfz_android.android.common.global.Variable;
import com.znfz.assur.znfz_android.android.common.task.GetAllOrderByScheduleTask;
import com.znfz.assur.znfz_android.android.common.task.GetScheduleTask;
import com.znfz.assur.znfz_android.android.main.publisher.activity.AddOrderActivity;
import com.znfz.assur.znfz_android.android.main.publisher.activity.OrderListActivity;
import com.znfz.assur.znfz_android.android.main.publisher.adapter.PublishOrderAdapter;
import com.znfz.assur.znfz_android.android.main.publisher.adapter.PublishOrderSchedulingAdapter;
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherOrderBean;
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherRoleBean;
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherSchedulingBean;
import java.util.ArrayList;
import java.util.List;
import butterknife.BindView;
import butterknife.ButterKnife;

/**
 * 订单界面
 * Created by assur on 2018/4/24.
 */
public class ApplicantFragmentOrder extends BaseFragment {


    @BindView(R.id.ll_empty_view)               // 无数据视图
    LinearLayout llEmptyView;
    @BindView(R.id.applicant_btn_order_today)   // 今日订单按钮
    Button btnTodayOrdre;
    @BindView(R.id.applicant_btn_order_history) // 历史订单按钮
    Button btnHistoryOrder;
    @BindView(R.id.applicant_listview_order)    // 订单列表
    RecyclerView orderRecyclerView;


    private List<PublisherOrderBean> orderListItems;       // 今日订单数据
    private PublishOrderAdapter orderAdapter;              // 今日订单列表适配器

    private List<PublisherSchedulingBean> historyScheduleListItems;   // 历史排班数据
    private PublishOrderSchedulingAdapter historyScheduleAdapter;     // 历史排班列表适配器



    @Override
    protected int setView() {
        return R.layout.frement_applicant_order;
    }

    @Override
    protected void init(View view) {
        ButterKnife.bind(this, view);

        // 列表视图
        LinearLayoutManager layoutmanager = new LinearLayoutManager(getActivity());
        orderRecyclerView.setLayoutManager(layoutmanager);
        orderRecyclerView.addItemDecoration(new RecyclerView.ItemDecoration() {
            @Override
            public void getItemOffsets(Rect outRect, View view, RecyclerView parent, RecyclerView.State state) {
                // 设置分割线
                outRect.set(10, 10, 10, 10);
            }
        });

        // 点击事件
        btnTodayOrdre.setOnClickListener(this);
        btnHistoryOrder.setOnClickListener(this);
    }

    @Override
    protected void initData(Bundle savedInstanceState) {

        // 今日订单列表适配器初始化
        orderAdapter = new PublishOrderAdapter(getActivity());
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

        // 历史订单的排班列表适配器初始化
        historyScheduleAdapter = new PublishOrderSchedulingAdapter(getActivity());
        historyScheduleAdapter.setOnItemClickListener(new PublishOrderSchedulingAdapter.OnItemClickListener() {
            @Override
            public void onClick(int position, int viewID) {
                UserInfo.currentHistorySchedule = historyScheduleListItems.get(position);
                startActivity(OrderListActivity.class);

            }
        });

        // 初始化今日列表数据
        //initTodayListItems();

        // 初始化历史排班列表数据
        //initHistoryScheduleListItems();

        // 获取当前排班请求
        getAllScheduleRequest();

        // 默认显示今天的订单
        showTodayListData();
    }



    @Override
    public void onClick(View v) {
        super.onClick(v);
        switch (v.getId()) {
            case R.id.applicant_btn_order_today:  // 今日订单
                btnTodayOrdre.setBackground(getResources().getDrawable(R.drawable.button_shape_corners_blue));
                btnTodayOrdre.setTextColor(getResources().getColor(R.color.white));
                btnHistoryOrder.setBackground(getResources().getDrawable(R.drawable.button_shape_corners_withe));
                btnHistoryOrder.setTextColor(getResources().getColor(R.color.black));
                showTodayListData();
                break;

            case R.id.applicant_btn_order_history:  // 历史订单
                btnTodayOrdre.setBackground(getResources().getDrawable(R.drawable.button_shape_corners_withe));
                btnTodayOrdre.setTextColor(getResources().getColor(R.color.black));
                btnHistoryOrder.setBackground(getResources().getDrawable(R.drawable.button_shape_corners_blue));
                btnHistoryOrder.setTextColor(getResources().getColor(R.color.white));
                showScheduleListData();
                break;

            default:

                break;

        }
    }

    /**
     * 显示今日列表数据
     */
    private void showTodayListData() {
        orderAdapter.setData(orderListItems);
        orderRecyclerView.setAdapter(orderAdapter);

        if(orderListItems == null || orderListItems.size() == 0) {
            llEmptyView.setVisibility(View.VISIBLE);
        } else {
            llEmptyView.setVisibility(View.GONE);
        }
    }


    /**
     * 显示历史排班列表数据
     */
    private void showScheduleListData() {
        historyScheduleAdapter.setData(historyScheduleListItems);
        orderRecyclerView.setAdapter(historyScheduleAdapter);

        if(historyScheduleListItems == null || historyScheduleListItems.size() == 0) {
            llEmptyView.setVisibility(View.VISIBLE);
        } else {
            llEmptyView.setVisibility(View.GONE);
        }
    }

//
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
//
//    /**
//     * 历史订单数据
//     */
//    private void initHistoryScheduleListItems() {
//        if(historyScheduleListItems != null) {
//            historyScheduleListItems.removeAll(historyScheduleListItems);
//        } else {
//            historyScheduleListItems = new ArrayList<PublisherSchedulingBean>();
//        }
//
//        for (int i = 0; i < 100; i++) {
//            PublisherSchedulingBean bean = new PublisherSchedulingBean();
//            bean.setScheduleCompany("公司" + i);
//            historyScheduleListItems.add(bean);
//        }
//    }

    /**
     * 获取所有排班请求
     */
    private void getAllScheduleRequest() {
        showLoading();
        GetScheduleTask task = new GetScheduleTask();
        protocol.Api.ReqGetSchedue.Builder requestBuilder = protocol.Api.ReqGetSchedue.newBuilder();
        requestBuilder.setCompanyName(Variable.CONPANY_NAME);
        requestBuilder.setUserAddress(UserInfo.USER_ADDRESS);
        task.setRequest(requestBuilder.build());
        task.setUpdateUIInterface(new GetScheduleTask.UpdateUIInterface() {
            @Override
            public void onSucceed(protocol.Api.RespGetSchedue result) {
                Logger.d("GetScheduleTask更新UI：" + result);
                //hideLoading();

                if(historyScheduleListItems != null) {
                    historyScheduleListItems.removeAll(historyScheduleListItems);
                } else {
                    historyScheduleListItems = new ArrayList<PublisherSchedulingBean>();
                }

                if(result.getStatusCode() == protocol.StatusOuterClass.Status.Success.ordinal()) {
                    List<protocol.Obj.Schedule> scheduleList = result.getSchedulesList();
                    for (int i = 0; i < scheduleList.size(); i++) {
                        protocol.Obj.Schedule schedule = scheduleList.get(i);
                        PublisherSchedulingBean bean = new PublisherSchedulingBean();
                        bean.setScheduleAddress(schedule.getJobaddr());
                        bean.setScheduleTimeSp(schedule.getTimeStamp());
                        bean.setScheduleCompany(schedule.getCompanyName());
                        List<PublisherRoleBean> roleBeanList = new ArrayList<>();
                        List<protocol.Obj.Job> jobList = schedule.getJobsList();
                        for (int j = 0; j < jobList.size(); j++) {
                            protocol.Obj.Job job = jobList.get(j);
                            PublisherRoleBean role = new PublisherRoleBean();
                            role.setRole(job.getRole() + "");
                            role.setRolePercentage(job.getRadio() + "");
                            role.setRoleCurNum(job.getCount() + "");
                            roleBeanList.add(role);
                        }
                        bean.setScheduleRoleList(roleBeanList);

                        if(i == 0) {
                            // 第一条数据是最后一次发布的排班，作为当前排班
                            UserInfo.currentSchedule = bean;
                        } else {
                            historyScheduleListItems.add(bean);
                        }
                    }

                    // 获取当前排班下的所有订单
                    getOrderByScheduleRequest(UserInfo.currentSchedule);

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
                    showTodayListData();

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
