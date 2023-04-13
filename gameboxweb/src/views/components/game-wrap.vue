<template>
  <template v-if="GameList.length > 0">
    <div class="container-fluid gallery-lightbox my-2">
      <div class="row main-box">
        <div class="col-lg-12 col-md-12 col-sm-12 col-12 p-0 u-may-like">
          <span class="title">{{TopicName}} Games</span>
          <span class="action" @click="goMore">MORE ></span>
        </div>
        <template v-if="ShowStyle === 'img'">
          <div class="col-lg-1 col-md-2 col-sm-3 col-4 p-0 snap-img" v-for="item, idx in GameList" :key="idx">
            <GameBlock :GameInfo="item" v-if="idx <= 11" />
          </div>
        </template>
        <template v-else>
          <div class="col-lg-4 col-md-4 col-sm-6 col-12 p-0 snap-img" v-for="item, idx in GameList" :key="idx">
            <GameAdvance :GameInfo="item" v-if="idx <= 23" />
          </div>
        </template>
      </div>
    </div>

    <DisplayAds :AdSlot="AdSlot" />
  </template>
</template>

<script setup>
import { defineProps } from "vue"
import { useRouter } from "vue-router"
import DisplayAds from "@/components/DisplayAds/index.vue"
import GameBlock from "@/components/Game/index.vue"
import GameAdvance from "@/components/GameAdvance/index.vue"

const router = useRouter()
const props = defineProps({
  ShowStyle: {
    type: String,
    default: "img",
  },
  AdSlot: {
    type: String,
    required: true,
  },
  TopicName: {
    type: String,
    required: true,
  },
  GameList: {
    type: Array,
    required: true,
    default: () => [],
  },
})

const goMore = () => {
  router.push({ name: "More", query: { c: props.TopicName } })
}
</script>
