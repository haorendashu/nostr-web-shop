<script setup>

import {UIStore} from "./stores/UIStore";
const uiStore = UIStore()
</script>

<template>
  <template v-if="uiStore.modalInfo != null">
    <div class="modal" style="display: block;">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h1 class="modal-title fs-5" id="exampleModalLabel">{{uiStore.modalInfo.title}}</h1>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="modal-body">
            <p>{{uiStore.modalInfo.body}}</p>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary" data-bs-dismiss="modal" v-on:click="uiStore.modalCancel">Cancel</button>
            <button type="button" class="btn btn-primary" v-on:click="uiStore.modalConfirm" >Confirm</button>
          </div>
        </div>
      </div>
    </div>
  </template>

  <div class="toast-container">
    <template v-for="toastInfo in uiStore.toastInfos">
      <div class="toast show" role="alert" aria-live="assertive" aria-atomic="true" style="display: block;">
        <div class="toast-header">
          <strong class="me-auto">NWS</strong>
        </div>
        <div class="toast-body">
          {{toastInfo.msg}}
        </div>
      </div>
    </template>
  </div>
</template>

<style scoped>
.toast-container {
  left: 50%;
  transform: translate(-50%, 0);
  top: 4rem;
}
</style>
