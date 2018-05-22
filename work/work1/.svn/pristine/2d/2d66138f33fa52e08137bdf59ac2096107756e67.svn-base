package com.znfz.assur.znfz_android.android.main.publisher.adapter;

import android.content.Context;
import android.support.v7.widget.RecyclerView;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ImageView;
import android.widget.RelativeLayout;
import android.widget.TextView;
import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.utils.TimeUtils;
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherSchedulingBean;
import java.util.List;
import butterknife.BindView;
import butterknife.ButterKnife;

/**
 * 历史带订单列表一级界面的排班列表配器
 * Created by assur on 2018/4/27.
 */
public class PublishOrderSchedulingAdapter extends RecyclerView.Adapter<PublishOrderSchedulingAdapter.ViewHolder> {

    private final Context mContext;
    private List<PublisherSchedulingBean> listItems;


    public PublishOrderSchedulingAdapter(Context context) {
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
        View view = LayoutInflater.from(parent.getContext()).inflate(R.layout.item_publisher_order_scheduling_list, parent, false);
        return new ViewHolder(view, parent.getContext());
    }

    @Override
    public void onBindViewHolder(ViewHolder holder, final int position) {
        PublisherSchedulingBean bean = listItems.get(position);
        // 设置视图数据
        holder.mTVSchedulingCompany.setText(bean.getScheduleCompany());
        if(bean.isCurSchedule()) {
            holder.mIVScheduling.setBackground(mContext.getResources().getDrawable(R.drawable.ic_red_round));
        } else {
            holder.mIVScheduling.setBackground(mContext.getResources().getDrawable(R.drawable.ic_gray_round));
        }

        holder.mRLContent.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                if(mOnItemClickListener != null) {
                    mOnItemClickListener.onClick(position, R.id.publisher_order_scheduling_list_item_rl_content);
                }
            }
        });
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

        @BindView(R.id.publisher_order_scheduling_list_item_rl_content)
        RelativeLayout  mRLContent;                  // 内容布局
        @BindView(R.id.publisher_order_scheduling_list_item_iv)
        ImageView  mIVScheduling;                    // 红色小圆点
        @BindView(R.id.publisher_order_scheduling_tv_item_company)
        TextView  mTVSchedulingCompany;              // 排班公司名称

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
