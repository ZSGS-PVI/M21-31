<template>
    <div>
      <el-table :data="visibleData" style="width: 100%">
        <el-table-column v-for="(field, index) in visibleFields" :key="index" :prop="field.key" :label="field.label"></el-table-column>
        <el-table-column label="MORE">
          <template slot-scope="scope">
            <el-button @click="showIndividualData(scope.row)" type="text" icon="el-icon-more"></el-button>
          </template>
        </el-table-column>
      </el-table>
  
      <el-drawer :visible="drawerVisible" direction="rtl" size="40%" @close="drawerVisible = false">
        <el-table :data="drawerData" style="width: 100%">
          <el-table-column prop="key" label="Element"></el-table-column>
          <el-table-column prop="value" label="Value"></el-table-column>
        </el-table>
      </el-drawer>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        visibleFields: [
          { key: "CLIENT IP", label: "CLIENT IP" },
          { key: "REMOTE USER", label: "REMOTE USER" },
          { key: "DATE & TIME", label: "DATE & TIME" },
          { key: "REQUEST", label: "REQUEST" },
          { key: "TARGET STATUS CODE", label: "TARGET STATUS CODE" },
        ],
        allFields: [
          {"CLIENT IP":"127.0.0.1","REMOTE USER":"-","DATE & TIME":"[01/Mar/2024:16:55:42 +0530]","REQUEST":"GET / HTTP/1.1","TARGET STATUS CODE":304,"BODY BYTES SENT":0,"REFERER":"-","USER AGENT":"Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Mobile Safari/537.36","SSL PROTOCOL":"-","SSL CIPHER":"-","REQUEST PROCESSING TIME":"0.002","TARGET PROCESSING TIME":"0.004","RECEIVED BYTES":802,"TARGET IP:PORT":"127.0.0.1:8080","STATUS CODE FROM LOAD BALANCER":304,"BYTES SENT":138,"SNI DOMAIN":"localhost","HOST":"localhost","URI":"/","HTTP UPGRADE":"-"},
          {"CLIENT IP":"127.0.0.1","REMOTE USER":"-","DATE & TIME":"[01/Mar/2024:16:55:42 +0530]","REQUEST":"GET / HTTP/1.1","TARGET STATUS CODE":304,"BODY BYTES SENT":0,"REFERER":"-","USER AGENT":"Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Mobile Safari/537.36","SSL PROTOCOL":"-","SSL CIPHER":"-","REQUEST PROCESSING TIME":"0.002","TARGET PROCESSING TIME":"0.004","RECEIVED BYTES":802,"TARGET IP:PORT":"127.0.0.1:8080","STATUS CODE FROM LOAD BALANCER":304,"BYTES SENT":138,"SNI DOMAIN":"localhost","HOST":"localhost","URI":"/","HTTP UPGRADE":"-"},
          {"CLIENT IP":"127.0.0.1","REMOTE USER":"-","DATE & TIME":"[01/Mar/2024:16:55:42 +0530]","REQUEST":"GET / HTTP/1.1","TARGET STATUS CODE":304,"BODY BYTES SENT":0,"REFERER":"-","USER AGENT":"Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Mobile Safari/537.36","SSL PROTOCOL":"-","SSL CIPHER":"-","REQUEST PROCESSING TIME":"0.002","TARGET PROCESSING TIME":"0.004","RECEIVED BYTES":802,"TARGET IP:PORT":"127.0.0.1:8080","STATUS CODE FROM LOAD BALANCER":304,"BYTES SENT":138,"SNI DOMAIN":"localhost","HOST":"localhost","URI":"/","HTTP UPGRADE":"-"},
          {"CLIENT IP":"127.0.0.1","REMOTE USER":"-","DATE & TIME":"[01/Mar/2024:16:55:42 +0530]","REQUEST":"GET / HTTP/1.1","TARGET STATUS CODE":304,"BODY BYTES SENT":0,"REFERER":"-","USER AGENT":"Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Mobile Safari/537.36","SSL PROTOCOL":"-","SSL CIPHER":"-","REQUEST PROCESSING TIME":"0.002","TARGET PROCESSING TIME":"0.004","RECEIVED BYTES":802,"TARGET IP:PORT":"127.0.0.1:8080","STATUS CODE FROM LOAD BALANCER":304,"BYTES SENT":138,"SNI DOMAIN":"localhost","HOST":"localhost","URI":"/","HTTP UPGRADE":"-"},
          {"CLIENT IP":"127.0.0.1","REMOTE USER":"-","DATE & TIME":"[01/Mar/2024:16:55:42 +0530]","REQUEST":"GET / HTTP/1.1","TARGET STATUS CODE":304,"BODY BYTES SENT":0,"REFERER":"-","USER AGENT":"Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Mobile Safari/537.36","SSL PROTOCOL":"-","SSL CIPHER":"-","REQUEST PROCESSING TIME":"0.002","TARGET PROCESSING TIME":"0.004","RECEIVED BYTES":802,"TARGET IP:PORT":"127.0.0.1:8080","STATUS CODE FROM LOAD BALANCER":304,"BYTES SENT":138,"SNI DOMAIN":"localhost","HOST":"localhost","URI":"/","HTTP UPGRADE":"-"},
          {"CLIENT IP":"127.0.0.1","REMOTE USER":"-","DATE & TIME":"[01/Mar/2024:16:55:42 +0530]","REQUEST":"GET / HTTP/1.1","TARGET STATUS CODE":304,"BODY BYTES SENT":0,"REFERER":"-","USER AGENT":"Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Mobile Safari/537.36","SSL PROTOCOL":"-","SSL CIPHER":"-","REQUEST PROCESSING TIME":"0.002","TARGET PROCESSING TIME":"0.004","RECEIVED BYTES":802,"TARGET IP:PORT":"127.0.0.1:8080","STATUS CODE FROM LOAD BALANCER":304,"BYTES SENT":138,"SNI DOMAIN":"localhost","HOST":"localhost","URI":"/","HTTP UPGRADE":"-"},
          {"CLIENT IP":"127.0.0.1","REMOTE USER":"-","DATE & TIME":"[01/Mar/2024:16:55:42 +0530]","REQUEST":"GET / HTTP/1.1","TARGET STATUS CODE":304,"BODY BYTES SENT":0,"REFERER":"-","USER AGENT":"Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Mobile Safari/537.36","SSL PROTOCOL":"-","SSL CIPHER":"-","REQUEST PROCESSING TIME":"0.002","TARGET PROCESSING TIME":"0.004","RECEIVED BYTES":802,"TARGET IP:PORT":"127.0.0.1:8080","STATUS CODE FROM LOAD BALANCER":304,"BYTES SENT":138,"SNI DOMAIN":"localhost","HOST":"localhost","URI":"/","HTTP UPGRADE":"-"},

            {"CLIENT IP":"127.0.0.1","REMOTE USER":"-","DATE & TIME":"[01/Mar/2024:16:55:42 +0530]","REQUEST":"GET / HTTP/1.1","TARGET STATUS CODE":304,"BODY BYTES SENT":0,"REFERER":"-","USER AGENT":"Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Mobile Safari/537.36","SSL PROTOCOL":"-","SSL CIPHER":"-","REQUEST PROCESSING TIME":"0.002","TARGET PROCESSING TIME":"0.004","RECEIVED BYTES":802,"TARGET IP:PORT":"127.0.0.1:8080","STATUS CODE FROM LOAD BALANCER":304,"BYTES SENT":138,"SNI DOMAIN":"localhost","HOST":"localhost","URI":"/","HTTP UPGRADE":"-"}
          // Add more data objects as needed
        ],
        visibleData: [],
        drawerVisible: false,
        drawerData: [],
      };
    },
    created() {
      this.visibleData = this.allFields; 
    },
    methods: {
      showIndividualData(row) {
        this.drawerData = Object.entries(row).map(([key, value]) => ({ key, value }));
        this.drawerVisible = true;
      },
    },
  };
  </script>
  
  <style>
  /* Add your custom styles here */
  </style>