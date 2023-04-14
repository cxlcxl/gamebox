<template>
  <div class="ad-position ad-horizontal" v-if="ad.showCustomAd">
    <div class="advertisement">Advertisement</div>
    <ins class="adsbygoogle" style="display:block" :data-ad-client="isGigame ? Settings.AdClientId : Settings.AdClientId1"
      :data-ad-slot="Settings.AdSlots[AdSlot]" data-ad-format="auto" data-full-width-responsive="true"></ins>
    <component :is="'script'">(adsbygoogle = window.adsbygoogle || []).push({});</component>
  </div>
</template>

<script setup>
import { defineProps } from "vue"
import Settings from "@/setting"
import { adStore } from "@/store/ad"
const ad = adStore()

defineProps({
  AdSlot: String,
})

const isGigame = location.host.indexOf("gigame.xyz") !== -1
</script>
