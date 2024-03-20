<template>
    <div class="">
        <h1 class="table-heading">Docker Logs</h1>
        <TableView :liveTableData="liveTableData" :logsTableData="logsTableData" :columns="columns"
            @fetch-logs-data="fetchLogsDataFromChild" @tab-change="logTabChange" @get-key="getPrimKey" />
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
                    field: 'ip',
                    label: 'IP',
                    width: '20%',
                    thClass: 'text-center'
                },
                {
                    field: 'method',
                    label: 'METHOD',
                    width: '20%',
                    thClass: 'text-center'
                },
                {
                    field: 'endpoint',
                    label: 'ENDPOINT',
                    width: '20%',
                    thClass: 'text-center'
                },
                {
                    field: 'status_code',
                    label: 'STATUS CODE',
                    width: '10%',
                    thClass: 'text-center'
                },
                {
                    field: 'user_agent',
                    label: 'USER AGENT',
                    width: '10%',
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
                    timestamp: jsonObject.timestamp,
                    ip: jsonObject.ip,
                    method: jsonObject.method,
                    endpoint: jsonObject.endpoint,
                    status_code: jsonObject.status_code,
                    user_agent: jsonObject.user_agent
                };
                if (tableType === 'live') {
                    if (this.liveTableData.length >= 10) {
                        this.liveTableData.pop();
                    }
                    this.liveTableData.unshift(objDet)
                    //console.log(this.liveTableData);
                } else {
                    // console.log(objDet)
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
                const jsonArray = await getLogs("docker_logs", offset);
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
        logTabChange() {
            this.liveTableData = [];
            this.logsTableData = [];
        },
        async getPrimKey(callback) {
            const pmKey = await getPrimaryKey("docker_logs");
            callback(pmKey);
        }
    },
    created(){
        console.log("docker created called");
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