<template>
  <div class="main-block" id="top">
    <h2 class="block__title block__title--default">
      <div class="block__titleWrap">{{ category_title }}</div>
    </h2>
    <div class="block-content">
      <div class="items-wrap" v-if="products.length > 0">
        <ProductItemBig :product="products[0]"/>
        <ProductSmallContent :products="products.slice(1)"/>
      </div>
      <div class="loading" v-else>
        <h2>Загрузка...</h2>
      </div>
    </div>
    <div class="items__moreButtonWrap">
      <router-link :to="`/categories/${categoryId}/products`">
        <button class="SectionButton items__moreButton">
          <div class="SectionButton__border">
            <div class="SectionButton__background">
              <div class="SectionButton__text">Перейти к товарам категории</div>
            </div>
          </div>
        </button>
      </router-link>
    </div>
  </div>
</template>

<script>
import ProductItemBig from "./ProductItemBig.vue";
import ProductSmallContent from "./ProductSmallContent.vue";
import api from '@/services/api'

export default {
  name: "ProductsSection",
  components: {
    ProductItemBig,
    ProductSmallContent
  },
  data() {
    return {
      products: [],
      category_title: ""
    };
  },
  props: {
    section: {
      type: Object,
      required: true
    },
    categoryId: {
      type: Number,
      required: true
    }
  },
  mounted() {
    this.fetchProducts();
  },
  methods: {
    async fetchProducts() {
      try {
        const productsCnt = 5;
        const response = await api.get(`/categories/${this.categoryId}/products?limit=${productsCnt}`);
        const data = await response.data;
        this.products = data.products;
        this.category_title = data.category_title;
      } catch (error) {
        console.error('Ошибка при получении списка товаров:', error);
      }
    }
  }
};
</script>

<style scoped lang="scss">
@import "@/assets/styles/button.scss";

.main-block {
  border-bottom: 3px dotted #eee;
  padding-top: 60px;
  padding-bottom: 30px;
}

.block-content {
  padding-top: 32px;
}

.items {
  &__moreButtonWrap {
    text-align: center;
    padding-top: 32px;
  }

  &__moreButton {
    display: inline-block;
  }
}

.items-wrap {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 32px 8px;
  padding: 0 24px;
}

.block {
  &__title {
    font-size: calc($section_font_size_offset + 24px);
    font-weight: 500;
    -webkit-font-smoothing: subpixel-antialiased;
    -moz-osx-font-smoothing: auto;
    color: $block_title;
    line-height: 30px;
    text-align: center;
    padding: 0 24px;
  }

  &__titleWrap, &__title {
    font-weight: $section_font_weight;
  }
}

@media (min-width: 1000px) {
  .block-content {
    max-width: 1200px;
    margin: 0 auto;
    padding-top: 40px;
  }

  .items-wrap {
    grid-template-columns: repeat(4, 1fr);
    gap: 24px 32px;
    padding: 0;
    align-items: center;
  }

  .items {
    &__moreButtonWrap {
      padding-top: 40px;
    }
  }

  .block {
    &__title {
      max-width: 1200px;
      margin: 0 auto;
      padding: 0;
      font-size: calc($section_font_size_offset + 34px);
      line-height: 44px;
      text-align: initial;
    }
  }
}
</style>
