<template>
  <CategoryItemSmall
      v-for="(category, index) in categories"
      :key="index"
      :category="category"
  />
</template>

<script>
import CategoryItemSmall from "./CategoryItemSmall.vue";

export default {
  name: "CategorySmallContent",
  components: {
    CategoryItemSmall
  },
  props: {
    categories: {
      type: Array,
      required: true
    }
  }
};
</script>