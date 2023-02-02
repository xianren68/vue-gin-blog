import req from "./req" 

// admin
// 登录
export const adminLogin = (data)=>req.post('/admin/login',data)
// 获取用户列表,默认为0
export const reqUserList = (param="",pagesize=0,pagenum=0)=>req.get(`/user/list?param=${param}&pagesize=${pagesize}&pagenum=${pagenum}`)
// 修改用户信息
export const editInfo = (id,data)=>req.put(`/user/edit/${id}`,data)
// 删除用户
export const deleteUser = (id)=>req.delete(`/user/delete/${id}`)
// 添加用户
export const addUser = (data)=>req.post("/user/add",data)