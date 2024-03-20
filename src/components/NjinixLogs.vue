<template>
    <div class="">
        <h1 class="table-heading">Nginx Logs</h1>
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
                    field: 'client_ip',
                    label: 'CLIENT IP',
                    thClass: 'text-left'
                },
                {
                    field: 'remote_user',
                    label: 'REMOTE USER',
                    thClass: 'text-left'
                },
                {
                    field: 'date_time',
                    label: 'DATE & TIME',
                    thClass: 'text-left'
                },
                {
                    field: 'request',
                    label: 'REQUEST',
                    thClass: 'text-left'
                },
                {
                    field: 'target_status_code',
                    label: 'TARGET STATUS CODE',
                    numeric: true,
                    thClass: 'text-left'
                },
                {
                    field: 'body_bytes_sent',
                    label: 'BODY BYTES SENT',
                    numeric: true,
                    thClass: 'text-left'
                },
                {
                    field: 'referer',
                    label: 'REFERER',
                    thClass: 'text-left'
                },
                {
                    field: 'user_agent',
                    label: 'USER AGENT',
                    thClass: 'text-left'
                },
                {
                    field: 'ssl_protocol',
                    label: 'SSL PROTOCOL',
                    thClass: 'text-left'
                },
                {
                    field: 'ssl_cipher',
                    label: 'SSL CIPHER',
                    thClass: 'text-left'
                },
                {
                    field: 'request_processingtime',
                    label: 'REQUEST PROCESSING TIME',
                    numeric: true,
                    thClass: 'text-left'
                },
                {
                    field: 'target_processingtime',
                    label: 'TARGET PROCESSING TIME',
                    numeric: true,
                    thClass: 'text-left'
                },
                {
                    field: 'received_bytes',
                    label: 'RECEIVED BYTES',
                    numeric: true,
                    thClass: 'text-left'
                },
                {
                    field: 'target_ip_port',
                    label: 'TARGET IP:PORT',
                    thClass: 'text-left'
                },
                {
                    field: 'statuscode_from_loadbalancer',
                    label: 'STATUS CODE FROM LOAD BALANCER',
                    numeric: true,
                    thClass: 'text-left'
                },
                {
                    field: 'bytes_sent',
                    label: 'BYTES SENT',
                    numeric: true,
                    thClass: 'text-left'
                },
                {
                    field: 'sni_domain',
                    label: 'SNI DOMAIN',
                    thClass: 'text-left'
                },
                {
                    field: 'host',
                    label: 'HOST',
                    thClass: 'text-left'
                },
                {
                    field: 'uri',
                    label: 'URI',
                    thClass: 'text-left'
                },
                {
                    field: 'http_upgrade',
                    label: 'HTTP UPGRADE',
                    thClass: 'text-left'
                }
            ]
        };
    },
    methods: {
        addData(jsonObject, tableType) {
            console.log(jsonObject)
            if (jsonObject && Object.keys(jsonObject).length > 0) {
                const objDet = {
                    client_ip: jsonObject.client_ip,
                    remote_user: jsonObject.remote_user,
                    date_time: jsonObject.date_time,
                    request: jsonObject.request,
                    target_status_code: jsonObject.target_status_code,
                    body_bytes_sent: jsonObject.body_bytes_sent,
                    referer: jsonObject.referer,
                    user_agent: jsonObject.user_agent,
                    ssl_protocol: jsonObject.ssl_protocol,
                    ssl_cipher: jsonObject.ssl_cipher,
                    request_processingtime: jsonObject.request_processingtime,
                    target_processingtime: jsonObject.target_processingtime,
                    received_bytes: jsonObject.received_bytes,
                    target_ip_port: jsonObject.target_ip_port,
                    statuscode_from_loadbalancer: jsonObject.statuscode_from_loadbalancer,
                    bytes_sent: jsonObject.bytes_sent,
                    sni_domain: jsonObject.sni_domain,
                    host: jsonObject.host,
                    uri: jsonObject.uri,
                    http_upgrade: jsonObject.http_upgrade,

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
                const jsonArray = await getLogs("nginx_logs", offset);
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
            const pmKey = await getPrimaryKey("nginx_logs");
            callback(pmKey);
        }
    },
    created(){
        console.log("nginx created called");
    connectWebSocket(this.handleDataReceived, "nginx_logs")
   }
};
</script>

<style scoped>
.table-heading {
    font-size: 25px;
    margin-bottom: 30px;
}
</style>