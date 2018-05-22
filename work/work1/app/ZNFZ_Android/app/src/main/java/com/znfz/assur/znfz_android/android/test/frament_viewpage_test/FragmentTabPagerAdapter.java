package com.znfz.assur.znfz_android.android.test.frament_viewpage_test;

import android.support.v4.app.Fragment;
import android.support.v4.app.FragmentManager;
import android.support.v4.app.FragmentPagerAdapter;

import java.util.List;

public class FragmentTabPagerAdapter extends android.support.v4.app.FragmentPagerAdapter {

    private FragmentManager mfragmentManager;
    private List<Fragment> mlist;

    //这是一段构造器，我没写的时候，第一次代码是报错的，在我做了下面这个构造器之后，没有报错！！！
    public FragmentTabPagerAdapter(FragmentManager fm, List<Fragment> list) {
        super(fm);
        this.mlist = list;
    }
    //显示第几个页面
    @Override
    public Fragment getItem(int position) {
        return mlist.get(position);
    }
    //一共有几个页面，注意，使用Fragment特有的构造器时，和ViewPager的原生构造器的方法不同
    @Override
    public int getCount() {
        return mlist.size();
    }
}
