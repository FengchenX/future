package com.znfz.assur.znfz_android.android.main.publisher.activity;

import android.graphics.Rect;
import android.support.v7.widget.LinearLayoutManager;
import android.support.v7.widget.RecyclerView;
import android.view.View;
import android.widget.AdapterView;
import android.widget.Button;

import com.bigkoo.pickerview.TimePickerView;
import com.flyco.dialog.entity.DialogMenuItem;
import com.flyco.dialog.listener.OnOperItemClickL;
import com.flyco.dialog.widget.ActionSheetDialog;
import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.app.BaseActivity;
import com.znfz.assur.znfz_android.android.common.utils.TimeUtils;
import com.znfz.assur.znfz_android.android.main.publisher.adapter.PublishAddSchedulingAdapter;
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherRoleBean;
import java.util.ArrayList;
import java.util.Date;
import java.util.List;
import butterknife.BindView;
import butterknife.ButterKnife;

/**
 * 添加员工界面
 * Created by assur on 2018/4/27
 */
public class AddEmployeeActivity extends BaseActivity {

    @BindView(R.id.publish_add_employee_btn_ok)    // 添加员工按钮
    Button addEmployeeBtn;


    @Override
    protected void initView() {
        setContentView(R.layout.activity_publisher_add_employee);
        ButterKnife.bind(this);
        initToolBar(getStringById(R.string.publisher_employees_add_title), true);

        addEmployeeBtn.setOnClickListener(this);
    }


    @Override
    protected void initData() {

    }

    @Override
    public void onClick(View v) {
        super.onClick(v);

        switch (v.getId()) {

            case R.id.publish_add_employee_btn_ok:      // 完成按钮
                finish();
                break;

            default:

                    break;
        }
    }



}
