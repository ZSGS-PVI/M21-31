<template>
    <div class="">
        <h1 class="table-heading">DNS Logs</h1>
        <TableView :liveTableData="liveTableData" :logsTableData="logsTableData" :columns="columns"
            @fetch-logs-data="fetchLogsDataFromChild" @tab-change="logTabChange" @get-key="getPrimKey"/>
    </div>
</template>
  
<script>
import TableView from "./TableView.vue";
import { connectWebSocket, getLogs, getPrimaryKey } from "../websoc";

export default {
    components: {
        TableView
    },
    data() {
        return {
            liveTableData: [],
            logsTableData: [],
            columns: [
                {
                    field: 'timestamp',
                    label: 'TIMESTAMP',
                    width: '20%',
                    numeric: true,
                    thClass: 'text-left'
                },
                {
                    field: 'client_ip',
                    label: 'CLIENT IP',
                    width: '20%',
                    thClass: 'text-center'
                },
                {
                    field: 'query',
                    label: 'QUERY',
                    width: '20%',
                    thClass: 'text-center'
                },
                {
                    field: 'client_port',
                    label: 'CLIENT PORT',
                    width: '40%',
                    thClass: 'text-center'
                }
            ],
        };
    },
    methods: {
        addData(jsonObject, tableType) {
            console.log(jsonObject)
            if (jsonObject && Object.keys(jsonObject).length > 0) {
                const objDet = {
                    timestamp: jsonObject.query,
                    client_ip: jsonObject.client_ip,
                    query: jsonObject.query,
                    client_port: jsonObject.client_port
                };
                if (tableType === 'live') {
                    this.liveTableData.unshift(objDet)
                } else {
                    console.log(objDet)
                    this.logsTableData.push(objDet)
                }

            }
        },


        handleDataReceived(data) {
           // console.log("Received data from WebSocket:", data);
            this.addData(data, "live");
        },

        async fetchLogsDataFromChild(offset) {
            try {
                const jsonArray = await getLogs("dns_logs", offset);
                console.log(jsonArray);

                if (!jsonArray || jsonArray === '') {
                    console.log("Empty JSON string comes.....");
                } else {
                    jsonArray.map(jsonString => this.addData(JSON.parse(jsonString), 'logs'));
                }
            } catch (error) {
                console.error(error);
            }
        },
        logTabChange(){
            this.liveTableData = [];
            this.logsTableData = [];
        },
        async getPrimKey(callback) {
            const pmKey = await getPrimaryKey("dns_logs");
            callback(pmKey);
        }
    },
   created(){
    console.log("called");
    connectWebSocket(this.handleDataReceived, "dns_logs")
   }
};
</script>
  
<style scoped>
.table-heading {
    font-size: 25px;
    margin-bottom: 30px;
}
</style>
  