<script setup>
import HeadUserImage from './HeadUserImage.vue'
import {onMounted, ref} from "vue";
import {UserStore} from "../stores/UserStore";
import {NdkStore} from "../stores/NdkStore";
import { useRoute } from 'vue-router'

const props = defineProps({
  title: {
    type: String,
    required: false,
  }
})

const isHome = ref(true)
const pubkey = ref("")

const userImageToPath = ref("/orders/")

onMounted(async () => {
  const route = useRoute()
  if (route.name == 'home') {
    isHome.value = true
  }

  const userStore = UserStore()
  if (await userStore.checkLogin()) {
    const ndkStore = NdkStore()
    const _pubkey = await ndkStore.pubkey()
    pubkey.value = _pubkey
    if (route.path.indexOf("\/seller\/") == 0) {
      userImageToPath.value = "/seller/productList"
    }
  } else {
    // un login, userImage jump to login page
    userImageToPath.value = "/login"
  }
})
</script>

<template>
  <div class="headComp container">
    <template v-if="title != null">
      <router-link to="#" @click.native="$router.back()">
        <div class="circle headBackBtn">
          <i class="bi bi-chevron-left"></i>
        </div>
      </router-link>
      <div class="headTitle">{{title}}</div>
    </template>
    <template v-if="title == null">
      <h1 class="">NWS</h1>
      <span class="subTitle">Nostr Web Shop</span>
    </template>

    <router-link :to="userImageToPath">
      <HeadUserImage :pubkey="pubkey"></HeadUserImage>
    </router-link>
  </div>
</template>

<style scoped>
.headComp {
  display: flex;
  align-items: center;
  height: 4.3rem;
  padding-top: 0.3rem;
  justify-content: space-between;
}
.headComp .subTitle {
  margin-left: 0.5rem;
  margin-right: auto!important;
  flex-grow: 1;
  /*margin-bottom: 0.5rem;*/
}
.headComp .userImageComp {
  /*margin-bottom: 0.5rem;*/
}
.headComp h1 {
  margin-bottom: 0;
}
.headBackBtn {
  flex-shrink: 0;
  width: 2.3rem;
  height: 2.3rem;
  line-height: 2.3rem;
  text-align: center;
  background-color: #999999;
  margin: 0.1rem 0.1rem 0.1rem 0.1rem;
  color: white;
}
.headTitle {
  margin-left: 0.5rem;
  margin-right: auto!important;
  font-weight: bold;
  overflow: hidden;
  max-height: 1.6rem;
  flex-grow: 1;
}
</style>
