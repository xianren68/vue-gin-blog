import { defineStore } from "pinia";
import { getCateList } from "@/api";
import { useRouter } from "vue-router";
import {reactive} from 'vue'
const router = useRouter()
// 创建分类仓库

const cateStore = defineStore('category',()=>{
    // 分类列表
    const cateList = reactive({
        list : []
    })
    // 获取方法
    const reqCateList = async (pagesize,pagenum)=>{
        const {data} = await getCateList(pagesize,pagenum)
        if(data.status == 200){
           cateList.list = data.data
        }else if(data.status >= 1004 && data.status < 2000){
            message.error(data.msg+'即将跳转到登录页');
            router.push('./login')
        }else {
            message.error(data.msg)
        }
    
    }
    return {cateList,reqCateList}
},{
    persist:{
        enabled:true
    }
})
export default cateStore
