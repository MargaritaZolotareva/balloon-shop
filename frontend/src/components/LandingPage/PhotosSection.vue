<template>
  <div class="Block Block--withoutPaddings Photos">
    <span class="Block__anchor" aria-hidden="true"></span>
    <h2 class="Block__title Block__title--default">
      <div class="Block__titleWrap">Наш отдел на Тургенева</div>
    </h2>
    <div class="Block__content Block__content--default">
      <div>
        <div class="Slider">
          <div class="Slider__touchableArea">
            <div class="Slider__itemsWrap Photos__sliderItemsWrap"
                 style="transform: translateX(0px); transition: transform 0.25s cubic-bezier(0.1, 0, 0.25, 1);">
              <div class="Slider__item Photos__sliderItem">
                <div class="Slider__itemContent Photos__sliderItemContent Photos__sliderItemContent--leftAlign">
                  <div class="Photos__photoWrap">
                    <div class="Photos__photo"
                         data-src="https://sun9-6.userapi.com/impg/ZJx4pK6PXtkUiaU3uNnfnwCINV_8gCMSLWu3sw/eoGtRJjUrX0.jpg?size=755x1080&amp;quality=95&amp;sign=eefc088726c4b85ba4044172272c2d2c&amp;type=album"
                         style="background-image: url(&quot;https://sun9-6.userapi.com/impg/ZJx4pK6PXtkUiaU3uNnfnwCINV_8gCMSLWu3sw/eoGtRJjUrX0.jpg?size=755x1080&amp;quality=95&amp;sign=eefc088726c4b85ba4044172272c2d2c&amp;type=album&quot;);"></div>
                  </div>
                  <div class="Photos__photoWrap">
                    <div class="Photos__photo"
                         data-src="https://sun9-15.userapi.com/impg/0tEqvaLiqaMA24khppGtXZjQ9Udbl-1ucEa-CA/qN4Oa0I83i4.jpg?size=1280x911&amp;quality=95&amp;sign=35ec17a82261d8d58da6eb96de7b3ddd&amp;type=album"
                         style="background-image: url(&quot;https://sun9-15.userapi.com/impg/0tEqvaLiqaMA24khppGtXZjQ9Udbl-1ucEa-CA/qN4Oa0I83i4.jpg?size=1280x911&amp;quality=95&amp;sign=35ec17a82261d8d58da6eb96de7b3ddd&amp;type=album&quot;);"></div>
                  </div>
                  <div class="Photos__photoWrap">
                    <div class="Photos__photo"
                         data-src="https://sun9-59.userapi.com/impg/_KESaphjFWZpg4ieBpv9SCLmK-AZ9kNXWlzDMg/46LLWPiWLww.jpg?size=1280x911&amp;quality=95&amp;sign=0b33fbdca0f67a189f4040d71201926d&amp;type=album"
                         style="background-image: url(&quot;https://sun9-59.userapi.com/impg/_KESaphjFWZpg4ieBpv9SCLmK-AZ9kNXWlzDMg/46LLWPiWLww.jpg?size=1280x911&amp;quality=95&amp;sign=0b33fbdca0f67a189f4040d71201926d&amp;type=album&quot;);"></div>
                  </div>
                </div>
              </div>
              <div class="Slider__item Photos__sliderItem">
                <div class="Slider__itemContent Photos__sliderItemContent">
                  <div class="Photos__photoWrap">
                    <div class="Photos__photo"
                         data-src="https://sun9-8.userapi.com/impg/dLxmuY6QUzD6UvdsLj4Y-9OyZZg10J6Xcd1LqQ/mtz7KVEHAVo.jpg?size=755x1080&amp;quality=95&amp;sign=3cf44eeaab9e3c94a5655088a759125c&amp;type=album"
                         loading="lazy"></div>
                  </div>
                  <div class="Photos__photoWrap">
                    <div class="Photos__photo"
                         data-src="https://sun9-8.userapi.com/impg/dLxmuY6QUzD6UvdsLj4Y-9OyZZg10J6Xcd1LqQ/sWGg3uygCS8.jpg?size=755x1080&amp;quality=95&amp;sign=a9daade161b1d857329419a3430c8d2f&amp;type=album"
                         loading="lazy"></div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>


<style scoped lang="scss">
.Photos {
  margin-top: 80px;
  margin-bottom: 80px;

  &__sliderItemsWrap {
    width: 100vw;
    margin: 0;
    padding-left: 24px;
  }

  &__sliderItem {
    width: 100%;
    max-width: 257px;
    margin-right: 8px;

    &:last-child {
      margin-right: 0;
    }
  }

  &__photoWrap {
    display: inline-block;
    border-radius: $elements_section_radius;
    width: 100%;
    max-width: 257px;
    height: 193px;
    overflow: hidden;
    -webkit-mask-image: -webkit-radial-gradient($white, $black);
    border: .5px solid $photo_border_color;
  }

  &__photo {
    background-position: center;
    background-size: cover;
    background-color: $photo_background_color;
    cursor: pointer;
    height: 100%;
  }
}

@media (min-width: 1000px) {
  .Photos {
    text-align: center;
    margin-top: 160px;
    margin-bottom: 160px;

    &__sliderItemsWrap {
      max-width: 1268px;
      margin: 0 auto;
      padding: 0;
    }

    &__sliderItem {
      max-width: initial;
      margin-right: 0;
    }

    &__sliderItemContent {
      display: flex;
      justify-content: center;
      padding: 0 34px;

      &--leftAlign {
        justify-content: flex-start;
      }
    }

    &__photoWrap {
      max-width: 389px;
      height: 292px;
      margin-right: 16px;
    }

    &__photo {
      height: 292px;
      transition: transform .3s;
    }
  }
}
</style>