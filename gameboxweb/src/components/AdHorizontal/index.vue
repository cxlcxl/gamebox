<template>
  <div class="ad-position ad-horizontal">
    <div class="advertisement">Advertisement</div>
    <!-- <component :is="'script'" async src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js?client=ca-pub-6230025462157662"
      crossorigin="anonymous" /> -->
    <ins class="adsbygoogle" style="display:block" :data-ad-client="Settings.AdClientId" :data-ad-slot="Settings.AdSlots[AdSlot]"
      data-ad-format="auto" data-full-width-responsive="true"></ins>
    <component :is="'script'">(adsbygoogle = window.adsbygoogle || []).push({});</component>
  </div>
</template>

<script setup>
import { defineProps } from "vue"
import Settings from "@/setting"

defineProps({
  AdSlot: String,
})
</script>
