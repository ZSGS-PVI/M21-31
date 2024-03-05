<template>
  <div class="">
    <h1 class="table-heading">Mysql Logs</h1>
    <TableView :liveTableData="liveTableData" :logsTableData="logsTableData" :columns="columns"
      @fetch-logs-data="fetchLogsDataFromChild" @tab-change="logTabChange" />
  </div>
</template>
  
<script>
import TableView from "./TableView.vue";
import { connectWebSocket, closeWebSocket, getLogs } from "../websoc";

export default {
  components: {
    TableView,
  },
  data() {
    return {
      logsTableData: [],
      liveTableData: [],
      columns: [
        {
          field: "timestamp",
          label: "TIMESTAMP",
          width: "20%",
          numeric: true,
          thClass: "text-left",
        },
        {
          field: "connection_id",
          label: "CONNECTION ID",
          width: "20%",
          thClass: "text-center",
        },
        {
          field: "query_type",
          label: "QUERY TYPE",
          width: "20%",
          thClass: "text-center",
        },
        {
          field: "action",
          label: "ACTION",
          width: "40%",
          thClass: "text-center",
        },
      ],
    };
  },
  methods: {
    addData(jsonObject, tableType) {
      if (jsonObject && Object.keys(jsonObject).length > 0) {
        const objDet = {
          timestamp: jsonObject.timestamp,
          connection_id: jsonObject.connection_id,
          query_type: jsonObject.query_type,
          action: jsonObject.action,
        };
        if (tableType === 'live') {
          this.liveTableData.unshift(objDet)
        } else {
          console.log(objDet)
          this.logsTableData.unshift(objDet)
        }
      }
    },

    handleDataReceived(data) {
      console.log("Received data from WebSocket:", data);
      this.addData(data, "live");
    },

    async fetchLogsDataFromChild() {
      try {
        const jsonArray = await getLogs("mysql_logs"); // Adjust the log type here
        console.log(jsonArray);

        if (!jsonArray || jsonArray === "") {
          console.log("Empty JSON string comes.....");
        } else {
          jsonArray.map((jsonString) => this.addData(JSON.parse(jsonString), "logs"));
        }
      } catch (error) {
        console.error(error);
      }
    },
    logTabChange() {
      this.liveTableData = [];
      this.logsTableData = [];
    }
  },
  created() {
    connectWebSocket(this.handleDataReceived, "mysql_logs"); // Adjust the log type here
  },
  beforeDestroy() {
    closeWebSocket();
  },
};
</script>
  
<style scoped>
.table-heading {
  font-size: 25px;
  margin-bottom: 30px;
}
</style>