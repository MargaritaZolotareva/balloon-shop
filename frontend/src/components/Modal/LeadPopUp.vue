<template>
  <div class="LeadFormWidget">
    <div class="LeadFormWidget__banner"
         :class="{ 'LeadFormWidget__banner--shown': isFormShown }">
      <div class="LeadFormWidget__avatarWrap">
        <img class="LeadFormWidget__avatar" :src="logo"></div>
      <div class="LeadFormWidget__content"
           @click="$emit('open-modal')">
        <div class="LeadFormWidget__header">Появятся вопросы — пишите! Или звоните <a :href="`tel:${contacts.phone}`" @click="trackPhoneClick">{{ contacts.phone }}</a></div>
        <div class="LeadFormWidget__subheader">Ответим в ближайшее время 😉</div>
        <button class="Button Button--secondary Button--defaultSize LeadFormWidget__button">Оставить заявку</button>
      </div>
      <div class="LeadFormWidget__cancelIconWrap" @click="closeWidget">
        <i class="far fa-xmark LeadFormWidget__cancelIcon"></i>
      </div>
    </div>
    <div class="LeadFormWidget__minimizedBanner"
         :class="{ 'LeadFormWidget__minimizedBanner--shown': !isFormShown }" @click="openWidget">
      <div class="LeadFormWidget__avatarWrap">
        <img class="LeadFormWidget__avatar" :src="logo">
        <div class="LeadFormWidget__chatIcon">
          <i class="far fa-message"></i>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import logo from '@/assets/images/logo.jpg';
import { contacts } from '@/assets/js/contacts';

const isFormShown = ref(true);

const closeWidget = () => {
  isFormShown.value = false;
};

const openWidget = () => {
  isFormShown.value = true;
};

const trackPhoneClick = () => {
  if (typeof window.ym === 'function') {
    window.ym(90974648, 'reachGoal', 'click_phone')
  }
}
</script>

<style scoped lang="scss">
.LeadFormWidget {
  &__banner {
    display: grid;
    grid-template-columns: 84px 1fr 52px;
    position: fixed;
    left: 0;
    right: 0;
    bottom: 0;
    background: $white;
    transition: transform $animation-speed, box-shadow $animation-speed;
    transform: translateY(150px);
    cursor: pointer;
    z-index: $second-layer-z-index;

    &--shown {
      transform: translateY(0);
    }
  }

  &__avatarWrap {
    position: relative;
    text-align: center;
    padding: 16px 0;
  }

  &__avatar {
    display: inline-block;
    vertical-align: middle;
    object-fit: cover;
    width: 56px;
    height: 56px;
    border: .5px solid $photo_border_color;
    border-radius: 50%;
  }

  &__content {
    padding: 16px 0;
    overflow: hidden;
  }

  &__header {
    padding-bottom: 2px;
    font-size: 16px;
    font-weight: 600;
    -webkit-font-smoothing: subpixel-antialiased;
    -moz-osx-font-smoothing: auto;
    line-height: 20px;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  &__subheader {
    padding-bottom: 12px;
    font-size: 14px;
    line-height: 18px;
    color: $gray_dark;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  &__button {
    display: inline-flex;
    align-items: center;
    padding: 5.5px 16px 6.5px;
    font-size: 14px;
    line-height: 18px;
    height: auto;
    color: $checkbox_color;
  }

  &__cancelIconWrap {
    padding: 16px 0;
    height: -moz-min-content;
    height: -webkit-min-content;
    height: min-content;
    box-sizing: border-box;
    text-align: center;
    cursor: pointer;
  }

  &__cancelIcon {
    color: $light_icon_tertiary;
    font-size: 20px;
  }

  &__minimizedBanner {
    position: fixed;
    bottom: 0;
    left: 16px;
    transform: translateY(100px);
    border-radius: 50%;
    box-shadow: 0 2px 4px $black_alpha8;
    transition: transform $animation-speed, box-shadow $animation-speed;
    overflow: hidden;
    cursor: pointer;
    z-index: $second-layer-z-index;
  }

  &__minimizedBanner {
    & .LeadFormWidget__avatarWrap {
      padding: 0;
    }

    &--shown {
      transform: translateY(-16px);
    }
  }

  &__chatIcon {
    position: absolute;
    display: inline-flex;
    align-items: center;
    justify-content: center;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    transition: background $animation-speed;
    background: $black_alpha8;
    color: $static_white;
    font-size: 24px;
  }

  a {
    color: $black;
    text-decoration: none;
  }
}

@media (min-width: 1000px) {
  .LeadFormWidget {
    &__header, &__subheader {
      max-width: 250px;
    }

    &__banner {
      left: auto;
      right: 16px;
      border-radius: $elements_section_radius;
      box-shadow: 0 2px 4px $black_alpha8;

      &--shown {
        transform: translateY(-16px);

        &:hover {
          transform: translateY(-20px);
          box-shadow: 0 0 2px $black_alpha8, 0 2px 24px $black_alpha8;
        }
      }
    }

    &__minimizedBanner {
      left: auto;
      right: 16px;
    }
  }
}
</style>