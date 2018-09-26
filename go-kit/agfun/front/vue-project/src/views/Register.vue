<template>
    <div class="register">
        <!-- <MyVue msg="Welcome to Your Vue.js App"/> -->
        <div class="center">
            <div id="userName">UserName:
                <input v-model="userName">
            </div>
            <div class="password1">Password:
                <input v-model="password1">
            </div>
            <div class="password2">Password:
                <input v-model="password2">
            </div>
            <div id="ensure">
                <button v-on:click="register">register</button>
                <button v-on:click="login">login</button>
            </div>
        </div>
    </div>
</template>
 
<script>
// import MyVue from '../components/MyVue.vue'

export default {
    name: 'register',
    components: {
        // MyVue
    },
    data() {
        return {
            userName: '',
            password1: '',
            password2: ''
        }
    },
    methods: {
        register: function(){
            console.log('register', this.userName, this.password1, this.password2, this);
            this.$router.push({path:'/view1'});
        },
        login: function(){
            console.log('login')
        }
    }
}
</script>
 
<style scoped>
.register{
    height: 50rem;
    
    display: flex;
    justify-content: center;
    align-items: center;
}
.center{
    width: 40rem;
    height: 25rem;
    /* background-color: yellow; */

}
</style>