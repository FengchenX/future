<template>
  <div class="hello">
    <!-- <h2>{{ msg }}</h2> -->
    <div class="body">
      <video width="320" height="240" controls>
        <source src="../assets/movie.mp4" type="video/mp4">
        <!-- source src="../assets/movie.ogg" type="video/ogg"-->
      </video>
      <video width="320" height="240" controls>
        <source src="../assets/movie.mp4" type="video/mp4">
        <!-- source src="../assets/movie.ogg" type="video/ogg"-->
      </video>
      <video width="320" height="240" controls>
        <source src="../assets/movie.mp4" type="video/mp4">
        <!-- source src="../assets/movie.ogg" type="video/ogg"-->
      </video>
      <video width="320" height="240" controls>
        <source src="../assets/movie.mp4" type="video/mp4">
        <!-- source src="../assets/movie.ogg" type="video/ogg"-->
      </video>
      <video width="320" height="240" controls>
        <source src="../assets/movie.mp4" type="video/mp4">
        <!-- source src="../assets/movie.ogg" type="video/ogg"-->
      </video>
      <video width="320" height="240" controls>
        <source src="../assets/movie.mp4" type="video/mp4">
        <!-- source src="../assets/movie.ogg" type="video/ogg"-->
      </video>
        <video width="320" height="240" controls>
        <source src="../assets/movie.mp4" type="video/mp4">
        <!-- source src="../assets/movie.ogg" type="video/ogg"-->
      </video>
      <video width="320" height="240" controls>
        <source src="../assets/movie.mp4" type="video/mp4">
        <!-- source src="../assets/movie.ogg" type="video/ogg"-->
      </video>
      <video width="320" height="240" controls>
        <source src="../assets/movie.mp4" type="video/mp4">
        <!-- source src="../assets/movie.ogg" type="video/ogg"-->
      </video>
      <button id="btn1" v-on:click="myClick">点我</button>
      <p> {{resp}} </p>
    </div>
    
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'MyVue',
  props: {
    msg: String,
  },
  data() {
    return {
      resp: ``
    }
  },
  methods: {
    myClick: function () {
      console.log("click");
      axios.post('/my-app/user', {
              FirstName: 'Fred',
              LastName: 'Flintstone'
            })
            .then(response => {
                console.log(response);
                this.resp = response.data;
            })
            .catch(response => {
                console.log(response);
            });
    }
  }
}

</script>


<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
video {
  margin: 10px 10px;
}

.body {
  width: 80%;
  margin-left: 10%;
  display: flex;
  flex-flow: row wrap;
  justify-content: space-between;
}
#btn1{
  flex: none;
}
</style>
