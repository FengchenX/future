package com.znfz.assur.znfz_android.android.main.publisher.frament;

import android.content.Intent;
import android.os.Bundle;
import android.util.Log;
import android.view.View;
import android.view.ViewGroup;
import android.widget.AdapterView;
import android.widget.EditText;
import android.widget.ImageView;
import android.widget.LinearLayout;
import android.widget.ListAdapter;
import android.widget.ListView;
import android.widget.RelativeLayout;
import android.widget.TextView;

import com.bigkoo.pickerview.TimePickerView;
import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.app.BaseFragment;
import com.znfz.assur.znfz_android.android.common.utils.TimeUtils;
import com.znfz.assur.znfz_android.android.main.publisher.activity.PublisherAddEmployeeActivity;
import com.znfz.assur.znfz_android.android.main.publisher.activity.publisherJobInformationActivity;
import com.znfz.assur.znfz_android.android.main.publisher.adapter.AddEmployeeListAdapter;
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherAddEmployeeBean;

import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Date;
import java.util.List;

import butterknife.BindView;
import butterknife.ButterKnife;

/**
 * Created by dansihan on 2018/5/21.
 */

public class PublisherFragmentAddEmployee extends BaseFragment {
    @BindView(R.id.ll_image_back)
    LinearLayout llImageBack;
    @BindView(R.id.tv_title)
    TextView tvTitle;
    @BindView(R.id.iv_right)
    ImageView ivRight;
    @BindView(R.id.tv_right)
    TextView tvRight;
    @BindView(R.id.rl_toolbar)
    RelativeLayout rlToolbar;
    @BindView(R.id.ed_publisher_company_name) //输入公司名称
            EditText edPublisherCompanyName;
    @BindView(R.id.ed_publisher_company_account) //输入公司账户
            EditText edPublisherCompanyAccount;
    @BindView(R.id.ed_publisher_store_name)  //输入门店名称
            EditText edPublisherStoreName;
    @BindView(R.id.ed_publisher_store_number) //输入门店编号
            EditText edPublisherStoreNumber;
    @BindView(R.id.tv_publisher_company_proportion) //选择公司比例
            TextView tvPublisherCompanyProportion;
    @BindView(R.id.tv_publisher_store_proportion)  //选择门店比例
            TextView tvPublisherStoreProportion;
    @BindView(R.id.tv_publisher_time_start)  //开始工作时间
            TextView tvPublisherTimeStart;
    @BindView(R.id.tv_publisher_time_end)  //结束工作时间
            TextView tvPublisherTimeEnd;
    @BindView(R.id.ed_publisher_address)  //工作大概区域
            EditText edPublisherAddress;
    @BindView(R.id.ed_publisher_address_detail)  //工作详细地址
            EditText edPublisherAddressDetail;
    @BindView(R.id.lv_add_employee_list)  //招聘工作人员列表
            ListView lvAddEmployeeList;
    @BindView(R.id.tv_add_employee)  //添加招聘工作人员
            TextView tvAddEmployee;

    private AddEmployeeListAdapter adapter;
    private List<PublisherAddEmployeeBean> listJobName = new ArrayList<>();


    @Override
    protected int setView() {
        return R.layout.activity_add_empioyee;
    }

    @Override
    protected void init(View view) {
        ButterKnife.bind(this, view);
        llImageBack.setVisibility(View.GONE);
        tvTitle.setText(getResources().getString(R.string.publisher_Release_schedule));
        tvAddEmployee.setOnClickListener(this);
        adapter = new AddEmployeeListAdapter(getActivity(), listJobName);
        lvAddEmployeeList.setAdapter(adapter);
        tvPublisherTimeStart.setOnClickListener(this);
        tvPublisherTimeEnd.setOnClickListener(this);
        lvAddEmployeeList.setOnItemClickListener(new AdapterView.OnItemClickListener() {
            @Override
            public void onItemClick(AdapterView<?> parent, View view, int position, long id) {
                Intent intent = new Intent();
                intent.putExtra("item", true);
                intent.putExtra("jobname", listJobName.get(position).getName());
                intent.putExtra("JobPersonNumber", listJobName.get(position).getJobPersonNumber());
                intent.putExtra("CommissionRate", listJobName.get(position).getCommissionRate());
                intent.putExtra("FriendName", listJobName.get(position).getFriendname());
                intent.setClass(getActivity(), publisherJobInformationActivity.class);
                Log.d("dandan", "CommissionRate" + listJobName.get(position).getCommissionRate());
                listJobName.remove(position);
                startActivityForResult(intent, 10);
            }
        });
    }

    @Override
    protected void initData(Bundle savedInstanceState) {

    }

    @Override
    public void onClick(View v) {
        super.onClick(v);
        switch (v.getId()) {

            case R.id.tv_add_employee:
                Intent intent = new Intent();
                intent.putExtra("item", false);
                intent.setClass(getActivity(), publisherJobInformationActivity.class);
                startActivityForResult(intent, 10);
                break;

            case R.id.tv_publisher_time_start:
                hideKeyboard();
                TimePickerView pvTime = new TimePickerView.Builder(getActivity(), new TimePickerView.OnTimeSelectListener() {
                    @Override
                    public void onTimeSelect(Date date, View v) {
                        tvPublisherTimeStart.setText(TimeUtils.getTime(date.getTime(), new SimpleDateFormat("yyyy-MM-dd HH:mm")));
                    }
                }).setType(new boolean[]{true, true, true, true, true, false})// 默认全部显示
                        .setCancelText(getStringById(R.string.cancel))//取消按钮文字
                        .setSubmitText(getStringById(R.string.ok))    //确认按钮文字
                        .setOutSideCancelable(false)//点击屏幕，点在控件外部范围时，是否取消显示
                        .setLabel("年", "月", "日", "时", "分", "秒")//默认设置为年月日时分秒
                        .isCenterLabel(false) //是否只显示中间选中项的label文字，false则每项item全部都带有label。
                        .build();
                pvTime.show();
                break;
            case R.id.tv_publisher_time_end:
                hideKeyboard();
                TimePickerView pvTime2 = new TimePickerView.Builder(getActivity(), new TimePickerView.OnTimeSelectListener() {
                    @Override
                    public void onTimeSelect(Date date, View v) {
                        tvPublisherTimeEnd.setText(TimeUtils.getTime(date.getTime(), new SimpleDateFormat("yyyy-MM-dd HH:mm")));
                    }
                }).setType(new boolean[]{true, true, true, true, true, false})// 默认全部显示
                        .setCancelText(getStringById(R.string.cancel))//取消按钮文字
                        .setSubmitText(getStringById(R.string.ok))    //确认按钮文字
                        .setOutSideCancelable(false)//点击屏幕，点在控件外部范围时，是否取消显示
                        .setLabel("年", "月", "日", "时", "分", "秒")//默认设置为年月日时分秒
                        .isCenterLabel(false) //是否只显示中间选中项的label文字，false则每项item全部都带有label。
                        .build();
                pvTime2.show();
                break;
        }
    }

    @Override
    public void onActivityResult(int requestCode, int resultCode, Intent data) {
        super.onActivityResult(requestCode, resultCode, data);

        if (requestCode == 10 && resultCode == 2) {
            PublisherAddEmployeeBean bean = new PublisherAddEmployeeBean();
            bean.setName(data.getStringExtra("jobname"));
            bean.setCommissionRate(data.getStringExtra("CommissionRate"));
            bean.setJobPersonNumber(data.getStringExtra("JobPersonNumber"));
            bean.setFriendname(data.getStringExtra("FriendName"));
            listJobName.add(bean);
            adapter.notifyDataSetChanged();
            setListViewHeightBasedOnChildren(lvAddEmployeeList);
        }
    }

    @Override
    public void onResume() {
        super.onResume();
        setListViewHeightBasedOnChildren(lvAddEmployeeList);
    }

    /**
     * 计算listview的高度
     *
     * @param listView
     */
    public void setListViewHeightBasedOnChildren(ListView listView) {
        ListAdapter listAdapter = listView.getAdapter();
        if (listAdapter == null) {
            return;
        }

        int totalHeight = 0;
        for (int i = 0; i < listAdapter.getCount(); i++) {
            View listItem = listAdapter.getView(i, null, listView);
            listItem.measure(0, 0);  // 获取item高度
            totalHeight += listItem.getMeasuredHeight();
        }

        ViewGroup.LayoutParams params = listView.getLayoutParams();
        // 最后再加上分割线的高度和padding高度，否则显示不完整。
        params.height = totalHeight + (listView.getDividerHeight() * (listAdapter.getCount() - 1)) + listView.getPaddingTop() + listView.getPaddingBottom();
        listView.setLayoutParams(params);
    }

}
