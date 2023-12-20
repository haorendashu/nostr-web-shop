<script setup>
import HeadComponent from '../components/HeadComponent.vue'
import ProductCardComponent from "../components/ProductCardComponent.vue";
</script>

<template>
  <HeadComponent title="Product List"></HeadComponent>
  <div class="homeComp container">
    <div class="row search">
      <div class="input-group">
        <input type="text" class="form-control" placeholder="Type to search" aria-describedby="button-addon2">
        <button class="btn btn-outline-secondary" type="button" id="button-addon2">Search</button>
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
