<template>
  <div id="app">
    <el-container>
      <el-header>
        <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect">
          <div class="headerLeft">
            <el-menu-item index="1">
              <router-link class="router" to="/view1">
                <div>主页</div>
              </router-link>
            </el-menu-item>
            <el-menu-item index="2">
              <router-link class="router" to="/view2">
                <div>处理中心</div>
              </router-link>
            </el-menu-item>
          </div>
          
          <div class="headerRight">
            <el-menu-item index="3">
              <router-link class="router" to="/login">
                <div>登录</div>
              </router-link>
            </el-menu-item>
            <el-menu-item index="4">
              <router-link class="router" to="/register">
                <div>注册</div>
              </router-link>
            </el-menu-item>
          </div>
        </el-menu>
      </el-header>

      <el-main>
        <!-- 路由出口 -->
        <!-- 路由匹配到的组件将渲染在这里 -->
        <router-view></router-view>
      </el-main>

      <el-footer>Footer</el-footer>
    </el-container>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'

export default {
  name: 'app',
  components: {
    HelloWorld,
  },
  data() {
      return {
        activeIndex: '1'
        // activeIndex2: '1'
      };
    },
  methods: {
    handleSelect(key, keyPath) {
      console.log(key, keyPath);
    }
  }
}
</script>

<style scoped>
/* #app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 0px;

} */
.router {
  color: black;
  text-decoration: none;
  /* line-height: 60px; */
}
.router.router-link-exact-active.active {
  color: #409EFF;
}

.el-header {
  background-color: #B3C0D1;
  color: #333;
  /* text-align: center; */
  line-height: 60px;
}

.el-footer {
  background-color: #B3C0D1;
  color: #333;
  text-align: center;
  line-height: 60px;
}

.el-main {
  background-color: #E9EEF3;
  color: #333;
  text-align: center;
  line-height: 60px;
}
.el-menu {
  background-color: #B3C0D1;
  color: #333;
  line-height: 60px;

  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: baseline;
  flex-wrap: nowrap;
}

.router div {
  width: 60px;
}
.headerLeft{
  width: 80%;
  height: 60px;
  /* display: inline-block; */
  display: flex;
  outline: none;
}

.headerRight{
  width: 5%;
  height: 60px;
  /* display: inline-block; */
  display: flex;
}


</style>
