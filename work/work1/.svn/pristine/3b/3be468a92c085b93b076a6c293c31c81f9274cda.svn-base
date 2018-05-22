package com.znfz.assur.znfz_android.android.common.app;

import android.app.Application;
import android.app.Dialog;
import android.content.Context;
import android.graphics.Bitmap;
import android.os.Handler;
import android.view.Gravity;

import com.facebook.cache.disk.DiskCacheConfig;
import com.facebook.common.internal.Supplier;
import com.facebook.common.util.ByteConstants;
import com.facebook.drawee.backends.pipeline.Fresco;
import com.facebook.imagepipeline.cache.MemoryCacheParams;
import com.facebook.imagepipeline.core.ImagePipelineConfig;
import com.flyco.dialog.listener.OnBtnClickL;
import com.flyco.dialog.widget.NormalDialog;
import com.orhanobut.logger.AndroidLogAdapter;
import com.orhanobut.logger.FormatStrategy;
import com.orhanobut.logger.Logger;
import com.orhanobut.logger.PrettyFormatStrategy;
import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.net.NetChangeObserver;
import com.znfz.assur.znfz_android.android.common.net.NetStateReceiver;
import com.znfz.assur.znfz_android.android.common.net.NetUtils;
import com.znfz.assur.znfz_android.android.common.utils.DialogUtils;

/**
 *  Application类
 * Created by assur on 2018/4/20 0900.
 */
public class ZnfzApplication  extends Application{

    //在整个应用执行的过程中，需要提供的变量
    public static Context context;   //需要使用的上下文对象
    public static Handler handler;   //需要使用的Handler
    public static Thread mainThread; //提供主线程对象
    public static int mainThreadId;  //提供主线程对象的Id

    // 网络观察者
    protected NetChangeObserver mNetChangeObserver = null;
    protected static boolean netConnected = false;
    public static boolean isNetConnected() {
        return netConnected;
    }

    @Override
    public void onCreate() {
        super.onCreate();

        context = this.getApplicationContext();
        handler = new Handler();
        mainThread = Thread.currentThread();       //实例化Application当前的线程为主线程
        mainThreadId = android.os.Process.myTid(); //获取当前线程的Id

        // 初始化fresco
        initFresco(this);
        // 初始化loger
        initLogger();
        // 开启网络广播监听
        NetStateReceiver.registerNetworkStateReceiver(this);
        netChangeObserver();

    }

    /**
     *  网络变化监听
     */
    private  void netChangeObserver() {
        mNetChangeObserver = new NetChangeObserver() {
            @Override
            public void onNetConnected(NetUtils.NetType type) {
                if(type != NetUtils.NetType.NONE) {
                    netConnected = true;
                } else {
                    netConnected = false;
                }

            }

            @Override
            public void onNetDisConnect() {
                netConnected = false;
                DialogUtils.hintDialog(AppManager.getAppManager().currentActivity(),
                        getResources().getString(R.string.no_net_message),
                        getResources().getString(R.string.i_known));
            }
        };
        NetStateReceiver.registerObserver(mNetChangeObserver);
    }


    /**
     * 配置fresco(图片加载)相关
     *
     * @param applicationContext
     */
    private void initFresco(Context applicationContext) {
        ImagePipelineConfig.Builder configBuilder = ImagePipelineConfig.newBuilder(applicationContext);
        //初始化MemoryCacheParams对象。从字面意思也能看出，这是内存缓存的Params
        final MemoryCacheParams bitmapCacheParams = new MemoryCacheParams(
                // 内存缓存最大值，这里用最大运行内存的三分之一
                (int) Runtime.getRuntime().maxMemory() / 3,
                // 内存缓存单个文件最大值，这里就用Integer的最大值了。
                Integer.MAX_VALUE,
                // 可以释放的内存缓存的最大值，当内存缓存到了上面设置的最大值，最多可以释放多少内存
                (int) Runtime.getRuntime().maxMemory() / 3,
                //可以被释放缓存的文件数
                Integer.MAX_VALUE,
                //Max cache entry size. 最大缓存入口大小，怎么理解不知道。。。
                Integer.MAX_VALUE);
        //初始化DiskCacheConfig对象
        DiskCacheConfig diskCacheConfig = DiskCacheConfig.newBuilder(applicationContext)
                //设置磁盘缓存路径
                .setBaseDirectoryPath(applicationContext.getExternalCacheDir())
                //设置磁盘缓存文件夹名
                .setBaseDirectoryName("imagecache")
                //设置磁盘缓存最大值
                .setMaxCacheSize(300 * ByteConstants.MB)
                .build();
        //getImagePipelineConfig设置属性
        configBuilder
                //设置内存缓存的Params
                .setBitmapMemoryCacheParamsSupplier(
                        new Supplier<MemoryCacheParams>() {
                            public MemoryCacheParams get() {
                                return bitmapCacheParams;
                            }
                        })
                //设置磁盘缓存的配置
                .setMainDiskCacheConfig(diskCacheConfig)
                //设置图片压缩质量，不设置默认为ARGB_8888
                .setBitmapsConfig(Bitmap.Config.RGB_565)
                //这个方法使得处理图片的速度比常规的裁剪更快，并且同时支持PNG，JPG以及WEP格式的图片（百度的）
                .setDownsampleEnabled(true);

        Fresco.initialize(this, configBuilder.build());
    }

    /**
     * 日志打印初始化
     */
    private void initLogger() {

        // Logger.addLogAdapter(new AndroidLogAdapter());
        FormatStrategy formatStrategy = PrettyFormatStrategy.newBuilder()
                .showThreadInfo(false)  //是否选择显示线程信息，默认为true
                .methodCount(2)         //方法数显示多少行，默认2行
                .methodOffset(1)        //隐藏方法内部调用到偏移量，默认5
                .tag("ZNFZ(Assur)")     //自定义TAG全部标签，默认PRETTY_LOGGER
                .build();
        Logger.addLogAdapter(new AndroidLogAdapter(formatStrategy));
    }


    @Override
    public void onLowMemory() {
        super.onLowMemory();

        NetStateReceiver.unRegisterNetworkStateReceiver(this);
        android.os.Process.killProcess(android.os.Process.myPid());
    }
}
