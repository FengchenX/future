package com.znfz.assur.znfz_android.android.common.app;

import android.app.Activity;
import android.content.Context;
import android.content.Intent;
import android.os.Bundle;
import android.support.v4.app.Fragment;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.view.inputmethod.InputMethodManager;
import android.widget.Toast;

import com.scwang.smartrefresh.layout.api.RefreshLayout;
import com.scwang.smartrefresh.layout.footer.ClassicsFooter;
import com.scwang.smartrefresh.layout.header.ClassicsHeader;
import com.scwang.smartrefresh.layout.listener.OnLoadMoreListener;
import com.scwang.smartrefresh.layout.listener.OnRefreshListener;
import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.app.BaseActivity;
import com.znfz.assur.znfz_android.android.common.kyloading.KyLoadingBuilder;
import com.znfz.assur.znfz_android.android.common.task.AddOrderAndPayTask;
import com.znfz.assur.znfz_android.android.common.utils.DialogUtils;
import com.znfz.assur.znfz_android.android.common.utils.SPUtils;
import com.znfz.assur.znfz_android.android.common.utils.ToastUtils;

import java.text.SimpleDateFormat;
import java.util.Date;


/**
 * Frament基类
 * Created by assur on 2018/4/20
 */

public abstract class BaseFragment extends Fragment implements View.OnClickListener{

    protected BaseActivity mActivity;
    public SPUtils spUtils;
    public Toast toast;
    // loading
    protected KyLoadingBuilder loading;
    // refresh
    protected RefreshLayout refreshLayout;


    protected abstract int setView();
    protected abstract void init(View view);
    protected abstract void initData(Bundle savedInstanceState);

    @Override
    public void onAttach(Activity activity) {
        super.onAttach(activity);
        mActivity = (BaseActivity) activity;
        spUtils = new SPUtils(mActivity);
    }

    @Override
    public void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
    }

    @Override
    public View onCreateView(LayoutInflater inflater, ViewGroup container,
                             Bundle savedInstanceState) {
        return inflater.inflate(setView(), container, false);
    }

    @Override
    public void onViewCreated(View view, Bundle savedInstanceState) {
        super.onViewCreated(view, savedInstanceState);
        init(view);
    }

    @Override
    public void onActivityCreated(Bundle savedInstanceState) {
        super.onActivityCreated(savedInstanceState);
        initData(savedInstanceState);
    }

    @Override
    public void onStart() {
        super.onStart();
    }

    @Override
    public void onResume() {
        super.onResume();
    }

    @Override
    public void onPause() {
        super.onPause();
    }


    @Override
    public void onStop() {
        super.onStop();
    }


    @Override
    public void onDestroyView() {
        super.onDestroyView();
    }

    @Override
    public void onDestroy() {
        super.onDestroy();
    }

    @Override
    public void onDetach() {
        super.onDetach();
    }

    @Override
    public void onClick(View v) {

    }

    /**
     * 初始化跳转
     */
    protected void startActivity(Class<?> cls) {
        Intent intent = new Intent(getActivity(), cls);
        startActivity(intent);
    }

    /**
     * 初始化跳转
     */
    public void startActivity(Class<?> cls, Bundle bundle) {
        Intent intent = new Intent(getActivity(), cls);
        intent.putExtras(bundle);
        startActivity(intent);
    }

    /**
     * 显示加载中...
     */
    public void showLoading() {
        if(loading != null) {
            return;
        }
        loading = new KyLoadingBuilder(mActivity.context);
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
     * 一般提示对话框
     * @param string
     */
    public void hintDialog(String string) {
        DialogUtils.hintDialog(mActivity.context, string, getStringById(R.string.dialog_ok));
    }

    /**
     * 加载失败对话框
     */
    public void loadErrorDialog() {
        DialogUtils.hintDialog(mActivity.context, getStringById(R.string.load_fail), getStringById(R.string.dialog_ok));
    }

    /**
     * toast提示
     * @param stringId
     */
    public void toast(int stringId) {
        ToastUtils.showToast(mActivity.context, getStringById(stringId));
    }

    /**
     * 关闭软键盘
     */
    public void hideKeyboard() {
        InputMethodManager imm = (InputMethodManager) mActivity.getSystemService(Context.INPUT_METHOD_SERVICE);
        if (imm.isActive() && mActivity.getCurrentFocus() != null) {
            if (mActivity.getCurrentFocus().getWindowToken() != null) {
                imm.hideSoftInputFromWindow(mActivity.getCurrentFocus().getWindowToken(), InputMethodManager.HIDE_NOT_ALWAYS);
            }
        }
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
        DialogUtils.hintDialog(mActivity.context, getStringById(stringId), getStringById(R.string.dialog_ok));
    }

    /**
     * Toast提示
     * @param text
     */
    public void showToast(String text) {
        if (toast == null) {
            toast = Toast.makeText(mActivity, text, Toast.LENGTH_SHORT);
        } else {
            toast.setText(text);
        }
        toast.show();
    }

    /**
     * 结束下啦刷新
     */
    public void finishRefresh() {
        if(refreshLayout != null) {
            refreshLayout.finishRefresh();
        }

    }

    /**
     * 下啦刷新回调接口
     */
    public interface RefreshLayoutCallBack {
        void onRefreshListenerBegin();
        void onLoadMoreBegin();
    }

    /**
     * 初始化下啦刷新
     */
    public void initRefreshLayout(final RefreshLayoutCallBack callBack) {

        // 上啦刷新/下啦刷新
        if(refreshLayout != null) {
            return;
        }
        refreshLayout = (RefreshLayout)mActivity.findViewById(R.id.refreshlayout);
        final ClassicsHeader header = (ClassicsHeader)mActivity.findViewById(R.id.classicsheader);
        header.REFRESH_HEADER_REFRESHING = "刷新中...";
        header.REFRESH_HEADER_RELEASE = "松手开始刷新";
        header.REFRESH_HEADER_PULLING = "请向下拉";
        header.REFRESH_HEADER_FAILED = "刷新失败！";
        header.REFRESH_HEADER_FINISH = "刷新成功";
        header.REFRESH_HEADER_LOADING = "加载中...";
        final ClassicsFooter footer = (ClassicsFooter)mActivity.findViewById(R.id.classicsFooter);
        footer.REFRESH_FOOTER_LOADING = "加载中...";
        footer.REFRESH_FOOTER_FAILED = "加载失败！";
        footer.REFRESH_FOOTER_FINISH = "加载成功";
        refreshLayout.setOnRefreshListener(new OnRefreshListener() {
            @Override
            public void onRefresh(RefreshLayout refreshlayout) {
                Date day=new Date();
                SimpleDateFormat df = new SimpleDateFormat("MM-dd HH:mm");
                header.setLastUpdateText("最后刷新时间: " + df.format(day));
                refreshlayout.finishRefresh(5000/*,false*/);//传入false表示刷新失败
                callBack.onRefreshListenerBegin();
            }
        });
        refreshLayout.setOnLoadMoreListener(new OnLoadMoreListener() {
            @Override
            public void onLoadMore(RefreshLayout refreshlayout) {
                refreshlayout.finishLoadMore(2000/*,false*/);//传入false表示加载失败
                callBack.onLoadMoreBegin();
            }
        });
    }




}