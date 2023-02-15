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
// 获取分类列表
export const getCateList = (pagesize=0,pagenum=0)=>req.get(`/category/list?pagesize=${pagesize}&pagenum=${pagenum}`)
// 添加分类
export const addCate = (data)=>req.post('/category/add',data)
// 删除分类
export const deleteCate = (id)=>req.delete(`/category/delete/${id}`)
// 编辑分类
export const editCate = (id,data)=>req.put(`/category/edit/${id}`,data)
// 获取文章列表
export const getArtList = (param,pagesize=0,pagenum=0)=>req.get(`/article/list?param=${param}&pagesize=${pagesize}&pagenum=${pagenum}`)
// 删除文章
export const deleteArt = (id)=>req.delete(`/article/delete/${id}`)
// 添加文章
export const addArticle = (data)=>req.post(`/article/add`,data)
// 编辑文章
export const editArticle = (id,data)=>req.put(`/article/edit/${id}`,data)
// 获取分类下的文章列表
export const reqCateArt = (id,pagesize,pagenum)=>req.get(`/article/catelist/${id}?pagesize=${pagesize}&pagenum=${pagenum}`)
// 获取单个文章数据
export const reqOnlyArt = (id)=>req.get(`article/only/${id}`)
// 上传图片
export const uploadfront = (data)=>req.post('/upload/img',data)