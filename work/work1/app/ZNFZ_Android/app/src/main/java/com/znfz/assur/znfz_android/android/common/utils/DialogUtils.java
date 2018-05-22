package com.znfz.assur.znfz_android.android.common.utils;

import android.content.Context;
import android.view.Gravity;
import com.flyco.dialog.listener.OnBtnClickL;
import com.flyco.dialog.widget.NormalDialog;

public class DialogUtils {

     /**
     * 提示对话框
     */
     public static void hintDialog(Context context,String msg, String btnText) {
        final NormalDialog dialog = new NormalDialog(context);
        dialog.content(msg)
                .btnNum(1)
                .style(NormalDialog.STYLE_TWO)
                .contentGravity(Gravity.CENTER)
                .btnText(btnText)
                .show();

        dialog.setOnBtnClickL(new OnBtnClickL() {
            @Override
            public void onBtnClick() {
                dialog.dismiss();
            }
        });
    }

//    /**
//     * ActionSheetDialog选择对话框
//     */
//    public void ActionSheetDialog() {
//        final String[] stringItems = {"接收消息并提醒", "接收消息但不提醒", "收进群助手且不提醒", "屏蔽群消息"};
//        final ActionSheetDialog dialog = new ActionSheetDialog(context, stringItems, null);
//        dialog.title("选择群消息提醒方式\r\n(该群在电脑的设置:接收消息并提醒)")//
//                .titleTextSize_SP(14.5f)//
//                .show();
//
//        dialog.setOnOperItemClickL(new OnOperItemClickL() {
//            @Override
//            public void onOperItemClick(AdapterView<?> parent, View view, int position, long id) {
//                dialog.dismiss();
//            }
//        });
//    }
//
//    /**
//     * 普通选择对话框
//     */
//    public void normalListDialog() {
//        DialogMenuItem item1 = new DialogMenuItem("收藏", 1);
//        DialogMenuItem item2 = new DialogMenuItem("分享", 2);
//        ArrayList<DialogMenuItem> list = new ArrayList<DialogMenuItem>();
//        list.add(item1);
//        list.add(item2);
//
//        final NormalListDialog dialog = new NormalListDialog(context, list);
//        dialog.title("请选择")
//                .titleTextSize_SP(18)
//                .titleBgColor(Color.parseColor("#409ED7"))
//                .itemPressColor(Color.parseColor("#85D3EF"))
//                .itemTextColor(Color.parseColor("#303030"))
//                .itemTextSize(14)
//                .cornerRadius(0)
//                .widthScale(0.8f)
//                .show(R.style.Theme_AppCompat_DayNight_Dialog);
//
//        dialog.setOnOperItemClickL(new OnOperItemClickL() {
//            @Override
//            public void onOperItemClick(AdapterView<?> parent, View view, int position, long id) {
//                dialog.dismiss();
//            }
//        });
//    }

}
