<template>
    <a-card title="用户列表" :bordered="false">
        <a-input-search v-model:value="param" placeholder="输入用户名搜索" style="width: 200px" @search="searchList" 
            allow-clear />
        <a-button type="primary" @click="modalShow">
            添加
        </a-button>
        <a-modal title="添加用户" style="top: 20px" v-model:visible="modalisShow" @ok="submitAdd">
            <a-form :model="addUserInfo" :rules='rules' :labelCol="labelCol" ref="addRef">
                <a-form-item  label="用户名" name="username">
                    <a-input v-model:value="addUserInfo.username"></a-input>
                </a-form-item>
                <a-form-item  label="密码" name="password">
                    <a-input v-model:value="addUserInfo.password"></a-input>
                </a-form-item>
                <a-form-item label="确认密码" name="rePassword">
                    <a-input v-model:value="addUserInfo.rePassword"></a-input>
                </a-form-item>
                <a-form-item  label="角色" name="role">
                    <a-input v-model:value="addUserInfo.role"></a-input>
                </a-form-item>
            </a-form>
        </a-modal>

        <a-table :dataSource="userList.list" :columns="columns" :pagination="pagination" @change="changePage">
            <template #bodyCell="{ column, text, record }">
                <template v-if="column.dataIndex == 'username'">
                    <div>
                        <a-input v-if="record.ID == editkey" placeholder="输入用户名" v-model:value="editdata.name" />
                        <template v-else>
                            {{ text }}
                        </template>
                    </div>
                </template>
                <template v-if="column.dataIndex == 'role'">
                    <div>
                        <a-input v-if="record.ID == editkey" placeholder="管理员：1，用户：2" v-model:value="editdata.r" />
                        <template v-else>
                            {{ customRender(text) }}
                        </template>
                    </div>
                </template>
                <template v-if="column.dataIndex == 'action'">
                    <span v-if='record.ID != editkey'>
                        <a-button type="primary" size="small" ghost @click="edit(record)">编辑</a-button>
                        <a-popconfirm title="确定删除吗？" @confirm="deleteuser(record.ID, record.role)">
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
import { reactive, onMounted, ref } from 'vue'
import { message } from 'ant-design-vue';
import 'ant-design-vue/es/message/style/css'
import { reqUserList, editInfo, deleteUser,addUser } from '@/api'
import { rulPassword,rulUsername} from '@/utils/rules';
import {useRouter} from 'vue-router'
let router = useRouter()
// 表格标签长度
const labelCol = reactive({
    style:{
        width:"80px"
    }
})
// 用户列表
const userList = reactive({
    list: []
})
// 表头
const columns = [
    {
        title: 'Id',
        dataIndex: 'ID',
        sorter: true,
        width: '10%',
        align: 'center',
    },
    {
        title: '用户名',
        dataIndex: 'username',
        width: '30%',
        align: 'center',
    },
    {
        title: '权限',
        dataIndex: 'role',
        width: "20%",
        align: 'center',
    },
    {
        title: "操作",
        dataIndex: 'action',
        align: 'center',
    }
];
// 权限信息
const customRender = (text) => {
    if (text == 0) {
        return "root"
    } else if (text == 1) {
        return "管理员"
    } else {
        return "用户"
    }
}
// 搜索关键字
let param = ref("")
// 请求用户列表
const reqList = async () => {
    const { current, pageSize } = pagination
    let res = await reqUserList(param.value, pageSize, current)
    let data = res.data
    if (data.status == 200) {
        userList.list = data.data
        // 
        pagination.total = data.total
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
    reqList()
}

// 控制输入框以及编辑按钮是否显示
const editkey = ref(0)
// 输入的编辑
const editdata = reactive({
    name: '', // 输入的用户名
    r: '' // 输入的角色
})
// 控制是在编辑哪一行
const edit = (record) => {
    editkey.value = record.ID
    // 方便在显示时有默认值
    editdata.name = record.username
    editdata.r = record.role
}
// 获取cookie中的role
const getCookie = () => {
    if (document.cookie){
        const list = document.cookie.split("=")
        let role = list[1].split(";")[0]
        return role
    }
    return 3
    
}
// 保存编辑
const saveEdit = async ({ ID, role, username }) => {
    // 获取cookie
    let userRole = getCookie()
    console.log(userRole);
    if (editdata.name.trim() === "") {
        message.error("请输入有效字符")
        return
    } else if (editdata.name.length < 2 || editdata.name.length > 10) {
        message.error("用户名应在2-10位")
        return
    } else if (editdata.r != '1' && editdata.r != '2') {
        message.error("权限应为1或2")
        return
    } else if (role <= userRole) {
        message.error('无此权限')
        return
    }
    let { name, r } = editdata
    // 转化为整型
    r = parseInt(r, 10)
    const data = {
        username: name,
        role: r
    }

    // 判断username是否没有改变
    if (username == name) {
        data.username = ' '
    }
    let res = await editInfo(ID, data)
    if (res.status == 200) {
        let data = res.data
        if (data.status == 200) {
            // 直接修改不用发送请求
            userList.list.forEach(ele => {
                if (ele.ID == ID) {
                    ele.username = name
                    ele.role = r
                }
            })
            editkey.value = 0
        } else if (data.status > 1004) {
            message.error(data.msg + "即将跳转到登录页")
            router.push('/login')
        } else {
            message.error(data.msg)
        }
    }
}
// 取消编辑
const cancelEdit = (id) => {
    editkey.value = 0
}
// 删除用户
const deleteuser = async (id, role) => {
    let userRole = getCookie()
    if (role <= userRole) {
        message.error("权限不够")
        return
    }
    let res = await deleteUser(id)
    if (res.status == 200) {
        let data = res.data
        if (data.status == 200) {
            // 重新请求列表
            reqList()
        } else if (data.status > 1004) {
            message.error(data.msg + "即将跳转到登录页")
            router.push('/login')
        } else {
            message.error(data.msg)
        }
    }

}
// 搜索用户
const searchList = () => {
    // 发送请求用户列表
    reqList()
}
// 控制添加用户的弹窗是否出现
let modalisShow = ref(false)
// 弹窗显示
const modalShow = () => {
    modalisShow.value = true
}
// 要添加的用户数据
const addUserInfo = reactive({
    username:"",
    password:"",
    rePassword:"",
    role:2
})
// 权限校验规则
const ruleRole = (_rule,value)=>{
    if(value <= getCookie){
        return Promise.reject("权限不够")
    }
    return Promise.resolve()
}
// 确认密码的校验
const ruleRepassword = async(_rule,value)=>{
    if(value !== addUserInfo.password){
        return Promise.reject("密码不一致")
    }
    return Promise.resolve()
}
// 设置校验规则
const rules = {
      username: [{ required: true, validator: rulUsername, trigger: 'blur' }],
      password: [{ required: true,validator: rulPassword, trigger: 'blur' }],
      rePassword:[{ required: true,validator: ruleRepassword, trigger: 'blur' }],
      role:[{ required: true,validator: ruleRole, trigger: 'blur' }],

    };
// 添加用户的请求
const submitUser = async ()=>{
    const {username,password,role} = addUserInfo
    let res = await addUser({username,password,role})
    let data = res.data
    if(data.status == 200){
        addUserInfo.username = ""
        addUserInfo.password = ""
        addUserInfo.rePassword = ""
        addUserInfo.role = 2
    }else if(data.status > 1004){
        message.error(data.msg + "即将跳转到登录页")
        router.push('/login')
    }else {
        message.error(data.msg)
    }
}
// 获取表单组件
let addRef = ref()
// 提交请求
const submitAdd = async ()=>{
    try {
        await  addRef.value.validate()
        await submitUser()
        modalisShow.value = false
        // 再次获取列表
        reqList()
    } catch (error) {
        message.error("输入不合法")
    }
}
// 生命周期钩子中发送请求
onMounted(() => {
    reqList()
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