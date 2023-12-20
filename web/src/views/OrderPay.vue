<script setup>
import HeadComponent from '../components/HeadComponent.vue'
import UserComponent from "../components/UserComponent.vue";
</script>

<template>
  <HeadComponent title="Order Pay"></HeadComponent>
  <div class="orderPay container">
    <UserComponent pubkey="29320975df855fe34a7b45ada2421e2c741c37c0136901fe477133a91eb18b07"></UserComponent>
    <h2>This is a product title</h2>
    <div class="productCode">CODE XXXX</div>
    <div class="qrCode">
      <vue-qrcode value="Hello, World!" :options="{ width: 200 }"></vue-qrcode>
    </div>
    <div class="productPrice">3,000 Sats</div>
    <div class="lnlink">lnbc660n1psdfdsfdslcfmewoifimcad4523fkmweofmdkf9234lkwmfdssfdsfdslcfmewoifimcad4523fkmweofmdkf9234lkwmfdssfdsfdslcfmewoifimcad4523fkmweofmdkf9234lkwmfdssf</div>
  </div>
</template>

<style scoped>
.orderPay {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding-top: 3rem;
}
.orderPay h2 {
  font-weight: bold;
}
.orderPay h2, .lnlink {
  margin-top: 1rem;
}
.lnlink {
  width: 30rem;
  word-break: break-all;
}

@media (max-width: 600px) {
  .lnlink {
    width: 90%;
  }
}
</style>