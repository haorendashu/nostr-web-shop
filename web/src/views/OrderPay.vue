<script setup>
import HeadComponent from '../components/HeadComponent.vue'
import UserComponent from "../components/UserComponent.vue";
import {onMounted, ref} from "vue";
import {useRouter} from "vue-router";
import {API} from "../api/api";
import {GetQuery} from "../utils/utils";
import axios from "axios";

const router = useRouter()
const api = new API()
const query = GetQuery()

const pr = ref("lnbc1u1pjc4yxepp5fg7l26vc3pmrcvyhdsmtlqy4lchsra4numh35feck3ultgs55fnqdq5g9kxy7fqd9h8vmmfvdjscqzzsxqyz5vqsp59s2e4fqplwvk7hamm9nrkgnr57pypzt2j8gve6k799xgyurpn2tq9qyyssqjh6t3npfqfdjemtm84lmr925yczs34tqcy5264w8exwxnrr3ulrq9c5ggzp0md9xhd7mh4n4r3dy7repdu8kuynuaf7dq0rrlxsactgpdy0yf4")
const price = ref(0)
const seller = ref("")
const name = ref("")
const code = ref("")

let VerifyUrl = ""
let Pid = ""

let checkTime = 1000 * 10

onMounted(async () => {
  let id = query["id"]
  if (id != null) {
    let result = await api.userOrderGet(id)
    if (result) {
      let order = result.data

      seller.value = order.Pubkey

      let p = 0
      let ns = []
      let cs = []
      for (let index in order.Skus) {
        let sku = order.Skus[index]
        if (sku) {
          ns.push(sku.Name)
          cs.push(sku.Code)

          p += sku.Price * sku.Num
        }
      }

      name.value = ns.join(" ")
      code.value = cs.join(" ")
      price.value = p
    }

    let payInfoResult = await api.userOrderPayGet(id)
    if (payInfoResult) {
      pr.value = payInfoResult.data.Pr

      VerifyUrl = payInfoResult.data.VerifyUrl
      Pid = payInfoResult.data.Id
      if (payInfoResult.data.PayStatus == 1 && payInfoResult.data.VerifyUrl) {
        setTimeout(verifyPayResult, checkTime)
      } else if (payInfoResult.data.PayStatus == 2) {
        handlePushInfo()
      }
    }
  }
})

async function verifyPayResult() {
  if (VerifyUrl) {
    const result = await axios.get(VerifyUrl)
    if (result.status == 200) {
      if (result.data && result.data.settled == true) {
        console.log("success")
        // paid success!
        doAfterPaid()
        return
      }
    }

    // paid unSuccess, check again!
    setTimeout(verifyPayResult, checkTime)
  }
}

async function doAfterPaid() {
  if (Pid && Pid != "") {
    let result = await api.userPayOrderCheck(Pid)
    if (result && result.data) {
      // pay order check success!
      // get push info
      handlePushInfo()
    }
  }
}

async function handlePushInfo() {
  let id = query["id"]
  if (id != null) {
    let pushResult = await api.userOrderPushInfoGet(id)
    console.log(pushResult)
    if (pushResult && pushResult.list && pushResult.list.length > 0) {
      let pushInfo = pushResult.list[0]
      if (pushInfo && pushInfo.PushType == 2 && pushInfo.PushUrl) {
        window.location.href = pushInfo.PushUrl
      }
    }
  }
}
</script>

<template>
  <HeadComponent title="Order Pay"></HeadComponent>
  <div class="orderPay container">
    <template v-if="seller != ''">
      <UserComponent :pubkey="seller"></UserComponent>
    </template>
    <template v-if="name != ''">
      <h2>{{name}}</h2>
    </template>
    <template v-if="code != ''">
      <div class="productCode">{{code}}</div>
    </template>
    <div class="qrCode">
      <vue-qrcode :value="pr" :options="{ width: 200 }"></vue-qrcode>
    </div>
    <div class="productPrice">{{price}} Sats</div>
    <div class="lnlink"><a :href="'lightning:' + pr">{{pr}}</a></div>
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