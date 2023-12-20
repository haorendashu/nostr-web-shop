import { defineStore } from "pinia"

export const UIStore = defineStore({
    id: 'UIStore',
    state: () => {
        return {
            toastInfos: [],
            modalInfo: null,
        }
    },
    actions: {
        toast(msg) {
            var toastInfo = new UIToastInfo(msg)
            this.toastInfos.push(toastInfo)
            setTimeout(() => {
                const index = this.toastInfos.indexOf(toastInfo)
                if (index > -1) {
                    this.toastInfos.splice(index, 1)
                }
            }, 5000)
        },
        modal(title, body, onConfirm) {
            let mi = new ModalInfo()
            mi.title = title
            mi.body = body
            mi.onConfirm = onConfirm
            this.modalInfo = mi
        },
        modalConfirm() {
            if (this.modalInfo.onConfirm) {
                this.modalInfo.onConfirm()
            }
            this.modalInfo = null;
        },
        modalCancel() {
            this.modalInfo = null;
        }
    }
})

export const UIToastInfo = class {
    msg = "";
    created_at = 0;

    constructor(msg) {
        this.msg = msg
        this.created_at = new Date().getTime()
    }
}

export const ModalInfo = class {
    title = "";
    body = "";
    onConfirm = function() {}
}