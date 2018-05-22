package com.znfz.assur.znfz_android.android.main.applicant.adapter;

import android.content.Context;
import android.support.v7.widget.RecyclerView;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ImageView;
import android.widget.RelativeLayout;
import android.widget.TextView;

import com.facebook.drawee.view.SimpleDraweeView;
import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.utils.RoleUtils;
import com.znfz.assur.znfz_android.android.main.applicant.bean.ApplicantMenuBean;
import com.znfz.assur.znfz_android.android.main.applicant.bean.ApplicantMyFriendBean;

import java.util.List;

import butterknife.BindView;
import butterknife.ButterKnife;

/**
 * Created by dansihan on 2018/5/18.
 */

public class ApplicantMyFriendAdapter extends RecyclerView.Adapter<ApplicantMyFriendAdapter.ViewHolder> {

    private final Context mContext;
    private List<ApplicantMyFriendBean> listItems;

    public ApplicantMyFriendAdapter(Context context) {
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
    public ApplicantMyFriendAdapter.ViewHolder onCreateViewHolder(ViewGroup parent, int viewType) {
        View view = LayoutInflater.from(parent.getContext()).inflate(R.layout.item_applicant_my_friend, parent, false);
        return new ApplicantMyFriendAdapter.ViewHolder(view, parent.getContext());
    }

    @Override
    public void onBindViewHolder(ApplicantMyFriendAdapter.ViewHolder holder, final int position) {
        ApplicantMyFriendBean bean = listItems.get(position);
        // 设置视图数据
        holder.mFrinedName.setText(bean.getName());
        holder.mFriendPhoneNumber.setText(bean.getPhonenumber());
        holder.mFriendImg.setImageURI(bean.getUrl());

        if (listItems.get(position).isIfshowcheck()) {
            holder.img_check_friend.setVisibility(View.VISIBLE);
            if (listItems.get(position).isChecked()) {
                holder.img_check_friend.setBackground(mContext.getResources().getDrawable(R.mipmap.friend_checked));

            } else {
                holder.img_check_friend.setBackground(mContext.getResources().getDrawable(R.mipmap.friend_uncheck));
            }
        }

        holder.Rll_friend_item.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                if (mOnItemClickListener != null)
                    mOnItemClickListener.onClick(position, R.id.Rll_friend_item);
            }
        });

        holder.img_check_friend.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                if (mOnItemClickListener != null)
                    mOnItemClickListener.onClick(position, R.id.img_check_friend);
            }
        });
    }


    /**
     * 给适配器填充数据
     *
     * @param listItems
     */
    public void setData(List<ApplicantMyFriendBean> listItems) {
        this.listItems = listItems;

        /**
         * 初始化所有朋友为未选中状态
         */
        for (ApplicantMyFriendBean myFriendBean : listItems) {
            myFriendBean.setChecked(false);
        }
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
        @BindView(R.id.tv_friend_name)// 好有名字
                TextView mFrinedName;
        @BindView(R.id.tv_friend_phone_number)// 好友电话
                TextView mFriendPhoneNumber;
        @BindView(R.id.iv_my_friend_image_small)
        SimpleDraweeView mFriendImg;         // 好友头像

        @BindView(R.id.Rll_friend_item)//好友列表item
                RelativeLayout Rll_friend_item;

        @BindView(R.id.img_check_friend)//好友是否被选中
                ImageView img_check_friend;

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

    private ApplicantMyFriendAdapter.onItemClickListener mOnItemClickListener;

    public void setOnItemClickListener(ApplicantMyFriendAdapter.onItemClickListener mOnItemClickListener) {
        this.mOnItemClickListener = mOnItemClickListener;
    }


}
