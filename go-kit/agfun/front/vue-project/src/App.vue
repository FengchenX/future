<template>
  <div id="app">
    <div id="header">
      <router-link to="/view1">Go to view1</router-link>
      <router-link to="/view2">Go to view2</router-link>
    </div>
    

    <!-- 路由出口 -->
    <!-- 路由匹配到的组件将渲染在这里 -->
    <router-view></router-view>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'

export default {
  name: 'app',
  components: {
    HelloWorld,
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 70px;
}
#header {
    position: fixed;
    /*固定定位*/
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    width: 100%;
    height: 40px;
    /* border: 1px solid blue; */
    background: gray;
    font-size: 30px;
    
    opacity: 0.5;   /* Firefox、Chorme、Opera等主流浏览器识别 */
    filter:alpha(opacity=50);   /* IE6及以上IE浏览器识别 */
}
router-link {
  
}
</style>
