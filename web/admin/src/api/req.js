import { Axios } from "axios";
// 重写axios方法
const req = Axios.creat({
    baseURL:"http://localhost:5000",
    timeout:1000
})
export default req