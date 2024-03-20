<template>
    <div class="">
        <h1 class="table-heading">Kvm Logs</h1>
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
                    width: '10%',
                    numeric: true,
                    thClass: 'text-left'
                },
                {
                    field: 'hostname',
                    label: 'HOSTNAME',
                    width: '10%',
                    thClass: 'text-center'
                },
                {
                    field: 'kernel',
                    label: 'KERNEL',
                    width: '10%',
                    thClass: 'text-center'
                },
                {
                    field: 'auditType',
                    label: 'AUDIT TYPE',
                    width: '10%',
                    thClass: 'text-center'
                },
                {
                    field: 'operation',
                    label: 'OPERATION',
                    width: '10%',
                    thClass: 'text-center'
                },
                {
                    field: 'profile',
                    label: 'PROFILE',
                    width: '10%',
                    thClass: 'text-center'
                },
                {
                    field: 'unconfined',
                    label: 'UNCONFINED',
                    width: '10%',
                    thClass: 'text-center'
                },
                {
                    field: 'pid',
                    label: 'PID',
                    width: '5%',
                    thClass: 'text-center'
                },
                {
                    field: 'comm',
                    label: 'COMM',
                    width: '10%',
                    thClass: 'text-center'
                },
                {
                    field: 'capability',
                    label: 'CAPABILITY',
                    width: '10%',
                    thClass: 'text-center'
                },
                {
                    field: 'capname',
                    label: 'CAPNAME',
                    width: '10%',
                    thClass: 'text-center'
                },
                {
                    field: 'auditEpoch',
                    label: 'AUDIT EPOCH',
                    width: '10%',
                    thClass: 'text-center'
                },
                {
                    field: 'timestampEpoch',
                    label: 'TIMESTAMP EPOCH',
                    width: '10%',
                    thClass: 'text-center'
                }
            ],
        };
    },
    methods: {
        addData(jsonObject, tableType) {
            console.log(jsonObject);
            if (jsonObject && Object.keys(jsonObject).length > 0) {
                const objDet = {
                    timestamp: jsonObject.timestamp,
                    hostname: jsonObject.hostname,
                    kernel: jsonObject.kernel,
                    auditType: jsonObject.audit_type,
                    operation: jsonObject.operation,
                    profile: jsonObject.profile,
                    unconfined: jsonObject.unconfined,
                    pid: jsonObject.pid,
                    comm: jsonObject.comm,
                    capability: jsonObject.capability,
                    capname: jsonObject.capname,
                    auditEpoch: jsonObject.audit_epoch,
                    timestampEpoch: jsonObject.timestamp_epoch
                };
                if (tableType === 'live') {
                    if (this.liveTableData.length >= 10) {
                        this.liveTableData.pop();
                    }
                    this.liveTableData.unshift(objDet)
                } else {
                    this.logsTableData.push(objDet)
                }
            }
        },

        handleDataReceived(data) {
            this.addData(data, "live");
        },

        async fetchLogsDataFromChild(offset) {
            try {
                const jsonArray = await getLogs("kvm_logs", offset);
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
            const pmKey = await getPrimaryKey("kvm_logs");
            callback(pmKey);
        }
    },
    created(){
        console.log("kvm created called");
    connectWebSocket(this.handleDataReceived, "kvm_logs")
   }
};
</script>

<style scoped>
.table-heading {
    font-size: 25px;
    margin-bottom: 30px;
}
</style>