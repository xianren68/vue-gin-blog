<template>
    <div class="container">
        <div class="login">
            <a-form :label-col="labelCol" :model="userInfo" :rules="rules" ref="formRef">
                <a-form-item label="用户名" name="username">
                    <a-input v-model:value="userInfo.username" placeholder="请输入用户名">
                        <template #prefix>
                            <UserOutlined class="site-form-item-icon" />
                          </template>                  
                    </a-input>
                </a-form-item>
                <a-form-item label="密码" name="password">
                    <a-input-password v-model:value="userInfo.password" placeholder="请输入密码">
                        <template #prefix>
                            <LockOutlined class="site-form-item-icon" />
                          </template>
                    </a-input-password>
                </a-form-item>
                <a-form-item name="remember">
                    <a-checkbox class="check" v-model:checked="userInfo.remember">已阅读</a-checkbox>
                </a-form-item>
                <a-form-item class="confirm">
                    <a-button type="primary" class="primar" @click="submit" >登录</a-button>
                    <a-button type="default" @click="resetFields">取消</a-button>
                </a-form-item>
            </a-form>
        </div>
    </div>
</template>

<script setup>
import {ref,reactive} from "vue"
import router from '@/router'
import {adminLogin} from '@/api'
import { UserOutlined, LockOutlined } from '@ant-design/icons-vue';
import { message } from 'ant-design-vue';
import 'ant-design-vue/es/message/style/css'
import { rulUsername,rulPassword } from "../utils/rules";
// 设置lable固定宽度
const labelCol = reactive({
    style:{
        width:"55px"
    }
})
// 设置用户信息
const userInfo = reactive({
    username:"",
    password:"",
    remember:true
})

// 设置校验规则
const rules = {
      username: [{ required: true, validator: rulUsername, trigger: 'blur' }],
      password: [{ required: true,validator: rulPassword, trigger: 'blur' }],
    };

// 表单重置以及提交表单
const formRef = ref()
    // 重置表单
const resetFields = ()=>{
    formRef.value.resetFields()
}
    // 提交表单
const submit = async ()=>{
    try {
        await  formRef.value.validate()
        login()
        // 发送请求
    } catch (error) {
        message.error("输入不合法")
    }  
}
    // 登录请求
const login = async ()=>{
    let res = await adminLogin(userInfo)
    const data = res.data
    if (res.status == 200) {
        if (data.status !== 200) {
            message.error(data.msg)
            return
        }
        // 保存token
        window.localStorage.setItem("token",data.token)
        // 路由跳转
        router.push('/')
    }else {
        message.error("请求出错")
    }
}

    





</script>

<style lang="scss" scoped>
    .container {
        height: 100%;
        background-color:#008c8c;
        .alert {
            position: absolute;
            top:40px;
            left:40%;
            width:300px;
            height: 20px;
        }
        .login {
            position: absolute;
            top: 50%;
            left: 70%;
            width: 400px;
            height: 300px;
            transform: translate(-50%,-50%);
            box-shadow: 0 0 10px #000;
            padding:60px 20px 40px 20px;
            background-color: azure;
            .confirm {
                display: flex;
                justify-content: flex-end;
                .primar {
                    margin: 0 20px 0 150px;
                }
            }
            .check {
                margin-left: 50px;
            }
        }
    }
    
</style>