<template>
  <div class="ContactsPage">
    <button class="Button Button--rounded Button--defaultSize ContactsPage__closeButton" @click="goBack">
      <svg
          class="CloseButtonIcon"
          viewBox="0 0 512 512"
          xmlns="http://www.w3.org/2000/svg"
          width="24"
          height="24"
          fill="currentColor">
        <path fill="currentColor" d="M73.4 297.4C60.9 309.9 60.9 330.2 73.4 342.7L233.4 502.7C245.9 515.2 266.2 515.2 278.7 502.7C291.2 490.2 291.2 469.9 278.7 457.4L173.3 352L544 352C561.7 352 576 337.7 576 320C576 302.3 561.7 288 544 288L173.3 288L278.7 182.6C291.2 170.1 291.2 149.8 278.7 137.3C266.2 124.8 245.9 124.8 233.4 137.3L73.4 297.3z"></path>
      </svg>
    </button>

    <h2 class="Block__title Block__title--default">
      <div class="Block__titleWrap">Контакты</div>
    </h2>

    <div class="Block__content">
      <p class="ContactsPage__intro">
        Вы можете связаться с нами по телефону, написать в мессенджерах или оставить заявку на сайте — мы обязательно ответим.
      </p>

      <div class="ContactsPage__blocks">
        <div class="ContactsPage__block">
          <h3>Реквизиты организации</h3>
          <p>ИП Золотарёва Е.Ю.</p>
          <p>ИНН: {{ contacts.inn }}</p>
          <p>ОГРНИП: {{ contacts.ogrnip }}</p>
          <p>Адрес: {{ contacts.address }}</p>
          <p>Email: <a :href="`mailto:${contacts.email}`">{{ contacts.email }}</a></p>
          <p>Телефон: <a :href="`tel:${contacts.phone}`" @click="trackPhoneClick">{{ contacts.phone }}</a></p>
        </div>
        <div class="ContactsPage__block">
          <h3 class="ContactsPage__block-title">Связаться с нами</h3>

          <div class="ContactsPage__phone">
            <span class="ContactsPage__icon">📞</span>
            <a :href="`tel:${contacts.phone}`" @click="trackPhoneClick">{{ contacts.phone }}</a>
          </div>

          <div class="ContactsPage__messengers">
            <a :href="contacts.whatsapp" target="_blank" class="ContactsPage__messenger whatsapp">
              WhatsApp
            </a>
            <a :href="contacts.telegram" target="_blank" class="ContactsPage__messenger telegram">
              Telegram
            </a>
            <a :href="contacts.viber" target="_blank" class="ContactsPage__messenger viber">
              Viber
            </a>
          </div>
        </div>
      </div>
    </div>

    <div class="ContactsPage__map">
      <div style="position:relative;overflow:hidden;"><a
          href="https://yandex.ru/maps/org/podarokk/161036879423/?utm_medium=mapframe&utm_source=maps"
          style="color:#eee;font-size:12px;position:absolute;top:0px;">Подарокк</a><a
          href="https://yandex.ru/maps/50/perm/category/gift_and_souvenir_shop/184108001/?utm_medium=mapframe&utm_source=maps"
          style="color:#eee;font-size:12px;position:absolute;top:14px;">Магазин подарков и сувениров в Перми</a>
        <iframe
            src="https://yandex.ru/map-widget/v1/org/podarokk/161036879423/?ll=56.293424%2C58.017340&utm_source=share&z=13"
            width="100%" height="480" frameborder="1" allowfullscreen="true" style="position:relative;"></iframe>
      </div>
    </div>
  </div>
</template>

<script setup>
import { contacts } from '@/assets/js/contacts';
import { useRouter } from 'vue-router'
const router = useRouter()
const goBack = () => {
  if (window.history.length > 1) {
    router.back();
  } else {
    router.push('/');
  }
}

const trackPhoneClick = () => {
  if (typeof window.ym === 'function') {
    window.ym(90974648, 'reachGoal', 'click_phone')
  }
}
</script>

<style scoped lang="scss">
.ContactsPage {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  padding: 100px 20px 30px 20px;

  &__closeButton {
    position: absolute;
    left: -5px;
    top: 87px;
    display: inline-flex;
    align-items: center;
    justify-content: center;
    background: 0 0;
    width: 56px;
    height: 56px;
    cursor: pointer;
    z-index: 1;
  }

  &__container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 20px;
  }

  &__title {
    font-size: 28px;
    margin-bottom: 12px;
  }

  &__intro {
    font-size: 15px;
    color: #555;
    margin-bottom: 32px;
  }

  &__blocks {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(260px, 1fr));
    gap: 32px;
    margin-bottom: 40px;
  }

  &__block {
    line-height: 1.8;
  }

  &__block-title {
    font-size: 16px;
    margin-bottom: 12px;
  }

  &__phone {
    display: flex;
    align-items: center;
    gap: 8px;
    font-size: 16px;
    font-weight: 600;
    margin-bottom: 12px;

    a {
      color: $black;
      text-decoration: none;

      &:hover {
        text-decoration: underline;
      }
    }
  }

  &__messengers {
    display: flex;
    gap: 10px;
    flex-wrap: wrap;
  }

  &__messenger {
    padding: 6px 12px;
    border-radius: 16px;
    font-size: 13px;
    text-decoration: none;
    border: 1px solid #ddd;

    &.whatsapp {
      color: #25d366;
      border-color: #25d366;
    }

    &.telegram {
      color: #229ed9;
      border-color: #229ed9;
    }

    &.viber {
      color: #7360f2;
      border-color: #7360f2;
    }

    &:hover {
      background: #f5f5f5;
    }
  }

  &__location {
    display: flex;
    gap: 6px;
    font-size: 14px;
    color: #444;
  }

  &__schedule {
    list-style: none;
    padding: 0;
    margin: 0;
    font-size: 14px;
    color: #555;

    li {
      margin-bottom: 6px;
    }
  }

  &__map {
    width: 100%;

    iframe {
      width: 100%;
      border: none;
    }
  }

  &__icon {
    line-height: 1;
  }

  h3 {
    font-size: 18px;
    margin-bottom: 12px;
    font-weight: bold;
  }

  p {
    font-size: 16px;
    line-height: 1.5;
  }
}

@media (min-width: 1000px) {
  .ContactsPage {
    padding: 120px 100px 0 100px;

    &__closeButton {
      left: 25px;
      top: 115px;
    }
  }
}
</style>
