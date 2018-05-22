package com.znfz.assur.znfz_android.android.test.listview_test;

//import android.net.wifi.hotspot2.pps.Credential;
import android.Manifest;
import android.accounts.Account;
import android.content.pm.PackageManager;
import android.os.Environment;
import android.provider.UserDictionary;
import android.support.v4.app.ActivityCompat;
import android.support.v7.widget.DialogTitle;
import android.support.v7.widget.LinearLayoutManager;
import android.support.v7.widget.RecyclerView;
import android.util.TypedValue;
import android.view.Gravity;
import android.widget.LinearLayout;
import android.widget.TextView;

import com.orhanobut.logger.Logger;
import com.znfz.assur.znfz_android.R;
import com.znfz.assur.znfz_android.android.common.app.BaseActivity;
//
import org.web3j.crypto.Credentials;
import org.web3j.crypto.ECKeyPair;
import org.web3j.crypto.RawTransaction;
import org.web3j.crypto.Wallet;
import org.web3j.crypto.WalletFile;
import org.web3j.crypto.WalletUtils;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.Web3jFactory;
import org.web3j.protocol.admin.AdminFactory;
import org.web3j.protocol.admin.methods.response.NewAccountIdentifier;
import org.web3j.protocol.core.DefaultBlockParameterName;
import org.web3j.protocol.core.methods.response.EthAccounts;
import org.web3j.protocol.core.methods.response.EthGetBalance;
import org.web3j.protocol.core.methods.response.EthGetTransactionCount;
import org.web3j.protocol.core.methods.response.TransactionReceipt;
import org.web3j.protocol.core.methods.response.Web3ClientVersion;
import org.web3j.protocol.http.HttpService;
import org.web3j.tx.Transfer;
import org.web3j.utils.Convert;

import java.io.File;
import java.math.BigDecimal;
import java.math.BigInteger;
import java.security.KeyPair;
import java.security.SecureRandom;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 *
 * 两级展开，字母排序，listview测试
 *
 */
public class TwoLevelListViewSortActivity extends BaseActivity {

    private RecyclerView recyclerView;
    private LinearLayout linearLayout;

    private List<FirstBean> listItems;         // 数据
    private FirstAdapter firstAdapter;         // 列表适配器

    private Map<Integer,List<String>> datas = new HashMap<>();//模拟服务器返回数据
    private List<String> list=new ArrayList<>();//adapter数据源
    private Map<Integer,String> keys=new HashMap<>();//存放所有key的位置和内容

    @Override
    protected void initView() {
        setContentView(R.layout.activity_test_twolevellistviewsort);


        // 字母视图
        linearLayout = (LinearLayout)findViewById(R.id.ll_right) ;
        for (int i = 0; i < 24; i++) {
            TextView textView = new TextView(this);
            textView.setText("A" + i);
            textView.setGravity(Gravity.CENTER);
            textView.setLayoutParams(new LinearLayout.LayoutParams(LinearLayout.LayoutParams.FILL_PARENT, LinearLayout.LayoutParams.FILL_PARENT, 1.0f));
            linearLayout.addView(textView);
        }

        // 列表视图
        recyclerView = (RecyclerView)findViewById(R.id.lv_first);
        LinearLayoutManager layoutmanager = new LinearLayoutManager(TwoLevelListViewSortActivity.this);
        layoutmanager.setOrientation(1);
        recyclerView.setLayoutManager(layoutmanager);
    }


    @Override
    protected void initData() {

        listItems = new ArrayList<>();
        for (int i = 0; i < Math.random()*10+5; i++) {//(5-15)
            List<String> list=new ArrayList<>();
            for (int j = 0; j < Math.random()*10+5; j++) {//(5-15)
                list.add("第"+(j+1)+"个item，我属于标题"+i);
            }
            datas.put(i,list);
        }
        for (int i = 0; i < datas.size(); i++) {
            keys.put(list.size(),"我是第"+i+"个标题");
            for (int j = 0; j < datas.get(i).size(); j++) {
                list.add(datas.get(i).get(j));

                FirstBean bean = new FirstBean();
                bean.setName(datas.get(i).get(j));
                if(j == 0 || j == 2) {
                    List<SecondBean> subList = new ArrayList<>();
                    for(int m = 0; m < 10; m++) {
                        SecondBean secBean = new SecondBean();
                        secBean.setName("二级菜单");
                        subList.add(secBean);
                    }
                    bean.setSubList(subList);
                    bean.setSpread(false);
                }
                listItems.add(bean);
            }
        }

        final FloatingItemDecoration floatingItemDecoration=new FloatingItemDecoration(this, getResources().getColor(R.color.bg_znfz_normal),1,1);
        floatingItemDecoration.setKeys(keys);
        floatingItemDecoration.setmTitleHeight((int) TypedValue.applyDimension(TypedValue.COMPLEX_UNIT_DIP,30,getResources().getDisplayMetrics()));
        recyclerView.addItemDecoration(floatingItemDecoration);
        recyclerView.setHasFixedSize(true);

        firstAdapter = new FirstAdapter(TwoLevelListViewSortActivity.this);
        firstAdapter.setOnItemClickListener(new FirstAdapter.OnItemClickListener() {
            @Override
            public void onClick(int position, int viewID) {
                switch (viewID) {
                    case R.id.rl_item_content:
                        listItems.get(position).setSpread(!listItems.get(position).isSpread());
                        firstAdapter.setData(listItems);
                        break;

                default:

                    break;

                }
            }

            @Override
            public void onStringSelected(String string) {
                showToast(string);
            }
        });
        recyclerView.setAdapter(firstAdapter);
        firstAdapter.setData(listItems);

    }


}
