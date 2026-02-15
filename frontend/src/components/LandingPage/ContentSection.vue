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
              Доставка по Перми от 200р в зависимости от адреса. Примерная стоимость доставки указана на карте, в
              неотмеченные зоны стоимость доставки рассчитывается индивидуально.
            </div>
          </div>
        </div>
        <div class="Content__attachments">
          <iframe src="https://www.google.com/maps/d/embed?mid=1l-B1xdWuZ1hSwNoD-tlMAFD7vCtgJeRL&hl=ru&ehbc=2E312F"
                  width="100%" height="480" loading="lazy"></iframe>
        </div>
      </div>
    </div>
  </div>
</template>
<style scoped lang="scss">
.Content {
  &__innerBlock {
    padding: 0;
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