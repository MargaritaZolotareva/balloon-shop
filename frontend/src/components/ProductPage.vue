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
            <a :href="contacts.telegram" target="_blank" rel="noopener noreferrer" class="Dropdown__item"
               @click="trackSocialClick('telegram')">
              Telegram
            </a>
            <a :href="contacts.whatsapp" target="_blank" rel="noopener noreferrer" class="Dropdown__item"
               @click="trackSocialClick('whatsapp')">
              WhatsApp
            </a>
            <a :href="contacts.viber" target="_blank" rel="noopener noreferrer" class="Dropdown__item"
               @click="trackSocialClick('viber')">
              Viber
            </a>
            <a :href="contacts.vk" target="_blank" rel="noopener noreferrer" class="Dropdown__item"
               @click="trackSocialClick('vk')">
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
              <i class="far fa-angle-left ProductPage__similarProdSliderNavButtonIcon--left SlideButton__icon SlideButton__icon--left"></i>
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
              <i class="far fa-angle-right ProductPage__similarProdSliderNavButtonIcon--right SlideButton__icon SlideButton__icon--right"></i>
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div class="loading" v-else>
    <h2>Загрузка...</h2>
  </div>
</template>

<script setup>
import { ref, onMounted, watch, computed, nextTick } from 'vue';
import api from '@/services/api';
import { useRoute, useRouter } from 'vue-router';
import { contacts } from '@/assets/js/contacts';

const product = ref(null);
const activeSlide = ref(0);
const translateX = ref(0);
const slideWidth = ref(0);
const startX = ref(0);
const endX = ref(0);
const similarActiveSlide = ref(0);
const similarTranslateX = ref(0);
const similarSlideWidth = ref(0);
const similarVisibleCount = ref(2);
const isDropdownOpen = ref(false);

const dropdownMenu = ref(null);
const slidesContainer = ref(null);
const similarSlidesContainer = ref(null);

const route = useRoute();
const router = useRouter();

const fetchProduct = async (id) => {
  try {
    const response = await api.get(`/products/${id}`);
    product.value = await response.data;
  } catch (error) {
    console.error('Ошибка при получении данных о товаре:', error);
  }
};

onMounted(() => {
  const productId = route.params.id;
  fetchProduct(productId).then(() => {
    nextTick(() => {
      calculateSlideWidth();
      calculateSimilarSlideWidth();
    });
  });

  window.addEventListener('resize', calculateSlideWidth);
  window.addEventListener('resize', calculateSimilarSlideWidth);
  document.addEventListener('click', handleClickOutside);
});

watch(route, (to) => {
  fetchProduct(to.params.id);
});

const calculateSlideWidth = () => {
  const container = slidesContainer.value;
  if (!container || !container.children.length) return;
  const slide = container.children[0];
  if (slide) {
    slideWidth.value = slide.clientWidth;
    translateX.value = -activeSlide.value * slideWidth.value;
  }
};

const calculateSimilarSlideWidth = () => {
  const container = similarSlidesContainer.value;
  if (!container || !container.children.length) return;
  const slide = container.children[0];
  if (slide) {
    similarSlideWidth.value = slide.clientWidth;
  }
};

const setActiveSlide = (index) => {
  activeSlide.value = index;
  translateX.value = -activeSlide.value * slideWidth.value;
};

const moveSlide = (direction) => {
  if (!product.value) return;
  let newIndex = activeSlide.value + direction;
  if (newIndex < 0) {
    newIndex = product.value.photos.length - 1;
  } else if (newIndex >= product.value.photos.length) {
    newIndex = 0;
  }
  setActiveSlide(newIndex);
};

const handleTouchStart = (event) => {
  startX.value = event.touches[0].clientX;
};

const handleTouchMove = (event) => {
  endX.value = event.touches[0].clientX;
};

const handleTouchEnd = () => {
  const swipeDistance = startX.value - endX.value;
  if (Math.abs(swipeDistance) > 50) {
    if (swipeDistance > 0) {
      moveSlide(1);
    } else {
      moveSlide(-1);
    }
  }
};

const handleTouchEndSimilar = () => {
  const swipeDistance = startX.value - endX.value;
  if (Math.abs(swipeDistance) > 20) {
    if (swipeDistance > 0) {
      moveSimilarSlide(1);
    } else {
      moveSimilarSlide(-1);
    }
  }
};

const moveSimilarSlide = (direction) => {
  if (!product.value) return;
  const maxIndex = product.value.similarProducts.length - similarVisibleCount.value;
  let newIndex = similarActiveSlide.value + direction * similarVisibleCount.value;
  if (newIndex < 0) {
    newIndex = 0;
  }
  if (newIndex > maxIndex) {
    newIndex = maxIndex;
  }
  similarActiveSlide.value = newIndex;
  similarTranslateX.value = -similarActiveSlide.value * similarSlideWidth.value;
};

const getImageUrl = (photoPath) => {
  const apiHost = process.env.VUE_APP_API_URL;
  return `${apiHost}/${photoPath}`;
};

const goBack = () => {
  router.back();
};

const toggleDropdown = () => {
  isDropdownOpen.value = !isDropdownOpen.value;
};

const handleClickOutside = (event) => {
  if (!isDropdownOpen.value || !dropdownMenu.value) return;
  if (!dropdownMenu.value.contains(event.target) && !event.target.closest('.Dropdown__toggle')) {
    isDropdownOpen.value = false;
  }
};

const formattedDescription = computed(() => {
  if (!product.value) return;
  return product.value?.description?.replace(/\n/g, '<br>') || '';
});

const trackSocialClick = (socialNetwork) => {
  if (typeof window.ym === 'function') {
    window.ym(90974648, 'reachGoal', 'click_social', {
      socialNetwork,
    });
  }
}
</script>

<style scoped lang="scss">
.ProductPage {
  display: flex;
  flex-direction: column;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  padding: 60px 0 30px 0;

  &__photo {
    width: 100vw;
    height: 100vw;
    background-size: contain;
    background-position: center;
    background-repeat: no-repeat;
    background-color: $product_cart_photo_background_color;
    //cursor: -webkit-zoom-in;
    //cursor: zoom-in;
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
    flex-direction: row;
    overflow: hidden;
    border-radius: $elements_section_radius;
    min-height: 520px;
    padding: 60px 100px 50px 100px;

    &__photo {
      display: inline-block;
      width: 560px;
      height: auto;
      min-height: 100%;
      min-width: 560px;
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
        right: 77px;
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
