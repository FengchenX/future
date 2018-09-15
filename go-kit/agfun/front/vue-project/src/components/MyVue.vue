<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <video width="320" height="240" controls>
      <source src="../assets/movie.mp4" type="video/mp4">
      <!-- source src="../assets/movie.ogg" type="video/ogg"-->
    </video>
    <button v-on:click="myClick">点我</button>
    <p> {{resp}} </p>
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
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
