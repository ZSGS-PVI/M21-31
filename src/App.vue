<template>
 <div class="container-fluid table-cont">
  <div class="row">
    <div class="col-lg-2 redis-log-sidebar">
        <SideBar />
    </div>
    <div class="col-lg-10 redis-log-main-content">
        <router-view />
    </div>
  </div>
</div>

</template>

<script>
import SideBar from './components/SideBar.vue'
import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap/dist/js/bootstrap.bundle.min.js';

export default {

  name: 'App',
  components: {
    SideBar
  }
}
</script>

<style >
body{
  margin: 0;
}

#Vapp {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  height: 100vh;
}

.redis-log-sidebar {
  height: 100%; 
  padding: 20px; 
}

.redis-log-main-content {
  height: 100%; 
  padding: 20px;
}

.app-heading {
  font-size: 24px;
  color: hsl(207, 35%, 74%);
  text-align: center;
}

.link-cont {
  margin-top: 20px;
}

.container-fluid {
  height: 100%; /* Ensure container takes full height of its parent */
}

.table-cont {
  background-color: hsl(207, 35%, 74%);
  padding: 20px;
  min-height: 100vh; 
}

.col-lg-3, .col-lg-9 {
  height: 100%; /* Ensure columns take full height of their parent */
}
</style>

