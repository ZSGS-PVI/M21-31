<template>
        <BTable :tableData="this.data" />
</template>

<script>
import BTable from "./BTable.vue";
import { connectWebSocket, closeWebSocket } from "../websoc";


export default {

    components: {
        BTable
    },
    data() {
        return {
            data: [
            ],
        }

    },
    methods: {
        addData(jsonObject) {

            const objDet = {
                timestamp: jsonObject.timestamp,
                ip_details: jsonObject.ip_details,
                command: jsonObject.command,
                content: `ID: ${jsonObject.content.id}, DEPT: ${jsonObject.content.dept}, NAME: ${jsonObject.content.name}, SALARY: ${jsonObject.content.salary}`
            }

            this.data.push(objDet);
            console.log(jsonObject);
        },
        handleDataReceived(data) {

            console.log("Received data from WebSocket:", data);

            this.addData(data);
        }
    },
    created() {
        connectWebSocket(this.handleDataReceived, "redis-log");
    },
    beforeDestroy() {
        closeWebSocket()
    }
}
</script>


<style scoped>
</style>