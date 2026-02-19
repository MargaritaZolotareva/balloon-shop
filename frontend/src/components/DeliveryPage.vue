<template>
  <div class="DeliveryPage">
    <button class="Button Button--rounded Button--defaultSize DeliveryPage__closeButton" @click="goBack">
      <i class="far fa-arrow-left CloseButtonIcon"></i>
    </button>

    <h2 class="Block__title Block__title--default">
      <div class="Block__titleWrap">Доставка и оплата</div>
    </h2>

    <div class="Block__content">
      <div class="DeliveryPage__textBlock">
        <h3>ОПЛАТА ЗАКАЗА</h3>
        <ul>
          <li>Наличными в отделе или курьеру</li>
          <li>Переводом на карту, реквизиты отправляем по запросу при оформлении заказа</li>
        </ul>
        <p><span class="icon">🚫</span> Платежи пластиковыми картами на данный момент не принимаются.</p>
        <p><span class="icon">❗</span> При заказе надутых гелием шаров необходима предоплата в размере не менее 50% от
          стоимости заказа.</p>
      </div>

      <div class="DeliveryPage__textBlock">
        <h3>ДОСТАВКА ТОВАРОВ</h3>
        <p><strong>Самовывоз</strong></p>
        <ul>
          <li>Без ограничений по минимальной сумме заказа.</li>
          <li>Самовывоз возможен в рабочие часы отдела. Желательно согласовывать примерное время самовывоза, т.к. мастер
            может быть на выездном оформлении и не находиться в отделе.
          </li>
        </ul>

        <p><strong>Курьером по Перми</strong></p>
        <ul>
          <li>Осуществляется при оформлении заказа до 20:00 (по местному времени), общей сумме заказа не менее 1000 руб.
            (отдалённые районы при заказе от 1500р) и предоплате не менее 50% суммы заказа.
          </li>
          <li>При заказе доставки курьером с 21:00 до 10:30 стоимость доставки может быть увеличена.</li>
          <li>Доставка осуществляется до подъезда. Время доставки согласовывается с клиентом при оформлении заказа.</li>
          <li>Максимальное время ожидания клиента курьером: 15 минут. Если клиент не появляется, курьер вправе увезти
            товар обратно в пункт самовывоза, т. е. клиенту необходимо забрать товар из отдела самому. Оплата доставки в
            этом случае клиенту не возвращается и включается в стоимость заказа.
          </li>
          <li>Стоимость доставки может отличаться от указанной на зонах доставки, точную стоимость необходимо уточнить
            при оформлении заказа. Стоимость доставки для неотмеченных зон также рассчитывается индивидуально.
          </li>
        </ul>

        <p>Остались вопросы? Можно связаться с нами по телефону <a :href="`tel:${contacts.phone}`" @click="trackPhoneClick">{{ contacts.phone }}</a>, написать в
          мессенджерах (<a href="https://t.me/podarokk_shop">Telegram</a>, <a
              :href="contacts.whatsapp">WhatsApp</a>, <a :href="contacts.viber">Viber</a>, <a
              :href="contacts.vk">VK</a>) или оставить заявку на сайте.</p>
      </div>

      <h3>ЗОНЫ ДОСТАВКИ</h3>
      <div class="DeliveryPage__map">
        <iframe src="https://www.google.com/maps/d/embed?mid=1l-B1xdWuZ1hSwNoD-tlMAFD7vCtgJeRL&hl=ru&ehbc=2E312F"
                width="100%" height="480" loading="lazy"></iframe>
      </div>
    </div>
  </div>
</template>

<script setup>
import {contacts} from '@/assets/js/contacts'
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
.DeliveryPage {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  padding: 100px 20px 30px 20px;

  &__closeButton {
    position: absolute;
    display: inline-flex;
    align-items: center;
    justify-content: center;
    background: 0 0;
    width: 56px;
    height: 56px;
    cursor: pointer;
    z-index: 1;
    left: -5px;
    top: 87px;
  }

  &__textBlock {
    margin-bottom: 24px;
  }

  &__map {
    margin-top: 24px;
  }

  h3 {
    font-size: 18px;
    margin-bottom: 12px;
    font-weight: bold;
  }

  ul {
    margin-left: 20px;
  }

  li {
    margin-bottom: 8px;
    font-size: 16px;
  }

  p {
    font-size: 16px;
    line-height: 1.5;
  }

  .icon {
    margin-right: 8px;
  }
}

@media (min-width: 1000px) {
  .DeliveryPage {
    padding: 120px 100px 0 100px;

    &__closeButton {
      left: 25px;
      top: 115px;
    }
    
    &__textBlock {
      margin-bottom: 32px;
    }
  }
}
</style>
