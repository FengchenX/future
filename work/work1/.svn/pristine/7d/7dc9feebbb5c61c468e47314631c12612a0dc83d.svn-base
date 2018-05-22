package com.znfz.assur.znfz_android.android.main.publisher.adapter;

import android.content.Context;
import android.support.v7.widget.RecyclerView;
import android.util.Log;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
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
 *  排班列表中的角色列表
 * Created by assur on 2018/4/24.
 */
public class PublishSchedulingRoleAdapter extends RecyclerView.Adapter<PublishSchedulingRoleAdapter.ViewHolder> {

    private final Context mContext;
    private List<PublisherRoleBean> listItems;


    public PublishSchedulingRoleAdapter(Context context) {
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
        View view = LayoutInflater.from(parent.getContext()).inflate(R.layout.item_publisher_scheduling_role_list, parent, false);
        return new ViewHolder(view, parent.getContext());
    }

    @Override
    public void onBindViewHolder(ViewHolder holder, final int position) {
        PublisherRoleBean bean = listItems.get(position);
        // 设置视图数据
        holder.mTVRoleName.setText(bean.getRole());
        holder.mTVRolePercentage.setText(bean.getRolePercentage()+"%");
        if(bean.getRoleCurNum() == null || bean.getRoleCurNum().length() == 0) {
            holder.mTVRoleNumCurNeed.setText(mContext.getResources().getString(R.string.publisher_home_scheduling_full_text));
            //holder.mTVRoleNumCurNeed.setTextColor(mContext.getResources().getColor(R.color.gray));
        } else {
            //holder.mTVRoleNumCurNeed.setTextColor(mContext.getResources().getColor(R.color.status_blue));
            holder.mTVRoleNumCurNeed.setText("" + Integer.valueOf(bean.getRoleCurNum()) + "人");
        }

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


        @BindView(R.id.publisher_scheduling_role_list_item_rolename)
        TextView  mTVRoleName;          // 角色名称
        // @BindView(R.id.publisher_scheduling_role_list_item_rolenum)
        //TextView  mTVRoleTotalNum;      // 角色总需数量
        @BindView(R.id.publisher_scheduling_role_list_item_roleper)
        TextView  mTVRolePercentage;    // 角色总百分比
        @BindView(R.id.publisher_scheduling_role_list_item_rolenum_remaining)
        TextView  mTVRoleNumCurNeed;    // 角色剩余需要人数

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
