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
          <i class="far fa-arrow-left SlideButton__icon SlideButton__icon--left"></i>
        </button>
      </div>
      <div class="Slider__navButtonWrap Slider__navButtonWrap--right">
        <button
            class="Button Button--rounded Button--bigSize SlideButton SlideButton--right Slider__navButton Advantages__sliderNavButton"
            @click="nextSlide">
          <i class="far fa-arrow-right SlideButton__icon SlideButton__icon--right"></i>
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
