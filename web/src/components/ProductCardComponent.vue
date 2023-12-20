
<script setup>
import UserComponent from "./UserComponent.vue";
import ProductPriceComponent from "./ProductPriceComponent.vue";
</script>

<template>
  <router-link to="/product/123">
    <div class="productCardComp">
      <div>
        <div class="productImage">
          <img src="https://via.assets.so/img.jpg?w=200&h=200&tc=white&bg=#cecece">
        </div>
        <div class="productName">
          This is a product name
        </div>
        <div class="seller">
          <UserComponent pubkey="29320975df855fe34a7b45ada2421e2c741c37c0136901fe477133a91eb18b07"></UserComponent>
        </div>
      </div>
      <ProductPriceComponent price="3000"></ProductPriceComponent>
    </div>
  </router-link>
</template>

<style scoped>
.productCardComp {
  display: inline-block;
  padding: 0.5rem;
  position: relative;
}
.productName {
  margin-top: 0.5rem;
  max-height: 1.5rem;
  overflow: hidden;
}
.productImage img {
  width: 100%;
}
.productCardComp .priceComp {
  position: absolute;
  top: 1rem;
  right: 1rem;
}
</style>
