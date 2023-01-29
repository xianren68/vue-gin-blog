import { createRouter ,createWebHashHistory} from "vue-router";
// 路由列表
const routes = [
    {
        path:'/',
        name:"admin",
        component:()=>import("@/views/admin.vue")
    }
    ,
    {
        path:'/login',
        name:"login",
        component:()=>import("@/views/login.vue") 
    }
]
// 创建路由
const router = createRouter({
    history:createWebHashHistory(),
    routes,
})

// 将路由对象暴露出去
export default router