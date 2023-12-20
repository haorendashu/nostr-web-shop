
<script setup>
import UserComponent from "./UserComponent.vue";
import {onMounted, ref} from "vue";

const props = defineProps({
  product: {
    type: Object,
    required: true,
  },
  orderStatus: {
    type: String,
    required: false
  },
  orderPrice: {
    type: Number,
    required: false
  },
  num: {
    type: Number,
    required: false
  },
  seller: {
    type: String,
    required: false
  }
})

const img = ref("")
const code = ref("")
const num = ref(1)
const title = ref("")
const price = ref(1)

onMounted(async () => {
  let imgs = props.product.Imgs.split(",")
  if (imgs != null && imgs.length > 0) {
    img.value = imgs[0]
  }

  if (props.product.Skus != null && props.product.Skus.length > 0) {
    code.value = props.product.Skus[0].Code

    if (props.num) {
      num.value = props.num
    } else {
      num.value = props.product.Skus[0].Stock
    }
  }

  title.value = props.product.Name
  price.value = props.product.Price
})

</script>

<template>
  <div class="orderProduct">
    <template v-if="seller">
      <div class="orderProductTop">
        <UserComponent :pubkey="props.seller"></UserComponent>
        <template v-if="orderStatus">
          <div class="orderStatus">
            {{orderStatus}}
          </div>
        </template>
      </div>
    </template>
    <div class="orderProductInfo">
      <img :src="img">
      <div class="orderProductInfoR">
        <div class="orderProductInfoRT">
          <h2>{{title}}</h2>
          <div class="productPrice">{{price}}</div>
        </div>
        <div class="orderProductInfoRC">
          <div class="orderCode">
            {{code}}
          </div>
          <div class="productNum">
            X {{num}}
          </div>
        </div>
      </div>
    </div>
    <template v-if="orderPrice">
      <div class="orderProductBottom">
        Total {{orderPrice}}
      </div>
    </template>
  </div>
</template>

<style scoped>
.orderProductTop {
  margin-bottom: 0.5rem;
  display: flex;
  align-items: center;
}
.orderStatus {
  margin-left: auto!important;
}
.orderProductInfo {
  display: flex;
  width: 100%;
}
.orderProductInfo img {
  width: 6rem;
  height: 6rem;
}
.orderProductInfoR {
  margin-left: 0.8rem;
  width: calc(100% - 6rem);
}
.orderProductInfoRT h2 {
  font-size: 1.4rem;
  overflow: hidden;
  max-height: 3.6rem;
}
.orderProductInfoRT {
  display: flex;
}
.productPrice {
  font-weight: bold;
  padding-left: 0.5rem;
  margin-left: auto!important;
}
.orderProductInfoRC {
  display: flex;
  color: #999999;
}
.productNum {
  margin-left: auto!important;
}
.orderProductBottom {
  display: flex;
  justify-content: end;
}
</style>
