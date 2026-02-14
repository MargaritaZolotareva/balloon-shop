<template>
  <div class="Top Top--withAdvantages" style="height: 667px;">
    <span class="Top__anchor" id="top" aria-hidden="true"></span>
    <div class="Top__background"
         style="background-image: linear-gradient(74.15deg, rgb(130, 204, 245) -7.5%, rgb(119, 202, 172) 49.69%, rgb(250, 217, 103) 105.7%);"></div>
    <div class="Top__content">
      <h1 class="Top__headerWrap">
        <span class="Top__header">{{ title }}</span>
      </h1>
      <h2 class="Top__descWrap">
        <span class="Top__desc">{{ description }}</span>
      </h2>
      <div class="Top__actionButtonWrap">
        <button class="Button ActionButton ActionButton--top" @click="openLeadModal">
          {{ buttonText }}
        </button>
      </div>
    </div>
    <div class="Top__callToScrollLabel" @click="scrollToNextSection">
      <i class="far fa-angle-down" style="font-size: 24px;"></i>
      <span class="Top__callToScrollLabelIconText">Подробнее</span>
    </div>
    <AdvantagesDesktop/>
  </div>
</template>

<script>
import AdvantagesDesktop from './AdvantagesDesktop.vue'

export default {
  name: 'TopContent',
  components: {
    AdvantagesDesktop
  },
  props: {
    title: {
      type: String,
      default: 'ПОДАРОКК | Гелиевые шары с доставкой в Перми'
    },
    description: {
      type: String,
      default: 'Заказать шары с гелием с доставкой/самовывозом в Перми, красиво упаковать подарок – всё для праздника в одном месте!'
    },
    buttonText: {
      type: String,
      default: 'Оставить заявку'
    }
  },
  inject: ['openLeadModal'],
  methods: {
    scrollToNextSection() {
      window.scrollBy({
        top: window.innerHeight,
        left: 0,
        behavior: "smooth",
      });
    }
  }
};
</script>

<style scoped lang="scss">
.Top {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  height: 100vh;
  background-color: transparent;

  &__content {
    position: relative;
  }

  &__background {
    background-size: cover;
    background-position: center;
    background-repeat: no-repeat;
    height: auto;
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    z-index: $background_layer-z-index;
  }

  &__header {
    padding-top: 0;
    padding-bottom: 16px;
    /*font-family: $section_font;*/
    font-weight: $section_font_weight;
    font-size: calc($section_font_size_offset + 32px);
    line-height: 40px;
    word-break: break-word;
    display: inline-block;
    padding: 0 24px;
    text-align: left;
    color: $static_white;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: pre-line;
  }

  &__desc {
    display: inline-block;
    padding: 0 24px;
    text-align: left;
    color: $static_white;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: pre-line;
  }

  &__headerWrap, &__descWrap {
    text-align: left;
  }

  &__desc {
    padding-bottom: 32px;
    font-size: 17px;
    line-height: 26px;
  }

  &__actionButtonWrap {
    padding: 0 24px;
    text-align: center;
    width: 100%;
    max-width: 600px;
    margin: 0 auto;
  }

  &__anchor {
    display: none;
  }

  &__callToScrollLabel {
    position: absolute;
    display: flex;
    justify-content: center;
    align-items: center;
    bottom: 0;
    left: 0;
    right: 0;
    padding-bottom: 48px;
    font-size: 17px;
    text-align: center;
    color: $static_white;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
  }

  &__callToScrollLabelIconText {
    padding-left: 6px;
  }
}

.ActionButton {
  display: inline-flex;
  justify-content: center;
  align-items: center;

  &--top {
    color: $static_black;
    background: $static_white;
    width: 100%;
  }
}

body, h1, h2, h3, h4, h5, h6, p, ul, li {
  font-weight: 400;
  -webkit-font-smoothing: subpixel-antialiased;
  -moz-osx-font-smoothing: auto;
  padding: 0;
  margin: 0;
}

@media (min-width: 1000px) {
  .Top {
    height: auto;
    min-height: 600px;
    position: relative;

    &__header {
      padding: 0 0 16px;
      margin: 0 auto;
      max-width: 789px;
      font-size: calc($section_font_size_offset + 44px);
      line-height: 58px;
    }

    &--withAdvantages {
      min-height: 764px;
      justify-content: space-between;

      .Top__content {
        padding-top: 172px;
      }
    }

    &__headerWrap, &__descWrap {
      text-align: center;
    }

    &__desc {
      padding: 0 0 48px;
      margin: 0 auto;
      max-width: 583px;
      font-size: 17px;
      font-weight: 400;
      -webkit-font-smoothing: subpixel-antialiased;
      -moz-osx-font-smoothing: auto;
      line-height: 22px;
    }

    &__actionButtonWrap {
      width: auto;
    }

    &__header, &__desc {
      text-align: center;
    }

    &__callToScrollLabel {
      display: none;
    }
  }

  .ActionButton {
    min-width: 196px;

    &--top {
      width: initial;
    }
  }
}

</style>
