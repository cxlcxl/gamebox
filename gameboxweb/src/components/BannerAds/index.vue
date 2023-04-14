<template>
  <div class="banner">
    <ins class="adsbygoogle" style="display:inline-block;width:320px;height:50px"
      :data-ad-client="isGigame ? Setting.AdClientId : Setting.AdClientId1" :data-ad-slot="Setting.AdSlots['amp-banner']"></ins>
    <component :is="'script'">
      (adsbygoogle = window.adsbygoogle || []).push({});
    </component>
  </div>
</template>

<script setup>
import Setting from "@/setting"
const isGigame = location.host.indexOf("gigame.xyz") !== -1
</script>

<style scoped lang="scss">
.banner {
  position: fixed;
  width: 320px;
  height: 50px;
  bottom: 0;
  left: calc(50% - 160px);
  overflow: hidden;
}
</style>