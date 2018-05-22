package com.znfz.assur.znfz_android.android.main.publisher.frament;

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
import com.znfz.assur.znfz_android.android.common.task.GetScheduleTask;
import com.znfz.assur.znfz_android.android.main.publisher.activity.OrderListActivity;
import com.znfz.assur.znfz_android.android.main.publisher.activity.ScheduleListActivity;
import com.znfz.assur.znfz_android.android.main.publisher.adapter.PublishHistorySchedulingAdapter;
import com.znfz.assur.znfz_android.android.main.publisher.adapter.PublishOrderSchedulingAdapter;
import com.znfz.assur.znfz_android.android.main.publisher.adapter.PublishSchedulingAdapter;
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherRoleBean;
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherSchedulingBean;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Set;

import butterknife.BindView;
import butterknife.ButterKnife;
import protocol.Obj;

/**
 * 排班界面
 * Created by assur on 2018/4/24.
 */

public class PublisherFragmentScheduling extends BaseFragment {

    @BindView(R.id.ll_empty_view)                      // 无数据视图
    LinearLayout llEmptyView;
    @BindView(R.id.publisher_btn_scheduling_current)   // 当前排班
    Button btnCurrentSchedule;
    @BindView(R.id.publisher_btn_scheduling_history)   // 历史排班
    Button btnHistorySchedule;
    @BindView(R.id.publisher_listview_scheduling)      // 排班列表
    RecyclerView schedulingRecyclerView;


    private boolean isShowingHistoryScheduling;               // 是否正在显示历史排班
    private List<PublisherSchedulingBean> curListItems;       // 当前排班数据
    private PublishSchedulingAdapter curChedulingAdapter;     // 当前排班列表适配器

    private Map<String, List<PublisherSchedulingBean>>  historyScheduleMap;  // 历史排班数据（键为公司名，值为该公司的所有排班列表）
    private List<String> keyList;
    private PublishHistorySchedulingAdapter historyScheduleAdapter;          // 历史排班列表适配器


    @Override
    protected int setView() {
        return R.layout.frement_publisher_scheduling;
    }

    @Override
    protected void init(View view) {

        ButterKnife.bind(this, view);
        btnCurrentSchedule.setOnClickListener(this);
        btnHistorySchedule.setOnClickListener(this);
    }

    @Override
    protected void initData(Bundle savedInstanceState) {

        // 列表视图初始化
        LinearLayoutManager layoutmanager = new LinearLayoutManager(getActivity());
        schedulingRecyclerView.setLayoutManager(layoutmanager);
        schedulingRecyclerView.addItemDecoration(new RecyclerView.ItemDecoration() {
            @Override
            public void getItemOffsets(Rect outRect, View view, RecyclerView parent, RecyclerView.State state) {
                // 设置分割线
                outRect.set(10, 10, 10, 10);
            }
        });


        // 当前排班列表适配器初始化
        curChedulingAdapter = new PublishSchedulingAdapter(getActivity());
        curChedulingAdapter.setOnItemClickListener(new PublishSchedulingAdapter.onItemClickListener() {
            @Override
            public void onClick(int position, int viewID) {
                switch (viewID) {
                    case R.id.publisher_scheduling_list_item_rl_btn_detail:
                        curListItems.get(position).setShowRoleList(!curListItems.get(position).isShowRoleList());
                        curChedulingAdapter.setData(curListItems);
                        break;
                }
            }
        });


        // 历史排班列表适配器初始化
        historyScheduleAdapter = new PublishHistorySchedulingAdapter(getActivity());
        historyScheduleAdapter.setOnItemClickListener(new PublishHistorySchedulingAdapter.OnItemClickListener() {
            @Override
            public void onClick(int position, int viewID) {
                UserInfo.scheduleistItemsOfOneCompany = historyScheduleMap.get(keyList.get(position));
                startActivity(ScheduleListActivity.class);
            }
        });


        // 默认显示当前的排版
        updateListView(isShowingHistoryScheduling);
    }

    /**
     * 当前排版/历史排版列表切换
     * @param showingHistory
     */
    private void updateListView(boolean showingHistory) {
        if(showingHistory) {

            if(historyScheduleMap == null || historyScheduleMap.size() == 0) {
                llEmptyView.setVisibility(View.VISIBLE);
            } else {
                llEmptyView.setVisibility(View.GONE);
                schedulingRecyclerView.setAdapter(historyScheduleAdapter);
                if(keyList != null) {
                    keyList.clear();
                    keyList = null;
                }
                keyList = new ArrayList<String>(historyScheduleMap.keySet());
                historyScheduleAdapter.setData(keyList);

            }

        } else {

            if(curListItems == null || curListItems.size() == 0) {
                llEmptyView.setVisibility(View.VISIBLE);
            } else {
                llEmptyView.setVisibility(View.GONE);
                schedulingRecyclerView.setAdapter(curChedulingAdapter);
                curChedulingAdapter.setData(curListItems);
            }
        }
    }

    @Override
    public void onResume() {
        super.onResume();
        // 获取所有排班
        getAllPublishRequest();
    }

    @Override
    public void onClick(View v) {
        super.onClick(v);
        switch (v.getId()) {
            case R.id.publisher_btn_scheduling_current:  // 今日排班
                btnCurrentSchedule.setBackground(getResources().getDrawable(R.drawable.button_shape_corners_blue));
                btnCurrentSchedule.setTextColor(getResources().getColor(R.color.white));
                btnHistorySchedule.setBackground(getResources().getDrawable(R.drawable.button_shape_corners_withe));
                btnHistorySchedule.setTextColor(getResources().getColor(R.color.black));
                isShowingHistoryScheduling = false;
                updateListView(isShowingHistoryScheduling);
                break;

            case R.id.publisher_btn_scheduling_history:  // 历史排班
                btnCurrentSchedule.setBackground(getResources().getDrawable(R.drawable.button_shape_corners_withe));
                btnCurrentSchedule.setTextColor(getResources().getColor(R.color.black));
                btnHistorySchedule.setBackground(getResources().getDrawable(R.drawable.button_shape_corners_blue));
                btnHistorySchedule.setTextColor(getResources().getColor(R.color.white));
                isShowingHistoryScheduling = true;
                updateListView(isShowingHistoryScheduling);
                break;
            default:

                break;

        }
    }


//    /**
//     * 当前排班数据
//     */
//    private void initCurListItems() {
//        if(curListItems != null) {
//            curListItems.removeAll(curListItems);
//        } else {
//            curListItems = new ArrayList<PublisherSchedulingBean>();
//        }
//
//        PublisherSchedulingBean bean1 = new PublisherSchedulingBean();
//        bean1.setCurSchedule(true);
//        for (int i = 0; i < 3; i++) {
//            if(i == 0) {
//                PublisherRoleBean role = new PublisherRoleBean();
//                role.setRole("经理");
//                role.setRoleCurNum("1");
//                role.setRolePercentage("40%");
//                role.setRoleCurNum("1");
//                bean1.getScheduleRoleList().add(role);
//            }
//            if(i == 1) {
//                PublisherRoleBean role = new PublisherRoleBean();
//                role.setRole("厨师");
//                role.setRoleCurNum("3");
//                role.setRolePercentage("20%");
//                role.setRoleCurNum("2");
//                bean1.getScheduleRoleList().add(role);
//            }
//            if(i == 2) {
//                PublisherRoleBean role = new PublisherRoleBean();
//                role.setRole("服务员");
//                role.setRoleCurNum("4");
//                role.setRolePercentage("40%");
//                role.setRoleCurNum("1");
//                bean1.getScheduleRoleList().add(role);
//            }
//
//        }
//        curListItems.add(bean1);
//
//    }

//    /**
//     * 历史排班数据
//     */
//    private void initHistoryListItems() {
//        if(historyListItems != null) {
//            historyListItems.removeAll(historyListItems);
//        } else {
//            historyListItems = new ArrayList<PublisherSchedulingBean>();
//        }
//
//        for (int i = 0; i < 100; i++) {
//            PublisherSchedulingBean bean = new PublisherSchedulingBean();
//            for (int j = 0; j < 3; j++) {
//                if(j == 0) {
//                    PublisherRoleBean role = new PublisherRoleBean();
//                    role.setRole("经理");
//                    role.setRoleCurNum("1");
//                    role.setRolePercentage("40%");
//                    role.setRoleCurNum("1");
//                    bean.getScheduleRoleList().add(role);
//                }
//                if(j == 1) {
//                    PublisherRoleBean role = new PublisherRoleBean();
//                    role.setRole("厨师");
//                    role.setRoleCurNum("3");
//                    role.setRolePercentage("20%");
//                    role.setRoleCurNum("2");
//                    bean.getScheduleRoleList().add(role);
//                }
//                if(j == 2) {
//                    PublisherRoleBean role = new PublisherRoleBean();
//                    role.setRole("服务员");
//                    role.setRoleCurNum("4");
//                    role.setRolePercentage("40%");
//                    role.setRoleCurNum("1");
//                    bean.getScheduleRoleList().add(role);
//                }
//            }
//            historyListItems.add(bean);
//        }
//    }


    /**
     * 获取所有排班请求
     */
    private void getAllPublishRequest() {
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
                hideLoading();
                if(result.getStatusCode() == protocol.StatusOuterClass.Status.Success.ordinal()) {
                    //hintDialog(result.toString());
                    //showToast("查询排班成功！");

                    if(curListItems != null) {
                        curListItems.removeAll(curListItems);
                    } else {
                        curListItems = new ArrayList<PublisherSchedulingBean>();
                    }


                    if(historyScheduleMap != null) {
                        historyScheduleMap.clear();
                    } else {
                        historyScheduleMap = new HashMap<String, List<PublisherSchedulingBean>>();
                    }
                    String companyName = "";
                    List<PublisherSchedulingBean> oneCompanyScheduleList = new ArrayList<>();


                    List<protocol.Obj.Schedule> scheduleList = result.getSchedulesList();
                    for (int i = 0; i < scheduleList.size(); i++) {
                        protocol.Obj.Schedule schedule = scheduleList.get(i);
                        PublisherSchedulingBean bean = new PublisherSchedulingBean();
                        bean.setScheduleAddress(schedule.getJobaddr());
                        bean.setScheduleFatherAddress(schedule.getFatherAddress());
                        bean.setScheduleCompanyPer(schedule.getCompanyRatio());
                        bean.setScheduleTimeSp(schedule.getTimeStamp());
                        bean.setScheduleCompany(schedule.getCompanyName());
                        bean.setScheduleCompanyPer(schedule.getCompanyRatio());
                        bean.setScheduleStorePer(100 - schedule.getCompanyRatio());

                        List<PublisherRoleBean> roleBeanList = new ArrayList<>();
                        List<Obj.Job> jobList = schedule.getJobsList();
                        for (int j = 0; j < jobList.size(); j++) {
                            Obj.Job job = jobList.get(j);
                            PublisherRoleBean role = new PublisherRoleBean();
                            role.setJobId(job.getJobId());
                            role.setCompanyName(job.getCompany());
                            role.setRoleCurNum(job.getCount() + "");
                            role.setRolePercentage(job.getRadio() + "");
                            role.setJobAddress(job.getJobAddress());
                            role.setRole(job.getRole());
                            role.setRoleWorkTimeSp(job.getTimeStamp());
                            role.setWhitelist(job.getWhiteListList());
                            role.setWhiteName(job.getWhiteName());
                            roleBeanList.add(role);
                        }
                        bean.setScheduleRoleList(roleBeanList);


                        if(i == 0) {
                            // 当前排版（第一个排版为当前的，后台已经排序）
                            curListItems.add(bean);
                        } else {
                            // 历史排版
                            if(!companyName.equals(bean.getScheduleCompany())) {
                                if(companyName.length() == 0) { // 第一次
                                    companyName = bean.getScheduleCompany();
                                } else {
                                    historyScheduleMap.put(companyName, new ArrayList<PublisherSchedulingBean>(oneCompanyScheduleList));
                                    oneCompanyScheduleList.clear();
                                    companyName = bean.getScheduleCompany();
                                    Logger.e("公司名：" + companyName + "，这个公司的排班数量：" + oneCompanyScheduleList.size());
                                }
                            } else {
                                oneCompanyScheduleList.add(bean);
                                if(i == scheduleList.size() - 1) {
                                    historyScheduleMap.put(companyName, new ArrayList<PublisherSchedulingBean>(oneCompanyScheduleList));
                                    Logger.e("公司名：" + companyName + "，这个公司的排班数量：" + oneCompanyScheduleList.size());
                                }
                            }
                        }

                    }

                    updateListView(isShowingHistoryScheduling);

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
