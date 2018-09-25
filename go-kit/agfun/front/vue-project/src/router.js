import Vue from 'vue'
import Router from 'vue-router'
import View1 from './views/view1.vue'
import View2 from './views/view2.vue'
import Register from './views/Register.vue'
import Login from './views/Login.vue'

Vue.use(Router)
 
export default new Router({
    linkActiveClass: 'active',
    // 路由配置
    routes: [
    {
        path: '/view1',
        component: View1
    }, {
        path: '/view2',
        component: View2
    }, {
        path: '/register',
        component: Register
    }, {
        path: '/login',
        component: Login
    }, {
        path: '/*',
        component: View1
    }
    ]
})