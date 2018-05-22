package com.znfz.assur.znfz_android.android.main.publisher.adapter;

import android.content.Context;
import android.support.v7.widget.RecyclerView;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;
import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.utils.RoleUtils;
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherEmployeeBean;

import java.util.List;

import butterknife.BindView;
import butterknife.ButterKnife;

/**
 * 员工列表配器
 * Created by assur on 2018/4/24.
 */
public class PublishEmployeesAdapter extends RecyclerView.Adapter<PublishEmployeesAdapter.ViewHolder> {

    private final Context mContext;
    private List<PublisherEmployeeBean> listItems;


    public PublishEmployeesAdapter(Context context) {
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
        View view = LayoutInflater.from(parent.getContext()).inflate(R.layout.item_publisher_employees_list, parent, false);
        return new ViewHolder(view, parent.getContext());
    }

    @Override
    public void onBindViewHolder(ViewHolder holder, final int position) {
        PublisherEmployeeBean bean = listItems.get(position);
        // 设置视图数据
        holder.mTVRole.setText(RoleUtils.getUserRoleString(mContext, bean.getPublisherEmployeeRole()));
        if(bean.getPublisherEmployeeRole().equals(mContext.getResources().getString(R.string.user_role_cook))) {
            holder.mTVRole.setBackground(mContext.getResources().getDrawable(R.drawable.bg_role_cook_text));
        } else if(bean.getPublisherEmployeeRole().equals(mContext.getResources().getString(R.string.user_role_waiter))) {
            holder.mTVRole.setBackground(mContext.getResources().getDrawable(R.drawable.bg_role_waiter_text));
        }
        holder.mTVName.setText(bean.getPublisherEmployeeName());
        holder.mTVPPer.setText(bean.getPublisherEmployeePer() + "%");
    }


    /**
     * 给适配器填充数据
     *
     * @param listItems
     */
    public void setData(List<PublisherEmployeeBean> listItems) {
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

        @BindView(R.id.publisher_employees_list_item_role)
        TextView  mTVRole;              // 角色
        @BindView(R.id.publisher_employees_list_item_name)
        TextView  mTVName;              // 姓名
        @BindView(R.id.publisher_employees_list_item_per)
        TextView  mTVPPer;             // 佣金分配百分比

        public ViewHolder(View itemView, final Context context) {
            super(itemView);
            this.context = context;
            ButterKnife.bind(this, itemView);

        }

    }


    /**
     * 点击事件监听接口
     */
    public interface onItemClickListener {
        void onClick(int position, int viewID);
    }
    private onItemClickListener mOnItemClickListener;
    public void setOnItemClickListener(onItemClickListener mOnItemClickListener) {
        this.mOnItemClickListener = mOnItemClickListener;
    }



}
