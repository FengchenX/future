package com.znfz.assur.znfz_android.android.main.publisher.frament;

import android.graphics.Rect;
import android.os.Bundle;
import android.support.v7.widget.LinearLayoutManager;
import android.support.v7.widget.RecyclerView;
import android.view.View;
import android.widget.LinearLayout;
import android.widget.RelativeLayout;
import android.widget.TextView;

import com.orhanobut.logger.Logger;
import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.app.BaseFragment;
import com.znfz.assur.znfz_android.android.common.global.UserInfo;
import com.znfz.assur.znfz_android.android.common.global.Variable;
import com.znfz.assur.znfz_android.android.common.task.AddOrderAndPayTask;
import com.znfz.assur.znfz_android.android.common.task.GetApplyTask;
import com.znfz.assur.znfz_android.android.common.utils.RoleUtils;
import com.znfz.assur.znfz_android.android.common.utils.TimeUtils;
import com.znfz.assur.znfz_android.android.main.publisher.activity.AddEmployeeActivity;
import com.znfz.assur.znfz_android.android.main.publisher.activity.PublisherAddEmployeeActivity;
import com.znfz.assur.znfz_android.android.main.publisher.adapter.PublishEmployeesAdapter;
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherEmployeeBean;
import java.util.ArrayList;
import java.util.List;
import butterknife.BindView;
import butterknife.ButterKnife;
import protocol.Obj;

/**
 * 员工界面
 * Created by assur on 2018/4/24.
 */
public class PublisherFragmentEmployees extends BaseFragment {


    @BindView(R.id.ll_empty_view)                       // 无数据视图
    LinearLayout llEmptyView;
    @BindView(R.id.publisher_employees_role_tv)         // 我的角色
    TextView tvUserRole;
    @BindView(R.id.publisher_employees_username_tv)     // 我的用户名
    TextView tvUserName;
    @BindView(R.id.publisher_employees_userphone_tv)    // 我的手机号
    TextView tvUserPhone;
    @BindView(R.id.publisher_employees_add_rl)           // 添加员工
    RelativeLayout rlAddEmployee;
    @BindView(R.id.publisher_employees_listview)        // 员工列表视图
    RecyclerView employeesRecyclerView;

    private List<PublisherEmployeeBean> listItems;           // 员工列表数据
    private PublishEmployeesAdapter publishEmployeesAdapter; // 员工列表适配器

    @Override
    protected int setView() {
        return R.layout.frement_publisher_employees;
    }

    @Override
    protected void init(View view) {
        ButterKnife.bind(this, view);

        rlAddEmployee.setOnClickListener(this);
    }

    @Override
    protected void initData(Bundle savedInstanceState) {
        //tvUserRole.setText(RoleUtils.getUserRoleString(getContext(), UserInfo.USER_ROLE));
        tvUserName.setText(UserInfo.USER_NAME);
        tvUserPhone.setText(UserInfo.USER_PHONE);


        LinearLayoutManager layoutmanager = new LinearLayoutManager(getActivity());
        employeesRecyclerView.setLayoutManager(layoutmanager);
        employeesRecyclerView.addItemDecoration(new RecyclerView.ItemDecoration() {
            @Override
            public void getItemOffsets(Rect outRect, View view, RecyclerView parent, RecyclerView.State state) {
                // 设置分割线
                outRect.set(10, 10, 10, 10);
            }
        });
        publishEmployeesAdapter = new PublishEmployeesAdapter(getActivity());
        employeesRecyclerView.setAdapter(publishEmployeesAdapter);
        //initListItemsData();
        getApplyRequest();
    }

    @Override
    public void onClick(View v) {
        super.onClick(v);

        switch(v.getId()) {
            case R.id.publisher_employees_add_rl:

                startActivity(PublisherAddEmployeeActivity.class);
                break;

            default:

                break;

        }
    }

    /**
     * 初始化员工列表数据
     */
    private void initListItemsData() {

        if(listItems != null) {
            listItems.removeAll(listItems);
        } else {
            listItems = new ArrayList<PublisherEmployeeBean>();
        }
        for (int i = 0; i < 10; i++){
            if(i%2 == 0) {
                PublisherEmployeeBean bean = new PublisherEmployeeBean();
                bean.setPublisherEmployeeRole("厨师");
                bean.setPublisherEmployeeName("丽丽" + i);
                bean.setPublisherEmployeePhone("15167269825");
                //listItems.add(bean);
            }else if(i%2 == 1) {
                PublisherEmployeeBean bean = new PublisherEmployeeBean();
                bean.setPublisherEmployeeRole("服务员");
                bean.setPublisherEmployeeName("刘俊落" + i);
                bean.setPublisherEmployeePhone("15111260025");
                //listItems.add(bean);
            }
        }

    }


    /**
     * 添加申请情况
     */
    private void getApplyRequest() {
        if(UserInfo.currentSchedule == null) {
            return;
        }
        showLoading();
        GetApplyTask task = new GetApplyTask();
        protocol.Api.ReqGetStaff.Builder requestBuilder = protocol.Api.ReqGetStaff.newBuilder();
        requestBuilder.setJobAddress(UserInfo.currentSchedule.getScheduleAddress());
        requestBuilder.setUserAddress(UserInfo.USER_ADDRESS);
        task.setRequest(requestBuilder.build());
        task.setUpdateUIInterface(new GetApplyTask.UpdateUIInterface() {
            @Override
            public void onSucceed(protocol.Api.RespGetStaff result) {
                Logger.d("GetApplyTask更新UI：" + result);
                hideLoading();

                if(listItems != null) {
                    listItems.removeAll(listItems);
                } else {
                    listItems = new ArrayList<PublisherEmployeeBean>();
                }

                List<Obj.scheduleStaff> list = result.getStaffsList();
                for (int i = 0; i < list.size(); i++) {
                    PublisherEmployeeBean bean = new PublisherEmployeeBean();
                    bean.setPublisherEmployeeAddress(list.get(i).getStaffAddr());
                    bean.setPublisherEmployeeName("员工");
                    bean.setPublisherEmployeeRole(list.get(i).getRole() + "");
                    bean.setPublisherEmployeePer(list.get(i).getRatio());
                    listItems.add(bean);
                }

                publishEmployeesAdapter.setData(listItems);
                if(listItems == null || listItems.size() == 0) {
                    llEmptyView.setVisibility(View.VISIBLE);
                } else {
                    llEmptyView.setVisibility(View.GONE);
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
