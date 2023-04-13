<template>
  <section class="content-main-w3" id="home">
    <SearchPage />

    <GameWrap TopicName="Played" :GameList="ds.playedGames" AdSlot="home-horizontal" />
    <GameWrap TopicName="Hot" :GameList="ds.games.Hot" AdSlot="home-horizontal-1" />
    <GameWrap TopicName="New" :GameList="ds.games.New" AdSlot="home-horizontal-2" />
    <GameWrap TopicName="All" :GameList="ds.games.All" ShowStyle="card" AdSlot="home-horizontal-2" />
  </section>
</template>

<script setup>
import { reactive, onMounted } from "vue"
import { gameList } from "@/utils/apis"
import { useRouter } from "vue-router"
import SearchPage from "@/components/Search/index.vue"
import GameWrap from "./components/game-wrap.vue"
import { getStorage } from "@/utils/cache"

const ds = reactive({
  games: { Hot: [], New: [], All: [] },
  playedGames: [],
})
const router = useRouter()
onMounted(() => {
  getGames()
  fetchPlayedGames()
})
const fetchPlayedGames = () => {
  const played = getStorage("played")
  if (Array.isArray(played)) {
    ds.playedGames = played
  }
}
const getGames = () => {
  gameList()
    .then((res) => {
      ds.games = res.data
    })
    .catch((e) => {
      console.log(e)
    })
}
const goMore = (t) => {
  router.push({ name: "More", query: { c: t } })
}
</script>
