<template>
    <div>
      <b-pagination v-model="currentPage" :total="totalRows" :per-page="perPage"></b-pagination>
      <b-tabs ref="tabs" v-model="activeTab" @input="handleTabChange">
        <b-tab-item label="Live">
          <b-table :data="slicedTableData" :columns="columns" class="custom-table"></b-table>
        </b-tab-item>
        <b-tab-item label="Logs">
          <b-table :data="slicedTableData" :columns="columns" class="custom-table"></b-table>
        </b-tab-item>
      </b-tabs>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        activeTab: 0,
        perPage: 10,
        currentPage: 1,
        totalRows: 0,
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
      };
    },
    props: ['tableData'],
    computed: {
      slicedTableData() {
        const start = (this.currentPage - 1) * this.perPage;
        return this.tableData.slice(start, start + this.perPage);
      }
    },
    watch: {
      activeTab() {
        this.currentPage = 1;
      },
      tableData(newVal) {
        this.totalRows = newVal.length; 
      }
    },
    methods: {
      handleTabChange(tab) {
        this.activeTab = tab;
      }
    }
  };
  </script>
  
  <style scoped>
  .custom-table th {
    text-align: center !important;
  }
  </style>
  