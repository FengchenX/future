package com.znfz.assur.znfz_android.android.main.applicant.adapter;

import android.content.Context;
import android.graphics.Rect;
import android.support.v7.widget.LinearLayoutManager;
import android.support.v7.widget.RecyclerView;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Button;
import android.widget.ImageView;
import android.widget.RelativeLayout;
import android.widget.TextView;
import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.utils.RoleUtils;
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherOrderBean;
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherRoleBean;
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherSchedulingBean;

import java.util.List;

import butterknife.BindView;
import butterknife.ButterKnife;

/**
 * 排班列表配器
 * Created by assur on 2018/4/25.
 */
public class ApplicantJobAdapter extends RecyclerView.Adapter<ApplicantJobAdapter.ViewHolder> {

    private final Context mContext;
    private List<PublisherRoleBean> listItems;


    public ApplicantJobAdapter(Context context) {
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
        View view = LayoutInflater.from(parent.getContext()).inflate(R.layout.item_applicant_job_list, parent, false);
        return new ViewHolder(view, parent.getContext());
    }

    @Override
    public void onBindViewHolder(ViewHolder holder, final int position) {
        PublisherRoleBean bean = listItems.get(position);
        // 设置视图数据
        holder.mTvRole.setText(RoleUtils.getUserRoleString(mContext,bean.getRole()));
        holder.mTvCompanyName.setText(bean.getCompanyName());
        holder.mTvNum.setText(bean.getRoleCurNum() + "人");
        holder.mTvPer.setText(bean.getRolePercentage() + "%");
        holder.mTvDate.setText(bean.getRoleWorkTimeSp() + "");
        if(bean.isBeenAppliciant()) {
            holder.mBtnApplicant.setEnabled(false);
            holder.mBtnApplicant.setTextColor(mContext.getResources().getColor(R.color.gray));
            holder.mBtnApplicant.setBackground(mContext.getResources().getDrawable(R.drawable.bg_btn_grap));
        } else {
            holder.mBtnApplicant.setEnabled(true);
            holder.mBtnApplicant.setTextColor(mContext.getResources().getColor(R.color.status_blue));
            holder.mBtnApplicant.setBackground(mContext.getResources().getDrawable(R.drawable.bg_btn_blue_line_two));
            holder.mBtnApplicant.setOnClickListener(new View.OnClickListener() {
                @Override
                public void onClick(View v) {
                    if(mOnItemClickListener != null) {
                        mOnItemClickListener.onClick(position, R.id.item_job_list_btn_applicant);
                    }
                }
            });
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

        @BindView(R.id.item_job_list_tv_role)
        TextView  mTvRole;              // 职位名称
        @BindView(R.id.item_job_list_tv_companyname)
        TextView  mTvCompanyName;       // 公司名称
        @BindView(R.id.item_job_list_tv_num_value)
        TextView  mTvNum;               // 剩余招聘人数
        @BindView(R.id.item_job_list_tv_per_value)
        TextView  mTvPer;               // 佣金比例
        @BindView(R.id.item_job_list_tv_date_value)
        TextView  mTvDate;               // 上班时间
        @BindView(R.id.item_job_list_btn_applicant)
        Button mBtnApplicant;            // 申请按钮


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
