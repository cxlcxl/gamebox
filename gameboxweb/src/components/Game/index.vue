<template>
  <div class="gallery_grid1 hover08" @click="handleRoute" style="cursor: pointer;">
    <div class="gallery_effect">
      <a>
        <figure><img v-lazy="props.GameInfo.Icon" /></figure>
      </a>
      <div class="game_name">{{props.GameInfo.Name}}</div>
    </div>
  </div>
</template>

<script>
import { defineComponent } from "vue"
import { useRouter } from "vue-router"

export default defineComponent({
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
