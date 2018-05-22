package com.znfz.assur.znfz_android.android.main.applicant.adapter;

import android.content.Context;
import android.support.v7.widget.RecyclerView;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;

import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.utils.RoleUtils;
import com.znfz.assur.znfz_android.android.main.applicant.bean.ApplicantherRoleBean;
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherRoleBean;

import java.util.List;

import butterknife.BindView;
import butterknife.ButterKnife;

/**
 * Created by dansihan on 2018/5/18.
 */

public class ApplicantOrderCommissionDetailAdapter extends RecyclerView.Adapter<ApplicantOrderCommissionDetailAdapter.ViewHolder> {

    private final Context mContext;
    private List<ApplicantherRoleBean> listItems;

    public ApplicantOrderCommissionDetailAdapter(Context context ) {
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
    public ApplicantOrderCommissionDetailAdapter.ViewHolder onCreateViewHolder(ViewGroup parent, int viewType) {
        View view = LayoutInflater.from(parent.getContext()).inflate(R.layout.item_applicant_commission_order_detail_list, parent, false);
        return new ApplicantOrderCommissionDetailAdapter.ViewHolder(view, parent.getContext());
    }

    @Override
    public void onBindViewHolder(ApplicantOrderCommissionDetailAdapter.ViewHolder holder, final int position) {
        ApplicantherRoleBean bean = listItems.get(position);
        // 设置视图数据
        holder.mTVRole.setText(RoleUtils.getUserRoleString(mContext, bean.getRole()));
        holder.mTVRoleNum.setText(bean.getRoleCurNum() + "人");
        holder.mTVRolePercentage.setText(bean.getRolePercentage() + "%");
        holder.mTVRoleMoney.setText(bean.getRoleMoney());

    }


    /**
     * 给适配器填充数据
     *
     * @param listItems
     */
    public void setData(List<ApplicantherRoleBean> listItems) {
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
        @BindView(R.id.order_role_list_item_role)
        TextView mTVRole;              // 佣金分配的角色名称
        @BindView(R.id.order_role_list_item_num)
        TextView mTVRoleNum;           // 佣金分配角色数量
        @BindView(R.id.order_role_list_item_per)
        TextView mTVRolePercentage;    // 佣金分配的角色百分比
        @BindView(R.id.order_role_list_item_money)
        TextView mTVRoleMoney;         // 佣金分配的角色金钱

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

    private ApplicantOrderCommissionDetailAdapter.onItemClickListener mOnItemClickListener;

    public void setOnItemClickListener(ApplicantOrderCommissionDetailAdapter.onItemClickListener mOnItemClickListener) {
        this.mOnItemClickListener = mOnItemClickListener;
    }


}
