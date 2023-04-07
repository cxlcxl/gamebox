<template>
  <div class="layout-container">
    <div class="slide-bar">
      <slide-bar />
    </div>

    <div class="main-container">
      <div class="main-body">
        <router-view v-slot="{ Component, route }">
          <component :is="Component" :key="route.fullPath" />
        </router-view>
      </div>
    </div>

    <footer-page />
  </div>
</template>

<script setup>
import SlideBar from "./slide-bar.vue"
import FooterPage from "./footer.vue"
</script>

<style>
.layout-container {
  padding: 0 0 20px 0;
}
</style>