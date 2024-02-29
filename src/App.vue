<template>
  <div id="app">
    <SideBar />
    <router-view/>
  </div>
</template>

<script>
import SideBar from '@/components/SideBar.vue'
export default {
  components:{
    SideBar
  },
  name: 'App',
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
