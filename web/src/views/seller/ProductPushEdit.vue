<script setup>
import HeadComponent from '../../components/HeadComponent.vue'
import { useRouter } from 'vue-router'
import {GetQuery} from "../../utils/utils";
import {onMounted, ref} from "vue";
import {API} from "../../api/api";

const query = GetQuery()
const router = useRouter()
const api = new API()

const status = ref(-1)
const name = ref("")
const noticePubkey = ref("")
const pushAddress = ref("")
const pushKey = ref("")
const pushType = ref(1)

async function save() {
  let info = {}
  info.Pid = query["id"]
  info.Status = status.value == true ? 1 : -1
  info.NoticePubkey = noticePubkey.value
  info.PushAddress = pushAddress.value
  info.PushKey = pushKey.value
  info.PushType = pushType.value - 1 + 1

  console.log(info)
  let result = await api.shopProductPushInfoSave(info)
  if (result["code"] == 200) {
    router.replace("/seller/productList")
  }
}

onMounted(async () => {
  let id = query["id"]
  let result = await api.shopProductPushInfoGet(id)
  if (result.data) {
    status.value = result.data.Status == 1
    name.value = result.data.Name
    noticePubkey.value = result.data.NoticePubkey
    pushAddress.value = result.data.PushAddress
    pushKey.value = result.data.PushKey
    pushType.value = result.data.PushType
  }
})
</script>

<template>
  <div class="productEdit">
    <HeadComponent title="Product Push Info Edit"></HeadComponent>
    <div class="container">
      <div class="row m_t_05">
        <label for="Name-url" class="form-label">Name</label>
        <div class="input-group">
          <input type="text" class="form-control" id="Name-url" disabled v-model="name" >
        </div>
      </div>

      <div class="row m_t_05">
        <label for="Order_Notice_Pubkey" class="form-label">Order Notice Pubkey</label>
        <div class="input-group">
          <input type="text" class="form-control" id="Order_Notice_Pubkey" v-model="noticePubkey">
        </div>
      </div>

      <div class="row m_t_1">
        <div class="input-group">
          <div class="form-check form-switch">
            <input class="form-check-input" type="checkbox" role="switch" id="flexSwitchCheckDefault" v-model="status">
            <label class="form-check-label" for="flexSwitchCheckDefault">Order push</label>
          </div>
        </div>
      </div>

      <template v-if="status">
        <div class="row m_t_05">
          <label for="Push_Address" class="form-label">Push Address</label>
          <div class="input-group">
            <input type="text" class="form-control" id="Push_Address" v-model="pushAddress">
          </div>
        </div>

        <div class="row m_t_05">
          <label for="Push_Address" class="form-label">Push Key</label>
          <div class="input-group">
            <input type="text" class="form-control" id="Push_Key" v-model="pushKey">
          </div>
        </div>

        <div class="row m_t_1">
          <div class="input-group">
            <div class="form-check m_r_05">
              <input class="form-check-input" type="radio" name="exampleRadios" id="exampleRadios1" value="1" v-model="pushType">
              <label class="form-check-label" for="exampleRadios1">
                Api push
              </label>
            </div>
            <div class="form-check">
              <input class="form-check-input" type="radio" name="exampleRadios" id="exampleRadios2" value="2" v-model="pushType">
              <label class="form-check-label" for="exampleRadios2">
                Front push
              </label>
            </div>
          </div>
        </div>
      </template>

      <div class="row m_t_1" style="padding: 0 0.8rem;">
        <button type="button" class="btn btn-dark" v-on:click="save">Submit</button>
      </div>

    </div>
  </div>
</template>

<style scoped>
@media (max-width: 600px) {
}
</style>