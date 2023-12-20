import { defineStore } from "pinia"
import { NdkStore } from "./NdkStore.js"

export const UserStore = defineStore({
    id: 'UserStore',
    state: () => {
        return {
            logined: false,
            token: "",
            ndkStore: NdkStore(),
        }
    },
    actions: {
        async checkToken() {
            if (!this.token || this.token == "") {
                const storage = window.localStorage
                const _token = storage["token"]
                this.token = _token
            }

            return this.token
        },
        async checkLogin() {
            const _token = await this.checkToken()
            if (_token) {
                return true
            } else {
                return false
            }
        },
        setToken(t) {
            this.token = t
            const storage = window.localStorage
            storage["token"] = t
        }
    }
})