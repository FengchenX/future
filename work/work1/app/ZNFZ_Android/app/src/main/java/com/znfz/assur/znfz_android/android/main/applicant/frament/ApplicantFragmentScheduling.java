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
import com.znfz.assur.znfz_android.android.common.task.ApplyJobTask;
import com.znfz.assur.znfz_android.android.common.task.GetJobTask;
import com.znfz.assur.znfz_android.android.common.utils.RoleUtils;
import com.znfz.assur.znfz_android.android.main.applicant.adapter.ApplicantJobAdapter;
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherRoleBean;
import java.util.ArrayList;
import java.util.List;
import butterknife.BindView;
import butterknife.ButterKnife;
import protocol.Obj;

/**
 * 排班界面
 * Created by assur on 2018/4/24.
 */

public class ApplicantFragmentScheduling extends BaseFragment {

    @BindView(R.id.ll_empty_view)                             // 无数据视图
    LinearLayout llEmptyView;
    @BindView(R.id.applicant_btn_scheduling_all)              // 全部按钮
    Button btnAllJob;
    @BindView(R.id.applicant_btn_scheduling_been_applicant)   // 已申请按钮
    Button btnBeenApplicantJob;
    @BindView(R.id.applicant_listview_scheduling)             // 工作列表
    RecyclerView schedulingRecyclerView;


    private List<PublisherRoleBean> allListItems;              // 全部工作数据
    private List<PublisherRoleBean> beenApplicantListItems;    // 已申请工作数据
    private ApplicantJobAdapter jobAdapter;                    // 工作列表适配器

    @Override
    protected int setView() {
        return R.layout.frement_applicant_scheduling;
    }

    @Override
    protected void init(View view) {

        ButterKnife.bind(this, view);
        btnAllJob.setOnClickListener(this);
        btnBeenApplicantJob.setOnClickListener(this);
    }

    @Override
    protected void initData(Bundle savedInstanceState) {


        LinearLayoutManager layoutmanager = new LinearLayoutManager(getActivity());
        schedulingRecyclerView.setLayoutManager(layoutmanager);
        schedulingRecyclerView.addItemDecoration(new RecyclerView.ItemDecoration() {
            @Override
            public void getItemOffsets(Rect outRect, View view, RecyclerView parent, RecyclerView.State state) {
                // 设置分割线
                outRect.set(10, 10, 10, 10);
            }
        });
        jobAdapter = new ApplicantJobAdapter(getActivity());
        jobAdapter.setOnItemClickListener(new ApplicantJobAdapter.OnItemClickListener() {
            @Override
            public void onClick(int position, int viewID) {
                switch (viewID) {
                    case R.id.item_job_list_btn_applicant:
                        applyJobRequest(position);
                        break;

                    default:

                        break;
                }
            }
        });
        schedulingRecyclerView.setAdapter(jobAdapter);


        if(allListItems == null || allListItems.size() == 0) {
            llEmptyView.setVisibility(View.VISIBLE);
        } else {
            llEmptyView.setVisibility(View.GONE);
        }
        //initAllListItems();
        jobAdapter.setData(allListItems);
    }

    @Override
    public void onResume() {
        super.onResume();

        // 获取所有工作
        getAllJobRequest();
    }

    @Override
    public void onClick(View v) {
        super.onClick(v);
        switch (v.getId()) {
            case R.id.applicant_btn_scheduling_all:  // 全部工作
                btnAllJob.setBackground(getResources().getDrawable(R.drawable.button_shape_corners_blue));
                btnAllJob.setTextColor(getResources().getColor(R.color.white));
                btnBeenApplicantJob.setBackground(getResources().getDrawable(R.drawable.button_shape_corners_withe));
                btnBeenApplicantJob.setTextColor(getResources().getColor(R.color.black));


                if(allListItems == null || allListItems.size() == 0) {
                    llEmptyView.setVisibility(View.VISIBLE);
                } else {
                    llEmptyView.setVisibility(View.GONE);
                }

                // initAllListItems();
                jobAdapter.setData(allListItems);
                break;

            case R.id.applicant_btn_scheduling_been_applicant:  // 已申请工作
                btnAllJob.setBackground(getResources().getDrawable(R.drawable.button_shape_corners_withe));
                btnAllJob.setTextColor(getResources().getColor(R.color.black));
                btnBeenApplicantJob.setBackground(getResources().getDrawable(R.drawable.button_shape_corners_blue));
                btnBeenApplicantJob.setTextColor(getResources().getColor(R.color.white));

                if(beenApplicantListItems == null || beenApplicantListItems.size() == 0) {
                    llEmptyView.setVisibility(View.VISIBLE);
                } else {
                    llEmptyView.setVisibility(View.GONE);
                }

                // initBeenApplicantListItems();
                jobAdapter.setData(beenApplicantListItems);
                break;

            default:

                break;

        }
    }


    /**
     * 全部工作数据
     */
    private void initAllListItems() {
        if(allListItems != null) {
            allListItems.removeAll(allListItems);
        } else {
            allListItems = new ArrayList<PublisherRoleBean>();
        }

        for (int i = 0; i < 3; i++) {
            PublisherRoleBean role = new PublisherRoleBean();
            if (i == 0) {
                role.setRole("经理");
                role.setRolePercentage("40%");
                role.setRoleCurNum("1");
                //role.setRoleWorkTime("2018-05-02");
            }
            if (i == 1) {
                role.setRole("厨师");
                role.setRolePercentage("20%");
                role.setRoleCurNum("2");
                //role.setRoleWorkTime("2018-05-03");
            }
            if (i == 2) {
                role.setRole("服务员");
                role.setRolePercentage("40%");
                role.setRoleCurNum("1");
                //role.setRoleWorkTime("2018-05-04");
            }

            allListItems.add(role);
        }
    }

    /**
     * 已申请工作的数据
     */
    private void initBeenApplicantListItems() {
        if(beenApplicantListItems != null) {
            beenApplicantListItems.removeAll(beenApplicantListItems);
        } else {
            beenApplicantListItems = new ArrayList<>();
        }

        for (int i = 0; i < 100; i++) {
            PublisherRoleBean role = new PublisherRoleBean();
            if (i == 0) {
                role.setRole("经理");
                role.setRolePercentage("40%");
                role.setRoleCurNum("1");
                //role.setRoleWorkTimeSp("2018-05-02");
            }
            beenApplicantListItems.add(role);
        }
    }

    /**
     * 获取所有工作请求
     */
    private void getAllJobRequest() {
        showLoading();
        GetJobTask task = new GetJobTask();
        protocol.Api.ReqGetCanApply.Builder requestBuilder = protocol.Api.ReqGetCanApply.newBuilder();
        requestBuilder.setCompanyName(Variable.CONPANY_NAME);
        requestBuilder.setUserAddress(UserInfo.USER_ADDRESS);
        task.setRequest(requestBuilder.build());
        task.setUpdateUIInterface(new GetJobTask.UpdateUIInterface() {
            @Override
            public void onSucceed(protocol.Api.RespGetCanApply result) {
                Logger.d("GetJobTask更新UI：" + result);
                hideLoading();
                if(result.getStatusCode() == protocol.StatusOuterClass.Status.Success.ordinal()) {
                    //hintDialog(result.toString());
                    //showToast("查询排班成功！");
                    if(allListItems != null) {
                        allListItems.removeAll(allListItems);
                    } else {
                        allListItems = new ArrayList<PublisherRoleBean>();
                    }

                    if(beenApplicantListItems != null) {
                        beenApplicantListItems.removeAll(beenApplicantListItems);
                    } else {
                        beenApplicantListItems = new ArrayList<PublisherRoleBean>();
                    }

                    List<protocol.Obj.Job> jobList = result.getJobsList();
                    for (int i = 0; i < jobList.size(); i++) {
                        protocol.Obj.Job job = jobList.get(i);
                        PublisherRoleBean bean = new PublisherRoleBean();
                        bean.setRole(job.getRole() + "");
                        bean.setCompanyName(job.getCompany());
                        bean.setRoleCurNum(job.getCount() + "");
                        bean.setRolePercentage(job.getRadio() + "");
                        bean.setRoleWorkTimeSp(job.getTimeStamp());
                        bean.setBeenAppliciant(job.getHasApply());
                        bean.setJobAddress(job.getJobAddress());
                        if(bean.isBeenAppliciant()) {
                            beenApplicantListItems.add(bean);
                        }
                        allListItems.add(bean);
                    }
                    jobAdapter.setData(allListItems);

                    if(allListItems == null || allListItems.size() == 0) {
                        llEmptyView.setVisibility(View.VISIBLE);
                    } else {
                        llEmptyView.setVisibility(View.GONE);
                    }

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
     * 申请工作请求
     */
    private void applyJobRequest(int index) {
        PublisherRoleBean bean = allListItems.get(index);
        if(!bean.getRole().equals(UserInfo.USER_ROLE)) {
            hintDialog("您不能申请该职位，因为您是" + RoleUtils.getUserRoleString(getContext(), UserInfo.USER_ROLE));
            return;
        }
        showLoading();
        ApplyJobTask task = new ApplyJobTask();
        protocol.Api.ReqFindJob.Builder requestBuilder = protocol.Api.ReqFindJob.newBuilder();
        requestBuilder.setUserAddress(UserInfo.USER_ADDRESS);
        requestBuilder.setAccountDescribe(UserInfo.USER_PRIVATE_KEY);
        requestBuilder.setPassWord(UserInfo.USER_PASSWORD);
        requestBuilder.setMyJob(Obj.Job.newBuilder()
//                .setRole(Integer.parseInt(bean.getRole()))
                .setTimeStamp(bean.getRoleWorkTimeSp())
                .setCompany(bean.getCompanyName())
                .setRadio(Long.parseLong(bean.getRolePercentage()))
                .setCount(Integer.parseInt(bean.getRoleCurNum()))
                .setJobAddress(bean.getJobAddress())
                .build());
        task.setRequest(requestBuilder.build());
        task.setUpdateUIInterface(new ApplyJobTask.UpdateUIInterface() {
            @Override
            public void onSucceed(protocol.Api.RespFindJob result) {
                Logger.d("ApplyJobTask更新UI：" + result);
                hideLoading();
                if(result.getStatusCode() == protocol.StatusOuterClass.Status.Success.ordinal()) {
                    //hintDialog(result.toString());
                    showToast("申请排班成功！");
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
