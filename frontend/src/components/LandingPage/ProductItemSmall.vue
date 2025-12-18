<template>
  <div class="product-item">
    <router-link :to="`/products/${product.id}`" div class="img-wrap">
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
  name: "ProductItemSmall",
  props: {
    product: Object
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

  .hover-content {
    position: absolute;
    top: 15px;
    right: 15px;
    left: 15px;
    bottom: 15px;
    text-align: center;
    background-color: rgba(42, 42, 42, .97);
    opacity: 0;
    visibility: hidden;
    transition: all 0.5s;

    .inner {
      position: absolute;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
      width: 100%;
    }

    h4 {
      color: #fff;
      font-size: 24px;
      font-weight: 700;
      margin-bottom: 15px;
    }

    p {
      color: #fff;
      padding: 0px 20px;
      margin-bottom: 20px;
    }

    .main-border-button a {
      color: #fff;
      text-decoration: none;
      border: 2px solid #fff;
      padding: 10px 20px;
      font-weight: 600;
      border-radius: 30px;
      transition: background-color 0.3s ease;

      &:hover {
        background-color: #fff;
        color: #333;
      }
    }
  }

  &:hover .hover-content {
    opacity: 1;
    visibility: visible;
  }
}
</style>
