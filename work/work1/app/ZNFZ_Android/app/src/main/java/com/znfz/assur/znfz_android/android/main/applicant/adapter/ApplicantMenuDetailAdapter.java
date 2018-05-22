package com.znfz.assur.znfz_android.android.main.applicant.adapter;

import android.content.Context;
import android.support.v7.widget.RecyclerView;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;

import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.utils.RoleUtils;
import com.znfz.assur.znfz_android.android.main.applicant.bean.ApplicantMenuBean;
import com.znfz.assur.znfz_android.android.main.applicant.bean.ApplicantherRoleBean;

import java.util.List;

import butterknife.BindView;
import butterknife.ButterKnife;

/**
 * Created by dansihan on 2018/5/18.
 */

public class ApplicantMenuDetailAdapter  extends RecyclerView.Adapter<ApplicantMenuDetailAdapter.ViewHolder> {
    private final Context mContext;
    private List<ApplicantMenuBean> listItems;

    public ApplicantMenuDetailAdapter(Context context ) {
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
    public ApplicantMenuDetailAdapter.ViewHolder onCreateViewHolder(ViewGroup parent, int viewType) {
        View view = LayoutInflater.from(parent.getContext()).inflate(R.layout.item_applicant_menu_order_detail_list, parent, false);
        return new ApplicantMenuDetailAdapter.ViewHolder(view, parent.getContext());
    }

    @Override
    public void onBindViewHolder(ApplicantMenuDetailAdapter.ViewHolder holder, final int position) {
        ApplicantMenuBean bean = listItems.get(position);
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
    public void setData(List<ApplicantMenuBean> listItems) {
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
        @BindView(R.id.order_role_list_item_dish)
        TextView mTVRole;              // 菜名称
        @BindView(R.id.order_role_list_item_num)
        TextView mTVRoleNum;           // 菜的份数
        @BindView(R.id.order_role_list_item_per)
        TextView mTVRolePercentage;    //
        @BindView(R.id.order_role_list_item_money)
        TextView mTVRoleMoney;         // 菜的单价

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

    private ApplicantMenuDetailAdapter.onItemClickListener mOnItemClickListener;

    public void setOnItemClickListener(ApplicantMenuDetailAdapter.onItemClickListener mOnItemClickListener) {
        this.mOnItemClickListener = mOnItemClickListener;
    }


}
