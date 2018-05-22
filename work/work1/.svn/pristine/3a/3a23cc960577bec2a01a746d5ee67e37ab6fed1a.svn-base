package com.znfz.assur.znfz_android.android.test.web3j_test;

import android.Manifest;
import android.content.pm.PackageManager;
import android.os.Environment;
import android.support.v4.app.ActivityCompat;

import com.orhanobut.logger.Logger;
import com.znfz.assur.znfz_android.android.common.app.BaseActivity;
import com.znfz.assur.znfz_android.android.common.global.UserInfo;
import com.znfz.assur.znfz_android.android.test.listview_test.TwoLevelListViewSortActivity;

import org.web3j.crypto.Credentials;
import org.web3j.crypto.ECKeyPair;
import org.web3j.crypto.WalletUtils;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.Web3jFactory;
import org.web3j.protocol.core.DefaultBlockParameterName;
import org.web3j.protocol.core.methods.response.EthAccounts;
import org.web3j.protocol.core.methods.response.EthGetBalance;
import org.web3j.protocol.core.methods.response.EthGetTransactionCount;
import org.web3j.protocol.core.methods.response.Web3ClientVersion;
import org.web3j.protocol.http.HttpService;

import java.io.File;
import java.math.BigInteger;
import java.util.List;

public class Web3jActivity extends BaseActivity{


    @Override
    protected void initView() {

    }

    @Override
    protected void initData() {

        //获取节点运行geth客户端的版本号
        // String url = "https://mainnet.infura.io/your api-key";
        String url = "http://39.108.80.66:8546/your api-key";
        final Web3j web3 = Web3jFactory.build(new HttpService(url));
        try {
            //获取节点运行geth客户端的版本号
            Web3ClientVersion web3ClientVersion = web3.web3ClientVersion().sendAsync().get();
            String clientVersion = web3ClientVersion.getWeb3ClientVersion();
            Logger.e("clientVersion(assur) :" + clientVersion);

            // 返回当前节点持有的帐户钱包地址列表
            EthAccounts ethAccounts = web3.ethAccounts().sendAsync().get();
            List<String> accountList = ethAccounts.getAccounts();
            Logger.e("accountList(assur) :" + accountList);


//            //Once you have obtained the next available nonce, the value can then be used to create your transaction object:
//            RawTransaction rawTransaction  = RawTransaction.createEtherTransaction(
//                    nonce, <gas price>, <gas limit>, <toAddress>, <value>);
//            //The transaction can then be signed and encoded:
//            byte[] signedMessage = TransactionEncoder.signMessage(rawTransaction, <credentials>);
//            String hexValue = Numeric.toHexString(signedMessage);

//            //The transaction is then sent using eth_sendRawTransaction:
//            EthSendTransaction ethSendTransaction = web3j.ethSendRawTransaction(hexValue).sendAsync().get();
//            String transactionHash = ethSendTransaction.getTransactionHash();

            new Thread(new Runnable() {
                @Override
                public void run() {
                    try {
                        // Credentials
                        String account = web3.ethAccounts().send().getAccounts().get(0);
                        Credentials credentials = Credentials.create(account);
                        ECKeyPair keyPair = credentials.getEcKeyPair();
                        BigInteger privateKey = keyPair.getPrivateKey();
                        BigInteger publicKey = keyPair.getPublicKey();
                        Logger.e("privateKey(assur) :" + privateKey);
                        Logger.e("publicKey(assur) :" + publicKey);

                        // 获取余额
                        EthGetBalance ethGetBalance = web3.ethGetBalance(account, DefaultBlockParameterName.LATEST).send();
                        BigInteger balance = ethGetBalance.getBalance();
                        Logger.e("balance(assur) :" + balance);


                        // 创建钱包1：获取权限
                        int REQUEST_EXTERNAL_STORAGE = 1;
                        String[] PERMISSIONS_STORAGE = {
                                Manifest.permission.READ_EXTERNAL_STORAGE,
                                Manifest.permission.WRITE_EXTERNAL_STORAGE
                        };
                        int permission = ActivityCompat.checkSelfPermission(Web3jActivity.this, Manifest.permission.WRITE_EXTERNAL_STORAGE);

                        if (permission != PackageManager.PERMISSION_GRANTED) {
                            // We don't have permission so prompt the user
                            ActivityCompat.requestPermissions(
                                    Web3jActivity.this,
                                    PERMISSIONS_STORAGE,
                                    REQUEST_EXTERNAL_STORAGE
                            );
                        }
                        // 创建钱包2：创建
                        String filePath = Environment.getExternalStorageDirectory().toString() + "/aaaaaa";
                        File file = new File(filePath);
                        if (!file.exists()) {
                            file.mkdir();
                        }
                        String password = UserInfo.USER_PASSWORD;
                        // String fileName = WalletUtils.generateFullNewWalletFile(password, file); // 模拟器会报内存溢出异常
                        String fileName = WalletUtils.generateNewWalletFile(password, new File(filePath), false);// true:模拟器会报内存溢出异常
                        Logger.e("fileName(assur) :" + fileName);

                        Credentials credentials_2 = WalletUtils.loadCredentials(password, filePath+"/"+fileName);
                        Logger.e("credentials_2.getAddress()(assur) :" + credentials_2.getAddress());
                        Logger.e("credentials_2.getEcKeyPair().getPrivateKey()(assur) :" + credentials_2.getEcKeyPair().getPrivateKey().toString());
                        Logger.e("credentials_2.getEcKeyPair().getPublicKey()(assur) :" + credentials_2.getEcKeyPair().getPublicKey().toString());

                        // 创建账户
                        EthGetTransactionCount ethGetTransactionCount = web3.ethGetTransactionCount(
                                credentials_2.getAddress(), DefaultBlockParameterName.LATEST).send();
                        BigInteger nonce = ethGetTransactionCount.getTransactionCount();
                        Logger.e("nonce(assur) :" + nonce);



                    } catch (Exception e1) {
                        Logger.e("e1(assur) :" + e1.toString());
                    }
                }
            }).start();



        } catch (Exception e) {
            Logger.e("Exception(assur) :" + e.toString());
        }



    }
}
