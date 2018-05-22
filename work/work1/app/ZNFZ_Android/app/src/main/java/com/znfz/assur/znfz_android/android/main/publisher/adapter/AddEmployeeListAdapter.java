package com.znfz.assur.znfz_android.android.main.publisher.adapter;

import android.content.Context;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.BaseAdapter;
import android.widget.TextView;

import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.main.publisher.bean.PublisherAddEmployeeBean;

import java.util.List;

/**
 * Created by dansihan on 2018/5/21.
 */

public class AddEmployeeListAdapter extends BaseAdapter {

    private Context mContext;
    private List<PublisherAddEmployeeBean> list;


    public AddEmployeeListAdapter(Context mContext, List<PublisherAddEmployeeBean> list) {
        this.mContext = mContext;
        this.list = list;
    }

    @Override
    public int getCount() {
        return list.size();
    }

    @Override
    public Object getItem(int position) {
        return position;
    }

    @Override
    public long getItemId(int position) {
        return position;
    }

    @Override
    public View getView(int position, View convertView, ViewGroup parent) {
        LayoutInflater inflater = LayoutInflater.from(mContext);
        ViewHolder holder = null;
        if (holder == null){
            convertView = inflater.inflate(R.layout.item_add_employee_job_name, null);
            holder = new ViewHolder();
            holder.name = (TextView) convertView.findViewById(R.id.tv_job_name);

        }else {
            holder = (ViewHolder) convertView.getTag();
        }

        holder.name.setText(list.get(position).getName());
        return convertView;
    }

    static class ViewHolder {
        TextView name;
        TextView toright;
    }
}
