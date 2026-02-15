<template>
  <CategoryItemSmall
      v-for="(category, index) in categories"
      :key="index"
      :category="category"
  />
</template>

<script setup>
import CategoryItemSmall from "./CategoryItemSmall.vue";

defineProps({
  categories: {
    type: Array,
    required: true
  }
})
</script>