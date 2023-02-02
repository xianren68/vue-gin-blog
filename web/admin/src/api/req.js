import  axios  from "axios";
// 重写axios方法
const req = axios.create({
    baseURL:"http://localhost:5000/v1",
    timeout:1000
})
req.defaults.withCredentials = true
// 设置请求拦截器携带tooken
req.interceptors.request.use(function (config){
    // 获取本地token
    let token = window.localStorage.getItem("token")
    if (token) {
        // 注意：token前边有 'Bearer ' 的信息前缀,API接口需求，Bearer后边有空格
        config.headers.Authorization = 'Bearer ' + token
      }
    return config
})
export default req