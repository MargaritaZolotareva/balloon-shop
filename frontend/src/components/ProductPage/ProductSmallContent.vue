<template>
  <ProductItemSmall
      v-for="(product, index) in products"
      :key="index"
      :product="product"
  />
</template>

<script setup>
import ProductItemSmall from "./ProductItemSmall.vue";

defineProps({
  products: {
    type: Array,
    required: true
  }
})
</script>