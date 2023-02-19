<template>
    <div class="home">
        <n-layout>
            <n-layout-header>
                <div class="top">
                    <div @click="routerHome">首页</div>
                    <n-menu  mode="horizontal" :options="menuOptions" @update:value="handelUp"/>
                <div class="search">
                    <n-input  placeholder="输入要搜索的文章" clearable v-model:value="param" @change="reset()">
                        <template #suffix>
                          <n-icon :component="FlashOutline" />
                        </template>
                      </n-input>
                      <n-button type="info" @click="searchArt()">
                        搜索
                      </n-button>
                </div>
                </div>
            </n-layout-header>
            <n-layout-content>
                <n-layout has-sider position="absolute" style="top: 64px; bottom: 64px">
                    <n-layout-sider bordered content-style="padding: 24px;">
                      <Nav></Nav>
                    </n-layout-sider>
                    <n-layout content-style="padding: 24px;">
                        <router-view></router-view>
                    </n-layout>
                  </n-layout>
            </n-layout-content>
            <n-layout-footer>{{new Date().getUTCFullYear()}}@xianren</n-layout-footer>
        </n-layout>
</div>
</template>

<script setup>
import Nav from '@/components/nav.vue'
import {useCateStore} from '@/store'
import {storeToRefs} from 'pinia'
import {useRoute,useRouter} from 'vue-router'
import { useArtStore } from '@/store';
const catestore = useCateStore()
const artstore = useArtStore()
// 分类仓库
const { cateList } = storeToRefs(catestore)
const { reqCateList} = catestore
// 文章仓库
const {id} = storeToRefs(artstore)
const {ArtList,ArtCat } = artstore
const route = useRoute()
const router = useRouter()

// 配置函数
const menuOPtion = (ele)=>{
    let x = ele
    x.label = ele.name
    return x
}
    
// 菜单配置
const menuOptions = reactive(cateList.value.list.map(menuOPtion))
// 通过监听来配置菜单
watch(cateList,(newValue)=>{
    Object.assign(menuOPtion,newValue.value.list.map(menuOPtion))
})

// 搜索文章参数
const param = ref('')
// 搜索文章
const searchArt = ()=>{
    router.push({
        name:'artlist'
    })
    ArtList(param.value)
}
// 重置搜索
const reset = ()=>{
    if(param.value == ''){
        ArtList(param.value)
    }
}
// 监听id变化
watch(id,(newValue)=>{
    if (newValue!=0) {
        router.push(`/detail/${id.value}`)
    }
})
const handelUp = (key,item)=>{
    ArtCat(item.id)
}
// 跳转首页
const routerHome = ()=>{
    if(route.path != '/'){
        router.push('/')
    }
    ArtList('')
}
// 在声明周期钩子中获取数据
onMounted(()=>{
    reqCateList()
})
</script>

<style lang="scss" scoped>
.home {
    height: 100%;
}

.n-layout--static-positioned {
    height: 100%;
}

.n-layout-header,
.n-layout-footer {
    padding: 24px;
    
}
.n-layout-header {
    height: 10%;
    background-color: rgb(142, 11, 122);
    padding-top:20px;
}
.n-layout-footer {
    text-align: center;
    background: rgba(128, 128, 128, 0.2);
}
.n-layout-content {
    height: 80%;
    background: #fff;
}
.top {
    display: flex;
    flex-direction: row;
    justify-content: space-around;
    align-items: center;
    font-size: 20px;
    .n-input--stateful{
        width: 300px;
    }
    .n-menu {
        font-size: 20px;
        font-weight: 500;
    }
}


</style>