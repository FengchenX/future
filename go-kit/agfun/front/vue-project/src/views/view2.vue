<template>
    <!-- <div id="box">
        <div class="inner item">1</div>
        <div class="inner">2</div>
        <div class="inner">3</div>
        <div class="inner">4</div>
        <div class="inner">5</div>
    </div> -->
    <div></div>
</template>
 
<script>
    export default {
        name: 'view2',
    }
</script>
 
<style scoped>
    #box{
        display: flex;
        /* flex-direction: row;
        flex-wrap: wrap; */
        flex-flow: row wrap;
        /* justify-content: flex-start; */
        /* align-items: baseline; */
        width: 500px;
        border: 1px solid red;
    }
    .inner{
        width: 100px;
        height: 100px;
        background: #8a4182;
        margin: 10px;
    }
    .inner2{
        width: 100px;
        height: 300px;
        background: #8a4182;
        margin: 10px;
    }
    .inner:nth-child(1){
      flex-grow: 1;
    }
</style>
>
</style>