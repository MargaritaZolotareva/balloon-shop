<template>
  <footer class="footer">
    <div class="footer__container">

      <div class="footer__nav">
        <ul>
          <li>
            <RouterLink to="/categories">Каталог</RouterLink>
          </li>
          <li>
            <RouterLink to="/delivery">Доставка и оплата</RouterLink>
          </li>
          <li>
            <RouterLink to="/categories/0/products">Акции</RouterLink>
          </li>
          <li>
            <RouterLink to="/contacts">Контакты</RouterLink>
          </li>
          <li>
            <RouterLink to="/privacy">Политика конфиденциальности</RouterLink>
          </li>
        </ul>
      </div>
      <div class="footer__address-block">
        <div class="footer__about">
          <p class="footer__text">
            Воздушные шары для праздников в Перми 🎈<br/>
          </p>
        </div>
        <div class="footer__location">
          <span class="footer__icon">📍</span>
          <span>Пермь, Тургенева, 21, 2 этаж, отдел 215</span>
        </div>

        <div class="footer__schedule">
        <div class="footer__schedule-title">
          <span class="footer__icon">⏰</span>
          <span>График работы</span>
        </div>

        <ul class="footer__schedule-list">
          <li><strong>Пн–Пт:</strong> 10:00 – 20:00</li>
          <li><strong>Сб:</strong> 10:00 – 19:00</li>
          <li><strong>Вс:</strong> выходной</li>
        </ul>
      </div>
      </div>
    </div>

    <div class="footer__bottom">
      © {{ year }} ИП Золотарёва Е.Ю.
    </div>
  </footer>
</template>

<script setup>
import {RouterLink} from 'vue-router';

const year = new Date().getFullYear()
</script>

<style scoped lang="scss">
.footer {
  background: #fff7fb;
  padding: 40px 20px 20px;
  border-top: 1px solid #eee;

  &__container {
    max-width: 1200px;
    margin: 0 auto;
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
    gap: 32px;
  }

  &__text {
    font-size: 14px;
    color: #555;
  }

  &__nav h4 {
    margin-bottom: 12px;
    font-size: 16px;
  }

  &__nav ul {
    list-style: none;
    padding: 0;
    margin: 0;
  }

  &__nav li {
    margin-bottom: 8px;
  }

  a {
    color: $link;
    text-decoration: none;
  }

  a:hover {
    text-decoration: underline;
  }

  &__bottom {
    text-align: center;
    font-size: 13px;
    color: #777;
    margin-top: 24px;
  }

  &__location {
    display: flex;
    margin: 12px 0;
    gap: 6px;
    font-size: 14px;
    color: #444;
  }

  &__schedule {
    margin-top: 14px;
    padding-top: 10px;
    border-top: 1px dashed #e0e0e0;
  }

  &__schedule-title {
    display: flex;
    align-items: center;
    gap: 6px;
    font-weight: 600;
    font-size: 14px;
    margin-bottom: 6px;
  }

  &__schedule-list {
    list-style: none;
    padding: 0;
    margin: 0;
    font-size: 14px;
    color: #555;
  }

  &__schedule-list li {
    margin-bottom: 4px;
  }

  &__icon {
    line-height: 1;
  }
  //&__contact-block {
  //  display: flex;
  //  flex-direction: column;
  //  gap: 10px;
  //  margin: 12px 0;
  //}

  &__phone {
    display: flex;
    align-items: center;
    gap: 8px;
    font-size: 15px;
    font-weight: 600;
  }

  &__phone a {
    color: $black;
    text-decoration: none;
  }

  &__phone a:hover {
    text-decoration: underline;
  }

  &__messengers {
    display: flex;
    flex-wrap: wrap;
    gap: 10px;
    align-items: center;
    font-size: 14px;

    &-title {
      color: #666;
      margin-right: 4px;
    }
  }

  &__messenger {
    padding: 6px 10px;
    border-radius: 16px;
    font-size: 13px;
    text-decoration: none;
    border: 1px solid #ddd;
    transition: all 0.2s ease;

    &:hover {
      background: #f5f5f5;
    }

    &--whatsapp {
      border-color: #25d366;
      color: #25d366;
    }
    &--telegram {
      border-color: #229ed9;
      color: #229ed9;
    }
  }
  &__address-block {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }
}
</style>
