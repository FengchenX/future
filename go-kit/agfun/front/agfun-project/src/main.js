

// 创建一个 Vue 实例或 "ViewModel"
// 它连接 View 与 Model
vm = new Vue({
    el: '#app',
    data: {
        people: [{
            name: 'Jack',
            age: 30,
            sex: 'Male'
        }, {
            name: 'Bill',
            age: 26,
            sex: 'Male'
        }, {
            name: 'Tracy',
            age: 22,
            sex: 'Female'
        }, {
            name: 'Chris',
            age: 36,
            sex: 'Male'
        }]
    }
})