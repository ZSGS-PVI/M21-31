<template>
    <div class="container">
        <div class="row">
            <div class="col-md-12">
                <div v-if="activeTab === 1" class="logs-count-para">
                    <p>{{ logsCount }} logs are generated at {{ logsGTime }}, want latest data? press
                        <a @click="getliveData" style="color: blue;">here</a>
                    </p>
                </div>
                <b-tabs ref="tabs" v-model="activeTab">
                    <TabItem label="Live" :data="filteredLiveTableData" :columns="filteredColumns" class="custom-table"
                        @more-info-click="showRowDetails" />
                    <TabItem label="Logs" :data="filteredLogsTableData" :columns="filteredColumns" class="custom-table"
                        @more-info-click="showRowDetails" />
                </b-tabs>
                <div v-if="activeTab === 1" class="button-container">
                    <button @click="prevPage" class="btn btn-primary button">Prev</button>
                    <button @click="nextPage" class="btn btn-primary button">Next</button>
                </div>
            </div>
        </div>
        <el-drawer v-if="drawerVisible" :visible.sync="drawerVisible" direction="rtl" size="40%"
            @close="drawerVisible = false">
            <div style="padding: 20px;">
                <h3>Row Details</h3>
                <div v-if="selectedRow">
                    <p v-for="(value, key) in selectedRow" :key="key">{{ key }}: {{ value }}</p>
                </div>
            </div>
        </el-drawer>
    </div>
</template>

<script>
import TabItem from '../components/TabItem.vue'

export default {
    components: { TabItem },
    props: ['columns', 'liveTableData', 'logsTableData'],
    data() {
        return {
            logsCount: '0',
            logsGTime: '',
            activeTab: 0,
            primaryKey:0,
            offset: 0,
            drawerVisible: false,
            selectedRow: null,
            columnsThreshold: 5,
        };
    },
    computed: {
        filteredColumns() {
            const filteredCols = this.columns.slice(0, this.columnsThreshold);
            if (this.hasMoreInfoColumn(this.columns)) {
                filteredCols.push({ label: 'More Info', field: 'more_info', slotName: 'moreInfoButton' });
            }
            return filteredCols;
        },
        filteredLiveTableData() {
            return this.filterTableData(this.liveTableData);
        },
        filteredLogsTableData() {
            return this.filterTableData(this.logsTableData);
        },
    },
    watch: {
        activeTab() {
            // this.emitTabChange();
            if (this.activeTab === 1) {
                console.log("ac")
                this.assignCurrrentTime();
                this.emitGetKey();
            }
        },
        offset() {
            if (this.activeTab === 1) {
                this.callParentM(); 
            }
        }
    },
    methods: {
        assignCurrrentTime() {
            const curDate = new Date();
            const hours = curDate.getHours();
            const amPm = hours >= 12 ? 'PM' : 'AM';
            const hour12 = hours % 12 || 12;
            const minutes = curDate.getMinutes();
            const seconds = curDate.getSeconds();
            const date = curDate.getDate();
            const month = curDate.getMonth() + 1;
            const year = curDate.getFullYear();
            this.logsGTime = `${date}-${month}-${year} / ${hour12}:${minutes < 10 ? '0' + minutes : minutes}:${seconds < 10 ? '0' + seconds : seconds} ${amPm}`;
        },
        getLogsCount() {
            this.logsCount = this.primaryKey >= 1000  ? Math.round(this.primaryKey / 1000) + "k" : this.primaryKey;
        },
        prevPage() {
            if (this.offset < this.primaryKey) {
                this.offset += 10;
            }
        },
        nextPage() {
            if (this.offset >= 10) {
                this.offset -= 10;
            }
        },
        callParentM() {
            this.emitTabChange();
            this.emitFetchLogsData(this.offset);
            this.getLogsCount();
        },
        getliveData() {
            this.assignCurrrentTime();
            this.emitGetKey();
        },
        emitTabChange() {
            this.$emit('tab-change');
        },
        emitGetKey() {
            this.$emit('get-key', (primaryKey) => {
                this.primaryKey = primaryKey;
                this.offset = this.primaryKey;
            });
        },

        emitFetchLogsData(offset) {
            this.$emit('fetch-logs-data', offset);
        },
        hasMoreInfoColumn(columns) {
            return columns.length > this.columnsThreshold;
        },
        filterTableData(data) {
            return data.map(row => {
                const newRow = {};
                this.columns.slice(0, this.columnsThreshold).forEach(col => {
                    newRow[col.field] = row[col.field];
                });
                return newRow;
            });
        },
        showRowDetails(row) {
            this.selectedRow = row;
            this.drawerVisible = true;
        }
    }
};
</script>

<style scoped>
.button-container {
    position: absolute;
    bottom: 20px;
    right: 20px;
}

.button {
    margin: 10px;
}

.logs-count-para {
    text-align: end;
}
</style>
