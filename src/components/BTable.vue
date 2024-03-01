<template>
        <b-table :data="tableData" :columns="columns" class="custom-table"></b-table>
</template>


<script>
 export default{
    data() {
        return {
            columns: [
                {
                    field: 'timestamp',
                    label: 'TIMESTAMP',
                    width: '20%',
                    numeric: true,
                    thClass: 'text-left'
                },
                {
                    field: 'ip_details',
                    label: 'IP DETAILS',
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
            ]
        }

    },
    props: ['tableData']
 }
</script>

<style scoped>
  .custom-table {
  width: 100%; 
}
</style>