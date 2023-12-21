<script setup>
import HeadComponent from '../components/HeadComponent.vue'
import ProductPriceComponent from "../components/ProductPriceComponent.vue";
import {useRoute, useRouter} from 'vue-router'
import {API} from "../api/api";
import {onMounted, reactive, ref} from "vue";
import UserComponent from "../components/UserComponent.vue";

const api = new API()
const router = useRouter()
const route = useRoute()

const imgs = reactive([])
const productName = ref("Product Name")
const productPrice = ref(1000)
const productDes = ref("")
const productContent = ref("")
const productStock = ref(0)
const seller = ref("")

const num = ref(1)

function minusNum() {
  if (num.value > 1) {
    num.value--
  }
}

function addNum() {
  num.value++
}

function placeOrder() {
  router.push("/orders/new?num=" + num.value + "&id=" + id)
}

const id = route.params.id

onMounted(async () => {
  let data = await api.baseProductGet(id)
  productName.value = data.data.Name
  productPrice.value = data.data.Skus[0].Price
  productDes.value = data.data.Des
  productContent.value = data.data.Content
  productStock.value = data.data.Skus[0].Stock
  seller.value = data.data.Pubkey
})
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
        {{productName}}
      </h3>
      <ProductPriceComponent :price="productPrice"></ProductPriceComponent>
      <div class="m_t_1">
        <UserComponent :pubkey="seller"></UserComponent>
      </div>
      <div class="productDetailDes">
        {{productDes}}
      </div>

      <div class="productOrder">
        <div class="productOrderNumComp">
          <i class="bi bi-dash-circle-fill" v-on:click="minusNum"></i>
          <div class="productNum badge text-bg-secondary">
            {{num}}
          </div>
          <i class="bi bi-plus-circle-fill" v-on:click="addNum"></i>
        </div>
        <button type="button" class="btn btn-dark productOrderBtn" v-on:click="placeOrder">Place Order</button>
      </div>

      <hr/>
      <div class="productDetailContent" v-html="productContent">
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