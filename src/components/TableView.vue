<template>
    <div>
        <b-tabs ref="tabs" v-model="activeTab" @input="handleTabChange">
            <TabItem label="Live" :data="liveTableData" :columns="columns" :perPage="perPage" :currentPage="currentPage" :totalRows="liveTotalRows" class="custom-table" />
            <TabItem label="Logs" :data="logsTableData" :columns="columns" :perPage="perPage" :currentPage="currentPage" :totalRows="logsTotalRows" class="custom-table" />
        </b-tabs>
        <b-pagination v-model="currentPage" :total="currentTotalRows" :per-page="perPage"></b-pagination>
    </div>
</template>

<script>
import TabItem from '../components/TabItem.vue'

export default {
    components: {
        TabItem
    },
    data() {
        return {
            activeTab: 0,
            perPage: 10,
            currentPage: 1,
            liveTotalRows: 0,
            logsTotalRows: 0,
        };
    },
    props: ['liveTableData', 'logsTableData', 'columns'],
    computed: {
        currentTotalRows() {
            return this.activeTab === 0 ? this.liveTotalRows : this.logsTotalRows;
        }
    },
    watch: {
        liveTableData(newVal) {
            if (newVal) {
                this.liveTotalRows = newVal.length;
            }
        },
        logsTableData(newVal) {
            if (newVal) {
                this.logsTotalRows = newVal.length;
            }
        },
        activeTab() {
            this.currentPage = 1;
        }
    },
    methods: {
        async handleTabChange() {
            this.$emit('tab-change')
            console.log("handle ")
            if (this.activeTab === 1 ) {
                this.$emit('fetch-logs-data');
            }
        }
    }
};
</script>

<style>

</style>