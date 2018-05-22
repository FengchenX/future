package com.znfz.assur.znfz_android.android.test.listview_test;

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
import java.util.List;

public class FirstAdapter extends RecyclerView.Adapter<FirstAdapter.ViewHolder>  {

    private final Context mContext;
    private List<FirstBean> listItems;


    public FirstAdapter(Context context) {
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
    public FirstAdapter.ViewHolder onCreateViewHolder(ViewGroup parent, int viewType) {
        View view = LayoutInflater.from(parent.getContext()).inflate(R.layout.item_test_twolevelsort_first_list, parent, false);
        return new FirstAdapter.ViewHolder(view, parent.getContext());
    }

    @Override
    public void onBindViewHolder(FirstAdapter.ViewHolder holder, final int position) {

        holder.rlContent.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                if(mOnItemClickListener != null) {
                    if(listItems.get(position).getSubList() == null || listItems.get(position).getSubList().size() == 0) {
                        mOnItemClickListener.onStringSelected(listItems.get(position).getName());
                    } else {
                        mOnItemClickListener.onClick(position, R.id.rl_item_content);
                    }
                }
            }
        });
        holder.tvName.setText(listItems.get(position).getName());
        if(listItems.get(position).getSubList() == null || listItems.get(position).getSubList().size() == 0) {
            holder.ivSpreadIcon.setVisibility(View.GONE);
        } else {
            holder.ivSpreadIcon.setVisibility(View.VISIBLE);
        }

        if(listItems.get(position).isSpread()) {
            holder.ivSpreadIcon.setImageDrawable(mContext.getResources().getDrawable(R.drawable.ic_arrow_down));
            holder.recyclerView.setVisibility(View.VISIBLE);
            holder.setData(listItems.get(position).getSubList());
        } else {
            holder.ivSpreadIcon.setImageDrawable(mContext.getResources().getDrawable(R.drawable.ic_arrow_right));
            holder.recyclerView.setVisibility(View.GONE);
        }


    }




    /**
     * 给适配器填充数据
     *
     * @param listItems
     */
    public void setData(List<FirstBean> listItems) {
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
        private RelativeLayout rlContent;
        private TextView tvName;
        private ImageView ivSpreadIcon;
        private RecyclerView recyclerView;

        private List<SecondBean> listItems;
        private SecondAdapter adapter;




        public ViewHolder(View itemView, final Context context) {
            super(itemView);
            this.context = context;

            rlContent = (RelativeLayout) itemView.findViewById(R.id.rl_item_content);
            tvName = (TextView) itemView.findViewById(R.id.item_name);
            ivSpreadIcon = (ImageView) itemView.findViewById(R.id.item_icon);
            // 列表视图
            recyclerView = (RecyclerView)itemView.findViewById(R.id.lv_second);
            LinearLayoutManager layoutmanager = new LinearLayoutManager(context);
            layoutmanager.setOrientation(1);
            recyclerView.setLayoutManager(layoutmanager);
            recyclerView.addItemDecoration(new RecyclerView.ItemDecoration() {
                @Override
                public void getItemOffsets(Rect outRect, View view, RecyclerView parent, RecyclerView.State state) {
                    // 设置分割线
                    outRect.set(0, 1, 0, 0);
                }
            });

            adapter = new SecondAdapter(context);
            adapter.setOnItemClickListener(new SecondAdapter.OnItemClickListener() {
                @Override
                public void onClick(int position, int viewID) {

                }

                @Override
                public void onStringSelected(String string) {
                    if(mOnItemClickListener != null) {
                        mOnItemClickListener.onStringSelected(string);
                    }
                }
            });
            recyclerView.setAdapter(adapter);

        }

        /**
         * 给适配器填充数据
         *
         * @param listItems
         */
        public void setData(List<SecondBean> listItems) {
           adapter.setData(listItems);
        }

    }


    /**
     * 点击事件监听接口
     */
    public interface OnItemClickListener {
        void onClick(int position, int viewID);
        void onStringSelected(String string);
    }
    private OnItemClickListener mOnItemClickListener;
    public void setOnItemClickListener(OnItemClickListener mOnItemClickListener) {
        this.mOnItemClickListener = mOnItemClickListener;
    }
}
