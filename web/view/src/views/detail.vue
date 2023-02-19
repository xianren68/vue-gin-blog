<template>
    <h2>{{artInfo.title}}</h2>
    <div v-html="artInfo.content">
       
    </div>
</template>

<script setup>
import { useRoute } from 'vue-router'
import { reqOnlyArt } from '@/api'
import { onMounted } from 'vue';
const route = useRoute()
const id = route.params.id
const artInfo = reactive({

})
const only = async () => {
    const { data } = await reqOnlyArt(id)
    if (data.status == 200) {
        Object.assign(artInfo,data.data)
    } else {
        alert(data.msg)
    }

}
// 获取文章数据
onMounted(() => {
    only()
})
</script>

<style lang="scss" scoped></style>