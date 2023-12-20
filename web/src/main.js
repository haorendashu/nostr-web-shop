import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'

const app = createApp(App)

import { createPinia, storeToRefs } from 'pinia'
import {NdkStore} from './stores/NdkStore.js'
const pinia = createPinia()
app.use(pinia)

import router from './router'
app.use(router)

import { QuillEditor } from '@vueup/vue-quill'
import '@vueup/vue-quill/dist/vue-quill.snow.css';
app.component('QuillEditor', QuillEditor)

const ndkStore = NdkStore()
ndkStore.initNdk()
storeToRefs(ndkStore)

import VueQrcode from '@chenfengyuan/vue-qrcode';
app.component(VueQrcode.name, VueQrcode);

app.mount('#app')
