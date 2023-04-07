<template>
  <div class="play-game-container">
    <div class="back-home" @click="backHome">
      <img src="/images/home.png" />
    </div>
    <iframe v-show="ds.game.GameUrl !== ''" ref="game" allowfullscreen="true" webkitallowfullscreen="true" mozallowfullscreen="true" @load="loaded"
      :src="ds.game.GameUrl" scrolling="auto" frameborder="0" id="game-iframe">
    </iframe>
  </div>
</template>

<script>
import { defineComponent, reactive, onMounted } from "vue"
import { useRoute, useRouter } from "vue-router"
import { getGameDetail } from "@/utils/apis"

export default defineComponent({
  setup() {
    const ds = reactive({
      more: true,
      game: {},
    })
    const route = useRoute()
    const router = useRouter()

    onMounted(() => {
      gameInfo(route.query.gid)
    })
    const gameInfo = (gid) => {
      if (!/^[A-Z0-9]{10,15}$/.test(gid)) {
        router.push({ name: "Home" })
        return
      }
      getGameDetail(gid)
        .then((res) => {
          ds.game = res.data
        })
        .catch((e) => {
          console.log(e)
        })
    }
    const backHome = () => {
      router.replace({ name: "Home" })
    }
    const loaded = () => {}
    return {
      ds,
      backHome,
      loaded,
    }
  },
})
</script>

<style scoped lang="scss">
.play-game-container {
  position: relative;
  width: 100%;
  height: 100%;
  background: grey;
  overflow: hidden;

  .back-home {
    width: 35px;
    height: 35px;
    position: fixed;
    background: rgba(255, 255, 255, 0.7);
    background-size: cover;
    border-top-right-radius: 5px;
    border-bottom-right-radius: 5px;
    top: 10%;
    left: 0;
    z-index: 9999;
    cursor: pointer;
    border: 1px solid #dedede;

    img {
      position: absolute;
      top: 3px;
      left: 3px;
      width: 25px;
      height: 25px;
    }
  }

  #game-iframe {
    position: relative;
    z-index: 9;
    width: 100%;
    height: 100%;
  }
}
</style>