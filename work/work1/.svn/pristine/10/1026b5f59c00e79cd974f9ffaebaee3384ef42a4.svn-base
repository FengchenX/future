package com.znfz.assur.znfz_android.android.main.publisher.activity;

import android.graphics.Rect;
import android.support.v7.widget.LinearLayoutManager;
import android.support.v7.widget.RecyclerView;
import android.view.View;
import android.widget.Button;
import android.widget.EditText;
import android.widget.ImageView;
import android.widget.TextView;
import com.bigkoo.pickerview.TimePickerView;
import com.orhanobut.logger.Logger;
import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.app.BaseActivity;
import com.znfz.assur.znfz_android.android.common.global.UserInfo;
import com.znfz.assur.znfz_android.android.common.global.Variable;
import com.znfz.assur.znfz_android.android.common.task.AddOrderAndPayTask;
import com.znfz.assur.znfz_android.android.common.utils.TimeUtils;
import com.znfz.assur.znfz_android.android.main.publisher.adapter.PublishSchedulingRoleAdapter;

import java.util.Date;
import butterknife.BindView;
import butterknife.ButterKnife;
import protocol.Obj;

/**
 * 添加订单界面
 * Created by assur on 2018/4/25
 */
public class AddOrderActivity extends BaseActivity {

    @BindView(R.id.publish_add_order_et_desk_num)     // 消费桌号
    EditText etDeskNum;
    @BindView(R.id.publish_add_order_et_money)        // 消费金额
    EditText etMoney;
    @BindView(R.id.publish_add_order_tv_time)         // 消费时间
    TextView tvTime;
    @BindView(R.id.publish_add_order_iv_timeselect)   // 选择时间按钮
    ImageView ivTimeselect;
    @BindView(R.id.publisher_add_order_btn_ok)        // 确定按钮
    Button btnOK;


    @BindView(R.id.publisher_add_order_role_list)
    RecyclerView roleList;                           // 角色列表
    private PublishSchedulingRoleAdapter roleAdapter;

    @Override
    protected void initView() {
        setContentView(R.layout.activity_publish_add_order);
        ButterKnife.bind(this);
        initToolBar(getStringById(R.string.publisher_add_order_title), true);

        ivTimeselect.setOnClickListener(this);
        btnOK.setOnClickListener(this);
    }


    @Override
    protected void initData() {

        // 默认显示当前时间
        tvTime.setText(TimeUtils.getCurrentTimeInString());

        // 角色列表
        LinearLayoutManager layoutmanager = new LinearLayoutManager(context);
        roleList.setLayoutManager(layoutmanager);
        roleList.addItemDecoration(new RecyclerView.ItemDecoration() {
            @Override
            public void getItemOffsets(Rect outRect, View view, RecyclerView parent, RecyclerView.State state) {
                // 设置分割线
                outRect.set(10, 10, 10, 10);
            }
        });
        roleAdapter = new PublishSchedulingRoleAdapter(context);
        roleList.setAdapter(roleAdapter);
        roleAdapter.setData(UserInfo.currentSchedule.getScheduleRoleList());
    }

    @Override
    public void onClick(View v) {
        super.onClick(v);

        switch (v.getId()) {
            case R.id.publish_add_order_iv_timeselect:  // 选择时间
                hideKeyboard();
                TimePickerView pvTime = new TimePickerView.Builder(this, new TimePickerView.OnTimeSelectListener() {
                    @Override
                    public void onTimeSelect(Date date, View v) {
                        tvTime.setText(TimeUtils.getTime(date.getTime()));
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


            case R.id.publisher_add_order_btn_ok:      // 确定


                if(etDeskNum.getText().toString() == null || etDeskNum.getText().toString().equals("")) {
                    hintDialog("请输入消费桌号");
                    return;
                }

                if(etMoney.getText().toString() == null || etMoney.getText().toString().equals("")) {
                    hintDialog("请输入消费金额");
                    return;
                }

                if(tvTime.getText().toString() == null || tvTime.getText().toString().equals("")) {
                    hintDialog("请选择消费时间");
                    return;
                }

                AddOrderAndPayRequest();

                break;



                default:

                    break;
        }
    }




    /**
     * 添加订单并付款请求
     */
    private void AddOrderAndPayRequest() {
        showLoading();
        AddOrderAndPayTask task = new AddOrderAndPayTask();
        protocol.Api.ReqPay.Builder requestBuilder = protocol.Api.ReqPay.newBuilder();
        Obj.Order order = Obj.Order.newBuilder()
                .setTable(Integer.parseInt(etDeskNum.getText().toString()))
                .setMoney(Long.parseLong(etMoney.getText().toString()))
                .setTimeStamp(TimeUtils.getTimeSp(tvTime.getText().toString()))
                .build();
        requestBuilder.setContent(order);
        requestBuilder.setJobAddress(UserInfo.currentSchedule.getScheduleAddress());
        requestBuilder.setMoney(Long.parseLong(etMoney.getText().toString()));
        requestBuilder.setPassWord(UserInfo.USER_PASSWORD);
        requestBuilder.setAccountDescribe(UserInfo.USER_PRIVATE_KEY);
        requestBuilder.setUserAddress(UserInfo.USER_ADDRESS);

        task.setRequest(requestBuilder.build());
        task.setUpdateUIInterface(new AddOrderAndPayTask.UpdateUIInterface() {
            @Override
            public void onSucceed(protocol.Api.RespPay result) {
                Logger.d("GetScheduleTask更新UI：" + result);
                hideLoading();
                if(result.getStatusCode() == protocol.StatusOuterClass.Status.Success.ordinal()) {
                    showToast("添加订单成功");
                    finish();
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
