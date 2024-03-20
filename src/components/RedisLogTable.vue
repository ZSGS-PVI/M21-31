<template>
    <div class="">
        <h1 class="table-heading">Redis Logs</h1>
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
            // primaryKey: 0,
            columns: [
                {
                    field: 'timestamp',
                    label: 'TIMESTAMP',
                    width: '20%',
                    numeric: true,
                    thClass: 'text-left'
                },
                {
                    field: 'addr_details',
                    label: 'ADDR DETAILS',
                    width: '20%',
                    thClass: 'text-center'
                },
                {
                    field: 'command',
                    label: 'COMMAND',
                    width: '20%',
                    thClass: 'text-center'
                },
                {
                    field: 'content',
                    label: 'CONTENT',
                    width: '40%',
                    thClass: 'text-center'
                }
            ],
        };
    },
    methods: {
        addData(jsonObject, tableType) {
            //  console.log(jsonObject)
            if (jsonObject && Object.keys(jsonObject).length > 0) {
                const objDet = {
                    timestamp: jsonObject.timestamp,
                    addr_details: `ip: ${jsonObject.addr_details.ip}, port:${jsonObject.addr_details.port}`,
                    command: jsonObject.command,
                    content: `ID: ${jsonObject.content.id}, DEPT: ${jsonObject.content.dept}, NAME: ${jsonObject.content.name}, SALARY: ${jsonObject.content.salary}`
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
                const jsonArray = await getLogs("redis_logs", offset);
                // console.log(jsonArray);

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
            const pmKey = await getPrimaryKey("redis_logs");
            callback(pmKey);
        }
    },
    created(){
        console.log("redis created called");
    connectWebSocket(this.handleDataReceived, "redis_logs")
   }
};
</script>

<style scoped>
.table-heading {
    font-size: 25px;
    margin-bottom: 30px;
}
</style>