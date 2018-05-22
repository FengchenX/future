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
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherOrderBean;
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherRoleBean;
import java.util.List;
import butterknife.BindView;
import butterknife.ButterKnife;

/**
 * 订单列表配器
 * Created by assur on 2018/4/24.
 */
public class PublishOrderAdapter extends RecyclerView.Adapter<PublishOrderAdapter.ViewHolder> {

    private final Context mContext;
    private List<PublisherOrderBean> listItems;


    public PublishOrderAdapter(Context context) {
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
        View view = LayoutInflater.from(parent.getContext()).inflate(R.layout.item_publisher_order_list, parent, false);
        return new ViewHolder(view, parent.getContext());
    }

    @Override
    public void onBindViewHolder(ViewHolder holder, final int position) {
        PublisherOrderBean bean = listItems.get(position);
        // 设置视图数据
        holder.mTVOrderTime.setText(TimeUtils.getTime(bean.getOrderTimeSp()));
        holder.mTVOrderDeskNum.setText("号桌:" + bean.getOrderDeskNum());
        holder.mTVOrderPrice.setText("¥" + bean.getOrderPrice());
        if(bean.isOrderShowAllocation()) {
            holder.rlDetail.setVisibility(View.VISIBLE);
            holder.mTVMore.setText(mContext.getResources().getString(R.string.publisher_home_order_detail_pickup));
            holder.mIVMore.setBackground(mContext.getResources().getDrawable(R.drawable.ic_arrow_up));
        } else {
            holder.rlDetail.setVisibility(View.GONE);
            holder.mTVMore.setText(mContext.getResources().getString(R.string.publisher_home_order_detail_more));
            holder.mIVMore.setBackground(mContext.getResources().getDrawable(R.drawable.ic_arrow_down));
        }

        holder.rlMore.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                if(mOnItemClickListener != null) {
                    mOnItemClickListener.onClick(position, R.id.publisher_order_list_item_rl_more);
                }
            }
        });

        // 每个角色获得的金钱
        for (int j = 0; j < bean.getOrderRoleList().size(); j++) {
            PublisherRoleBean role = bean.getOrderRoleList().get(j);
            role.setRoleMoney((Double.parseDouble(bean.getOrderPrice())*0.01 * Long.parseLong(role.getRolePercentage())) + "");
        }
        holder.setRoleData(bean.getOrderRoleList());
    }


    /**
     * 给适配器填充数据
     *
     * @param listItems
     */
    public void setData(List<PublisherOrderBean> listItems) {
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
        private PublishOrderRoleAdapter publishOrderRoleAdapter;

        @BindView(R.id.publisher_order_list_item_timevalue)
        TextView  mTVOrderTime;         // 订单时间
        @BindView(R.id.publisher_order_list_item_desknum)
        TextView  mTVOrderDeskNum;      // 订单桌号
        @BindView(R.id.publisher_order_list_item_pricevalue)
        TextView  mTVOrderPrice;        // 订单价格
        @BindView(R.id.publisher_order_list_item_rl_detail)
        RelativeLayout rlDetail;        // 订单详情
        @BindView(R.id.publisher_order_role_list)
        RecyclerView roleList;          // 角色分配列表
        @BindView(R.id.publisher_order_list_item_rl_more)
        RelativeLayout rlMore;            // 更多详情/收起布局
        @BindView(R.id.publisher_order_list_item_more_text)
        TextView  mTVMore;              // 更多详情/收起文字
        @BindView(R.id.publisher_order_list_item_more_image)
        ImageView mIVMore;              // 更多详情/收起图标


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
            publishOrderRoleAdapter = new PublishOrderRoleAdapter(context);
            roleList.setAdapter(publishOrderRoleAdapter);
        }

        /**
         * 设置角色列表数据
         *
         * @param listItems
         */
        public void setRoleData(List<PublisherRoleBean> listItems) {
            publishOrderRoleAdapter.setData(listItems);
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
