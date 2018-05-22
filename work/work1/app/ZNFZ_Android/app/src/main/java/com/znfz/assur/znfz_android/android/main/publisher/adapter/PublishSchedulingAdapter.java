package com.znfz.assur.znfz_android.android.main.publisher.adapter;

import android.content.Context;
import android.graphics.Rect;
import android.support.v7.widget.LinearLayoutManager;
import android.support.v7.widget.RecyclerView;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ImageView;
import android.widget.RelativeLayout;
import android.widget.TextView;
import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.utils.TimeUtils;
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherRoleBean;
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherSchedulingBean;
import java.util.List;
import butterknife.BindView;
import butterknife.ButterKnife;

/**
 * 排班列表配器
 * Created by assur on 2018/4/25.
 */
public class PublishSchedulingAdapter extends RecyclerView.Adapter<PublishSchedulingAdapter.ViewHolder> {

    private final Context mContext;
    private List<PublisherSchedulingBean> listItems;


    public PublishSchedulingAdapter(Context context) {
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
        View view = LayoutInflater.from(parent.getContext()).inflate(R.layout.item_publisher_scheduling_list, parent, false);
        return new ViewHolder(view, parent.getContext());
    }

    @Override
    public void onBindViewHolder(ViewHolder holder, final int position) {
        PublisherSchedulingBean bean = listItems.get(position);
        // 设置视图数据
        holder.mTVSchedulingCompanyName.setText(bean.getScheduleCompany());
        holder.mTVSchedulingTime.setText(TimeUtils.getTime(bean.getScheduleTimeSp()) + "");
        holder.mTVSchedulingCompanyPer.setText("公司(" + bean.getScheduleCompanyPer() + "%)");
        holder.mTVSchedulingStorePer.setText("门店(" + bean.getScheduleStorePer() + "%)");
        holder.mRlSchedulingDetail.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                if(mOnItemClickListener != null) {
                    mOnItemClickListener.onClick(position, R.id.publisher_scheduling_list_item_rl_btn_detail);
                }
            }
        });
        holder.roleList.setVisibility(bean.isShowRoleList() ? View.VISIBLE : View.GONE);
        holder.mIvSchedulingDetailArrow.setImageDrawable(bean.isShowRoleList() ? mContext.getResources().getDrawable(R.drawable.ic_arrow_down) : mContext.getResources().getDrawable(R.drawable.ic_arrow_right));
        holder.setRoleData(bean.getScheduleRoleList());
    }


    /**
     * 给适配器填充数据
     *
     * @param listItems
     */
    public void setData(List<PublisherSchedulingBean> listItems) {
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
        private PublishSchedulingRoleAdapter publishSchedulingRoleAdapter;

        @BindView(R.id.publisher_scheduling_list_item_iv)
        ImageView  mIVScheduling;                    // 公司图标
        @BindView(R.id.publisher_scheduling_tv_item_companyname)
        TextView  mTVSchedulingCompanyName;          // 公司名称
        @BindView(R.id.publisher_scheduling_tv_item_time_value)
        TextView  mTVSchedulingTime;                 // 排班时间
        @BindView(R.id.publisher_scheduling_tv_item_per_value_company)
        TextView  mTVSchedulingCompanyPer;           // 公司比例
        @BindView(R.id.publisher_scheduling_tv_item_per_value_store)
        TextView  mTVSchedulingStorePer;             // 门店比例
        @BindView(R.id.publisher_scheduling_list_item_rl_btn_detail)
        RelativeLayout mRlSchedulingDetail;          // 查看详情
        @BindView(R.id.publisher_scheduling_iv_item_detail_arrow)
        ImageView mIvSchedulingDetailArrow;     // 查看详情箭头
        @BindView(R.id.publisher_scheduling_role_list)
        RecyclerView roleList;                       // 角色列表

        public ViewHolder(View itemView, final Context context) {
            super(itemView);
            this.context = context;
            ButterKnife.bind(this, itemView);

            LinearLayoutManager layoutmanager = new LinearLayoutManager(context);
            roleList.setLayoutManager(layoutmanager);
            roleList.addItemDecoration(new RecyclerView.ItemDecoration() {
                @Override
                public void getItemOffsets(Rect outRect, View view, RecyclerView parent, RecyclerView.State state) {
                    // 设置分割线
                    outRect.set(10, 10, 10, 10);
                }
            });
            publishSchedulingRoleAdapter = new PublishSchedulingRoleAdapter(context);
            roleList.setAdapter(publishSchedulingRoleAdapter);
        }

        /**
         * 设置角色列表数据
         *
         * @param listItems
         */
        public void setRoleData(List<PublisherRoleBean> listItems) {
            publishSchedulingRoleAdapter.setData(listItems);
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
