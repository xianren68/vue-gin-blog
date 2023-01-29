import { createApp } from 'vue'
import "./style.css"
import App from './App.vue'
import router from './router'
import {Button,Form,Input,Checkbox} from "ant-design-vue"

createApp(App).use(router).mount('#app')
