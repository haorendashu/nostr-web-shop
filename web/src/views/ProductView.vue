<script setup>
import HeadComponent from '../components/HeadComponent.vue'
import ProductPriceComponent from "../components/ProductPriceComponent.vue";
import { useRouter } from 'vue-router'

const router = useRouter()
function placeOrder() {
  router.push("/orders/new")
}
</script>

<template>
  <div class="productDetail">
    <HeadComponent title=""></HeadComponent>
    <div class="container" id="productDetailTopContainer">
      <div class="productDetailTop">
        <img src="https://via.assets.so/img.jpg?w=200&h=200&tc=white&bg=#cecece">
      </div>
    </div>
    <div class="mainComp container productDetailMain">
      <h3>
        This is product name
      </h3>
      <ProductPriceComponent price="3000"></ProductPriceComponent>
      <div class="productDetailDes">
        This is product description. This is product description. This is product description. This is product description.
      </div>

      <div class="productOrder">
        <div class="productOrderNumComp">
          <i class="bi bi-dash-circle-fill"></i>
          <div class="productNum badge text-bg-secondary">
            1
          </div>
          <i class="bi bi-plus-circle-fill"></i>
        </div>
        <button type="button" class="btn btn-dark productOrderBtn" v-on:click="placeOrder">Place Order</button>
      </div>

      <hr/>
      <div class="productDetailContent">
        图文详情<br/>
        <img src="https://via.assets.so/img.jpg?w=200&h=200&tc=white&bg=#cecece">
      </div>
    </div>
  </div>
</template>

<style scoped>
.productDetailMain h3 {
  margin-top: 0.6rem;
}
.productDetailDes {
  margin-top: 0.6rem;
  color: #adb5bd;
}
.productOrder {
  display: flex;
  align-items: center;
  margin-top: 1rem;
}
.productOrderNumComp{
  height: 2rem;
  line-height: 2rem;
  display: flex;
}
.productNum {
  line-height: 1.6rem;
  margin: 0 1rem;
}
.productOrderNumComp i {
  font-size: 1.8rem;
}
.productOrderBtn {
  margin-left: auto!important;
}

@media (max-width: 600px) {
  #productDetailTopContainer {
    padding: 0;
  }
  .productDetailTop {
    position: relative;
    aspect-ratio: 1 / 1;
  }
  .productDetailTop img {
    width: 100%;
  }
  .productDetailContent img {
    width: 100%;
  }
}
</style>