package com.znfz.assur.znfz_android.android.main.publisher.adapter;

import android.content.Context;
import android.graphics.Rect;
import android.support.v7.widget.LinearLayoutManager;
import android.support.v7.widget.RecyclerView;
import android.text.Editable;
import android.text.TextWatcher;
import android.view.KeyEvent;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.EditText;
import android.widget.ImageView;
import android.widget.RelativeLayout;
import android.widget.TextView;

import com.orhanobut.logger.Logger;
import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.utils.RoleUtils;
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherOrderBean;
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherRoleBean;

import java.util.List;

import butterknife.BindView;
import butterknife.ButterKnife;

/**
 * 订单列表配器
 * Created by assur on 2018/4/24.
 */
public class PublishAddSchedulingAdapter extends RecyclerView.Adapter<PublishAddSchedulingAdapter.ViewHolder> {

    private final Context mContext;
    private List<PublisherRoleBean> listItems;


    public PublishAddSchedulingAdapter(Context context) {
        this.mContext = context;
    }

    /**
     * 创建视图ViewHolder
     *
     * @param parent
     * @param viewType
     * @return
     */
    @Override
    public ViewHolder onCreateViewHolder(ViewGroup parent, int viewType) {
        View view = LayoutInflater.from(parent.getContext()).inflate(R.layout.item_publisher_add_scheduling_list, parent, false);
        return new ViewHolder(view, parent.getContext());
    }

    @Override
    public void onBindViewHolder(final ViewHolder holder, final int position) {
        PublisherRoleBean bean = listItems.get(position);
        // 设置视图数据
        String roleString = RoleUtils.getUserRoleString(mContext, bean.getRole());
        holder.mETRole.setText((roleString == null || roleString.length() == 0) ? "": roleString);
        holder.mETRoleNum.setText((bean.getRoleCurNum() == null || bean.getRoleCurNum().length() == 0) ? "" : bean.getRoleCurNum());
        holder.mETRolePer.setText((bean.getRolePercentage() == null || bean.getRolePercentage().length() == 0) ? "" : bean.getRolePercentage());
        holder.mIVRoleNumChoice.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                if(mOnItemClickListener != null) {
                    mOnItemClickListener.onClick(position, R.id.publish_add_scheduling_iv_rolenumselect);
                }
            }
        });

        holder.mIVRolePerChoice.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                if(mOnItemClickListener != null) {
                    mOnItemClickListener.onClick(position, R.id.publish_add_scheduling_iv_perselect);
                }
            }
        });
    }


    /**
     * 给适配器填充数据
     *
     * @param listItems
     */
    public void setData(List<PublisherRoleBean> listItems) {
        this.listItems = listItems;
        notifyDataSetChanged();
    }


    /**
     * 返回item个数
     */
    @Override
    public int getItemCount() {
        if (listItems != null && listItems.size() > 0) {
            return listItems.size();
        }
        return 0;
    }

    /**
     * ViewHolder实例化控件
     */
    public class ViewHolder extends RecyclerView.ViewHolder {

        private Context context;

        @BindView(R.id.publish_add_scheduling_et_role)
        TextView  mETRole;                 // 角色
        @BindView(R.id.publish_add_scheduling_et_rolenum)
        TextView  mETRoleNum;              // 角色数量
        @BindView(R.id.publish_add_scheduling_iv_rolenumselect)
        ImageView  mIVRoleNumChoice;       // 角色数量选择
        @BindView(R.id.publish_add_scheduling_et_per)
        TextView  mETRolePer;              // 角色分配的佣金比例
        @BindView(R.id.publish_add_scheduling_iv_perselect)
        ImageView  mIVRolePerChoice;  // 上角色分配的佣金比例选择


        public ViewHolder(View itemView, final Context context) {
            super(itemView);
            this.context = context;
            ButterKnife.bind(this, itemView);
        }

    }


    /**
     * 点击事件监听接口
     */
    public interface OnItemClickListener {
        void onClick(int position, int viewID);
    }
    private OnItemClickListener mOnItemClickListener;
    public void setOnItemClickListener(OnItemClickListener mOnItemClickListener) {
        this.mOnItemClickListener = mOnItemClickListener;
    }



}
