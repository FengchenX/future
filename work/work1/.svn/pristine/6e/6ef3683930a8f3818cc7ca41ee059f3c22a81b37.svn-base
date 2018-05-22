package com.znfz.assur.znfz_android.android.common.utils;

import android.content.Context;

import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.global.Variable;

public class RoleUtils {

    /**
     * 获取用户角色字符串
     * @return
     */
    public static String getUserRoleString(Context context, String role) {
        String userRoleString = "" ;
        if(role.equals(Variable.USER_ROLE_PUBLISHER_MANAGER)) {
            userRoleString = context.getResources().getString(R.string.user_role_manager);
        } else if(role.equals(Variable.USER_ROLE_APPLICANT_COOK)) {
            userRoleString = context.getResources().getString(R.string.user_role_cook);
        } else if(role.equals(Variable.USER_ROLE_APPLICANT_WAITER)) {
            userRoleString = context.getResources().getString(R.string.user_role_waiter);
        }
        return userRoleString;
    }
}
