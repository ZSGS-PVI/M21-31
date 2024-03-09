<template>
    <b-tab-item :label="label">
        <b-table :data="data" :columns="columns" class="custom-table" ></b-table>
    </b-tab-item>
</template>

<script>
export default {
    name: 'TabItem',
    props: ['columns', 'data', 'label'],
    
};
</script>
<style>

</style>