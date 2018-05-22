package com.znfz.assur.znfz_android.android.common.utils;

import android.content.Context;
import android.widget.Toast;

public class ToastUtils {

    private static Toast toast;


    public static void showToast(Context context, String content) {
        if (toast == null) {
            toast = Toast.makeText(context.getApplicationContext(), content, 0);
        } else {
            toast.setText(content);
        }

        toast.show();
    }

    public static void showToast(Context context, int strID) {
        showToast(context, context.getString(strID));
    }

}
