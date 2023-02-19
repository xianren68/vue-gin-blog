import { createRouter ,createWebHashHistory} from "vue-router";
// 路由列表
const routes = [
    {
        name:'login',
        path:'/login',
        component:()=>import('@/views/login.vue')
    },
    {
        name:'home',
        path:'/',
        component:()=>import('@/views/Home.vue'),
        children:[
            {
                name:'artlist',
                path:'/',
                component:()=>import('@/views/artlist.vue')

            },
            {
                name:'detail',
                path:'detail/:id',
                component:()=>import('@/views/detail.vue')
            }
        ]

    }
]
// 创建路由对象
const router = createRouter({
    history:createWebHashHistory(),
    routes
})
export default router