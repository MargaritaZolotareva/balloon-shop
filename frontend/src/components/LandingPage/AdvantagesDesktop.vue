<template>
  <div class="Advantages Advantages--desktop">
    <div class="Slider Advantages__slider">
      <div class="Slider__touchableArea">
        <div
            class="Slider__itemsWrap Advantages__sliderItemsWrap"
            :style="{ transform: `translateX(-${currentSlide * 100}%)`, transition: 'transform 0.25s cubic-bezier(0.1, 0, 0.25, 1)' }"
        >
          <div
              v-for="(advantage, index) in advantages"
              :key="index"
              class="Slider__item Advantages__sliderItem"
          >
            <div class="Slider__itemContent Advantages__sliderItemContent">
              <div class="AdvantageItem">{{ advantage }}</div>
            </div>
          </div>
        </div>
      </div>

      <div class="SliderIndexButtons">
        <div
            v-for="(advantage, index) in advantages"
            :key="index"
            class="SliderIndexButtons__buttonWrap"
        >
          <div
              class="SliderIndexButtons__button SliderIndexButtons__button--roundedWhite"
              :class="{ 'SliderIndexButtons__button--active': index === currentSlide }"
          ></div>
        </div>
      </div>

      <div class="Slider__navButtonWrap Slider__navButtonWrap--left">
        <button
            class="Button Button--rounded Button--bigSize SlideButton SlideButton--left Slider__navButton Advantages__sliderNavButton"
            @click="prevSlide">
          <svg
              class="SlideButton__icon SlideButton__icon--left"
              viewBox="0 0 580 580"
              xmlns="http://www.w3.org/2000/svg"
              width="24"
              height="24"
              fill="currentColor">
            <path fill="currentColor" d="M73.4 297.4C60.9 309.9 60.9 330.2 73.4 342.7L233.4 502.7C245.9 515.2 266.2 515.2 278.7 502.7C291.2 490.2 291.2 469.9 278.7 457.4L173.3 352L544 352C561.7 352 576 337.7 576 320C576 302.3 561.7 288 544 288L173.3 288L278.7 182.6C291.2 170.1 291.2 149.8 278.7 137.3C266.2 124.8 245.9 124.8 233.4 137.3L73.4 297.3z"></path>
          </svg>
        </button>
      </div>
      <div class="Slider__navButtonWrap Slider__navButtonWrap--right">
        <button
            class="Button Button--rounded Button--bigSize SlideButton SlideButton--right Slider__navButton Advantages__sliderNavButton"
            @click="nextSlide">
          <svg
              class="SlideButton__icon SlideButton__icon--right"
              viewBox="0 0 580 580"
              xmlns="http://www.w3.org/2000/svg"
              width="24"
              height="24"
              fill="currentColor">
            <path fill="currentColor" d="M566.6 342.6C579.1 330.1 579.1 309.8 566.6 297.3L406.6 137.3C394.1 124.8 373.8 124.8 361.3 137.3C348.8 149.8 348.8 170.1 361.3 182.6L466.7 288L96 288C78.3 288 64 302.3 64 320C64 337.7 78.3 352 96 352L466.7 352L361.3 457.4C348.8 469.9 348.8 490.2 361.3 502.7C373.8 515.2 394.1 515.2 406.6 502.7L566.6 342.7z"></path>
          </svg>
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'AdvantagesDesktop',
  data() {
    return {
      currentSlide: 0,
      advantages: [
        'Своевременная доставка – доставляем в оговорённое время, от 200р по Перми',
        '-10% на шары при предоплате заранее и самовывозе (-5% при предоплате и доставке)',
        'С любовью собираем каждую композицию независимо от её стоимости',
        'Гелиевые шары от 80р/шт',
        'Шары цифры от 600р (с гелием)',
      ],
    };
  },
  methods: {
    prevSlide() {
      if (this.currentSlide > 0) {
        this.currentSlide -= 1;
      } else {
        this.currentSlide = this.advantages.length - 1;
      }
    },
    nextSlide() {
      if (this.currentSlide < this.advantages.length - 1) {
        this.currentSlide += 1;
      } else {
        this.currentSlide = 0;
      }
    },
  },
};
</script>

<style scoped lang="scss">
.Advantages {
  text-align: center;

  &--desktop {
    display: none;
  }

  &__slider {
    padding-bottom: 56px;
  }

  &__sliderItemsWrap {
    max-width: 740px;
  }

  &__sliderItem {
    padding: 32px 24px 24px;
  }

  &__sliderItemContent {
    font-size: 16px;
  }

  &__sliderNavButton {
    .SlideButton__icon {
      color: $static_white;
    }
  }
}

.AdvantageItem {
  color: $dark;
  font-size: 20px;
  /*font-family: $section_font; */
}


@media (min-width: 1000px) {
  .Advantages {
    border-bottom: 1px solid $block_border;

    &--desktop {
      display: block;
      border-bottom-width: 0;
      padding-top: 80px;
    }

    &__sliderItem {
      padding-top: 56px;
      padding-left: 0;
      padding-right: 0;
    }

    &__sliderItemContent {
      font-size: 20px;
      line-height: 28px;
      max-width: 696px;
      /*font-family: $section_font;*/
    }

    &__sliderNavButton {
      background-color: transparent;
      box-shadow: none;

      .SlideButton__icon--left,
      .SlideButton__icon--right {
        transition: transform .25s;
      }

      &:not(.SlideButton--disabled):hover .SlideButton__icon--left {
        transform: translateX(-4px);
      }

      &:not(.SlideButton--disabled):hover .SlideButton__icon--right {
        transform: translateX(4px);
      }

      &:not(.SlideButton--disabled):hover .SlideButton__icon {
        color: $static_white;
        opacity: .8;
      }

      &:not(.SlideButton--disabled):hover {
        transform: none;
        background-color: transparent;
      }
    }


  }

  .AdvantageItem {
    color: $static_white;
    padding: 0 56px;
  }
}
</style>
