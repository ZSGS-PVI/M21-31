<template>
    <b-table :data="tableData" :columns="columns" class="custom-table"></b-table>
  </template>
  
  <script>
  export default {
    props: ['tableData'],
    data() {
      return {
        columns: [
          {
            field: "timestamp",
            label: "Timestamp".toUpperCase(),
            width: "20%",
            numeric: true,
            thClass: "text-left",
          },
          // {
          //   field: "type",
          //   label: "Type".toUpperCase(),
          //   width: "20%",
          //   thClass: "text-center",
          // },
          {
            field: "connection_id",
            label: "Connection ID".toUpperCase(), // Capitalize and modify label
            width: "20%",
            thClass: "text-center",
          },
          {
            field: "query_type",
            label: "Query Type".toUpperCase(),
            width: "40%",
            thClass: "text-center",
          },
  
          {
            field: "action",
            label: "Action".toUpperCase(),
            width: "20%",
            thClass: "text-center",
          },
        ],
      };
    },
  };
  </script>
  
  <style scoped>
  .custom-table {
    width: 100%;
  }
  </style>