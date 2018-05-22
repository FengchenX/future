package com.znfz.assur.znfz_android.android.main.publisher.adapter;

import android.content.Context;
import android.support.v7.widget.RecyclerView;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;
import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.utils.TimeUtils;
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherIncomeBean;
import java.util.List;
import butterknife.BindView;
import butterknife.ButterKnife;

/**
 * 收入详情列表配器
 * Created by assur on 2018/4/24.
 */
public class PublishIncomeAdapter extends RecyclerView.Adapter<PublishIncomeAdapter.ViewHolder> {

    private final Context mContext;
    private List<PublisherIncomeBean> listItems;


    public PublishIncomeAdapter(Context context) {
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
        View view = LayoutInflater.from(parent.getContext()).inflate(R.layout.item_publisher_income_detail_list, parent, false);
        return new ViewHolder(view, parent.getContext());
    }

    @Override
    public void onBindViewHolder(ViewHolder holder, final int position) {
        PublisherIncomeBean bean = listItems.get(position);
        // 设置视图数据
        holder.mTVTitle.setText(bean.getIncomeTitle());
        holder.mTVTime.setText(TimeUtils.getTime(bean.getIncomeTimeSp()));
        holder.mTVCount.setText(bean.getIncomeType() == 1 ? "+ " + bean.getIncomeCount() : "- " + bean.getIncomeCount() );
    }


    /**
     * 给适配器填充数据
     *
     * @param listItems
     */
    public void setData(List<PublisherIncomeBean> listItems) {
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

        @BindView(R.id.publisher_income_detail_list_item_title)
        TextView  mTVTitle;             // 积分标题
        @BindView(R.id.publisher_income_detail_list_item_time)
        TextView  mTVTime;              // 积分时间
        @BindView(R.id.publisher_income_detail_list_item_count)
        TextView  mTVCount;             // 积分数量

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
