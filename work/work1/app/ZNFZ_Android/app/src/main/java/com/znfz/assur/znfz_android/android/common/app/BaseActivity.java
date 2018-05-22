package com.znfz.assur.znfz_android.android.common.app;


import android.app.Fragment;
import android.content.Context;
import android.content.Intent;
import android.content.pm.ActivityInfo;
import android.os.Build;
import android.os.Bundle;
import android.support.v7.app.AppCompatActivity;
import android.view.View;
import android.view.inputmethod.InputMethodManager;
import android.widget.ImageView;
import android.widget.LinearLayout;
import android.widget.RelativeLayout;
import android.widget.TextView;
import android.widget.Toast;
import com.jaeger.library.StatusBarUtil;
import com.orhanobut.logger.Logger;
import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.kyloading.KyLoadingBuilder;
import com.znfz.assur.znfz_android.android.common.utils.DialogUtils;
import com.znfz.assur.znfz_android.android.common.utils.SPUtils;
import com.znfz.assur.znfz_android.android.common.utils.StatusUiTextUtils;
import com.znfz.assur.znfz_android.android.common.utils.ToastUtils;

import javax.annotation.Nullable;

/**
 *  Activity基类
 * Created by assur on 2018/4/20 0900.
 */
public abstract class BaseActivity extends AppCompatActivity implements View.OnClickListener {

    public Context context;
    public SPUtils spUtils;
    public Toast toast;

    // toolbar
    protected TextView tv_title;            //toolbar标题
    protected LinearLayout ll_image_back;   //toolbar返回
    protected RelativeLayout toolbar;
    protected ImageView iv_right;           //toolbar右选键
    protected TextView tv_right;

    // loading
    protected KyLoadingBuilder loading;


    @Override
    protected void onCreate(@Nullable Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        Logger.i("ActivityName:"+getClass().getSimpleName());
        context = this;
        spUtils = new SPUtils(context);
        //解决后台程序 点击桌面图标重新打开应用的问题
        if ((getIntent().getFlags() & Intent.FLAG_ACTIVITY_BROUGHT_TO_FRONT) != 0) {
            finish();
            return;
        }
        // 将Activity实例添加到AppManager的堆栈中
        AppManager.getAppManager().addActivity(this);
        // 禁止横屏
        setRequestedOrientation(ActivityInfo.SCREEN_ORIENTATION_PORTRAIT);
        // 初始化视图
        initView();
        // 初始化数据
        initData();

    }

    @Override
    protected void onResume() {
        super.onResume();
    }

    @Override
    protected void onDestroy() {
        super.onDestroy();

        //将Activity实例从AppManager的堆栈中移除
        AppManager.getAppManager().finishActivity(this);

    }


    /**
     * Toast提示
     * @param text
     */
    public void showToast(String text) {
        if (toast == null) {
            toast = Toast.makeText(context, text, Toast.LENGTH_SHORT);
        } else {
            toast.setText(text);
        }
        toast.show();
    }

    @Override
    public void onClick(View v) {
        switch (v.getId()) {
            case R.id.ll_image_back:  // 返回按钮点击
                AppManager.getAppManager().finishActivity(this);
                break;
        }
    }


    /**
     * 顶部导航栏 使用Toolbar
     * @param titleName
     */
    public void initToolBar(String titleName, boolean showBackBtn) {
        this.setStatusBarColor(getResources().getColor(R.color.white), 66);
        toolbar = (RelativeLayout) findViewById(R.id.rl_toolbar);
        tv_title = (TextView) findViewById(R.id.tv_title);
        ll_image_back = (LinearLayout) findViewById(R.id.ll_image_back);
        tv_right = (TextView) findViewById(R.id.tv_right);
        iv_right = (ImageView) findViewById(R.id.iv_right);
        toolbar.setBackgroundColor(getResources().getColor(R.color.toolbar));//设置toolbar颜色
        ll_image_back.setOnClickListener(this);
        tv_right.setOnClickListener(this);
        iv_right.setOnClickListener(this);
        setToolBarTitle(titleName);

        if(!showBackBtn) {
            ll_image_back.setVisibility(View.GONE);
        }
    }

    public void setToolbarBackground(int backgroundColor) {
        toolbar.setBackgroundColor(backgroundColor);//设置toolbar颜色
    }

    public void setToolbarRightImg(int img) {
        iv_right.setImageResource(img);
        iv_right.setVisibility(View.VISIBLE);
    }

    public void setToolbarRightTv(String str) {
        tv_right.setText(str);
        tv_right.setVisibility(View.VISIBLE);
    }

    public void setToolbarRightTv(String str, int color) {
        this.setToolbarRightTv(str);
        tv_right.setTextColor(getResources().getColor(color));
    }



    /**
     * 设置非全屏状态栏，图片或文字在状态栏下
     */
    public void setStatusBarColor(int colorId, int statusbarApha) {
        boolean b = setUiTextColor(false);
        if (!b) {
            StatusBarUtil.setColor(this, colorId, statusbarApha);
        } else {
            StatusBarUtil.setColor(this, colorId, 0);
        }
        setUiTextColor(false);
    }

    /**
     * 设置状态栏
     */
    public boolean setUiTextColor(boolean isFull) {
        boolean isAdaptation = false;
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.KITKAT && Build.VERSION.SDK_INT < Build.VERSION_CODES.M) {//4.4以上系统仅对魅族，小米做过适配
            if (!StatusUiTextUtils.FlymeSetStatusBarLightMode(this.getWindow(), true)) {
                isAdaptation = StatusUiTextUtils.MIUISetStatusBarLightMode(this.getWindow(), true);
            } else {
                isAdaptation = true;
            }
        } else if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.M) {//6.0系统改变字体颜色
            if (isFull) {
                this.getWindow().getDecorView().setSystemUiVisibility(View.SYSTEM_UI_FLAG_LAYOUT_FULLSCREEN | View.SYSTEM_UI_FLAG_LIGHT_STATUS_BAR);
            } else {
                this.getWindow().getDecorView().setSystemUiVisibility(View.SYSTEM_UI_FLAG_LIGHT_STATUS_BAR);
            }
            isAdaptation = true;
        }
        return isAdaptation;
    }

    public void setStatusBarFull(Object object, int statusbarApha, View view) {
        if (object instanceof Fragment) {
            StatusBarUtil.setTranslucentForImageViewInFragment(this, statusbarApha, view);//顶部图片展示在顶部
        } else if (object instanceof BaseActivity) {
            StatusBarUtil.setTranslucentForImageView(this, 0, null);//顶部图片展示在顶部
        }
        StatusBarUtil.setTranslucentForImageView(this, 0, null);//顶部图片展示在顶部
        setUiTextColor(true);
    }

    public void setStatusBarBgFull() {
        StatusBarUtil.setColor(this, getResources().getColor(R.color.white), 0);
        setUiTextColor(true);
    }

    public void setStatusBarBgFull1() {
        StatusBarUtil.setTransparent(this);//顶部背景色展示在顶部
        setUiTextColor(true);
    }



    /**
     * 设置toolbar标题
     * @param toolBarTitle
     */
    public void setToolBarTitle(String toolBarTitle) {
        tv_title.setText(toolBarTitle);
    }

    /**
     * 设置toolbar标题颜色
     * @param toolBarTitleColor
     */
    public void setToolBarTitleColor(int toolBarTitleColor) {
        tv_title.setTextColor(toolBarTitleColor);
    }

    /**
     * 初始化控件
     */
    protected abstract void initView();


    /**
     * 初始化数据
     */
    protected abstract void initData();

    /**
     * 初始化跳转
     */
    protected void startActivity(Class<?> cls) {
        Intent intent = new Intent(this, cls);
        startActivity(intent);
    }

    /**
     * 初始化跳转
     */
    public void startActivity(Class<?> cls, Bundle bundle) {
        Intent intent = new Intent(this, cls);
        intent.putExtras(bundle);
        startActivity(intent);
    }

    protected void startActivityForResult(Class<?> cls, int requestCode) {
        Intent intent = new Intent(this, cls);
        startActivityForResult(intent, requestCode);
    }

    protected void startActivityForResult(Class<?> cls, Bundle bundle, int requestCode) {
        Intent intent = new Intent(this, cls);
        intent.putExtras(bundle);
        startActivityForResult(intent, requestCode);
    }

    protected void startActivityForResult(Class<?> cls, Bundle bundle) {
        Intent intent = new Intent(this, cls);
        intent.putExtras(bundle);
        startActivityForResult(intent, 0);
    }

    @Override
    public void onBackPressed() {
        hideLoading();
        finish();
        overridePendingTransition(android.R.anim.slide_in_left, android.R.anim.slide_out_right);
    }


    /**
     * 显示加载中...
     */
    public void showLoading() {
        if(loading != null) {
           return;
        }
        loading = new KyLoadingBuilder(this);
        loading.setIcon(R.drawable.loading);
        loading.setText(getResources().getString(R.string.loading));
        loading.setOutsideTouchable(false);
        loading.setBackTouchable(false);
        loading.show();
    }

    /**
     * 隐藏加载中...
     */
    public void hideLoading() {
        if(loading != null) {
            loading.dismiss();
            loading = null;
        }
    }

    /**
     * 网络是否连接
     * @return
     */
    public boolean isNetConnected() {
        return ZnfzApplication.isNetConnected();
    }


    /**
     * 获取字符串
     * @param stringId
     * @return
     */
    public String getStringById(int stringId) {
       return getResources().getString(stringId);
    }


    /**
     * 一般提示对话框
     * @param stringId
     */
    public void hintDialog(int stringId) {
        DialogUtils.hintDialog(BaseActivity.this, getStringById(stringId), getStringById(R.string.dialog_ok));
    }

    /**
     * 一般提示对话框
     * @param string
     */
    public void hintDialog(String string) {
        DialogUtils.hintDialog(BaseActivity.this, string, getStringById(R.string.dialog_ok));
    }

    /**
     * 加载失败对话框
     */
    public void loadErrorDialog() {
        DialogUtils.hintDialog(BaseActivity.this, getStringById(R.string.load_fail), getStringById(R.string.dialog_ok));
    }

    /**
     * toast提示
     * @param stringId
     */
    public void toast(int stringId) {
        ToastUtils.showToast(BaseActivity.this, getStringById(stringId));
    }

    /**
     * 关闭软键盘
     */
    public void hideKeyboard() {
        InputMethodManager imm = (InputMethodManager) this.getSystemService(Context.INPUT_METHOD_SERVICE);
        if (imm.isActive() && this.getCurrentFocus() != null) {
            if (this.getCurrentFocus().getWindowToken() != null) {
                imm.hideSoftInputFromWindow(this.getCurrentFocus().getWindowToken(), InputMethodManager.HIDE_NOT_ALWAYS);
            }
        }
    }
}
