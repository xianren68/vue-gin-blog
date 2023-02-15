// 校验规则
// 用户名的校验规则
export const rulUsername = async (_rule,value)=>{
    if ((value.length < 2) || (value.length > 10)) {
        return Promise.reject("用户名的长度在2-10")
    }else if (value.trim() == ""){
        return Promise.reject("请输入有效字符")
    }
    return Promise.resolve()
}
// 密码的校验规则
export const rulPassword = async (_rule,value)=> {
    if ((value.length < 8) || (value.length > 18)) {
        return Promise.reject("密码的长度在6-18")
    }else if (value.trim() == ""){
        return Promise.reject("请输入有效字符")
    }else{
        return Promise.resolve()
    }
    
}
// 不能为空的值的校验规则
export const rulNotnull = (_rule,value)=>{
    if (value.trim() == ""){
        return Promise.reject("请输入值")
    }
    return Promise.resolve()
}

