<template>
  <div class="Modal Modal" aria-modal="true">
    <div class="Modal__backdrop"></div>
    <div class="Modal__content">
      <button class="Button Button--rounded Button--defaultSize Modal__closeButton">
        <svg aria-hidden="true" display="block" class="Modal__closeButtonIcon" width="16" height="16"
             viewBox="0 0 16 16" fill="none" style="width: 16px; height: 16px;">
          <path fill="currentColor"
                d="M12.736 3.264a.9.9 0 0 1 0 1.272L9.273 8l3.463 3.464a.9.9 0 0 1 .081 1.18l-.08.092a.9.9 0 0 1-1.273 0L8 9.273l-3.464 3.463a.9.9 0 0 1-1.272-1.272L6.727 8 3.264 4.536a.9.9 0 0 1-.08-1.18l.08-.092a.9.9 0 0 1 1.272 0L8 6.727l3.464-3.463a.9.9 0 0 1 1.272 0"></path>
        </svg>
      </button>
      <div class="ProductDescMore__content">
        <div class="ProductDescMore__header">Информация о товаре</div>
        <div class="Tabs__adaptiveContainer">
          <div class="Tabs__adaptiveTabsWrap">
            <nav class="Tabs">
              <ul class="Tabs__list">
                <li class="Tabs__item Tabs__item--active"><span id="0">Описание</span></li>
              </ul>
              <div class="Tabs__divider" style="transform: translateX(-10px) scaleX(0.6);"></div>
            </nav>
          </div>
        </div>
        <div class="ProductDescMore__desc"><span class="SpanWithAnchors">В композицию входит:<br>- фонтан из 2 шаров с конфетти, 1 фольгированного и 4 шаров латекс<br>- грузик<br><br>Цена указана с обработкой Hi-Float для долгого полёта.</span>
        </div>
        <div class="ProductDescMore__help">Остались вопросы?
          <button class="Button Button--link Button--defaultSize ProductDescMore__helpLink">Оставьте свои контакты
          </button>
          , и мы перезвоним вам.
        </div>
      </div>
    </div>
  </div>
</template>

<script>

</script>


<style scoped lang="scss">
.ProductDescMore {
  &__content {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    padding: 17px 24px 24px;
  }

  &__header {
    font-size: 17px;
    padding-bottom: 8px;
    text-align: center;
  }

  &__desc {
    padding-top: 24px;
    font-size: 14px;
    line-height: 160%;
    max-height: 70vh;
    overflow-y: auto;
  }

  &__help {
    padding-top: 24px;
    font-size: 14px;
    line-height: 160%;
  }

  &__helpLink {
    font-size: 14px;
  }
}

.Tabs {
  position: relative;
  padding: 0;
  border-bottom: 1px solid $steel_gray_80;
  font-size: 13px;
  line-height: 15px;
  display: flex;
  align-items: center;
  justify-content: space-between;

  &__adaptiveContainer {
    .Tabs {
      border-bottom-color: $tabs_item_border;

      &__item {
        margin: 0 16px;

        &--active, &--active:hover {
          border-bottom: 2px solid $tabs_item_active;
        }

        &:first-child {
          margin-left: 0;
        }

        &:last-child {
          margin-right: 0;
        }

        & > * {
          color: $tabs_item_inactive;
          padding-left: 0;
          padding-right: 0;
        }

        &--active > * {
          color: $black;
        }
      }

      &__divider {
        background: $tabs_item_active;
      }
    }
  }

  &__list {
    margin: 0 10px -1px;
    padding: 0;
    display: flex;
  }

  &__item {
    margin: 0 4px;
    list-style: none;
    flex: 0 0 auto;
    border-bottom: 2px solid transparent;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
    white-space: nowrap;

    &--active, &--active:hover {
      border-bottom: 2px solid $light_blue_500;
    }

    & > * {
      display: inline-block;
      vertical-align: top;
      padding: 16px 6px 14px;
      text-decoration: none;
      color: $gray_600;
      cursor: pointer;
    }

    &--active > * {
      color: $black;
    }
  }

  &__divider {
    visibility: hidden;
    width: 100px;
    height: 2px;
    pointer-events: none;
    position: absolute;
    bottom: -1px;
    left: 0;
    background: $light_blue_500;
    transform-origin: center;
    transition: transform .3s ease, visiblity .3s ease;
  }
}

@media (min-width: 1000px) {
  .ProductDescMore {
    &__content {
      position: static;
      padding: 40px;
      width: 680px;
    }

    &__header {
      font-size: 24px;
      line-height: 32px;
      text-align: left;
    }
  }
}
</style>