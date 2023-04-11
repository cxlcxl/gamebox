<template>
  <div class="search-container">
    <img class="icon-search" src="/images/search.png" />
    <input type="text" placeholder="Search Games" v-model="search.k" @blur="doSearch" />
  </div>
</template>

<script>
import { defineComponent, onMounted, reactive } from "vue"
import { useRoute, useRouter } from "vue-router"

export default defineComponent({
  props: {},
  setup() {
    const search = reactive({
      k: "",
    })
    const route = useRoute()
    const router = useRouter()

    onMounted(() => {
      if (route.query.k) {
        search.k = route.query.k.trim()
      }
    })

    const doSearch = () => {
      if (search.k.trim() !== "") {
        router.push({ name: "Search", query: { k: search.k } })
      }
    }
    return {
      search,
      doSearch,
    }
  },
})
</script>
