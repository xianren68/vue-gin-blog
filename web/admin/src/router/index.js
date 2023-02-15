import { createRouter ,createWebHashHistory} from "vue-router";
import { message } from 'ant-design-vue';
import 'ant-design-vue/es/message/style/css'
// 路由列表
const routes = [
    {
        path:'/',
        name:"admin",
        component:()=>import("@/views/Admin/index.vue"),
        children:[
            {
                path:'/',
                name:'home',
                component:()=>import("@/views/Admin/home/index.vue"),
                meta:{
                    name:'仪表盘',
                    isShow:true
                }
            },
            {
                path:'/write/:id',
                name:'write',
                component:()=>import("@/views/Admin/article/write.vue"),
                meta:{
                    name:'写文章',
                    isShow:true
                },
            },
            {
                path:'/artlist',
                name:'artlist',
                component:()=>import("@/views/Admin/article/artlist.vue"),
                meta:{
                    name:'文章列表',
                    isShow:true
                }
            },
            {
                path:'/category',
                name:'category',
                component:()=>import("@/views/Admin/category/index.vue"),
                meta:{
                    name:'分类',
                    isShow:true
                }
            },
            {
                path:'/user',
                name:'user',
                component:()=>import("@/views/Admin/user/index.vue"),
                meta:{
                    name:'用户列表',
                    isShow:true
                }
            },


        ]
    }
    ,
    {
        path:'/login',
        name:"login",
        component:()=>import("@/views/login.vue") 
    },

]
// 创建路由
const router = createRouter({
    history:createWebHashHistory(),
    routes,
})
// 设置路由守卫
router.beforeEach((to,from)=>{
    // 获取token
    let token = window.localStorage.getItem("token")
    if(to.path != "/login") {
        if (!token) {
            message.error("验证失效，即将跳转到登录页面")
            alert("验证失效，即将跳转到登录页面")
            router.push('/login')
        }
    }
})


// 将路由对象暴露出去
export default router