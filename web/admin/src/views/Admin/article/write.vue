<template>
    <a-form :model="aOrEArt" :label-col="labelCol" ref="formRef" :rules="rules">
        <a-form-item label="文章标题" name="title">
            <a-input v-model:value="aOrEArt.title"></a-input>
        </a-form-item>
        <a-form-item label="选择文章分类" name="cid">
            <a-select v-model:value="aOrEArt.cid" :options="cateList.list.map(item => ({ value: item.id, label: item.name }))"
                placeholder="请选择分类">

            </a-select>
        </a-form-item>
        <a-form-item label="输入文章介绍" name="desc">
            <a-input v-model:value="aOrEArt.desc"></a-input>
        </a-form-item>
        <a-form-item name="img">
            <a-upload list-type="picture" v-model:file-list="fileList" v-if="upShow" :custom-request="customrequest"
                class="upload-list-inline">
                <a-button>
                    <upload-outlined></upload-outlined>
                    选择封面
                </a-button>
            </a-upload>

            <div v-else>
                <img :src="aOrEArt.img" alt="">
                <a-button @click="upShow = true" size="small">修改封面</a-button>
            </div>
        </a-form-item>
        <a-form-item name="content">
            <Editor :init="init" v-model="aOrEArt.content"></Editor>
        </a-form-item>
        <a-form-item>
            <a-button @click="submit">保存</a-button>
        </a-form-item>
    </a-form>
</template>

<script setup>
import { useRoute, useRouter } from 'vue-router';
import { reactive, onMounted, ref } from 'vue'
import { reqOnlyArt, uploadfront, addArticle, editArticle } from '@/api'
import { message } from 'ant-design-vue';
import { useCateStore } from '@/store'
import { storeToRefs } from 'pinia';
import { rulNotnull } from '@/utils/rules'
import Editor from "@tinymce/tinymce-vue";
// 富文本编辑器
const init = reactive({
    plugins: "lists link image table code help wordcount",
    content_css: "writer", //主题tinymce-5-dark || tinymce-5 || default || writer || document || dark
    custom_undo_redo_levels: 50, //回退数量
    end_container_on_empty_block: true, //块级文本是否换行
    keep_styles: true, //回车是否保存原有样式，例如code块回车是否截断
    menubar: false,//是否开启顶部菜单 > false 关闭菜单 |  'edit insert view format table tools help' 菜单按照这里排序 | 参考:https://www.tiny.cloud/docs/tinymce/6/menus-configuration-options/
    toolbar_mode: 'wrap',//功能栏是否换行 > | wrap 换行  | scrolling 滚动 | sliding 省略
    height: 200,
    toolbar_location: 'top',//菜单栏位置 > bottom 底部 | top 顶部
    style_formats_merge: true,//是否开启默认功能
    elementpath: false,//是否展示编辑层级  > p span
    resize: false,//调整宽高 > true 调整高 | false 不可调整宽高 | both 宽高可调
    language: 'zh_CN',//中文

    // 自定义快捷将
    text_patterns: [
        { start: "---", replacement: "<hr/>" },
        { start: "--", replacement: "—" },
        { start: "-", replacement: "—" },
        { start: "(c)", replacement: "©" },
        { start: "//brb", replacement: "Be Right Back" },
        {
            start: "//h",
            replacement:
                '<h1 style="color: blue">Heading here</h1> <h2>Author: Name here</h2> <p><em>Date: 01/01/2000</em></p> <hr />',
        },
    ],

    // 自定义指令
    text_patterns_lookup: (ctx) => {
        const parentTag = ctx.block.nodeName.toLowerCase();
        if (parentTag === "pre" || parentTag === "p") {
            return [{ start: "`", end: "`", format: "code" }];
        } else if (parentTag === "p") {
            return [{ start: "*", end: "*", format: "bold" }];
        } else if (parentTag === "span") {
            return [
                // ctx.text is the string from the start of the block to the cursor
                { start: "brb", replacement: ctx.text + ": Be Right Back" },
            ];
        } else {
            return [];
        }
    },
})
const store = useCateStore()
const { cateList } = storeToRefs(store)
const { reqCateList } = store
const route = useRoute()
const router = useRouter()
// 设置标签宽度
labelCol: { style: { width: '150px' } }
// 绑定数据
const aOrEArt = reactive({
    title: '',
    desc: '',
    cid: undefined,
    img: '',
    content: '',
})
// 控制上传图片是否显示
const upShow = ref(false)
// 获取路由传递的文章参数
const artId = route.params
// 获取文章
const reqArt = async (id) => {
    const { data } = await reqOnlyArt(id)
    if (data.status == 200) {
        console.log(data.data);
        Object.assign(aOrEArt, data.data)
    } else if (data.status >= 1004 && data.status <= 2000) {
        message.error(data.msg + '即将跳转到登录页面')
        router.push('/login')
    } else {
        message.error(data.msg)
    }
}
const fileList = ref([])
// 自定义上传方法
const customrequest = async ({ file }) => {
    const formData = new FormData()
    formData.append('file', file)
    const { data } = await uploadfront(formData)
    if (data.status == 200) {
        upShow.value = false
        aOrEArt.img = data.url
    } else if (data.status >= 1004 && data.status <= 2000) {
        message.error(data.msg + "即将跳转到登录页")
        router.push('/login')
    } else {
        message.error(data.msg)
    }

}
// 获取表单对象
const formRef = ref()
// 校验规则
const rules = {
    title: [{ required: true, validator: rulNotnull, trigger: 'blur' }],
    cid: [{ required: true, trigger: 'blur' }],
    desc: [{ required: true, validator: rulNotnull, trigger: 'blur' }],
    content: [{ required: true, validator: rulNotnull, trigger: 'blur' }],
};
// 添加文章
const addArt = async () => {
    const { data } = await addArticle(aOrEArt)
    if (data.status == 200) {
        message.info("添加成功")
        router.push('/artlist')
    } else if (data.status >= 1004 && data.status <= 2000) {
        message.error(data.msg + "，即将跳转到登录页")
        router.push('/login')
    } else {
        message.error(data.msg)
    }
}
// 编辑文章
const editArt = async () => {
    const { data } = await editArticle(artId.id, aOrEArt)
    if (data.status == 200) {
        message.info("编辑成功")
        router.push('/artlist')
    } else if (data.status >= 1004 && data.status <= 2000) {
        message.error(data.msg + "，即将跳转到登录页")
        router.push('/login')
    } else {
        message.error(data.msg)
    }
}
// 提交
const submit = async () => {
    try {
        await formRef.value.validate()
        if (artId.id > 0) {
            editArt()
        } else {
            addArt()
        }
    } catch (error) {
        message.error('输入不合法')
    }
}
// 在生命周期钩子中获取数据
onMounted(() => {
    if (artId.id > 0) {
        reqArt(artId.id)
    } else {
        upShow.value = true
    }
    // 如果store中已有数据则直接用，若没有则请求
    if (cateList.value.list.length == 0) {
        reqCateList()
    }
})
</script>

<style lang="scss" scoped>
img {
    height: 60px;
    margin-right: 20px;
}
</style>