<script setup>
import { RouterLink, RouterView } from 'vue-router'
import FootComponent from './components/FootComponent.vue'
import UIUtil from "./UIUtil.vue";
</script>

<template>
  <UIUtil></UIUtil>
  <div class="appContainer">
    <div class="main">
      <RouterView />
    </div>
    <foot-component></foot-component>
  </div>
</template>

<style scoped>
</style>
