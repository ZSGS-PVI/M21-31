<template>
    <div class="custom-table">
        <b-table :data="this.tableData" :columns="columns"></b-table>
    </div>
     
</template>


<script>
 export default{
    data() {
        return {
            columns: [
                {
                    field: 'timestamp',
                    label: 'TIMESTAMP',
                    width: '200',
                    numeric: true
                },
                {
                    field: 'ip_details',
                    label: 'IP DETAILS',
                    width: '200'
                },
                {
                    field: 'command',
                    label: 'COMMAND',
                    width: '200'
                },
                {
                    field: 'content',
                    label: 'CONTENT',
                    width: '500',
                    center: true
                }
            ]
        }

    },
    props: ['tableData']
 }
</script>

<style scoped>
.custom-table{
    background-color:  hsl(207, 35%, 74%);
    color: red;
}

</style>