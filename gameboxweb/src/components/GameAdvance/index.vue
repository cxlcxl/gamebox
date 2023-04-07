<template>
  <div class="gallery-game">
    <div class="img-box"><img v-lazy="props.GameInfo.Icon" /></div>
    <div class="flex-width">
      <p class="gallery-game__name">{{props.GameInfo.Name}}</p>
      <p class="gallery-game__score">
        <show-stars :star="props.GameInfo.Star" />
      </p>
    </div>
    <div class="gallery-game__play">
      <span class="play-btn" style="cursor: pointer;" @click="handleRoute">PLAY</span>
    </div>
  </div>
</template>

<script>
import { defineComponent } from "vue"
import ShowStars from "../Stars/index.vue"
import { useRouter } from "vue-router"

export default defineComponent({
  components: { ShowStars },
  props: {
    GameInfo: {
      type: Object,
      required: true,
    },
  },
  setup(props) {
    const router = useRouter()
    const handleRoute = () => {
      router.push({ name: "GameInfo", query: { gid: props.GameInfo.GameId } })
    }
    return {
      props,
      handleRoute,
    }
  },
})
</script>
