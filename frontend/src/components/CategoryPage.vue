<template>
  <div v-if="category_title" class="CategoryPage">
    <button class="Button Button--rounded Button--defaultSize CategoryPage__closeButton" @click="goBack">
      <i class="far fa-arrow-left CloseButtonIcon"></i>
    </button>
    <h2 class="Block__title Block__title--default">
      <div class="Block__titleWrap">{{ category_title }}</div>
    </h2>
    <div class="Block__content">
      <div class="items-wrap" v-if="products.length > 0">
        <ProductSmallContent :products="products"/>
      </div>
      <div v-else>
        <p>Загрузка...</p>
      </div>
    </div>
  </div>
  <div v-else class="loading">Загрузка...</div>
</template>

<script>
import axios from 'axios';
import ProductSmallContent from "@/components/LandingPage/ProductSmallContent.vue";


export default {
  name: 'CategoryPage',
  components: {ProductSmallContent},
  data() {
    return {
      category_title: "",
      products: []
    }
  },
  mounted() {
    const categoryId = this.$route.params.id;
    this.fetchCategoryProducts(categoryId)
  },
  methods: {
    async fetchCategoryProducts(categoryId) {
      try {
        const apiUrl = process.env.VUE_APP_API_URL;
        const response = await axios.get(`${apiUrl}/categories/${categoryId}/products`);
        const data = await response.data;
        this.products = data.products;
        this.category_title = data.category_title;
      } catch (error) {
        console.error("Ошибка при загрузке товаров категории:", error);
      }
    },
    goBack() {
      this.$router.back();
    },
  },
};
</script>

<style scoped lang="scss">
.CategoryPage {
  display: flex;
  flex-direction: column;
  padding: 100px 0 30px 0;

  &__photo {
    width: 100vw;
    height: 100vw;
    background-size: contain;
    background-position: center;
    background-repeat: no-repeat;
    background-color: $product_cart_photo_background_color;
    cursor: -webkit-zoom-in;
    cursor: zoom-in;
  }

  &__content {
    padding: 16px 24px 24px;
  }

  &__buttons {
    display: grid;
    grid-auto-flow: column;
    grid-template-columns: 1fr 50px;
    gap: 10px;
    padding: 32px 0;
  }

  &__name {
    color: $black;
    font-size: 20px;
    line-height: 28px;
    margin-bottom: 16px;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
    display: -webkit-box;
    overflow: hidden;
  }

  &__descTitle {
    font-weight: 500;
    -webkit-font-smoothing: subpixel-antialiased;
    -moz-osx-font-smoothing: auto;
    font-size: 14px;
    line-height: 18px;
    padding-bottom: 8px;
  }

  &__desc {
    display: -webkit-box;
    -webkit-box-orient: vertical;
    overflow: hidden;
    font-size: 14px;
    line-height: 18px;
    color: $gray_dark;
  }

  &__photoSliderIndexButtons {
    width: 100vw;
    bottom: 8px;
    position: absolute;
    left: 0;
    right: 0;
  }
  &__closeButton {
    position: absolute;
    top: 87px;
    left: -5px;
    display: inline-flex;
    align-items: center;
    justify-content: center;
    background: 0 0;
    width: 56px;
    height: 56px;
    cursor: pointer;
    z-index: 1;
  }
  &__price {
    display: inline-flex;
    align-items: center;
    font-weight: 500;
    -webkit-font-smoothing: subpixel-antialiased;
    -moz-osx-font-smoothing: auto;
    font-size: 28px;
  }
}

.ProductItem {
  display: inline-block;
  position: relative;
  cursor: pointer;

  &__photoWrap {
    display: flex;
    overflow: hidden;
    border-radius: $elements_section_radius;
    -webkit-mask-image: -webkit-radial-gradient($white, $black);
    border: .5px solid $photo_border_color;
  }

  &__name {
    color: $gray_dark;
    font-size: 14px;
    line-height: 18px;
    padding-top: 12px;
    padding-bottom: 2px;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
    display: -webkit-box;
    overflow: hidden;
  }

  &__price {
    font-size: 15px;
    line-height: 20px;
    color: $dark;
  }

  &__photo {
    position: relative;
    display: inline-block;
    width: 100%;
    height: 0;
    padding-bottom: 100%;
    background-size: cover;
    background-position: center;
    background-color: $photo_background_color;
    color: $gray;
  }
}

.ActionButton {
  &--productCartButton {
    width: 100%;
  }
}

.Link {
  &--buttonPrimary {
    color: $white;
    background: $landing_button_primary_background;
  }

  &--button, &--buttonPrimary {
    display: inline-block;
    position: relative;
    padding: 10.5px 16px 11.5px;
    height: 44px;
    outline: 0;
    border-width: 0;
    border-radius: $elements_section_radius;
    font-weight: 500;
    -webkit-font-smoothing: subpixel-antialiased;
    -moz-osx-font-smoothing: auto;
    font-size: 16px;
    cursor: pointer;
    text-align: center;
    vertical-align: middle;
  }
}
.items-wrap {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 32px 8px;
  padding: 0 24px;
}

@media (min-width: 1000px) {
  .CategoryPage {
    padding: 120px 100px 50px 100px;

    &__photo {
      display: inline-block;
      width: 560px;
      height: auto;
      min-height: 100%;
    }

    &__content {
      display: inline-block;
      //width: 388px;
      padding: 48px 48px 40px 48px;
    }

    &__name {
      color: $black;
      font-size: 20px;
      line-height: 28px;
      margin-bottom: 16px;
      -webkit-line-clamp: 2;
      -webkit-box-orient: vertical;
      display: -webkit-box;
      overflow: hidden;
    }

    &__price {
      display: inline-flex;
      align-items: center;
      font-weight: 500;
      -webkit-font-smoothing: subpixel-antialiased;
      -moz-osx-font-smoothing: auto;
      font-size: 28px;
    }

    &__buttons {
      display: flex;
      gap: initial;
      padding: 40px 0;
    }

    &__photoSlider {
      width: 560px;
    }

    &__photoSliderTouchableArea {
      height: 100%;
      padding: 0;
    }

    &__photoSliderItemsWrap {
      height: 100%;
    }

    &__photoSliderItem {
      height: 100%;
    }

    &__photoSliderItemContent {
      height: 100%;
    }

    &__photoSliderIndexButtons {
      display: none;
    }

    &__photoSliderNavButtonWrap {
      top: 0;
      bottom: 0;
      margin-top: 0;
      margin-left: 0;

      &--left {
        left: 0;
      }

      &--right {
        right: 0;
      }
    }
    &__closeButton {
      position: absolute;
      top: 115px;
      left: 25px;
      display: inline-flex;
      align-items: center;
      justify-content: center;
      background: 0 0;
      width: 56px;
      height: 56px;
      cursor: pointer;
      z-index: $second-layer-z-index;
    }
    .Block {
      &__title {
        text-align: center;
      }
    }
  }

  .ActionButton {
    &--productCartButton {
      display: flex;
      align-items: center;
      width: auto;
      height: 44px;
      padding-top: 0;
      padding-bottom: 0;
    }
  }

  .ProductItem {
    &__photo {
      padding-bottom: 96%;
      transition: transform .3s;
      transform: scale(1);
    }

    &:hover .ProductItem__photo {
      transform: scale(1.05);
    }
  }

  .items-wrap {
    grid-template-columns: repeat(4, 1fr);
    gap: 24px 32px;
    padding: 0;
    align-items: center;
  }
}
</style>
