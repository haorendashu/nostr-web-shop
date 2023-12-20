<script setup>
import HeadComponent from '../components/HeadComponent.vue'
import ProductCardComponent from "../components/ProductCardComponent.vue";
import {API} from "../api/api";
import {onMounted} from "vue";
function doSearch() {
}

const api = new API()
onMounted(async () => {
  const data = await api.baseProductList()
  router.push("/orders/new")
})
</script>

<template>
  <HeadComponent></HeadComponent>
  <div class="homeComp container">
    <div class="row search">
      <div class="input-group">
        <input type="text" class="form-control" placeholder="Type to search" aria-describedby="button-addon2">
        <button v-on:click="doSearch" class="btn btn-outline-secondary" type="button" id="button-addon2">Search</button>
      </div>
    </div>
    <div class="productList">
      <div class="row">
        <div class="col-lg-3 col-6">
          <ProductCardComponent></ProductCardComponent>
        </div>
        <div class="col-lg-3 col-6">
          <ProductCardComponent></ProductCardComponent>
        </div>
        <div class="col-lg-3 col-6">
          <ProductCardComponent></ProductCardComponent>
        </div>
        <div class="col-lg-3 col-6">
          <ProductCardComponent></ProductCardComponent>
        </div>
      </div>
    </div>
  </div>
</template>
