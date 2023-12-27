<script setup>
import HeadComponent from '../components/HeadComponent.vue'
import OrderProductComponent from "../components/OrderProductComponent.vue"
import {API} from "../api/api"
import {useRouter} from "vue-router"
import {onMounted, reactive} from "vue"

const router = useRouter()
const api = new API()

const orders = reactive([])

onMounted(async () => {
  let result = await api.userOrderList()
  orders.splice(0, orders.length)
  for (let index in result.list) {
    let order = result.list[index]

    let products = []
    for (let skuIndex in order.Skus) {
      let sku = order.Skus[skuIndex]
      products.push({
        "Id": sku.Id,
        "Name": sku.Name,
        "Imgs": sku.Img,
        "Price": order.Price,
        "Num": sku.Num,
        "Skus": [
          {
            "Code": sku.Code,
            "Price": sku.Price,
          }
        ]
      })
    }

    orders.push({
      "Order": order,
      "Products": products,
    })
  }
})

function payStatus(ps) {
  if (ps == 2) {
    return "Paied"
  } else {
    return "Unpay"
  }
}
</script>

<template>
  <HeadComponent title="Orders"></HeadComponent>
  <div class="ordersComp container" v-for="order in orders">
    <router-link :to="'/orders/pay?id=' + order.Order.Id">
      <div class="orderItemComp">
        <OrderProductComponent :pay-status="payStatus(order.Order.PayStatus)" :order-price="order.Order.Price" :seller="order.Order.Seller" :product="order.Products[0]" :num="order.Products[0].Num" :comment="order.Order.Comment"></OrderProductComponent>
      </div>
    </router-link>
  </div>
</template>

<style scoped>
a {
  text-decoration: none;
}
.orderItemComp {
  border: 1px solid #dee2e6;
  padding: 0.6rem;
  border-radius: 0.5rem;
  margin-bottom: 1rem;
}
</style>
