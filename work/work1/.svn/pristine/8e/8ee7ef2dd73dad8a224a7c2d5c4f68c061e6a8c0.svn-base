package com.znfz.assur.znfz_android.android.main.applicant.frament;
import android.graphics.Rect;
import android.os.Bundle;
import android.support.v7.widget.LinearLayoutManager;
import android.support.v7.widget.RecyclerView;
import android.view.View;
import android.widget.LinearLayout;
import android.widget.TextView;

import com.orhanobut.logger.Logger;
import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.app.BaseFragment;
import com.znfz.assur.znfz_android.android.common.global.UserInfo;
import com.znfz.assur.znfz_android.android.common.global.Variable;
import com.znfz.assur.znfz_android.android.common.task.GetAllIncomeTask;
import com.znfz.assur.znfz_android.android.common.utils.RoleUtils;
import com.znfz.assur.znfz_android.android.main.publisher.adapter.PublishIncomeAdapter;
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherIncomeBean;
import java.util.ArrayList;
import java.util.List;
import butterknife.BindView;
import butterknife.ButterKnife;

/**
 * 收入界面
 * Created by assur on 2018/4/24.
 */

public class ApplicantFragmentIncome extends BaseFragment {

    @BindView(R.id.applicant_income_tv_username)      // 用户名
            TextView tvUserName;
    @BindView(R.id.applicant_income_tv_userrole)      // 用户角色
            TextView tvUserRole;
    @BindView(R.id.applicant_income_tv_userphone)     // 用户手机号
            TextView tvUserPhone;

    @BindView(R.id.ll_empty_view)                       // 无数据视图
    LinearLayout llEmptyView;
    @BindView(R.id.applicant_income_detail_listview)    // 收入明细列表
    RecyclerView incomeDetailRecyclerView;

    private List<PublisherIncomeBean> listItems;         // 收入列表数据
    private PublishIncomeAdapter applicantIncomeAdapter; // 收入列表适配器


    @Override
    protected int setView() {
        return R.layout.frement_applicant_income;
    }

    @Override
    protected void init(View view) {
        ButterKnife.bind(this, view);
    }

    @Override
    protected void initData(Bundle savedInstanceState) {

        tvUserName.setText(UserInfo.USER_PHONE);
        tvUserPhone.setText("手机号: " + UserInfo.USER_PHONE);
        tvUserRole.setText(RoleUtils.getUserRoleString(getContext(), UserInfo.USER_ROLE));

        LinearLayoutManager layoutmanager = new LinearLayoutManager(getActivity());
        incomeDetailRecyclerView.setLayoutManager(layoutmanager);
        incomeDetailRecyclerView.addItemDecoration(new RecyclerView.ItemDecoration() {
            @Override
            public void getItemOffsets(Rect outRect, View view, RecyclerView parent, RecyclerView.State state) {
                // 设置分割线
                outRect.set(0, 0, 0, 0);
            }
        });
        applicantIncomeAdapter = new PublishIncomeAdapter(getActivity());
        incomeDetailRecyclerView.setAdapter(applicantIncomeAdapter);
        //initListItemsData();

        getAllScheduleRequest();
    }

    /**
     * 初始化积分列表数据
     */
    private void initListItemsData() {

//        if(listItems != null) {
//            listItems.removeAll(listItems);
//        } else {
//            listItems = new ArrayList<PublisherIncomeBean>();
//        }
//        for (int i = 0; i < 20; i++){
//            PublisherIncomeBean bean = new PublisherIncomeBean();
//            bean.setIncomeTitle("服务佣金" + (i + 1));
//            bean.setIncomeTime("2018-04-0" + (i+1));
//            bean.setIncomeType(1);
//            bean.setIncomeCount("" + (i+1));
//            listItems.add(bean);
//        }
//
//        if(listItems == null || listItems.size() == 0) {
//            llEmptyView.setVisibility(View.VISIBLE);
//        } else {
//            llEmptyView.setVisibility(View.GONE);
//        }

    }


    /**
     * 获取所有排班请求
     */
    private void getAllScheduleRequest() {
        showLoading();
        GetAllIncomeTask task = new GetAllIncomeTask();
        protocol.Api.ReqGetAllIncome.Builder requestBuilder = protocol.Api.ReqGetAllIncome.newBuilder();
        requestBuilder.setCompanyName(Variable.CONPANY_NAME);
        requestBuilder.setUserAddress(UserInfo.USER_ADDRESS);
        task.setRequest(requestBuilder.build());
        task.setUpdateUIInterface(new GetAllIncomeTask.UpdateUIInterface() {
            @Override
            public void onSucceed(protocol.Api.RespGetAllIncome result) {
                Logger.d("GetScheduleTask更新UI：" + result);
                hideLoading();
                if(result.getStatusCode() == protocol.StatusOuterClass.Status.Success.ordinal()) {

                    if(listItems != null) {
                        listItems.removeAll(listItems);
                    } else {
                        listItems = new ArrayList<>();
                    }

                    List<protocol.Obj.Order> orderList = result.getOrdersList();
                    for (int i = 0; i < orderList.size(); i++) {
                        protocol.Obj.Order order = orderList.get(i);

                        PublisherIncomeBean bean = new PublisherIncomeBean();
                        bean.setIncomeTitle("服务佣金");
                        bean.setIncomeTimeSp(order.getTimeStamp());
                        bean.setIncomeType(1);
                        bean.setIncomeCount("" + order.getGetMoney());
                        listItems.add(bean);
                    }
                    showListData();

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
     * 显示列表数据
     */
    private void showListData() {
        if(listItems == null || listItems.size() == 0) {
            llEmptyView.setVisibility(View.VISIBLE);
        } else {
            llEmptyView.setVisibility(View.GONE);
        }
        applicantIncomeAdapter.setData(listItems);
    }

}
