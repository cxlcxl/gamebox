<template>
  <div class="gallery_grid1 hover08" @click="handleRoute" style="cursor: pointer;">
    <div class="gallery_effect">
      <a>
        <figure><img v-lazy="props.GameInfo.icon" /></figure>
      </a>
      <div class="game_name">{{props.GameInfo.name}}</div>
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
      router.push({ name: "GameInfo", query: { gid: props.GameInfo.game_id } })
    }
    return {
      props,
      handleRoute,
    }
  },
})
</script>
