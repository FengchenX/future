import Vue from 'vue'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import App from './App.vue'
import router from './router.js'
import axios from 'axios'; //引入axios


Vue.use(ElementUI);
// Vue.use(ElementUI, { size: 'small', zIndex: 3000 });

Vue.prototype.$ajax  = axios; //修改Vue的原型属性
Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
