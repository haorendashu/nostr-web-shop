<script setup>
import HeadComponent from '../components/HeadComponent.vue'
import OrderProductComponent from "../components/OrderProductComponent.vue";
</script>

<template>
  <HeadComponent title="Orders"></HeadComponent>
  <div class="ordersComp container">
    <div class="orderItemComp">
        <OrderProductComponent order-status="未支付" order-price="3000"  seller="29320975df855fe34a7b45ada2421e2c741c37c0136901fe477133a91eb18b07"></OrderProductComponent>
    </div>
  </div>
  <div class="ordersComp container">
    <div class="orderItemComp">
      <OrderProductComponent order-status="完成" order-price="2000" seller="29320975df855fe34a7b45ada2421e2c741c37c0136901fe477133a91eb18b07"></OrderProductComponent>
    </div>
  </div>
  <div class="ordersComp container">
    <div class="orderItemComp">
      <OrderProductComponent order-status="完成" order-price="6000" seller="29320975df855fe34a7b45ada2421e2c741c37c0136901fe477133a91eb18b07"></OrderProductComponent>
    </div>
  </div>
  <div class="ordersComp container">
    <div class="orderItemComp">
      <OrderProductComponent order-status="完成" order-price="3000" seller="29320975df855fe34a7b45ada2421e2c741c37c0136901fe477133a91eb18b07"></OrderProductComponent>
    </div>
  </div>
</template>

<style scoped>
.orderItemComp {
  border: 1px solid #dee2e6;
  padding: 0.6rem;
  border-radius: 0.5rem;
  margin-bottom: 1rem;
}
</style>
