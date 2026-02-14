<template>
  <div class="Block Content">
    <span class="Block__anchor" aria-hidden="true"></span>
    <div class="Block__content Block__content--default">
      <div class="Content__blockContent">
        <div class="Block Content__innerBlock">
          <span class="Block__anchor" aria-hidden="true"></span>
          <h2 class="Block__title Block__title--default Content__innerBlockTitle">
            <div class="Block__titleWrap">Часто задаваемые вопросы</div>
          </h2>
          <div class="Block__content Block__content--default Content__innerBlockContent">
            <div>
              <div class="Slider Content__attachmentsSlider">
                <div class="Slider__touchableArea">
                  <div class="Slider__itemsWrap"
                       style="transform: translateX(0px); transition: transform 0.25s cubic-bezier(0.1, 0, 0.25, 1);">
                    <div class="Slider__item">
                      <div class="Slider__itemContent">
                        <div class="Content__attachmentImage"
                             :data-src="delivery_zones"
                             :style="{backgroundImage: `url(${delivery_zones})`}"
                             loading="lazy"></div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="Content__desc">
              - Как получить скидку на гелиевые шары?
              Скидка предоставляется при заказе не позднее, чем за сутки, через группу, по телефону или в отделе и 100%
              предоплате.
              <br/><br>
              - Можно поменять цвета шаров в композициях или составить свою?
              Конечно можно! Композиции представлены как примеры, их состав и цветовую гамму можно менять. Также
              составим композицию с нуля по вашим пожеланиям.
              <br/><br>
              - Вы обрабатываете шары?
              Да, все латексные шары в композициях мы обрабатываем пропиткой для долгого полёта.
              <br/><br>
              - Сколько стоит доставка?
              Доставка по Перми от 100р в зависимости от адреса. Примерная стоимость доставки указана на картинке, в
              неотмеченные зоны стоимость доставки рассчитывается индивидуально.
            </div>
          </div>
        </div>
        <div class="Content__attachments">
          <div class="Content__attachmentImage"
               :data-src="delivery_zones"
               :style="{backgroundImage: `url(${delivery_zones})`}"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import delivery_zones from '@/assets/images/delivery_zones.jpg';

export default {
  data() {
    return {
      delivery_zones
    };
  },
}
</script>
<style scoped lang="scss">
.Content {
  &__innerBlock {
    padding: 0;
  }

  &__attachmentImage {
    height: 56vw;
    background-position: center;
    background-size: cover;
    background-color: $photo_background_color;
    background-repeat: no-repeat;
  }

  &__desc {
    font-size: 17px;
    line-height: 26px;
    color: $gray_dark;
    word-break: break-word;
    white-space: pre-line;
  }

  &__desc, &__actionButtonWrap {
    padding: 0 24px;
  }

  &__attachments {
    display: none;
  }

  &__attachmentsSlider {
    padding-bottom: 32px;
  }
}

@media (min-width: 1000px) {
  .Content {
    &__blockContent {
      max-width: 1200px;
      margin: 0 auto;
      display: grid;
      grid-template-columns: repeat(2, 1fr);
      gap: 30px;
    }

    &__innerBlockContent {
      padding-top: 24px;
    }

    &__attachmentsSlider {
      display: none;
    }

    &__attachmentsSlider {
      padding-bottom: 32px;
    }

    &__attachmentImage {
      min-width: 585px;
      height: 329px;
      border-radius: $elements_section_radius;
      border: .5px solid $photo_border_color;
      cursor: pointer;
    }

    &__desc, &__actionButtonWrap {
      padding: 0;
    }

    &__attachments {
      display: inline-grid;
      gap: 16px;
      padding-top: 8px;
    }
  }
}
</style>