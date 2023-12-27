import {UIStore} from "../stores/UIStore";
import axios from "axios";
import {UserStore} from "../stores/UserStore";
import G from "../utils/globals";
import {useRouter} from "vue-router";
import {encodeURI} from "js-base64";

export const API = class {
    uiStore = null;
    userStore = null;
    router = null;

    constructor() {
        this.uiStore = UIStore()
        this.userStore = UserStore()
        this.router = useRouter()
    }

    resultCheck = function(result) {
        if (result.status != 200) {
            this.uiStore.toast("http error " + result.status)
            return false
        }

        const resultData = result.data
        if (resultData["code"] != 200) {
            const msg = resultData["msg"]
            if (msg) {
                this.uiStore.toast(msg)
            }
            if (resultData["code"] == 403) {
                this.router.push("/login?u=" + encodeURIComponent(location.href))
            }
            return false
        }

        return true
    }

    baseHttpConfig = async function() {
        const _token = await this.userStore.checkToken()
        return {
            headers: {
                "token": _token,
            }
        }
    }

    baseProductList = async function () {
        const result = await axios.get(G.BasePath + "/base/product/list", await this.baseHttpConfig())

        if (!this.resultCheck(result)) {
            return
        }

        return result.data
    }

    baseProductGet = async function (id) {
        const result = await axios.get(G.BasePath + "/base/product/" + id, await this.baseHttpConfig())

        if (!this.resultCheck(result)) {
            return
        }

        return result.data
    }

    userOrderAdd = async function (order) {
        const result = await axios.post(G.BasePath + "/user/order/add", order, await this.baseHttpConfig())

        if (!this.resultCheck(result)) {
            return
        }

        return result.data
    }

    userOrderGet = async function (id) {
        const result = await axios.get(G.BasePath + "/user/order/" + id, await this.baseHttpConfig())

        if (!this.resultCheck(result)) {
            return
        }

        return result.data
    }

    userOrderPayGet = async function (id) {
        const result = await axios.get(G.BasePath + "/user/orderPay/" + id, await this.baseHttpConfig())

        if (!this.resultCheck(result)) {
            return
        }

        return result.data
    }

    userOrderList = async function () {
        const result = await axios.get(G.BasePath + "/user/order/list", await this.baseHttpConfig())

        if (!this.resultCheck(result)) {
            return
        }

        return result.data
    }

    shopProductAdd = async function (product) {
        const result = await axios.post(G.BasePath + "/shop/product/", product, await this.baseHttpConfig())

        if (!this.resultCheck(result)) {
            return
        }

        return result.data
    }

    shopProductGet = async function (id) {
        const result = await axios.get(G.BasePath + "/shop/product/" + id, await this.baseHttpConfig())

        if (!this.resultCheck(result)) {
            return
        }

        return result.data
    }

    shopProductDel = async function (id) {
        const result = await axios.delete(G.BasePath + "/shop/product/" + id, await this.baseHttpConfig())

        if (!this.resultCheck(result)) {
            return
        }

        return result.data
    }

    shopProductList = async function (id) {
        const result = await axios.get(G.BasePath + "/shop/product/list", await this.baseHttpConfig())

        if (!this.resultCheck(result)) {
            return
        }

        return result.data
    }

    shopProductPushInfoGet = async function (id) {
        const result = await axios.get(G.BasePath + "/shop/productPushInfo/" + id, await this.baseHttpConfig())

        if (!this.resultCheck(result)) {
            return
        }

        return result.data
    }

    shopProductPushInfoSave = async function (info) {
        const result = await axios.post(G.BasePath + "/shop/productPushInfo/", info, await this.baseHttpConfig())

        if (!this.resultCheck(result)) {
            return
        }

        return result.data
    }
}