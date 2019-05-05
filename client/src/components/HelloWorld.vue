<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <input type="text" v-model="Name">
    <button @click="searchName">kensaku</button>
    {{SearchResult}}
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'HelloWorld',
  data () {
    return {
      msg: 'Welcome to Your Vue.js App',
      Name: 'null',
      SearchResult: []
    }
  },
  methods: {
    async searchName () {
      const resp = await axios.get('http://localhost:1323?name=' + this.Name)
      this.SearchResult = resp.data
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
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
