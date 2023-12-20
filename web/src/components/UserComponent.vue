<script setup>
import {NdkStore, pubkeyShort} from '../stores/NdkStore.js'
import {onMounted, reactive, ref} from "vue"

const props = defineProps({
  pubkey: {
    type: String,
    required: true
  }
})

const user = reactive({})

onMounted(async () => {
  const profile = await NdkStore().fetchProfile(props.pubkey)
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
})
</script>

<template>
  <div class="userComp">
    <template v-if="user.image != null">
      <img class="userImage img-circle" :src="user.image" width="30" height="30">
    </template>
    <template v-if="user.image == null">
      <img class="userImage" src="@/assets/head.svg" width="30" height="30">
    </template>
    <template v-if="user.name != null">
      <span class="userName">{{user.name}}</span>
    </template>
    <template v-if="user.name == null">
      <span class="userName">{{pubkeyShort(pubkey)}}</span>
    </template>
  </div>
</template>

<style scoped>
</style>
