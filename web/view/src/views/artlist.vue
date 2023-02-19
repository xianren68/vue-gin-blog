<template>
    <n-list>
        <template #header>文章列表</template>
        <n-list-item v-for='item of artList.list' :key="item.ID">
            <template #prefix>
                <img :src="item.img" alt="" style="width: 80px; height: 50px" @click="detail(item.ID)">
              </template>
              <n-thing :title="item.title">
                <p>{{item.desc}}</p>
                <p style="font-size:12px; color:blue;">{{format(item.UpdateAt)}}</p>
              </n-thing>
            <template #suffix>
                <p style="font-size:12px; color:#008c8c">#{{item.Category.name}}</p>
            </template>
        </n-list-item>
    </n-list>
</template>

<script setup>
import { useArtStore } from '@/store';
import { storeToRefs, } from 'pinia';
import { onMounted } from 'vue';
import moment from 'moment';
import { useRouter } from 'vue-router';
const router = useRouter()
const artstore = useArtStore()
// 文章仓库
const {artList,id} = storeToRefs(artstore)
const {ArtList} = artstore
// 格式化时间
const format = (value)=>{
    return moment(value).format('YYYY-MM-DD HH:mm')
}
const detail = (Id)=>{
    id.value = Id
}
onMounted(()=>{
    ArtList('')
})
</script>

<style lang="sass" scoped>

</style>