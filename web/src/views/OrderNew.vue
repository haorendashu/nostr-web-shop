<script setup>
import HeadComponent from '../components/HeadComponent.vue'
import OrderProductComponent from "../components/OrderProductComponent.vue";
import { useRouter } from 'vue-router'

const router = useRouter()
function placeOrder() {
  router.push("/orders/pay")
}
</script>

<template>
  <HeadComponent title="Place Order"></HeadComponent>
  <div class="orderNew container">

    <OrderProductComponent seller="29320975df855fe34a7b45ada2421e2c741c37c0136901fe477133a91eb18b07"></OrderProductComponent>

    <div class="orderPlaceInfo">
      <hr/>
      <div class="orderTotal input-group">
        <div class="orderTotalL">
          Total
        </div>
        <div class="orderTotalR">
          3000
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
.orderProduct {
  flex-grow: 1;
}
.orderPlaceInfo {
  flex-shrink: 0;
}
.orderNew textarea {
  min-height: 6rem;
}
</style>