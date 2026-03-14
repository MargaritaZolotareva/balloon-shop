<template>
  <header :class="['Header', {'Header--scrolled': isScrolled }]">
    <div class="Header__groupWrap" @click="redirectToHome">
      <div class="Avatar Avatar--s Header__groupAvatar">
        <div class="Avatar__wrapper">
          <img class="Avatar__img" :src="logo" alt="ПОДАРОКК | Воздушные и гелиевые шары"/>
        </div>
      </div>

      <div class="Header__groupInfo">
        <div :class="['Header__groupName',
                        { 'Header__groupName--light': isLightTheme,
                          'Header__groupName--dark': !isLightTheme}]">
          ПОДАРОКК | Воздушные и гелиевые шары
        </div>

        <a
            :class="['Contact', 'Contact--phone', 'HeaderContact', 'HeaderContact--phone', 'HeaderContact__headerPhone--mobile',
                {'HeaderContact--light': isLightTheme,
                 'HeaderContact--dark': !isLightTheme}]"
            :href="`tel:${contacts.phone}`"
            @click="trackPhoneClick"
        >
          <svg
              class="HeaderContact__icon--phone"
              viewBox="0 0 580 580"
              xmlns="http://www.w3.org/2000/svg"
              width="24"
              height="24"
              fill="currentColor">
            <path fill="currentColor" d="M224.2 89C216.3 70.1 195.7 60.1 176.1 65.4L170.6 66.9C106 84.5 50.8 147.1 66.9 223.3C104 398.3 241.7 536 416.7 573.1C493 589.3 555.5 534 573.1 469.4L574.6 463.9C580 444.2 569.9 423.6 551.1 415.8L453.8 375.3C437.3 368.4 418.2 373.2 406.8 387.1L368.2 434.3C297.9 399.4 241.3 341 208.8 269.3L253 233.3C266.9 222 271.6 202.9 264.8 186.3L224.2 89z"></path>
          </svg>
          <span class="Contact__name HeaderContact__name HeaderContact__name--phone">+7&nbsp;(950)&nbsp;445-48-84</span>
        </a>
      </div>
    </div>

    <nav :class="['Nav', { 'Nav--light': isLightTheme,
                           'Nav--dark': !isLightTheme }]">
      <ul class="Nav__elemsWrap">
        <li class="Nav__elem">
          <a :class="['Nav__link', 'Nav__link--active', { 'Nav__link--light': isLightTheme,
                                                            'Nav__link--dark': !isLightTheme }]" :href="ROUTES.CATEGORIES">Каталог</a>
        </li>
        <li class="Nav__elem"><a :class="['Nav__link', { 'Nav__link--light': isLightTheme,
                                                             'Nav__link--dark': !isLightTheme }]" :href="ROUTES.DELIVERY">Доставка и оплата</a>
        </li>
        <li class="Nav__elem"><a :class="['Nav__link', { 'Nav__link--light': isLightTheme,
                                                             'Nav__link--dark': !isLightTheme }]"
                                 :href="ROUTES.PROMO_URL">Акции</a></li>
        <li class="Nav__elem"><a :class="['Nav__link', { 'Nav__link--light': isLightTheme,
                                                             'Nav__link--dark': !isLightTheme }]"
                                 :href="ROUTES.CONTACTS">Контакты</a></li>
      </ul>
    </nav>

    <div class="Header__contactsWrap">
      <a
          :class="['Contact', 'Contact--phone', 'HeaderContact', 'HeaderContact--phone', 'HeaderContact__headerPhone--desktop',
                 { 'HeaderContact--light': isLightTheme,
                   'HeaderContact--dark': !isLightTheme } ]"
          :href="`tel:${contacts.phone}`"
          @click="trackPhoneClick"
      >
        <svg
            class="HeaderContact__icon--phone"
            viewBox="0 0 580 580"
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            fill="currentColor">
          <path fill="currentColor" d="M224.2 89C216.3 70.1 195.7 60.1 176.1 65.4L170.6 66.9C106 84.5 50.8 147.1 66.9 223.3C104 398.3 241.7 536 416.7 573.1C493 589.3 555.5 534 573.1 469.4L574.6 463.9C580 444.2 569.9 423.6 551.1 415.8L453.8 375.3C437.3 368.4 418.2 373.2 406.8 387.1L368.2 434.3C297.9 399.4 241.3 341 208.8 269.3L253 233.3C266.9 222 271.6 202.9 264.8 186.3L224.2 89z"></path>
        </svg>
        <span class="Contact__name HeaderContact__name HeaderContact__name--phone">{{ contacts.phone }}</span>
      </a>
      <div class="Header__actionButton">
        <button
            :class="['Button', 'Button--default', 'Button--defaultSize', 'ActionButton', 'ActionButton__headerButton',
                    {'ActionButton__headerButton--light': isLightTheme,
                     'ActionButton__headerButton--dark': !isLightTheme}]"
            @click="$emit('open-modal')"
            aria-label="Оставить заявку"
        >
          Оставить заявку
        </button>
      </div>
    </div>
  </header>
</template>

<script setup>
import { ref, onMounted, onUnmounted, computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'

import logo from '@/assets/images/logo.jpg';
import {contacts, ROUTES} from '@/assets/js/const';

const router = useRouter()
const route = useRoute()

const isLightTheme = ref(false)
const isScrolled = ref(false)

onMounted(() => {
  window.addEventListener('scroll', handleScroll)
})

onUnmounted(() => {
  window.removeEventListener('scroll', handleScroll)
})
const handleScroll = () => {
  if (isHomePage.value) {
    if (window.scrollY > 0) {
      isLightTheme.value = false
      isScrolled.value = true
    } else {
      isLightTheme.value = true
      isScrolled.value = false
    }
  } else {
    isLightTheme.value = false
    isScrolled.value = window.scrollY > 0
  }
}

const trackPhoneClick = () => {
  if (typeof window.ym === 'function') {
    window.ym(90974648, 'reachGoal', 'click_phone')
  }
}

const redirectToHome = () => {
  router.push('/');
}

const isHomePage = computed(() => {
  return route.meta.isHomePage
})
</script>

<style scoped lang="scss">
.Header {
  display: grid;
  grid-template-columns: 1fr minmax(auto, 362px);
  align-items: center;
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  padding: 12px 16px;
  z-index: $foreground-layer-z-index;


  &__groupWrap {
    display: inline-flex;
    align-items: center;
    padding-right: 16px;
    cursor: pointer;
  }

  &__groupInfo {
    display: grid;
    gap: 3px;
    padding-left: 4px;
  }

  &__groupAvatar {
    margin-right: 4px;
  }

  &__groupName {
    display: none;

    &--light {
      color: $static_white;
    }

    &--dark {
      color: $dark;
    }
  }

  &__groupMembers {
    &--light {
      color: $static_white;
    }

    &--dark {
      color: $dark;
    }
  }

  &__contactsWrap {
    display: flex;
    align-items: center;
    justify-content: flex-end;
  }

  &__actionButton {
    display: initial;
  }

  &--scrolled {
    background-color: $block_light_background;
    transition: background-color .3s;
  }
}

.Nav {
  display: none;

  &__elemsWrap {
    display: flex;
    justify-content: space-around;
    padding: 10px;
    margin-bottom: 0;
  }
}

.modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
}

.modal-content {
  background-color: white;
  padding: 20px;
  border-radius: 10px;
  width: 80%;
}

.modal input {
  margin: 10px 0;
  padding: 10px;
  width: 100%;
  border-radius: 5px;
  border: 1px solid #ddd;
}

.Contact {
  display: inline-block;
  text-align: center;
  text-decoration: none;
  color: $link;
  font-size: 14px;


}

.HeaderContact {
  display: inherit;
  padding: 0 5px;

  &--light {
    color: $static_white;
  }

  &--dark {
    color: $dark;
  }

  &__headerPhone {
    &--desktop {
      display: none;
    }
  }

  &__name {
    display: none;
  }
  &__icon {
    &--phone {
      font-size: 28px;
    }
  }
}

.ActionButton {
  display: inline-flex;
  justify-content: center;
  align-items: center;

  &__headerButton {
    min-width: initial;
    height: 36px;
    margin-right: 10px;

    &--light {
      background-color: $white_alpha16;
      color: $static_white;
    }

    &--dark {
      background-color: $block_light_background;
      color: $black;

      &:hover {
        opacity: .8;
      }
    }
  }
}

@media (min-width: 1000px) {
  .Header {
    display: grid;
    grid-template-columns: minmax(auto, 370px) 1fr minmax(auto, 370px);

    &__groupAvatar {
      margin-right: 0;
    }

    &__groupInfo {
      padding-left: 10px;
    }

    &__groupName {
      display: inline-block;
      font-size: 14px;
      font-weight: 500;
      -webkit-font-smoothing: subpixel-antialiased;
      -moz-osx-font-smoothing: auto;
      max-width: 312px;
      color: $white;

      &--light {
        color: $static_white;
      }

      &--dark {
        color: $dark;
      }
    }

    &__actionButton {
      display: initial;
      white-space: nowrap;
    }
  }

  .HeaderContact {
    &__headerPhone {
      &--mobile {
        display: none;
      }

      &--desktop {
        display: inline-block;
        padding: 0;
      }
    }

    &__name {
      display: none;

      &--phone {
        display: inline;
        font-weight: 500;
        -webkit-font-smoothing: subpixel-antialiased;
        -moz-osx-font-smoothing: auto;
      }
    }

    &__icon {
      &--phone {
        display: none;
      }
    }
  }

  .Nav {
    display: block;
    text-align: center;

    &__link {
      padding: 0 10px;
      text-decoration: none;

      &--light {
        color: $static_white;
      }

      &--dark {
        color: $gray;

        .Nav__link {
          &--active,
          &:hover {
            color: $dark;
          }
        }
      }
    }

    &__elemsWrap {
      display: inline-grid;
      grid-auto-flow: column;
      gap: 12px;
    }

    &__elem {
      display: inline-flex;
      align-items: center;
      cursor: pointer;
    }

    &--light {
      color: $static_white;

      .Nav__link {
        opacity: .64;

        &--active, &:hover {
          opacity: 1;
        }
      }
    }

    &--dark {
      .Nav__link--active, .Nav__link:hover {
        color: $dark;
      }
    }
  }

  .ActionButton {
    min-width: 196px;

    &__headerButton {
      margin-left: 18px;

      &--light {
        &:hover {
          background-color: $white_alpha24;
        }
      }

      &--dark {
        &:hover {
          opacity: .8;
        }
      }
    }
  }
}
</style>
