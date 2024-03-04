<template>
    <b-tab-item :label="label">
        <b-table :data="slicedTableData" :columns="columns" class="custom-table" ></b-table>
    </b-tab-item>
</template>

<script>
export default {
    name: 'TabItem',
    props: ['columns', 'data', 'label', 'perPage', 'currentPage'],
    computed: {
        slicedTableData() {
            const start = (this.currentPage - 1) * this.perPage;
            return this.data.slice(start, start + this.perPage);
        }
    }
};
</script>
<style>

</style>