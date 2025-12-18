<template>
  <ProductItemSmall
      v-for="(product, index) in products"
      :key="index"
      :product="product"
  />
</template>

<script>
import ProductItemSmall from "./ProductItemSmall.vue";

export default {
  name: "ProductSmallContent",
  components: {
    ProductItemSmall
  },
  props: {
    products: {
      type: Array,
      required: true
    }
  }
};
</script>