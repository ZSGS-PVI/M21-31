import Vue from 'vue'
import VueRouter from 'vue-router'
import RedisLogTable from '@/components/RedisLogTable'
import MysqlLogs from '@/components/MysqlLogs'
import NjinixLogs from '@/components/NjinixLogs'
import KVMLogs from '@/components/KVMLogs.vue'
import DNSLogs from "@/components/DNSLogs"
import DockerLogs from "@/components/DockerLogs.vue"

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
        path: '/kvm-logs',
        name: 'VirtualBoxLogs',
        component: KVMLogs
      },
      {
        path: '/dns-logs',
        name: 'DNSLogs',
        component: DNSLogs
      },
      {
        path: '/docker-logs',
        name: 'DockerLogs',
        component: DockerLogs
      },
  ]
})
