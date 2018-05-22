package com.znfz.assur.znfz_android.android.common.utils;

import android.content.Context;
import android.content.SharedPreferences;
import android.text.TextUtils;

public class SPUtils {
    /**
     * sq 构造 持有context 对象
     * @param context
     */
    public SharedPreferences sp;
    public Context context;


    public SPUtils(Context context) {
        this.context = context;
        if (sp==null){
            sp = context.getSharedPreferences("", Context.MODE_PRIVATE);
        }
    }


    /**
     * getString/char
     * @param key 字段名
     * @return value 返回值
     */
    public String getString(String key){
        String value = sp.getString(key, "");
        if (!TextUtils.isEmpty(value)){
            return value;
        } else {
            return "";
        }
    }

    /**
     * getInt方法
     * @param key 字段名
     *
     * @return value int返回值
     */
    public int getInt(String key){
        int value = sp.getInt(key, 0);
        if (value!=0){
            return value;
        }else {
            return -1;
        }
    }

    /**
     * getLong 方法
     *
     * @param key 字段名
     * @return Long返回值
     */
    public Long getLong(String key){
        Long value = sp.getLong(key, 0);
        if (value!=0){
            return value;
        }else {
            return Long.valueOf(0);
        }
    }
    /**
     *  putString 方法
     *
     * @param key 字段名
     * @param value 字段值
     */
    public void putString(String key,String value){
        sp.edit().putString(key,value).commit();

    }
    /**
     *  putInt 方法
     *
     * @param key 字段名
     * @param value 字段值
     */
    public void putInt(String key,int value){
        sp.edit().putInt(key,value).commit();

    }
    /**
     *  putLong 方法
     *
     * @param key 字段名
     * @param value 字段值
     */
    public void putLong(String key,Long value){
        sp.edit().putLong(key,value).commit();

    }
    // 清除数据
    public void clearSpSpace(){
        sp.edit().clear().commit();
    }

}
