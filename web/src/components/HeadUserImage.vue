<script setup>
import {NdkStore} from '../stores/NdkStore.js'
import {onMounted, reactive, watch} from "vue"

const props = defineProps({
  pubkey: {
    type: String,
    required: false,
  }
})

const user = reactive({})

onMounted(async () => {
  if (props.pubkey && props.pubkey != "") {
    await updateProfile(props.pubkey)
  }
})

async function updateProfile(_pubkey) {
  const profile = await NdkStore().fetchProfile(_pubkey)
  if (profile) {
    user.about = profile.about
    user.banner = profile.banner
    user.displayName = profile.displayName
    user.image = profile.image
    user.lud16 = profile.lud16
    user.name = profile.name
    user.nip05 = profile.nip05
    user.website = profile.website
  }
}

watch(props, async (newProps, oldProps) => {
  if (newProps.pubkey != oldProps.pubkey) {
    await updateProfile(newProps.pubkey)
  }
})
</script>

<template>
  <div class="userImageComp">
    <template v-if="user.image != null">
      <img class="userImage circle" :src="user.image">
    </template>
    <template v-if="user.image == null">
      <img class="circle" src="@/assets/head.svg">
    </template>
  </div>
</template>

<style scoped>
</style>
