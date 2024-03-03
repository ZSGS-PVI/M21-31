<template>
    <MysqlTable :tableData="data" />
  </template>
  
  <script>
  import MysqlTable from "./MysqlTable.vue";
  import { connectWebSocket, closeWebSocket } from "../websoc";
  
  export default {
    components: {
      MysqlTable,
    },
    data() {
      return {
        data: [],
      };
    },
    methods: {
      addData(jsonObject) {
        const objDet = {
          timestamp: jsonObject.timestamp,
          connection_id: jsonObject.connection_id,
          query_type: jsonObject.query_type,
          action: jsonObject.action,
        };
  
        this.data.push(objDet);
      },
      handleDataReceived(data) {
        console.log("Received data from WebSocket:", data);
        this.addData(data);
      },
    },
    created() {
      connectWebSocket(this.handleDataReceived, "mysql_logs");
    },
    beforeDestroy() {
      closeWebSocket();
    },
  };
  </script>
  
  <style scoped>
  </style>