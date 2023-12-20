<script setup>
const props = defineProps({
  price: {
    price: Number,
    required: true,
  }
})
</script>

<template>
  <div class="priceComp">{{props.price}} Sats</div>
</template>

<style scoped>
.priceComp {
  display: inline-block;
  background-color: orange;
  padding: 0.2rem 0.4rem;
  border-radius: 0.3rem;
  color: white;
  font-size: 0.85rem;
}
</style>
