<template>
    <div class="container">
        <div class="login">
            <a-form :label-col="labelCol" :model="userInfo" :rules="rules" >
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
                    <a-checkbox class="check" v-model:value="userInfo.remember">已阅读</a-checkbox>
                </a-form-item>
                <a-form-item class="confirm">
                    <a-button type="primary" class="primar">登录</a-button>
                    <a-button type="default">取消</a-button>
                </a-form-item>
            </a-form>
        </div>
    </div>
</template>

<script setup>
import {ref,reactive} from "vue"
import { UserOutlined, LockOutlined } from '@ant-design/icons-vue';
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
// 校验规则
let rulUsername = async (_rule,value)=>{
    console.log(value.length);
    if ((value.length < 4) || (value.length > 10)) {
        return Promise.reject("用户名的长度在4-10")
    }else if (value.trim() == ""){
        return Promise.reject("请输入有效字符")
    }
    return Promise.resolve()
}
let rulPassword = async (_rule,value)=> {
    if ((value.length <= 4) || (value.length >= 10)) {
        return Promise.reject("密码的长度在4-10")
    }else if (value.trim() == ""){
        return Promise.reject("请输入有效字符")
    }else{
        return Promise.resolve()
    }
    
}
// 设置校验规则
const rules = {
      username: [{ required: true, validator: rulUsername, trigger: 'blur' }],
      password: [{ validator: rulPassword, trigger: 'blur' }],
    };
</script>

<style lang="scss" scoped>
    .container {
        height: 100%;
        background-color:#008c8c;
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