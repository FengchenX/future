package com.znfz.assur.znfz_android.android.test.listview_test;

import android.content.Context;
import android.support.v7.widget.RecyclerView;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.RelativeLayout;
import android.widget.TextView;
import com.znfz.assur.znfz_android.R;
import java.util.List;

public class SecondAdapter extends RecyclerView.Adapter<SecondAdapter.ViewHolder>  {

    private final Context mContext;
    private List<SecondBean> listItems;


    public SecondAdapter(Context context) {
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
    public SecondAdapter.ViewHolder onCreateViewHolder(ViewGroup parent, int viewType) {
        View view = LayoutInflater.from(parent.getContext()).inflate(R.layout.item_test_twolevelsort_second_list, parent, false);
        return new SecondAdapter.ViewHolder(view, parent.getContext());
    }

    @Override
    public void onBindViewHolder(SecondAdapter.ViewHolder holder, final int position) {

        holder.rlContent.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                if(mOnItemClickListener != null) {
                    mOnItemClickListener.onStringSelected(listItems.get(position).getName());
                }
            }
        });
        holder.tvName.setText(listItems.get(position).getName());
    }




    /**
     * 给适配器填充数据
     *
     * @param listItems
     */
    public void setData(List<SecondBean> listItems) {
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

        public ViewHolder(View itemView, final Context context) {
            super(itemView);
            this.context = context;

            rlContent = (RelativeLayout) itemView.findViewById(R.id.rl_item_content);
            tvName = (TextView) itemView.findViewById(R.id.item_name);

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
