<script setup>
import HeadComponent from '../../components/HeadComponent.vue'
import { useRouter } from 'vue-router'
import OrderProductComponent from "../../components/OrderProductComponent.vue"
import {API} from "../../api/api"
import {onMounted, reactive} from "vue"
import {UIStore} from "../../stores/UIStore";

const api = new API()
const router = useRouter()
const uiStore = UIStore()

const products = reactive([])

function toNewProduct() {
  router.push("/seller/productNew")
}

function productDel(id) {
  uiStore.modal("Delete confirm", "Delete this product?", async function () {
    await api.shopProductDel(id)
    await updateProducts()
  })
}

async function updateProducts() {
  let result = await api.shopProductList()

  products.splice(0, products.length)
  for (let index in result.list) {
    products.push(result.list[index])
  }
}

onMounted(async () => {
  await updateProducts()
})
</script>

<template>
  <HeadComponent title="Product List"></HeadComponent>
  <div class="container m_t_05">
    <button type="button" class="btn btn-dark" v-on:click="toNewProduct">
      <i class="bi bi-plus"></i>
    </button>
  </div>
  <div class="productList container m_t_05">
    <template v-for="product in products">
      <div class="productListItemContainer">
        <router-link :to="'/seller/productEdit?id=' + product.Id">
          <div class="productListItem m_b_1">
            <OrderProductComponent :product="product"></OrderProductComponent>
          </div>
        </router-link>
        <div class="productListItemDel" v-on:click="productDel(product.Id)">
          <i class="bi bi-trash-fill"></i>
        </div>
      </div>
    </template>
  </div>
</template>

<style scoped>
.productList a {
  text-decoration: none;
}
.productListItemContainer {
  position: relative;
}
.productListItemDel {
  position: absolute;
  right: 0;
  bottom: 0;
  font-size: 1.4rem;
  color: red;
}
@media (max-width: 600px) {
}
</style>