<template>
    <b-tab-item :label="label">
        <b-table  :data="data" :columns="columns" class="custom-table">
            <!-- <b-table-column v-for="column in columns" :field="column.field" :label="column.label" :width="column.width"
                :key="column.name">
                    <template v-if="column.field === 'more_info'">
                        <button @click="handleMoreInfoClick(data.item)">...</button>
                    </template>
                    <template v-else> 
                        {{  row[column.field] }}
                    </template>
            </b-table-column> -->
        </b-table>
    </b-tab-item>
</template>

<script>
export default {
    name: 'TabItem',
    props: ['columns', 'data', 'label'],

    watch:{
        data(){
            //console.log(this.data);
        }
    },

    methods: {
        handleMoreInfoClick(row) {
            // Emit an event with the clicked row data
            this.$emit('more-info-click', row);
        }
    }

};
</script>
<style></style>