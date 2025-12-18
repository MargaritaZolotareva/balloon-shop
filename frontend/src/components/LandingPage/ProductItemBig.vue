<template>
  <div class="product-item product-item--big">
    <router-link :to="`/products/${product.id}`" class="img-wrap">
      <div :data-src="getImageUrl(product.photo)" class="product-item__img"
           :style="{ backgroundImage: `url(${getImageUrl(product.photo)})` }">
      </div>
    </router-link>
    <div class="product-item__name">
      <span>{{ product.title }}</span>
    </div>
    <div class="product-item__price">
      <span>{{ product.price }} ₽</span>
    </div>
  </div>
</template>

<script>
export default {
  name: "ProductItemBig",
  props: {
    product: {
      type: Object,
      required: true
    }
  },
  methods: {
    getImageUrl(photoPath) {
      const apiHost = process.env.VUE_APP_API_URL;
      return `${apiHost}/${photoPath}`;
    }
  }
};
</script>

<style scoped lang="scss">
.product-item {
  display: inline-block;
  position: relative;
  cursor: pointer;

  &--big {
    grid-column-start: auto;
    grid-column-end: auto;
    grid-row-start: auto;
    grid-row-end: auto;

    .product-item__img {
      padding-bottom: 100%;
    }
  }

  .img-wrap {
    display: flex;
    overflow: hidden;
    border-radius: $elements_section_radius;
    -webkit-mask-image: -webkit-radial-gradient($white, $black);
    border: .5px solid $photo_border_color;
  }

  &__img {
    position: relative;
    display: inline-block;
    width: 100%;
    height: 0;
    padding-bottom: 100%;
    background-size: cover;
    background-position: center;
    background-color: $photo_background_color;
    color: $gray;

    &:hover {
      transition: transform .3s;
      transform: scale(1.05);
    }
  }

  &__name {
    color: $gray_dark;
    font-size: 14px;
    line-height: 18px;
    padding-top: 12px;
    padding-bottom: 2px;
    -webkit-line-clamp: 1;
    -webkit-box-orient: vertical;
    display: -webkit-box;
    overflow: hidden;
  }

  &__price {
    font-size: 15px;
    line-height: 20px;
    color: $dark;
  }
}

@media (min-width: 1000px) {
  .product-item {
    &--big {
      grid-column-start: 1;
      grid-column-end: 3;
      grid-row-start: 1;
      grid-row-end: 3;

      .product-item__img {
        padding-bottom: 108%;
      }
    }
  }
}
</style>
