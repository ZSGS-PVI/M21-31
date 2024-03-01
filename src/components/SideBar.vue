<template>
    <div class="sidebar">
        <h1 class="app-heading">
            M21-31
        </h1>
        <div class="link-cont">
            <router-link to="/">REDIS-LOGS</router-link>
            <router-link to="/njinix-logs">NginX-LOGS</router-link>
            <router-link to="/mysql-logs">MySql-LOGS</router-link>
            <router-link to="/vbox-logs">VirtualBox-LOGS</router-link>
        </div>


    </div>
</template>


<script>
</script>


<style>
.sidebar {
    height: 100%;
    width: 15%;
    position: fixed;
    top: 0;
    left: 0;
    background-color: #333;
    padding-top: 20px;
}

.sidebar a {
    padding: 10px 15px;
    text-decoration: none;
    font-size: 18px;
    color: white;
    display: block;
}

.sidebar a:hover {
    background-color: #555;
}

.app-heading {
    font-size: 24px;
    color: hsl(207, 35%, 74%);
    text-align: center;
    margin-top: 20px;
}

.link-cont {
    margin-top: 20px;
}
</style>