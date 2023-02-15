<template>
    <a-card title="文章列表" :bordered="false"  style="height:100">
        <a-input-search v-model:value="param" placeholder="输入文章名搜索" style="width: 200px" @search="searchList" 
            allow-clear />
        <a-button type="primary">
            添加
        </a-button>
        <a-select v-model:value="cateValue" :options="cateList.list.map(item=>({value:item.id,label:item.name}))" placeholder="请选择分类"
        @change="handleChange" allowClear>
           
        </a-select>
        <a-table :dataSource="artList.list" :columns="columns" :pagination="pagination" @change="changePage" scroll>
            <template #bodyCell="{ column, text, record }">
                <template v-if="column.dataIndex == 'img'">
                    <img :src="text" alt="">
                </template>
                <template v-if="column.dataIndex == 'action'">
                        <a-button type="primary" size="small" ghost @click="edit(record.ID)">编辑</a-button>
                        <a-popconfirm title="确定删除吗？" @confirm="deleteart(record.ID)">
                            <template #icon><question-circle-outlined style="color: red" /></template>
                            <a-button type="danger" size="small" ghost>删除</a-button>
                        </a-popconfirm>

                </template>
            </template>
        </a-table>


    </a-card>
</template>

<script setup>
import { reactive, onMounted, ref } from 'vue'
import { message } from 'ant-design-vue';
import 'ant-design-vue/es/message/style/css'
import {getArtList,deleteArt,reqCateArt} from '@/api'
import { useRouter } from 'vue-router';
import {useCateStore} from '@/store'
import { storeToRefs } from 'pinia';
const store = useCateStore()
const {cateList} = storeToRefs(store)
const {reqCateList} = store
const router = useRouter()
// 文章列表
const artList = reactive({
    list:[]
})
// 获取文章列表
const reqList = async()=>{
    let res = await getArtList(param.value,pagination.pageSize,pagination.current)
    const data = res.data
    if (data.status == 200){
        artList.list = data.data
    }else {
        message.error(data.msg)
    }
}
// 关键字
const param = ref('')
// 表头
const columns = [
    {
        title: 'Id',
        dataIndex: 'ID',
        sorter: true,
        width: '5%',
        align: 'center',
    },
    {
        title: '标题',
        dataIndex: 'title',
        width: '10%',
        align: 'center',
    },
    {
        title: '简介',
        dataIndex: 'desc',
        width: "20%",
        align: 'center',
    },
    {
        title:'更新时间',
        dataIndex:'UpdatedAt',
        with:"10%",
        align:'center',
        // customRender:({text,record,columns})=>{
        //     text = new Date(text).format('YYYY年MM月DD日 HH:mm')
        // }

    },
    {
        title: "分类",
        dataIndex: ['Category','name'],
        width: "10%",
        align: 'center',
    },
    {
        title:"缩略图",
        dataIndex:'img',
        width: "20%",
        align:'center'
    },
    {
        title:"操作",
        dataIndex:'action',
        algin:'center'
    }
];
// 分页信息
const pagination = reactive({
    size: 'large',
    current: 1, // 当前页数
    pageSize: 10,
    total: 0,
    pageSizeOptions: ['10', '20', '30'],//可选的页面显示条数                   
    hideOnSinglePage: false, // 只有一页时是否隐藏分页器
    showQuickJumper: true, //是否可以快速跳转至某页
    showSizeChanger: true, //是否可以改变pageSize
})
// 分页发生修改
const changePage = (pag) => {
    pagination.current = pag.current
    pagination.pageSize = pag.pagesize
    // 重新发送请求
   
}
// 删除文章
const deleteart = async (id) => {
    let res = await deleteArt(id)
    const data = res.data
    if(data.status == 200){
        reqList()
    }else if(data.status >= 1004 && data.status < 2000){
        message.error(data.msg+",即将跳转到登录页")
        router.push("/login")
    }else {
        message.error(data.msg)
    }
}
// 搜索文章
const searchList = ()=>{
    reqList()
}
// 选择的分类
const cateValue = ref(undefined)
// 获取分类下的文章列表
const handleChange = async ()=>{
    if(cateValue.value == undefined){
        reqList()
        return 
    }
    const {data} = await reqCateArt(cateValue.value,pagination.pageSize,pagination.current)
    if(data.status == 200){
        artList.list = data.data
    }else if(data.status >= 1004 && data.status <2000 ){
        message.error(data.msg + "即将跳转到登录页")
        router.push('./login')
    }else {
        message.error(data.msg)
    }
}
// 编辑文章,页面跳转
const edit = (id)=>{
    router.push(`/write/${id}`)
}
// 在生命周期钩子中发送请求
onMounted(()=>{
    reqList()
    // 如果store中已有数据则直接用，若没有则请求
    if(cateList.value.list.length == 0){
        reqCateList()
    }
})
</script>

<style lang="scss" scoped>
img {
    height: 60px;
}
.ant-table-wrapper {
    margin-top: 20px;
}

.ant-btn {
    margin-left: 10px;
}
.ant-select{
    margin-left: 100px;
    width: 200px;
}
.ant-table-body{
    height: 500px;
  }
</style>