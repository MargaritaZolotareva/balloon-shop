<template>
  <div v-if="product" class="ProductPage">
    <button class="Button Button--rounded Button--defaultSize ProductPage__closeButton" @click="goBack">
      <i class="far fa-arrow-left CloseButtonIcon"></i>
    </button>
    <div v-if="product.photos && product.photos.length === 1"
         class="ProductPage__photo"
         :data-src="getImageUrl(product.photos[0].photo_path)"
         :style="{ backgroundImage: `url(${getImageUrl(product.photos[0].photo_path)})` }">
    </div>
    <div v-else class="Slider ProductPage__photoSlider">
      <div class="Slider__touchableArea ProductPage__photoSliderTouchableArea"
           @touchstart="handleTouchStart"
           @touchmove="handleTouchMove"
           @touchend="handleTouchEnd">
        <div class="Slider__itemsWrap ProductPage__photoSliderItemsWrap"
             ref="slidesContainer"
             :style="{ transform: `translateX(${translateX}px)`, transition: 'transform 0.25s cubic-bezier(0.1, 0, 0.25, 1)' }">
          <div
              v-for="(photo, index) in product.photos"
              :key="index"
              class="Slider__item ProductPage__photoSliderItem">
            <div class="Slider__itemContent ProductPage__photoSliderItemContent">
              <div class="ProductPage__photo"
                   :data-src="getImageUrl(photo.photo_path)"
                   :style="{ backgroundImage: `url(${getImageUrl(photo.photo_path)})` }">
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="SliderIndexButtons ProductPage__photoSliderIndexButtons">
        <div v-for="(photo, index) in product.photos" :key="index"
             class="SliderIndexButtons__buttonWrap SliderIndexButtons__buttonWrap--lined">
          <div
              :class="[
                'SliderIndexButtons__button',
                'SliderIndexButtons__button--lined',
                { 'SliderIndexButtons__button--active': activeSlide === index }
              ]"
              @click="setActiveSlide(index)">
          </div>
        </div>
      </div>

      <div
          class="Slider__navButtonWrap Slider__navButtonWrap--left ProductPage__photoSliderNavButtonWrap ProductPage__photoSliderNavButtonWrap--left">
        <button
            class="Button Button--rounded Button--bigSize SlideButton SlideButton--leftGradient Slider__navButton"
            @click="moveSlide(-1)">
          <i class="far fa-arrow-left SlideButton__icon SlideButton__icon--leftGradient" style="font-size: 36px;"></i>
        </button>
      </div>
      <div
          class="Slider__navButtonWrap Slider__navButtonWrap--right ProductPage__photoSliderNavButtonWrap ProductPage__photoSliderNavButtonWrap--right">
        <button
            class="Button Button--rounded Button--bigSize SlideButton SlideButton--rightGradient Slider__navButton"
            @click="moveSlide(1)">
          <i class="far fa-arrow-right SlideButton__icon SlideButton__icon--rightGradient" style="font-size: 36px;"></i>
        </button>
      </div>
    </div>

    <div class="ProductPage__content">
      <div class="ProductPage__name">
        <span>{{ product.title  }}</span>
      </div>
      <div class="ProductPage__price">
        <span>{{ product.price }} ₽</span>
      </div>
      <div class="ProductPage__buttons">
        <div class="Dropdown">
          <button class="Dropdown__toggle" @click="toggleDropdown">
            Написать
          </button>
          <div v-if="isDropdownOpen" class="Dropdown__menu" ref="dropdownMenu">
            <a href="https://t.me/yourusername" target="_blank" rel="noopener noreferrer" class="Dropdown__item">
              Telegram
            </a>
            <a href="https://wa.me/yourphonenumber" target="_blank" rel="noopener noreferrer" class="Dropdown__item">
              WhatsApp
            </a>
            <a href="viber://chat?number=yourphonenumber" target="_blank" rel="noopener noreferrer" class="Dropdown__item">
              Viber
            </a>
            <a href="https://vk.com/im?sel=-27526995" target="_blank" rel="noopener noreferrer" class="Dropdown__item">
              VK
            </a>
          </div>
        </div>
      </div>
      <div class="ProductPage__descWrap">
        <div class="ProductPage__descTitle">Информация:</div>
        <span class="SpanWithAnchors ProductPage__desc" v-html="formattedDescription"></span>
      </div>
      <div class="ProductPage__similarProdWrap">
        <div class="ProductPage__similarProdTitle">Другие товары</div>
        <div class="Slider">
          <div class="Slider__touchableArea ProductPage__similarProdTouchableArea"
               @touchstart="handleTouchStart"
               @touchmove="handleTouchMove"
               @touchend="handleTouchEndSimilar">
            <div ref="similarSlidesContainer" class="Slider__itemsWrap ProductPage__similarProdItemsWrap"
                 :style="{
                    transform: `translateX(${similarTranslateX}px)`,
                    transition: 'transform 0.25s cubic-bezier(0.1, 0, 0.25, 1)'
                  }">
              <router-link v-for="item in product.similarProducts" :key="item.id"
                           :to="`/products/${item.id}`"
                           class="Slider__item ProductPage__similarProdSliderItem">
                <div class="Slider__itemContent">
                  <div class="ProductItem ProductPage__similarProdItem">
                    <div class="ProductItem__photoWrap">
                      <div :style="{ backgroundImage: `url(${getImageUrl(item.photo_path)})` }"
                           class="ProductItem__photo ProductPage__similarProdItemPhoto"></div>
                    </div>
                    <div class="ProductItem__name ProductPage__similarProdItemName">
                      <span>{{ item.title }}</span>
                    </div>
                    <span class="ProductItem__price ProductPage__similarProdItemPrice">{{ item.price }} ₽</span>
                  </div>
                </div>
              </router-link>
            </div>
          </div>
          <div
              class="Slider__navButtonWrap Slider__navButtonWrap--left ProductPage__similarProdSliderNavButtonWrap ProductPage__similarProdSliderNavButtonWrap--left">
            <button
                class="Button Button--rounded Button--bigSize SlideButton SlideButton--left Slider__navButton ProductPage__similarProdSliderNavButton"
                :class="{ 'SlideButton--disabled': similarActiveSlide === 0 }"
                @click="moveSimilarSlide(-1)">
              <i class="fat fa-angle-left ProductPage__similarProdSliderNavButtonIcon--left SlideButton__icon SlideButton__icon--left"></i>
            </button>
          </div>
          <div
              class="Slider__navButtonWrap Slider__navButtonWrap--right ProductPage__similarProdSliderNavButtonWrap ProductPage__similarProdSliderNavButtonWrap--right">
            <button
                class="Button Button--rounded Button--bigSize SlideButton SlideButton--right Slider__navButton ProductPage__similarProdSliderNavButton"
                :class="{
                  'SlideButton--disabled':
                    similarActiveSlide >= product.similarProducts.length - similarVisibleCount
                }"
                @click="moveSimilarSlide(1)">
              <i class="fat fa-angle-right ProductPage__similarProdSliderNavButtonIcon--right SlideButton__icon SlideButton__icon--right"></i>
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div v-else class="loading">Загрузка...</div>
</template>

<script>
import axios from 'axios';


export default {
  name: 'ProductPage',
  data() {
    return {
      product: null,
      isLoading: false,
      activeSlide: 0,
      translateX: 0,
      slideWidth: 0,
      startX: 0,
      endX: 0,
      similarActiveSlide: 0,
      similarTranslateX: 0,
      similarSlideWidth: 0,
      similarVisibleCount: 3,
      isDropdownOpen: false,
    }
  },
  mounted() {
    const productId = this.$route.params.id;
    this.fetchProduct(productId).then(() => {
      this.$nextTick(() => {
        this.calculateSlideWidth();
        this.calculateSimilarSlideWidth();
      });
    });
    window.addEventListener('resize', this.calculateSlideWidth);
    window.addEventListener('resize', this.calculateSimilarSlideWidth);
    document.addEventListener('click', this.handleClickOutside);
  },
  methods: {
    async fetchProduct(id) {
      try {
        const apiUrl = process.env.VUE_APP_API_URL;
        const response = await axios.get(`${apiUrl}/products/${id}`);
        this.product = await response.data;
      } catch (error) {
        console.error('Ошибка при получении данных о товаре:', error);
      }
    },
    setActiveSlide(index) {
      console.log('activeSlide', this.activeSlide);
      this.activeSlide = index;
      this.translateX = -this.activeSlide * this.slideWidth;
      console.log(index);
    },
    moveSlide(direction) {
      let newIndex = this.activeSlide + direction;
      if (newIndex < 0) {
        newIndex = this.product.photos.length - 1;
      } else if (newIndex >= this.product.photos.length) {
        newIndex = 0;
      }
      this.setActiveSlide(newIndex);
    },
    handleTouchStart(event) {
      this.startX = event.touches[0].clientX;
    },
    handleTouchMove(event) {
      this.endX = event.touches[0].clientX;
    },
    handleTouchEnd() {
      const swipeDistance = this.startX - this.endX;

      if (Math.abs(swipeDistance) > 50) {
        if (swipeDistance > 0) {
          this.moveSlide(1);
        } else {
          this.moveSlide(-1);
        }
      }
    },
    handleTouchEndSimilar() {
      const swipeDistance = this.startX - this.endX;

      if (Math.abs(swipeDistance) > 20) {
        if (swipeDistance > 0) {
          this.moveSimilarSlide(1);
        } else {
          this.moveSimilarSlide(-1);
        }
      }
    },
    calculateSlideWidth() {
      const slidesContainer = this.$refs.slidesContainer;
      const slide = slidesContainer ? slidesContainer.children[0] : null;
      if (slide) {
        this.slideWidth = slide.clientWidth;
        this.translateX = -this.activeSlide * this.slideWidth;
      }
    },
    calculateSimilarSlideWidth() {
      const container = this.$refs.similarSlidesContainer;
      const slide = container ? container.children[0] : null;

      if (slide) {
        this.similarSlideWidth = slide.clientWidth;
        console.log('Similar slide width: ', this.similarSlideWidth);
      }
    },
    moveSimilarSlide(direction) {
      const maxIndex =
          this.product.similarProducts.length - 2;

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
    getImageUrl(photoPath) {
      const apiHost = process.env.VUE_APP_API_URL;
      return `${apiHost}/${photoPath}`;
    },
    goBack() {
      this.$router.back();
    },
    toggleDropdown() {
      this.isDropdownOpen = !this.isDropdownOpen;
    },
    handleClickOutside(event) {
      const dropdown = this.$refs.dropdownMenu;

      if (this.isDropdownOpen && dropdown && !dropdown.contains(event.target) && !event.target.closest('.Dropdown__toggle')) {
        this.isDropdownOpen = false;
      }
    },
  },
  computed: {
    formattedDescription() {
      return this.product.description.replace(/\n/g, '<br>');
    },
  },
};
</script>

<style scoped lang="scss">
.ProductPage {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  padding: 60px 0 0 0;

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
  &__closeButton {
    position: absolute;
    top: 80px;
    left: 0;
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
  &__similarProdTouchableArea {
    width: 90vw;
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

@media (min-width: 1000px) {
  .ProductPage {
    position: static;
    display: flex;
    overflow: hidden;
    border-radius: $elements_section_radius;
    min-height: 520px;
    padding: 60px 100px 0 100px;

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
        right: 13px;
      }
    }

    &__similarProdSliderNavButton {
      width: 32px;
      height: 32px;
    }

    &__similarProdSliderNavButtonIcon {
      &--left {
        position: absolute;
        right: 3px;
      }

      &--right {
        position: absolute;
        right: 1px;
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
    &__closeButton {
      position: relative;
      top: 0;
      left: -60px;
      display: inline-flex;
      align-items: center;
      justify-content: center;
      background: 0 0;
      width: 56px;
      height: 56px;
      cursor: pointer;
      z-index: $foreground-layer-z-index;
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
