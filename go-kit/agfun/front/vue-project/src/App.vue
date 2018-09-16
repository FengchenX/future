<template>
  <div id="app">
    <!--img alt="Vue logo" src="./assets/logo.png"-->
    <MyVue msg="Welcome to Your Vue.js App"/>
    <router-link to="/view1">Go to view1</router-link>
    <router-link to="/view2">Go to view2</router-link>
    <!-- 路由出口 -->
    <!-- 路由匹配到的组件将渲染在这里 -->
    <router-view></router-view>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'
import MyVue from './components/MyVue.vue'


export default {
  name: 'app',
  components: {
    HelloWorld,
    MyVue
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
  margin-top: 60px;
}
</style>
