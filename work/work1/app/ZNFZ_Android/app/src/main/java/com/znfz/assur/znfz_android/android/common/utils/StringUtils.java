package com.znfz.assur.znfz_android.android.common.utils;

import android.app.ActivityManager;
import android.content.Context;
import android.content.pm.PackageInfo;
import android.content.pm.PackageManager;
import android.net.ConnectivityManager;
import android.net.NetworkInfo;
import android.text.TextUtils;

import org.json.JSONObject;

import java.io.UnsupportedEncodingException;
import java.lang.reflect.Field;
import java.net.URLEncoder;
import java.text.DecimalFormat;
import java.util.HashMap;
import java.util.Iterator;
import java.util.List;
import java.util.Map;
import java.util.Random;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

/**
 * 字符串工具类
 */
public class StringUtils {
    private static PackageInfo sPackageInfo;
    public static Random random = new Random(100);

    /**
     * 字符串长度大于int a时保留前a个+...
     *
     * @param string
     * @return
     */
    public static String cutString(String string, int limitNum) {
        if (string.length() > limitNum) {
            string = string.substring(0, limitNum) + "...";
        }
        return string;
    }

    /**
     * 车牌号格式：汉字 + A-Z + 5位A-Z或0-9（只包括了普通车牌号，教练车和部分部队车等车牌号不包括在内）正则表达式有局限性，比如第一位只限定是汉字，
     * 没限定只有34个省汉字缩写；车牌号不存在字母I和O，防止和1、0混淆；部分车牌无法分辨等等
     *
     * @param carnumber
     * @return
     */
    public static boolean isCarnumberNO(String carnumber) {
        return !TextUtils.isEmpty(carnumber) && carnumber.matches("[\u4e00-\u9fa5]{1}[A-Z]{1}[A-Z_0-9]{5}");
    }

    /**
     * 身份证号正则判断
     *
     * @param idNum
     * @return
     */
    public static boolean isIdNum(String idNum) {
        return !TextUtils.isEmpty(idNum) && idNum.matches("(\\d{14}[0-9a-zA-Z])|(\\d{17}[0-9a-zA-Z])");
    }

    /**
     * 判断是否为+86手机号码
     *
     * @param number
     * @return
     */
    public static boolean is86Num(String number) {
//        return !TextUtils.isEmpty(number) && number.matches("^((\\+{0,1}86){0,1})1[0-9]{10}");
        return !TextUtils.isEmpty(number) && number.matches("^[+][8][6][1][3-8]\\d{9}$");
    }

    /**
     * 判断是否为邮箱
     */
    public static boolean isEmail(String Email) {
        return !TextUtils.isEmpty(Email) && Email.matches("^([a-z0-9A-Z]+[-|\\\\.]?)+[a-z0-9A-Z]@([a-z0-9A-Z]+(-[a-z0-9A-Z]+)?\\\\.)+[a-zA-Z]{2,}$");
    }

    /**
     * 判断是否为手机号码
     *
     * @param number
     * @return
     */
    public static boolean isMobilePhoneNumber(String number) {
        return !TextUtils.isEmpty(number) && number.matches("^[1][3-8]\\d{9}$");
    }

    public static String getRandom() {
        return random.nextInt() + "";
    }

    public static String getApkVersionName(Context context) {
        return getPackageInfo(context).versionName;
    }

    public synchronized static PackageInfo getPackageInfo(Context context) {
        if (sPackageInfo == null) {
            try {
                sPackageInfo = context.getPackageManager().getPackageInfo(
                        context.getPackageName(), 0);
            } catch (PackageManager.NameNotFoundException e) {
                assert false;
                return null;
            }
        }
        return sPackageInfo;
    }

    public static String encodeJson(JSONObject json) {
        return encodeJson(json, "utf-8");
    }

    public static String encodeJson(JSONObject json, String paramsEncoding) {
        try {
            StringBuilder encodedParams = new StringBuilder();
            Iterator<String> itr = json.keys();
            while (itr.hasNext()) {
                String key = itr.next();
                encodedParams.append(URLEncoder.encode(key, paramsEncoding));
                encodedParams.append('=');
                encodedParams.append(URLEncoder.encode(json.opt(key).toString(), paramsEncoding));
                encodedParams.append('&');
            }
            return encodedParams.toString();
        } catch (UnsupportedEncodingException uee) {
            throw new RuntimeException("Encoding not supported: " + paramsEncoding, uee);
        }
    }

    /**
     * 得到设备屏幕的宽度
     */
    public static int getScreenWidth(Context context) {
        return context.getResources().getDisplayMetrics().widthPixels;
    }

    /**
     * 得到设备屏幕的高度
     */
    public static int getScreenHeight(Context context) {
        return context.getResources().getDisplayMetrics().heightPixels;
    }

    /**
     * 得到设备的密度
     */
    public static float getScreenDensity(Context context) {
        return context.getResources().getDisplayMetrics().density;
    }

    /**
     * 把密度转换为像素
     */
    public static int dip2px(Context context, float px) {
        final float scale = getScreenDensity(context);
        return (int) (px * scale + 0.5);
    }

    /**
     * 网络是否可用
     *
     * @param context
     * @return
     */
    public static boolean isNetworkAvailable(Context context) {
        if (context != null) {
            ConnectivityManager mConnectivityManager = (ConnectivityManager) context
                    .getSystemService(Context.CONNECTIVITY_SERVICE);
            NetworkInfo mNetworkInfo = mConnectivityManager.getActiveNetworkInfo();
            if (mNetworkInfo != null) {
                return mNetworkInfo.isAvailable();
            }
        }
        return false;
    }

    public static Map<String, Object> PO2Map(Object o) throws Exception {
        Map<String, Object> map = new HashMap<String, Object>();
        Field[] fields = null;
        String clzName = o.getClass().getSimpleName();
        fields = o.getClass().getDeclaredFields();
        for (Field field : fields) {
            field.setAccessible(true);
            String proName = field.getName();
            Object proValue = field.get(o);
            map.put(proName, proValue);
        }
        return map;
    }

    /**
     * 判断某个服务是否正在运行的方法
     *
     * @param mContext
     * @param serviceName 是包名+服务的类名（例如：net.loonggg.testbackstage.TestService）
     * @return true代表正在运行，false代表服务没有正在运行
     */
    public static boolean isServiceWork(Context mContext, String serviceName) {
        boolean isRunning = false;
        ActivityManager activityManager = (ActivityManager)
                mContext.getSystemService(Context.ACTIVITY_SERVICE);
        List<ActivityManager.RunningServiceInfo> serviceList
                = activityManager.getRunningServices(Integer.MAX_VALUE);
        if (!(serviceList.size() > 0)) {
            return false;
        }
        for (int i = 0; i < serviceList.size(); i++) {
            if (serviceList.get(i).service.getClassName().equals(serviceName)) {
                isRunning = true;
                break;
            }
        }
        return isRunning;
    }

    /**
     * 将 07月07日改为  7月7日
     *
     * @param date
     * @return
     */
    public static String formatDate(String date) {
        if ("0".equals(String.valueOf(date.charAt(3)))) {
            date = date.substring(0, 3) + date.substring(4, date.length());
        }
        if (date.startsWith("0")) {
            date = date.substring(1, date.length());
        }
        return date;
    }

    /**
     * 路劲规划中时间换算
     *
     * @param second
     * @return
     */
    public static String getFriendlyTime(int second) {
        if (second > 3600) {
            int hour = second / 3600;
            int miniate = (second % 3600) / 60;
            return hour + "h" + miniate + "min";
        }
        if (second >= 60) {
            int miniate = second / 60;
            return miniate + "min";
        }
        return second + "s";
    }


    /**
     * 验证是否是URL
     *
     * @param url
     * @return
     * @author YOLANDA
     */
    public static boolean isTrueURL(String url) {
        String regex = "^(https?|ftp|file)://[-a-zA-Z0-9+&@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&@#/%=~_|]";
        Pattern patt = Pattern.compile(regex);
        Matcher matcher = patt.matcher(url);
        return matcher.matches();
    }

    public static double doubleParse(String stringDouble) {
        return Double.parseDouble(!TextUtils.isEmpty(stringDouble) ? stringDouble : "-1");
    }

    /**
     * 将string类型的数字转为两位小数的string
     *
     * @param stringDouble
     * @return
     */
    public static String doubleStringParse(String stringDouble) {
        String format = new DecimalFormat("0.00").format(doubleParse(stringDouble));
        return format;
    }

    /**
     * 将string类型的数字转为两位小数的string
     * 带人命币符号前缀
     *
     * @param stringDouble
     * @return
     */
    public static String doubleRMBStringParse(String stringDouble) {
        String format = new DecimalFormat("0.00").format(doubleParse(stringDouble));
        return "¥" + format;
    }

    /**
     * 将double类型的数字转为两位小数的string
     * 带人命币符号前缀
     *
     * @param stringDouble
     * @return
     */
    public static String doubleRMBStringParse(double stringDouble) {
        String format = new DecimalFormat("0.00").format(stringDouble);
        return "¥" + format;
    }

    /**
     * m转换km
     *
     * @param mValue
     * @return
     */
    public static String mToKm(String mValue) {
        if (!TextUtils.isEmpty(mValue)) {
            if (doubleParse(mValue) > 0) {
                String format = new DecimalFormat("0.0").format(doubleParse(mValue) / 1000);
                return format + " km";
            } else {
                return "未知";
            }
        } else {
            return "未知";
        }
    }
}
