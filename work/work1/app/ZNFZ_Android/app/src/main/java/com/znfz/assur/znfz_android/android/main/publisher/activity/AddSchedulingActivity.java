package com.znfz.assur.znfz_android.android.main.publisher.activity;

import android.graphics.Rect;
import android.support.v7.widget.LinearLayoutManager;
import android.support.v7.widget.RecyclerView;
import android.view.Gravity;
import android.view.View;
import android.widget.AdapterView;
import android.widget.ImageView;
import android.widget.TextView;
import com.bigkoo.pickerview.TimePickerView;
import com.flyco.dialog.entity.DialogMenuItem;
import com.flyco.dialog.listener.OnBtnClickL;
import com.flyco.dialog.listener.OnOperItemClickL;
import com.flyco.dialog.widget.ActionSheetDialog;
import com.flyco.dialog.widget.NormalDialog;
import com.orhanobut.logger.Logger;
import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.app.BaseActivity;
import com.znfz.assur.znfz_android.android.common.global.UserInfo;
import com.znfz.assur.znfz_android.android.common.global.Variable;
import com.znfz.assur.znfz_android.android.common.task.AddScheduleTask;
import com.znfz.assur.znfz_android.android.common.utils.RoleUtils;
import com.znfz.assur.znfz_android.android.common.utils.TimeUtils;
import com.znfz.assur.znfz_android.android.main.publisher.adapter.PublishAddSchedulingAdapter;
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherRoleBean;
import java.util.ArrayList;
import java.util.Date;
import java.util.List;
import butterknife.BindView;
import butterknife.ButterKnife;
import protocol.Obj;

/**
 * 添加排班界面
 * Created by assur on 2018/4/26
 */
public class AddSchedulingActivity extends BaseActivity {

    @BindView(R.id.publish_add_scheduling_et_company)    // 上班公司名称
    TextView mETCompany;
    @BindView(R.id.publish_add_scheduling_et_time)       // 上班时间
    TextView mETWorkTime;
    @BindView(R.id.publish_add_scheduling_iv_timeselect) // 上班时间选择
    ImageView mIVWorkTimeChoice;
    @BindView(R.id.publisher_add_scheduling_listview)    // 添加排班职位信息列表
    RecyclerView addSchedulingRecyclerView;

    private List<PublisherRoleBean> listItems;                        // 排班数据列表
    private PublishAddSchedulingAdapter publishAddSchedulingAdapter;  // 排班数据列表适配器


    @Override
    protected void initView() {
        setContentView(R.layout.activity_publish_add_scheduling);
        ButterKnife.bind(this);
        initToolBar(getStringById(R.string.publisher_add_scheduling_title), true);
        setToolbarRightTv(getStringById(R.string.publisher_add_scheduling_ok));

        mETCompany.setText(Variable.CONPANY_NAME);
        mIVWorkTimeChoice.setOnClickListener(this);
    }


    @Override
    protected void initData() {
        publishAddSchedulingAdapter = new PublishAddSchedulingAdapter(AddSchedulingActivity.this);
        publishAddSchedulingAdapter.setOnItemClickListener(new PublishAddSchedulingAdapter.OnItemClickListener() {
            @Override
            public void onClick(final int position, int viewID) {
                switch (viewID) {
                    case R.id.publish_add_scheduling_iv_rolenumselect:  // 选择人数
                        final ArrayList<DialogMenuItem> numItemArray = new ArrayList<>();
                        for (int i = 1; i < 21; i ++) {
                            DialogMenuItem item = new DialogMenuItem(i + "" , i);
                            numItemArray.add(item);
                        }
                        final ActionSheetDialog numDialog = new ActionSheetDialog(context, numItemArray, null);
                        numDialog.title(getStringById(R.string.publisher_add_scheduling_rolenum_hint))//
                                .titleTextSize_SP(14.5f)
                                .show();
                        numDialog.setOnOperItemClickL(new OnOperItemClickL() {
                            @Override
                            public void onOperItemClick(AdapterView<?> parent, View view, int p, long id) {
                                numDialog.dismiss();
                                listItems.get(position).setRoleCurNum(numItemArray.get(p).mOperName);
                                publishAddSchedulingAdapter.setData(listItems);
                            }
                        });
                        break;

                    case R.id.publish_add_scheduling_iv_perselect:  // 选择百分比
                        final ArrayList<DialogMenuItem> perItemArray = new ArrayList<>();
                        for (int i = 1; i < 100; i ++) {
                            DialogMenuItem item = new DialogMenuItem(i + "" , i);
                            perItemArray.add(item);
                        }
                        final ActionSheetDialog perDialog = new ActionSheetDialog(context, perItemArray, null);
                        perDialog.title(getStringById(R.string.publisher_add_scheduling_per_hint))//
                                .titleTextSize_SP(14.5f)
                                .show();
                        perDialog.setOnOperItemClickL(new OnOperItemClickL() {
                            @Override
                            public void onOperItemClick(AdapterView<?> parent, View view, int p, long id) {
                                perDialog.dismiss();
                                listItems.get(position).setRolePercentage(perItemArray.get(p).mOperName);
                                publishAddSchedulingAdapter.setData(listItems);
                            }
                        });
                        break;

                    default:

                            break;
                }
            }
        });

        LinearLayoutManager layoutmanager = new LinearLayoutManager(AddSchedulingActivity.this);
        addSchedulingRecyclerView.setLayoutManager(layoutmanager);
        addSchedulingRecyclerView.addItemDecoration(new RecyclerView.ItemDecoration() {
            @Override
            public void getItemOffsets(Rect outRect, View view, RecyclerView parent, RecyclerView.State state) {
                // 设置分割线
                outRect.set(10, 20, 10, 20);
            }
        });
        addSchedulingRecyclerView.setAdapter(publishAddSchedulingAdapter);
        initListItems();
        publishAddSchedulingAdapter.setData(listItems);
    }

    @Override
    public void onClick(View v) {
        super.onClick(v);

        switch (v.getId()) {

            case R.id.tv_right:      // 完成按钮
                check();
                break;

            case R.id.publish_add_scheduling_iv_timeselect:
                hideKeyboard();
                TimePickerView pvTime = new TimePickerView.Builder(AddSchedulingActivity.this, new TimePickerView.OnTimeSelectListener() {
                    @Override
                    public void onTimeSelect(Date date, View v) {
                        mETWorkTime.setText(TimeUtils.getTime(date.getTime()));
                    }
                }).setType(new boolean[]{true, true, true, true, true, true})// 默认全部显示
                        .setCancelText(getStringById(R.string.cancel))//取消按钮文字
                        .setSubmitText(getStringById(R.string.ok))    //确认按钮文字
                        .setOutSideCancelable(false)//点击屏幕，点在控件外部范围时，是否取消显示
                        .setLabel("年","月","日","时","分","秒")//默认设置为年月日时分秒
                        .isCenterLabel(false) //是否只显示中间选中项的label文字，false则每项item全部都带有label。
                        .build();
                pvTime.show();

                break;

            default:

                    break;
        }
    }

    /**
     * 角色数据初始化
     */
    private void initListItems() {

        if(listItems != null) {
            listItems.removeAll(listItems);
        } else {
            listItems = new ArrayList<PublisherRoleBean>();
        }

        PublisherRoleBean bean2 = new PublisherRoleBean();
        bean2.setRole(Variable.USER_ROLE_APPLICANT_COOK);
        listItems.add(bean2);

        PublisherRoleBean bean3 = new PublisherRoleBean();
        bean3.setRole(Variable.USER_ROLE_APPLICANT_WAITER);
        listItems.add(bean3);
    }

    /**
     * 检查填写信息正确性
     */
    private void check() {

        if (!isNetConnected()) {
            // 网络没有连接
            toast(R.string.no_net_message);
            return;
        }
        // 判断所有角色的时间需要一致
        int myPercentage = 100;
        if(mETCompany.getText().toString() == null ||mETCompany.getText().toString().length() == 0) {
            hintDialog("请填写公司名称");
            return;
        }
        if(mETWorkTime.getText().toString() == null ||mETWorkTime.getText().toString().length() == 0) {
            hintDialog("请选择上班日期");
            return;
        }
        for (int i = 0; i < listItems.size(); i++) {
            PublisherRoleBean bean = listItems.get(i);

            if(bean.getRoleCurNum().length() == 0) {
                hintDialog(RoleUtils.getUserRoleString(this,bean.getRole())  + "人数未选择");
                return;
            }

            if(bean.getRolePercentage().length() == 0) {
                hintDialog(RoleUtils.getUserRoleString(this,bean.getRole())  + "佣金比例未选择");
                return;
            }

            // 剩余我的佣金分配比例
            myPercentage = myPercentage - Integer.valueOf(bean.getRolePercentage()) * Integer.valueOf(bean.getRoleCurNum());
        }

        if(myPercentage < 0) {
            hintDialog("分配比例不对，请您检查");
            return;
        }

        final NormalDialog dialog = new NormalDialog(context);
        dialog.content("您的佣金比例为: " + myPercentage  + "%, 确定添加排班吗？")
                .btnNum(2)
                .style(NormalDialog.STYLE_TWO)
                .contentGravity(Gravity.CENTER)
                .btnText("确定", "继续修改")
                .show();

        dialog.setOnBtnClickL(
                new OnBtnClickL() {
                    @Override
                    public void onBtnClick() {
                        dialog.dismiss();
                        addScheduleRequest();
                    }
                },
                new OnBtnClickL() {
                    @Override
                    public void onBtnClick() {
                        dialog.dismiss();
                    }
                });

    }


    /**
     * 添加排班请求（以最后一次发布的作为当前排班）
     */
    private void addScheduleRequest() {
//        showLoading();
//        AddScheduleTask task = new AddScheduleTask();
//        protocol.Api.ReqScheduling.Builder requestBuilder = protocol.Api.ReqScheduling.newBuilder();
//        requestBuilder.setUserAddress(UserInfo.USER_ADDRESS);
//        requestBuilder.setCompany(mETCompany.getText().toString());
//        requestBuilder.setPassWord(UserInfo.USER_PASSWORD);
//        requestBuilder.setTimeStamp(TimeUtils.getTimeSp(mETWorkTime.getText().toString()));
//        requestBuilder.setAccountDescribe(UserInfo.USER_PRIVATE_KEY);
//        for(int i = 0; i < listItems.size(); i++) {
//            PublisherRoleBean bean = listItems.get(i);
//            Obj.Job job =  Obj.Job.newBuilder()
//                    .setRole(Integer.valueOf(bean.getRole()))
//                    .setCount(Integer.valueOf(bean.getRoleCurNum()))
//                    .setRadio(Integer.valueOf(bean.getRolePercentage()))
//                    .setCompany(mETCompany.getText().toString())
//                    .setTimeStamp(TimeUtils.getTimeSp(mETWorkTime.getText().toString()))
//                    .build();
//            requestBuilder.addJobs(job);
//        }
//        task.setRequest(requestBuilder.build());
//        task.setUpdateUIInterface(new AddScheduleTask.UpdateUIInterface() {
//            @Override
//            public void onSucceed(protocol.Api.RespScheduling result) {
//                Logger.d("AddScheduleTask更新UI：" + result);
//                hideLoading();
//                if(result.getStatusCode() == protocol.StatusOuterClass.Status.Success.ordinal()) {
//                    if(result.getJobAddress() == null || result.getJobAddress().equals("")) {
//                        hintDialog("添加排班失败，排班地址为空！");
//                    } else {
//                        showToast("添加排班成功！");
//                        finish();
//                    }
//                } else {
//                    // 请求失败
//                    loadErrorDialog();
//                }
//            }
//
//            @Override
//            public void onError() {
//                // 请求失败
//                hideLoading();
//                loadErrorDialog();
//            }
//        });
//        task.execute(Variable.SERVER_IP, Variable.SERVER_PORT);
    }




}
