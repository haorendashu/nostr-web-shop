<script setup>
import HeadComponent from '../components/HeadComponent.vue'
import OrderProductComponent from "../components/OrderProductComponent.vue";
import { useRouter } from 'vue-router'
import {GetQuery} from "../utils/utils";
import {onMounted, reactive, ref} from "vue";
import {API} from "../api/api";

const query = GetQuery()
const router = useRouter()
const api = new API()

const product = reactive({})
const num = ref(query["num"] - 1 + 1)
const seller = ref("")
const price = ref(1)

function placeOrder() {
  router.push("/orders/pay")
}

onMounted(async () => {
  let id = query["id"]

  let result = await api.baseProductGet(id)
  if (result && result.data) {
    product.value = result.data
    seller.value = result.data.Pubkey

    if (result.data.Skus && result.data.Skus.length > 0) {
      let sku = result.data.Skus[0]
      price.value = sku.Price
    }
  }
})
</script>

<template>
  <HeadComponent title="Place Order"></HeadComponent>
  <div class="orderNew container">

    <div class="orderProducts m_t_1">
      <OrderProductComponent :product="product" :num="num" :seller="seller"></OrderProductComponent>
    </div>

    <div class="orderPlaceInfo">
      <hr/>
      <div class="orderTotal input-group">
        <div class="orderTotalL">
          Total
        </div>
        <div class="orderTotalR">
          {{num * price}}
        </div>
      </div>
      <div style="margin-top: 0.6rem">
        <textarea class="form-control" aria-label="With textarea"></textarea>
      </div>
      <div style="display: flex; margin-top: 1rem;">
        <div style="margin-right: auto!important;"></div>
        <button type="button" class="btn btn-dark productOrderBtn" v-on:click="placeOrder">Order</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.orderTotal {
  display: flex;
}
.orderTotalL {
  margin-right: auto!important;
}
.orderNew {
  min-height: calc(100% - 4rem);
  display: flex;
  flex-direction: column;
  align-items: stretch;
}
.orderProducts {
  flex-grow: 1;
}
.orderPlaceInfo {
  flex-shrink: 0;
}
.orderNew textarea {
  min-height: 6rem;
}
</style>