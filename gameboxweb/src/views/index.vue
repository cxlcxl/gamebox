<template>
  <section class="content-main-w3" id="home">
    <div class="container-fluid gallery-lightbox my-2">
      <div class="row main-box">
        <div class="col-lg-12 col-md-12 col-sm-12 col-12 p-0 u-may-like">Hot Games</div>
        <div class="col-lg-1 col-md-2 col-sm-3 col-4 p-0 snap-img" v-for="item, idx in ds.games.Hot" :key="idx">
          <GameBlock :GameInfo="item" />
        </div>
      </div>
    </div>

    <div class="ad-position">
      <!-- TODO 广告位 -->
    </div>

    <div class="container-fluid gallery-lightbox my-2">
      <div class="row main-box">
        <div class="col-lg-12 col-md-12 col-sm-12 col-12 p-0 u-may-like">New Games</div>
        <div class="col-lg-1 col-md-2 col-sm-3 col-4 p-0 snap-img" v-for="item, idx in ds.games.New" :key="idx">
          <GameBlock :GameInfo="item" />
        </div>
      </div>
    </div>

    <div class="ad-position">
      <!-- TODO 广告位 -->
    </div>

    <div class="container-fluid gallery-lightbox my-2">
      <div class="row main-box">
        <div class="col-lg-12 col-md-12 col-sm-12 col-12 p-0 u-may-like">All Games</div>
        <div class="col-lg-4 col-md-4 col-sm-6 col-12 p-0 snap-img" v-for="item, idx in ds.games.All" :key="idx">
          <GameAdvance :GameInfo="item" />
        </div>
      </div>
    </div>

    <div class="ad-position">
      <!-- TODO 广告位 -->
    </div>
  </section>

</template>


<script>
import { defineComponent, reactive, onMounted } from "vue"
import { gameList } from "@/utils/apis"
import GameAdvance from "@/components/GameAdvance/index.vue"
import GameBlock from "@/components/Game/index.vue"
import { getStorage, setStorage } from "@/utils/cache"

export default defineComponent({
  components: { GameBlock, GameAdvance },
  setup() {
    const ds = reactive({
      games: { Hot: [], New: [], All: [] },
      stars: [1, 2, 3, 4, 5],
    })
    onMounted(() => {
      let d = getStorage("games")
      if (d) {
        ds.games = d
      } else {
        getGames()
      }
    })
    const getGames = () => {
      gameList()
        .then((res) => {
          ds.games = res.data
          setStorage("games", res.data)
        })
        .catch((e) => {
          console.log(e)
        })
    }
    return {
      ds,
    }
  },
})
</script>
