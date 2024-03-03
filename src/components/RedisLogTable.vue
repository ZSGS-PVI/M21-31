<template>
    <div class="">
        <h1 class="table-heading">Redis Logs</h1>
        <BTable :tableData="data" @fetch-logs-data="fetchLogsDataFromChild" />
    </div>
</template>
  
<script>
import BTable from "./BTable.vue";
import { connectWebSocket, closeWebSocket, getLogs } from "../websoc";

export default {
    components: {
        BTable
    },
    data() {
        return {
            data: []
        };
    },
    methods: {
        addData(jsonObject) {
            if (jsonObject !== null) {
                const objDet = {
                    timestamp: jsonObject.timestamp,
                    ip_details: jsonObject.ip_details,
                    command: jsonObject.command,
                    content: `ID: ${jsonObject.content.id}, DEPT: ${jsonObject.content.dept}, NAME: ${jsonObject.content.name}, SALARY: ${jsonObject.content.salary}`
                };
                if (this.data) {
                    this.data.unshift(objDet);
                } else {
                    this.data = [objDet]; 
                }
                console.log(jsonObject);
            }
        },
        handleDataReceived(data) {
            console.log("Received data from WebSocket:", data);
            this.addData(data);
        },
        fetchLogsDataFromChild() {
            this.data = getLogs("redis_logs");
        }
    },
    created() {
        connectWebSocket(this.handleDataReceived, "redis_logs");
    },
    beforeDestroy() {
        closeWebSocket();
    }
};
</script>
  
<style scoped>
.table-heading {
    color: rgba(115, 42, 224, 0.654);
    font-size: 50px;
    margin-bottom: 30px;
}
</style>
  