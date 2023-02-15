<template>
    <a-card title="分类列表" :bordered="false">
        <a-button type="primary" @click="modalShow">
            添加
        </a-button>
        <a-modal title="添加分类" style="top: 20px" v-model:visible="modalisShow" @ok="submitAdd">
                    <a-input v-model:value="cateName"></a-input>
        </a-modal>
        <a-table :dataSource="cateList.list" :columns="columns" :pagination="pagination" @change="changePage">
            <template #bodyCell="{ column, text, record }">
                <template v-if="column.dataIndex == 'name'">
                    <div>
                        <a-input v-if="record.id == editkey" placeholder="输入分类名" v-model:value="editdata.name" />
                        <template v-else>
                            {{ text }}
                        </template>
                    </div>
                </template>
                <template v-if="column.dataIndex == 'action'">
                    <span v-if='record.id != editkey'>
                        <a-button type="primary" size="small" ghost @click="edit(record)">编辑</a-button>
                        <a-popconfirm title="确定删除吗？" @confirm="deleteuser(record.id)">
                            <template #icon><question-circle-outlined style="color: red" /></template>
                            <a-button type="danger" size="small" ghost>删除</a-button>
                        </a-popconfirm>
                    </span>
                    <span v-else>
                        <a-button type="primary" size="small" @click="saveEdit(record)">保存</a-button>
                        <a-button size="small" @click="cancelEdit">取消</a-button>
                    </span>

                </template>
            </template>
        </a-table>


    </a-card>
</template>

<script setup>
import {ref,reactive,onMounted} from 'vue'
import {editCate,deleteCate,addCate} from '@/api'
import { message } from 'ant-design-vue';
import { useCateStore } from '@/store';
import { useRouter } from 'vue-router';
import { storeToRefs } from 'pinia';
// 路由
const router = useRouter()
// 仓库
const store = useCateStore()
const {reqCateList} = store
const {cateList} = storeToRefs(store)
// 表头
const columns = [
    {
        title: 'Id',
        dataIndex: 'id',
        sorter: true,
        width: '10%',
        align: 'center',
    },
    {
        title: '类名',
        dataIndex: 'name',
        width: '30%',
        align: 'center',
    },
    {
        title: "操作",
        dataIndex: 'action',
        align: 'center',
    }
];
// 控制添加分类的弹窗是否出现
let modalisShow = ref(false)
// 弹窗显示
const modalShow = () => {
    modalisShow.value = true
}
// 要添加的分类数据
let cateName = ref("")
// 添加分类
const submitAdd = async ()=>{
    // 判断输入是否为空
    if (cateName.value.trim()===""){
        message.error("请输入分类名")
        return 
    }
    let res = await addCate({name:cateName.value})
    const data = res.data
    if(data.status == 200){
        message.info("添加成功")
        modalisShow = false
        reqCateList(pagination.pageSize,pagination.current)
    }else if(data.status >= 1004 && data.status < 2000){
        message.error(data.msg+",即将跳转到登录页")
        router.push("/login")
    }else {
        message.error(data.msg)
    }
}
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

// 控制输入框以及编辑按钮是否显示
const editkey = ref(0)
// 输入的编辑
const editdata = reactive({
    name: '', // 输入的用户名
})
// 控制是在编辑哪一行
const edit = (record) => {
    editkey.value = record.id
}
// 保存编辑
const saveEdit = async ({id}) => { 
    // 判断输入是否为空
    if (editdata.name.trim()===""){
        message.error("请输入分类名")
        return 
    }
    let res = await editCate(id,{name:editdata.name})
    const data = res.data
    if(data.status == 200){
        message.info("编辑成功")
        modalisShow = false
        cateList.list.forEach(ele=>{if(ele.id == id){ele.name = editdata.name}})
        editkey.value=0
    }else if(data.status >= 1004 && data.status < 2000){
        message.error(data.msg+",即将跳转到登录页")
        router.push("/login")
    }else {
        message.error(data.msg)
    }
}
// 取消编辑
const cancelEdit = (id) => {
    editkey.value = 0
}
// 删除分类
const deleteuser = async (id) => {
    let res = await deleteCate(id)
    const data = res.data
    if(data.status == 200){
        reqCateList(pagination.pageSize,pagination.current)
    }else if(data.status >= 1004 && data.status < 2000){
        message.error(data.msg+",即将跳转到登录页")
        router.push("/login")
    }else {
        message.error(data.msg)
    }

}
// 生命周期钩子中获取分类列表
onMounted(()=>{
    reqCateList(pagination.pageSize,pagination.current)
})
</script>

<style lang="scss" scoped>
.ant-table-wrapper {
    margin-top: 20px;
}

.ant-btn {
    margin-left: 10px;
}

</style>