<template>
    <a-layout class="layout">
        <a-layout-header>
          <div class="logo" />
          <a-menu class="nav" v-for="nav in routerList"
            v-model:selectedKeys="selectedKeys"
            theme="dark"
            mode="horizontal"
            :style="{ lineHeight: '64px'}"
          >
            <a-menu-item :key="nav.name" @click="turnRouter(nav.path)">{{nav.meta.name}}</a-menu-item>
          </a-menu>
          <div class="out">
            <a-button type="danger" size="small" @click="outLogin">退出</a-button>
          </div>
        </a-layout-header>
        <a-layout-content style="padding: 0 40px">
          <a-breadcrumb style="margin: 16px 0">
            <a-breadcrumb-item>Home</a-breadcrumb-item>
            <a-breadcrumb-item>List</a-breadcrumb-item>
            <a-breadcrumb-item>App</a-breadcrumb-item>
          </a-breadcrumb>
          <div :style="{ background: '#fff', padding: '24px', minHeight: '400px' }">
            <router-view></router-view>
          </div>
        </a-layout-content>
        <a-layout-footer style="text-align: center">
          Ant Design ©2018 Created by Ant UED
        </a-layout-footer>
      </a-layout>
</template>

<script setup>
import router from '@/router';
import { useRouter,useRoute } from 'vue-router';
import {reactive,ref} from 'vue'

// 退出登录
const outLogin = ()=>{
  // 删除token
  window.localStorage.clear()
  // 路由跳转
  router.push('/login')
}
// 获取路由列表，动态决定导航栏
const routerList = reactive(useRouter().getRoutes().filter(ele=>ele.meta.isShow))
// // 导航栏选择
const selectedKeys = ref([useRoute().name])
// 路由跳转
const turnRouter = (path)=>{
  router.push(path)
}
</script>

<style lang="scss" scoped>
    .layout {
        height: 100%;
        background-color: darkcyan;
        .logo {
            float: left;
            margin: 0 ;
            width: 100px;
            height: 50px;
            
        }
        .nav {
            float: left;
        }
        .out {
            float :right;  
        }
    }
</style>