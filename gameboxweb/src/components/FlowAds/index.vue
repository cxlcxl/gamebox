<template>
  <ins class="adsbygoogle" style="display:block" data-ad-format="fluid" data-ad-layout-key="-h8+8+o-6h+ah"
    :data-ad-client="isGigame ? Settings.AdClientId : Settings.AdClientId1" :data-ad-slot="Settings.AdSlots['flow-ads']"></ins>
  <component :is="'script'">(adsbygoogle = window.adsbygoogle || []).push({});</component>
</template>

<script setup>
import { defineProps } from "vue"
import Settings from "@/setting"
// import { adStore } from "@/store/ad"
// const ad = adStore()

defineProps({
  AdSlot: String,
})

const isGigame = location.host.indexOf("gigame.xyz") !== -1
</script>

<style scoped lang="scss">
</style>