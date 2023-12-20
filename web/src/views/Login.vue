<script setup>
import HeadComponent from '../components/HeadComponent.vue'
import {NdkStore} from "../stores/NdkStore";
import { Base64 } from 'js-base64';
import axios from 'axios';
import {UserStore} from "../stores/UserStore";
import {useRouter} from "vue-router/dist/vue-router";
import G from "../utils/globals";
import {GetQuery} from "../utils/utils";

const query = GetQuery()
const isShop = query["isShop"]

const router = useRouter()
const userStore = UserStore()
const ndkStore = NdkStore()
async function loginByExtension() {
  const nip98EventText = await ndkStore.genNIP98EventText(G.LoginEventU, "GET")
  const authStr = Base64.encodeURL(nip98EventText)

  const result = await axios.get(G.BasePath + "/base/login", {
    headers: {
      "Authorization": "Nostr " + authStr,
    }
  })

  console.log(result.data)
  if (result.data["code"] == 200 && result.data["token"]) {
    userStore.setToken(result.data["token"])
  } else {
    console.log(result.data["msg"])
  }

  let u = query["u"]
  if (u) {
    location.href = decodeURIComponent(u)
  } else {
    if (isShop) {
      router.push("/seller/productList")
    } else {
      router.push("/")
    }
  }
}
</script>

<template>
  <HeadComponent title="Login"></HeadComponent>
  <div class="loginComp container">
    <h1>NWS - Nostr Web Shop</h1>
    <div>
      <div class="input-group mb-3">
        <input type="password" class="form-control" placeholder="Private Key" >
      </div>
      <div class="row" style="padding: 0px 0.8rem;">
        <button type="button" class="btn btn-dark">Login By Private Key</button>
      </div>
      <hr/>
      <div class="row" style="padding: 0px 0.8rem;">
        <button type="button" class="btn btn-dark" v-on:click="loginByExtension">Login By Nostr Extensions (NIP-07)</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.loginComp {
  margin-top: 5rem;
}
.loginComp h1 {
  text-align: center;
  font-weight: bold;
  margin-bottom: 3rem;
}
</style>
