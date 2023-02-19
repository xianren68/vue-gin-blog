import { defineStore } from "pinia";
import { getCateList } from "@/api";
import { useRouter } from "vue-router";
import {reactive} from 'vue'
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
        }else {
            alert(data.msg)
        }
    
    }
    return {cateList,reqCateList}
},{
    persist:{
        enabled:true
    }
})
export default cateStore