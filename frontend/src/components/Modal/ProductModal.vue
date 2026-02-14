<template>
  <div class="Modal Modal" aria-modal="true">
    <div class="Modal__backdrop" @click="close"></div>
    <div class="Modal__content">
      <button class="Button Button--rounded Button--defaultSize Modal__closeButton">
        <svg aria-hidden="true" display="block" class="Modal__closeButtonIcon" width="16" height="16"
             viewBox="0 0 16 16" fill="none" style="width: 16px; height: 16px;">
          <path fill="currentColor"
                d="M12.736 3.264a.9.9 0 0 1 0 1.272L9.273 8l3.463 3.464a.9.9 0 0 1 .081 1.18l-.08.092a.9.9 0 0 1-1.273 0L8 9.273l-3.464 3.463a.9.9 0 0 1-1.272-1.272L6.727 8 3.264 4.536a.9.9 0 0 1-.08-1.18l.08-.092a.9.9 0 0 1 1.272 0L8 6.727l3.464-3.463a.9.9 0 0 1 1.272 0"></path>
        </svg>
      </button>
      <div v-if="product" class="ProductCart">
        <div v-if="product.images.length === 1"
             class="ProductCart__photo"
             :style="{ backgroundImage: `url(${product.images[0]})` }">
        </div>
        <div v-else class="Slider ProductCart__photoSlider">
          <div class="Slider__touchableArea ProductCart__photoSliderTouchableArea">
            <div class="Slider__itemsWrap ProductCart__photoSliderItemsWrap"
                 ref="slidesContainer"
                 :style="{ transform: `translateX(${translateX}px)`, transition: 'transform 0.25s cubic-bezier(0.1, 0, 0.25, 1)' }">
              <div
                  v-for="(image, index) in product.images"
                  :key="index"
                  class="Slider__item ProductCart__photoSliderItem">
                <div class="Slider__itemContent ProductCart__photoSliderItemContent">
                  <div class="ProductCart__photo"
                       :style="{ backgroundImage: `url(${image})` }">
                  </div>
                </div>
              </div>
            </div>
          </div>

          <div class="SliderIndexButtons ProductCart__photoSliderIndexButtons">
            <div v-for="(image, index) in product.images" :key="index"
                 :class="['SliderIndexButtons__buttonWrap SliderIndexButtons__buttonWrap--lined', { 'SliderIndexButtons__button--active': activeSlide === index }]">
              <div
                  class="SliderIndexButtons__button SliderIndexButtons__button--lined"
                  @click="setActiveSlide(index)">
              </div>
            </div>
          </div>

          <div
              class="Slider__navButtonWrap Slider__navButtonWrap--left ProductCart__photoSliderNavButtonWrap ProductCart__photoSliderNavButtonWrap--left">
            <button
                class="Button Button--rounded Button--bigSize SlideButton SlideButton--leftGradient Slider__navButton"
                @click="moveSlide(-1)">
              <svg aria-hidden="true" display="block" class="SlideButton__icon SlideButton__icon--leftGradient"
                   width="36" height="36" viewBox="0 0 36 36" fill="currentColor">
                <path
                    d="M22.884 7.116a1.25 1.25 0 0 1 0 1.768L13.768 18l9.116 9.116a1.25 1.25 0 0 1-1.768 1.768l-10-10a1.25 1.25 0 0 1 0-1.768l10-10a1.25 1.25 0 0 1 1.768 0"></path>
              </svg>
            </button>
          </div>
          <div
              class="Slider__navButtonWrap Slider__navButtonWrap--right ProductCart__photoSliderNavButtonWrap ProductCart__photoSliderNavButtonWrap--right">
            <button
                class="Button Button--rounded Button--bigSize SlideButton SlideButton--rightGradient Slider__navButton"
                @click="moveSlide(1)">
              <svg aria-hidden="true" display="block" class="SlideButton__icon SlideButton__icon--rightGradient"
                   width="36" height="36" viewBox="0 0 36 36" fill="currentColor">
                <path
                    d="M13.116 7.116a1.25 1.25 0 0 1 1.768 0l10 10a1.25 1.25 0 0 1 0 1.768l-10 10a1.25 1.25 0 0 1-1.768-1.768L22.232 18l-9.116-9.116a1.25 1.25 0 0 1 0-1.768"></path>
              </svg>
            </button>
          </div>
        </div>

        <div class="ProductCart__content">
          <div class="ProductCart__name">
            <span>{{ product.name }}</span>
          </div>
          <div class="ProductCart__price">
            <span>{{ product.price }} ₽</span>
          </div>
          <div class="ProductCart__buttons">
            <a
                href="https://vk.me/podarokk.shop"
                target="_blank"
                rel="noreferrer noopener"
                class="Link Link--buttonPrimary ActionButton--productCartButton"
            >
              <svg
                  aria-hidden="true"
                  display="block"
                  class="ActionButton"
                  width="28"
                  height="28"
                  viewBox="0 0 28 28"
                  fill="none"
                  style="width: 28px; height: 28px;"
              >
                <path
                    fill="currentColor"
                    fill-rule="evenodd"
                    d="M4.546 4.546C3 6.093 3 8.582 3 13.56v.88c0 4.978 0 7.467 1.546 9.013S8.582 25 13.56 25h.88c4.978 0 7.467 0 9.013-1.547S25 19.418 25 14.44v-.88c0-4.978 0-7.467-1.547-9.014S19.418 3 14.44 3h-.88C8.582 3 6.093 3 4.546 4.546M6.7 9.7c.12 5.747 2.984 9.2 8.006 9.2h.285v-3.288c1.846.184 3.241 1.538 3.801 3.288H21.4c-.716-2.615-2.598-4.061-3.774-4.614 1.175-.681 2.828-2.339 3.223-4.586H18.48c-.514 1.823-2.038 3.481-3.489 3.638V9.7h-2.369v6.373c-1.469-.369-3.324-2.155-3.406-6.373z"
                    clip-rule="evenodd"
                ></path>
              </svg>
              <span>Написать</span>
            </a>
          </div>
          <div class="ProductCart__descWrap">
            <div class="ProductCart__descTitle">Информация:</div>
            <span class="SpanWithAnchors ProductCart__desc" v-html="product.description"></span>
            <a class="ProductCart__descMoreLink">Смотреть всю информацию</a>
          </div>
          <div class="ProductCart__similarProdWrap">
            <div class="ProductCart__similarProdTitle">Другие товары</div>
            <div class="Slider">
              <div class="Slider__touchableArea ProductCart__similarProdTouchableArea">
                <div ref="similarSlidesContainer" class="Slider__itemsWrap ProductCart__similarProdItemsWrap"
                     :style="{
                      transform: `translateX(${similarTranslateX}px)`,
                      transition: 'transform 0.25s cubic-bezier(0.1, 0, 0.25, 1)'
                    }">
                  <div v-for="item in similarProducts" :key="item.id"
                       class="Slider__item ProductCart__similarProdSliderItem">
                    <div class="Slider__itemContent">
                      <div class="ProductItem ProductCart__similarProdItem">
                        <div class="ProductItem__photoWrap">
                          <div :style="{ backgroundImage: `url(${item.image})` }"
                               class="ProductItem__photo ProductCart__similarProdItemPhoto"></div>
                        </div>
                        <div class="ProductItem__name ProductCart__similarProdItemName">
                          <span>{{ item.name }}</span>
                        </div>
                        <span class="ProductItem__price ProductCart__similarProdItemPrice">{{ item.price }} ₽</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
              <div
                  class="Slider__navButtonWrap Slider__navButtonWrap--left ProductCart__similarProdSliderNavButtonWrap ProductCart__similarProdSliderNavButtonWrap--left">
                <button
                    class="Button Button--rounded Button--bigSize SlideButton SlideButton--left Slider__navButton ProductCart__similarProdSliderNavButton"
                    :class="{ 'SlideButton--disabled': similarActiveSlide === 0 }"
                    @click="moveSimilarSlide(-1)">
                  <svg aria-hidden="true" display="block"
                       class="ProductCart__similarProdSliderNavButtonIcon--left SlideButton__icon SlideButton__icon--left"
                       width="28" height="28" viewBox="0 0 28 28" fill="none" style="width: 28px; height: 28px;">
                    <path fill="currentColor"
                          d="m12.414 14 5.793-5.793a1 1 0 0 0-1.414-1.414l-6.5 6.5a1 1 0 0 0 0 1.414l6.5 6.5a1 1 0 0 0 1.414-1.414z"></path>
                  </svg>
                </button>
              </div>
              <div
                  class="Slider__navButtonWrap Slider__navButtonWrap--right ProductCart__similarProdSliderNavButtonWrap ProductCart__similarProdSliderNavButtonWrap--right">
                <button
                    class="Button Button--rounded Button--bigSize SlideButton SlideButton--right Slider__navButton ProductCart__similarProdSliderNavButton"
                    :class="{
                    'SlideButton--disabled':
                      similarActiveSlide >= similarProducts.length - similarVisibleCount
                  }"
                    @click="moveSimilarSlide(1)">
                  <svg aria-hidden="true" display="block"
                       class="ProductCart__similarProdSliderNavButtonIcon--right SlideButton__icon SlideButton__icon--right"
                       width="28" height="28" viewBox="0 0 28 28" style="width: 28px; height: 28px;">
                    <g fill="none" fill-rule="evenodd">
                      <path d="M0 0h28v28H0z"></path>
                      <path stroke="#8bc1ff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                            d="m11 7.5 6.5 6.5-6.5 6.5"></path>
                      <path fill="currentColor" fill-rule="nonzero"
                            d="m16.086 14-5.793 5.793a1 1 0 0 0 1.414 1.414l6.5-6.5a1 1 0 0 0 0-1.414l-6.5-6.5a1 1 0 0 0-1.414 1.414z"></path>
                    </g>
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div v-else class="loading">Загрузка...</div>
    </div>
  </div>
</template>

<script>

export default {
  name: 'ProductPage',
  data() {
    return {
      product: {},
      similarProducts: [],
      isLoading: false,
      activeSlide: 0,
      translateX: 0,
      slideWidth: 0,
      similarActiveSlide: 0,
      similarTranslateX: 0,
      similarSlideWidth: 0,
      similarVisibleCount: 3
    };
  },
  mounted() {
    this.fetchProductData();
    this.calculateSlideWidth();
    this.calculateSimilarSlideWidth();
  },
  methods: {
    async fetchProductData() {
    },
    setActiveSlide(index) {
      this.activeSlide = index;
      this.translateX = -this.activeSlide * this.slideWidth;
    },
    moveSlide(direction) {
      let newIndex = this.activeSlide + direction;
      if (newIndex < 0) {
        newIndex = this.product.images.length - 1;
      } else if (newIndex >= this.product.images.length) {
        newIndex = 0;
      }
      this.activeSlide = newIndex;
      this.translateX = -this.activeSlide * this.slideWidth;
    },
    calculateSlideWidth() {
      const slidesContainer = this.$refs.slidesContainer;
      const slide = slidesContainer ? slidesContainer.children[0] : null;
      if (slide) {
        this.slideWidth = slide.clientWidth;
      }
    },
    calculateSimilarSlideWidth() {
      const container = this.$refs.similarSlidesContainer;
      const slide = container ? container.children[0] : null;

      if (slide) {
        this.similarSlideWidth = slide.clientWidth;
      }
    },
    moveSimilarSlide(direction) {
      const maxIndex =
          this.similarProducts.length - 1;

      let newIndex = this.similarActiveSlide + direction;

      if (newIndex < 0) {
        newIndex = 0;
      }

      if (newIndex > maxIndex) {
        newIndex = maxIndex;
      }

      this.similarActiveSlide = newIndex;
      this.similarTranslateX = -this.similarActiveSlide * this.similarSlideWidth;
    },
  },
};
</script>

<style scoped lang="scss">
.ProductCart {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;

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
    -webkit-line-clamp: 4;
    -webkit-box-orient: vertical;
    overflow: hidden;
    font-size: 14px;
    line-height: 18px;
    color: $gray_dark;
  }

  &__descMoreLink {
    display: inline-block;
    margin-top: 4px;
    font-size: 14px;
    line-height: 18px;
    color: $checkbox_color;
    cursor: pointer;
  }

  &__similarProdWrap {
    padding-top: 24px;
    display: flex;
    flex-direction: column;
  }

  &__similarProdTitle {
    font-weight: 500;
    -webkit-font-smoothing: subpixel-antialiased;
    -moz-osx-font-smoothing: auto;
    font-size: 14px;
    line-height: 18px;
  }

  &__similarProdItemsWrap {
    padding-top: 8px;
    align-items: flex-start;
  }

  &__similarProdSliderItem {
    max-width: 141px;
    padding-right: 8px;
  }

  &__similarProdItem {
    width: 100%;
    text-align: left;
  }

  &__similarProdItemPhoto {
    padding-bottom: 98%;
  }

  &__similarProdItemName {
    -webkit-line-clamp: 2;
    color: $gray_dark;
  }

  &__similarProdItemPrice {
    font-weight: 500;
    -webkit-font-smoothing: subpixel-antialiased;
    -moz-osx-font-smoothing: auto;
    display: inline-block;
  }

  &__photoSliderIndexButtons {
    width: 100vw;
    bottom: 8px;
    position: absolute;
    left: 0;
    right: 0;
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

@media (min-width: 1000px) {
  .ProductCart {
    position: static;
    display: flex;
    overflow: hidden;
    border-radius: $elements_section_radius;
    min-height: 520px;

    &__photo {
      display: inline-block;
      width: 560px;
      height: auto;
      min-height: 100%;
    }

    &__content {
      display: inline-block;
      width: 388px;
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

    &__desc {
      -webkit-line-clamp: 3;
    }

    &__similarProdTouchableArea {
      width: -moz-fit-content;
      width: -webkit-fit-content;
      width: fit-content;
      padding-top: 2px;
      padding-bottom: 0;
      min-width: 388px;
    }

    &__similarProdSliderItem {
      max-width: 129px;
      padding-right: 8px;
    }

    &__similarProdItemPhoto {
      padding-bottom: 100%;
    }

    &__similarProdItemName {
      font-size: 13px;
      padding-top: 10px;
    }

    &__similarProdItemPrice {
      font-size: 14px;
      line-height: 18px;
    }

    &__similarProdSliderNavButtonWrap {
      top: 82px;

      &--left {
        left: 38px;
      }

      &--right {
        right: -34px;
      }
    }

    &__similarProdSliderNavButton {
      width: 32px;
      height: 32px;
    }

    &__similarProdSliderNavButtonIcon {
      &--left {
        position: absolute;
      }

      &--right {
        position: absolute;
      }
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

  .Slider__navButtonWrap {
    opacity: 0;
    pointer-events: none;
    transition: opacity 0.2s ease;
  }

  .Slider__touchableArea:hover ~ .Slider__navButtonWrap,
  .Slider:hover .Slider__navButtonWrap {
    opacity: 1;
    pointer-events: auto;
  }
}
</style>
