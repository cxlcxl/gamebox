<template>
  <template v-for="item, idx in GameList" :key="idx">
    <div class="col-lg-4 col-md-4 col-sm-6 col-12 p-0 snap-img">
      <div class="gallery-game">
        <div class="img-box"><img v-lazy="item.icon" /></div>
        <div class="flex-width">
          <p class="gallery-game__name">{{item.name}}</p>
          <p class="gallery-game__score">
            <show-stars :star="item.star" />
          </p>
        </div>
        <div class="gallery-game__play">
          <span class="play-btn" style="cursor: pointer;" @click="handleRoute(item.game_id)">PLAY</span>
        </div>
      </div>
    </div>

    <div class="col-lg-4 col-md-4 col-sm-6 col-12 p-0 snap-img" v-if="idx > 0 && (idx % Setting.FlowAdShowRate) === 0">
      <FlowAds />
    </div>
  </template>
</template>

<script setup>
import { defineProps } from "vue"
import ShowStars from "../Stars/index.vue"
import FlowAds from "../FlowAds/index.vue"
import { useRouter } from "vue-router"
import Setting from "@/setting"

const props = defineProps({
  GameList: {
    type: Object,
    required: true,
  },
})

const router = useRouter()
const handleRoute = (gid) => {
  router.push({ name: "GameInfo", query: { gid } })
}
</script>
