<template>
  <section class="content-main-w3" id="home">
    <div class="logo-wthree text-center">
      <a class="navbar-brand" href="#">
        {{ds.gameInfo.game.Name}}
      </a>
    </div>

    <div class="container-fluid gallery-lightbox my-2">
      <div class="row m-0 mb-3">
        <div class="col-lg-9 col-md-8 col-sm-12 col-12 p-0">
          <div class="game-detail-ad-bg">
            <div class="game-detail-icon mb-2">
              <img :src="ds.gameInfo.game.Icon">
            </div>
            <div class="game-desc-title p-2">
              <h6>DESCRIPTION</h6>
            </div>
            <div class="game-desc p-2">{{ds.gameInfo.game.Description}}</div>
            <div class="game-play" @click="play(ds.gameInfo.game.GameId)"><a>PLAY NOW</a></div>
            <!-- TODO 广告位 -->
          </div>
        </div>

        <div class="col-lg-3 col-md-4 p-0 snap-img">
          <!-- TODO 广告位 -->
        </div>
      </div>

      <div class="row m-0">

        <div class="col-lg-9 col-md-8 col-sm-12 col-12 p-0">
          <div class="row m-0 may-like-box">
            <div class="col-lg-12 col-md-12 col-sm-12 col-12 p-0 u-may-like">May Like</div>
            <div class="col-lg-4 col-md-6 col-sm-6 col-12 p-0 snap-img" v-for="item, idx in ds.gameInfo.games" :key="idx">
              <GameAdvance :GameInfo="item" />
            </div>
          </div>
        </div>

        <div class="col-lg-3 col-md-4 p-0 snap-img">
          <!-- TODO 广告位 -->
        </div>
      </div>

    </div>

    <div class="ad-position">
      <!-- TODO 广告位 -->
    </div>

  </section>
</template>

<script>
import { defineComponent, onMounted, reactive } from "vue"
import { useRoute, useRouter } from "vue-router"
import GameAdvance from "@/components/GameAdvance/index.vue"
import { getGameInfo } from "@/utils/apis"

export default defineComponent({
  components: { GameAdvance },
  setup() {
    const route = useRoute()
    const router = useRouter()

    const ds = reactive({
      gameInfo: { game: {}, games: [] },
    })

    onMounted(() => {
      gameInfo(route.query.gid)
    })
    const gameInfo = (gid) => {
      if (!/^[A-Z0-9]{10,15}$/.test(gid)) {
        router.push({ name: "Home" })
        return
      }
      getGameInfo({ gid })
        .then((res) => {
          ds.gameInfo = res.data
        })
        .catch((e) => {
          console.log(e)
        })
    }
    const play = (gid) => {
      router.push({ name: "PlayAd", query: { gid } })
    }

    return {
      ds,
      play,
    }
  },
})
</script>
