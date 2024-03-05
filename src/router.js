import Vue from 'vue'
import VueRouter from 'vue-router'
import RedisLogTable from '@/components/RedisLogTable'
import MysqlLogs from '@/components/MysqlLogs'
import NjinixLogs from '@/components/NjinixLogs'
import VirtualBoxLogs from '@/components/VirtualBoxLogs'
import DNSLogs from "@/components/DNSLogs"


Vue.use(VueRouter)

export default new VueRouter({
  routes: [
    {
      path: '/',
      name: 'redis-logs',
      component: RedisLogTable
    },
    {
      path: '/njinix-logs',
      name: 'NjinixLogs',
      component: NjinixLogs
    },
    {
        path: '/mysql-logs',
        name: 'MySqlLogs',
        component: MysqlLogs
      },
      {
        path: '/vbox-logs',
        name: 'VirtualBoxLogs',
        component: VirtualBoxLogs
      },
      {
        path: '/dns-logs',
        name: 'DNSLogs',
        component: DNSLogs
      },
  ]
})
