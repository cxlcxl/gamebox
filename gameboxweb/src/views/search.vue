<template>
  <section class="content-main-w3" id="home">
    <div class="container-fluid gallery-lightbox my-2">
      <div class="row main-box">
        <div class="col-lg-12 col-md-12 col-sm-12 col-12 p-0 u-may-like">Search Result</div>
        <div class="col-lg-4 col-md-4 col-sm-6 col-12 p-0 snap-img" v-for="item,idx in ds.games" :key="idx">
          <GameAdvance :GameInfo="item" />
        </div>
        <div class="empty" v-show="ds.games.length === 0">Query result is empty</div>
      </div>
    </div>

    <div class="ad-position">
      <!-- TODO 广告位 -->
    </div>

  </section>
</template>
<script>
import { defineComponent, onMounted, reactive } from "vue"
import { searchGames } from "@/utils/apis"
import { useRoute } from "vue-router"
import GameAdvance from "@/components/GameAdvance/index.vue"

export default defineComponent({
  components: { GameAdvance },
  setup() {
    const route = useRoute()
    const ds = reactive({
      games: [],
    })

    onMounted(() => {
      if (route.query.k.trim() !== "") {
        fetchGames(route.query.k.trim())
      }
    })
    const fetchGames = (k) => {
      searchGames({ k })
        .then((res) => (ds.games = res.data))
        .catch((e) => console.log(e))
    }
    return {
      ds,
    }
  },
})
</script>

<style lang="scss" scoped>
.main-box {
  .empty {
    text-align: center;
    color: #606266;
    height: 50px;
    line-height: 50px;
    color: var(--danger);
    text-shadow: 0 1px 2px var(--danger);
  }
}
</style>