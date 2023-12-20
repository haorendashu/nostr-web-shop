<script setup>
import HeadComponent from '../../components/HeadComponent.vue'
import { useRouter } from 'vue-router'
import {onMounted, reactive, ref} from "vue"
import {API} from "../../api/api"
import {GetQuery} from "../../utils/utils";
import {UIStore} from "../../stores/UIStore";

const router = useRouter()
const api = new API()
const query = GetQuery()
const uiStore = UIStore()

let product = {}

const productName = ref("")
const skus = reactive([{
  "Code": "",
  "Price": 0,
  "Stock": 0,
}])
const imgs = reactive([])
const productLNWallet = ref("")
const productDescription = ref("")
const productContent = ref("")

function addImage() {
  imgs.push("https://via.assets.so/img.jpg?w=200&h=200&tc=white&bg=#cecece")
}

var quill = null
function onEditorRealy(_quill) {
  quill = _quill
  editor.value.setHTML(productContent.value)
}

const editor = ref(null)
async function productSave() {
  const content = editor.value.getHTML()
  productContent.value = content

  if (!productName.value) {
    uiStore.toast("Product Name can't be null")
    return
  }

  if (!skus[0].Code) {
    uiStore.toast("Product Code can't be null")
    return
  }
  if (!skus[0].Price && skus[0].Price < 1) {
    uiStore.toast("Product Price can't be null")
    return
  }
  if (!skus[0].Stock && skus[0].Stock < 1) {
    uiStore.toast("Product Stock can't be null")
    return
  }
  if (imgs.length <= 0) {
    uiStore.toast("Image can't be null")
    return
  }

  if (!productLNWallet.value) {
    uiStore.toast("LNWallet can't be null")
    return
  }
  if (productLNWallet.value.indexOf("@getalby.com") < 0) {
    uiStore.toast("LNWallet must be getalby address")
    return
  }

  let id = query["id"]
  if (id) {
    product.Id = id - 1 + 1
  }
  product.Name = productName.value
  product.Des = productDescription.value
  product.Content = productContent.value
  product.Lnwallet = productLNWallet.value
  product.Imgs = imgs.join(",")
  product.Skus = skus

  let maxPrice = 0
  let skusLength = skus.length
  for (let i = 0; i < skusLength; i++) {
    let sku = skus[i]
    if (sku["Price"] > maxPrice) {
      maxPrice = sku.Price
    }
  }
  product.Price = maxPrice

  const result = await api.shopProductAdd(product)
  if (result && result["code"] == 200 && result["pid"]) {
    router.replace("/seller/productPushEdit?id=" + result["pid"])
  }
}

onMounted(async () => {
  var id = query["id"]
  if (id != null) {
    var result = await api.shopProductGet(id)
    if (result) {
      let p = result.data

      productName.value = p.Name
      imgs.splice(0, imgs.length)
      if (p.Imgs) {
        let imgStrs = p.Imgs.split(",")
        for (let i in imgStrs) {
          imgs.push(imgStrs[i])
        }
      }
      productLNWallet.value = p.Lnwallet
      productDescription.value = p.Des
      productContent.value = p.Content

      skus.splice(0, skus.length)
      if (p.Skus != null) {
        for (let index in p.Skus) {
          skus.push(p.Skus[index])
        }
      }

      if (quill != null) {
        editor.value.setHTML(p.Content)
      }
    }
  }
})
</script>

<template>
  <div class="productEdit">
    <HeadComponent title="Product Edit"></HeadComponent>
    <div class="container">
      <div class="row m_t_05">
        <label for="Name-url" class="form-label">Name</label>
        <div class="input-group">
          <input type="text" class="form-control" id="Name-url" v-model="productName">
        </div>
      </div>

      <template v-for="sku in skus">
        <div class="row m_t_05">
          <div class="col">
            <label for="Product_Code" class="form-label">Product Code</label>
            <div class="input-group">
              <input type="text" class="form-control" id="Product_Code" v-model="sku.Code">
            </div>
          </div>
          <div class="col">
            <label for="Price" class="form-label">Price</label>
            <div class="input-group">
              <input type="number" class="form-control" id="Price" v-model.number="sku.Price">
              <span class="input-group-text" id="Price-addon2">Sats</span>
            </div>
          </div>
        </div>

        <div class="row m_t_05">
          <div class="col">
            <label for="Stock" class="form-label">Stock</label>
            <div class="input-group">
              <input type="number" class="form-control" id="Stock" v-model.number="sku.Stock">
            </div>
          </div>
          <div class="col">
          </div>
        </div>
      </template>

      <div class="row m_t_05 m_b_05">
        <label class="form-label">Images</label>
        <div class="productEditImages">
          <div class="imgBox m_r_05" v-on:click="addImage">
            <i class="bi bi-plus"></i>
          </div>
          <div class="imgBox m_r_05" v-for="img in imgs">
            <img :src="img">
          </div>
        </div>
      </div>

      <div class="row m_t_05">
        <label for="LNWallet" class
            ="form-label">LNWallet</label>
        <div class="input-group">
          <input type="text" class="form-control" id="LNWallet" v-model="productLNWallet">
        </div>
      </div>

      <div class="row m_t_05">
        <label for="Product_Description" class="form-label">Product Description</label>
        <div class="input-group">
          <textarea  type="text" class="form-control" id="Product_Description"  v-model="productDescription"></textarea>
        </div>
      </div>

      <div class="row m_t_05 productContentEditor">
        <label class="form-label">Product Content</label>
        <div>
          <QuillEditor ref="editor" theme="snow" toolbar="#myToolbar" :onReady="onEditorRealy">
            <template #toolbar>
              <div id="myToolbar">
                <button class="ql-header" value="1"></button>
                <button class="ql-header" value="2"></button>
                <button class="ql-list" value="ordered"></button>
                <button class="ql-list" value="bullet"></button>

                <button class="ql-bold"></button>
                <button class="ql-italic"></button>
                <button class="ql-underline"></button>
                <button class="ql-strike"></button>

                <button class="ql-blockquote"></button>
                <button class="ql-code-block"></button>

                <button class="ql-image"></button>
                <select class="ql-color"></select>
                <select class="ql-background"></select>
                <button class="ql-clean"></button>
              </div>
            </template>
          </QuillEditor>
        </div>
      </div>

      <div class="row" style="margin-top: 5rem; padding: 0 0.8rem;">
        <button type="button" class="btn btn-dark" v-on:click="productSave">Save</button>
      </div>

    </div>
  </div>
</template>

<style scoped>
.imgBox {
  width: 5rem;
  height: 5rem;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: grey;
  font-size: 2rem;
  font-weight: bold;
  color: white;
}
.productEditImages {
  display: flex;
  align-items: center;
}
.imgBox img {
  width: 100%;
  height: 100%;
}
.productContentEditor {
  margin-bottom: 3rem;
}
@media (max-width: 600px) {
}
</style>