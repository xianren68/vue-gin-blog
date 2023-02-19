import { defineStore } from "pinia";
import { getArtList,reqCateArt } from "@/api";
import { useRouter } from "vue-router";
import {reactive} from 'vue'
const router = useRouter()
// 创建分类仓库

const ArtStore = defineStore('art',()=>{
    // 分类列表
    const artList = reactive({
        list : []
    })
    const id = ref(0)
    // 获取方法
    const ArtList = async (param)=>{
        const {data} = await getArtList(param)
        if(data.status == 200){
            artList.list = data.data
        }else {
            alert(data.msg)
        }
    }
    // 获取指定分类的
    const ArtCat = async (id)=>{
        const {data} = await reqCateArt(id)
        if(data.status == 200){
            artList.list = data.data
        }else {
            alert(data.msg)
        }
    }
    return {id,artList,ArtList,ArtCat}
},{
    persist:{
        enabled:true
    }
})
export default ArtStore