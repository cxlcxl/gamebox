<template>
  <img class="star" :src="i <= props.star ? '/images/star-fill.png' : '/images/star.png'" v-for="i, idx1 in stars" :key="idx1" />
</template>

<script>
import { defineComponent } from "vue"

export default defineComponent({
  props: {
    star: {
      type: Number,
      required: true,
    },
  },
  setup(props) {
    const stars = [1, 2, 3, 4, 5]
    return {
      props,
      stars,
    }
  },
})
</script>

<style scoped>
.star {
  width: 15px;
  height: 15px;
}
</style>