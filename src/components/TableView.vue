<template>
    <div class="container">
        <div class="row">
            <div class="col-md-12">
                <b-tabs ref="tabs" v-model="activeTab">
                    <TabItem label="Live" :data="liveTableData" :columns="columns" class="custom-table" />
                    <TabItem label="Logs" :data="logsTableData" :columns="columns" class="custom-table" />
                </b-tabs>
                <div v-if="activeTab === 1" class="button-container">
                    <button @click="prevPage" class="btn btn-primary button">Prev</button>
                    <button @click="nextPage" class="btn btn-primary button">Next</button>
                </div>
            </div>
        </div>
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
            offset:0,
        };
    },
    props: ['columns', 'liveTableData', 'logsTableData'],
    watch: {
        activeTab() {
            this.$emit('tab-change');
            if (this.activeTab === 1 ) {
                this.$emit('fetch-logs-data', this.offset);
            }
        }
    },
    methods: {
        prevPage() {
            if(this.offset>10){
                this.offset -= 10;
                this.callParentM();
            }

        },
        nextPage() {
            this.offset += 10;
            this.callParentM();
        },
        callParentM(){
            this.$emit('tab-change');
            this.$emit('fetch-logs-data', this.offset);
        }
    }
};
</script>

<style>
.button-container {
    position: absolute;
    bottom: 20px;
    right: 20px;
}
.button{
    margin: 10px;
}
</style>
