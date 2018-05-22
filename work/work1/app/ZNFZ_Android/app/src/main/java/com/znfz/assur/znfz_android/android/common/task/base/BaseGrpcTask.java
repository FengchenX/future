package com.znfz.assur.znfz_android.android.common.task.base;


import android.os.AsyncTask;
import android.os.Handler;
import android.os.Looper;
import android.util.Log;

import com.orhanobut.logger.Logger;

import java.util.HashMap;
import java.util.concurrent.TimeUnit;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

/**
 *  Grpc 的异步AsyncTask处理类的基类
 * Created by assur on 2018/4/20 0920.
 */
public abstract class BaseGrpcTask <T>  extends AsyncTask<String, String, T> {

    private ManagedChannel mChannel;


    @Override
    protected void onPreExecute() {

    }

    /**
     * 异步操作
     * @param channel
     * @return
     */
    protected abstract T doInback(ManagedChannel channel);

    /**
     * 更新UI操作
     * @param result
     */
    protected abstract void onSucceed(T result);
    protected abstract void onError();

    @Override
    protected T doInBackground(String... params) {
        try {
            mChannel = ManagedChannelBuilder.forAddress(params[0], Integer.parseInt(params[1]))
                    .usePlaintext(true)
                    .build();
            return doInback(mChannel);
        } catch (Exception e) {
            Logger.e(e.getMessage());
            e.printStackTrace();
        }
        return null;
    }

    @Override
    protected void onPostExecute(T result) {
        try {
            if(result != null) {
                onSucceed(result);
            } else {
                Logger.e("result == null");
                error();
            }
            mChannel.shutdown().awaitTermination(1, TimeUnit.SECONDS);
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
            Logger.e(e.getMessage());
            error();
        }
        Log.i("TAG", "onPostExecute:  result = " + result);
    }



    protected void error() {
        Handler mainHandler = new Handler(Looper.getMainLooper());
        mainHandler.post(new Runnable() {
            @Override
            public void run() {
                onError();
            }
        });

    }

}