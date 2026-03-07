<template>
  <div class="Modal Modal" aria-modal="true">
    <div class="Modal__backdrop" @click="close"></div>
    <div class="Modal__content">
      <button class="Button Button--rounded Button--defaultSize Modal__closeButton" @click="close">
        <svg
            class="Modal__closeButton"
            viewBox="0 0 512 512"
            xmlns="http://www.w3.org/2000/svg"
            width="20"
            height="20"
            fill="currentColor">
          <path fill="currentColor" d="M183.1 137.4C170.6 124.9 150.3 124.9 137.8 137.4C125.3 149.9 125.3 170.2 137.8 182.7L275.2 320L137.9 457.4C125.4 469.9 125.4 490.2 137.9 502.7C150.4 515.2 170.7 515.2 183.2 502.7L320.5 365.3L457.9 502.6C470.4 515.1 490.7 515.1 503.2 502.6C515.7 490.1 515.7 469.8 503.2 457.3L365.8 320L503.1 182.6C515.6 170.1 515.6 149.8 503.1 137.3C490.6 124.8 470.3 124.8 457.8 137.3L320.5 274.7L183.1 137.4z"></path>
        </svg>
      </button>
      <div class="LeadFormSentModal__leadFormWrap">
        <div class="LeadFormSent">
          <svg
              class="LeadFormSent__successAnimation"
              viewBox="0 0 576 576"
              xmlns="http://www.w3.org/2000/svg"
              width="96"
              height="96"
              fill="currentColor">
            <path fill="currentColor" d="M320 576C178.6 576 64 461.4 64 320C64 178.6 178.6 64 320 64C461.4 64 576 178.6 576 320C576 461.4 461.4 576 320 576zM320 112C205.1 112 112 205.1 112 320C112 434.9 205.1 528 320 528C434.9 528 528 434.9 528 320C528 205.1 434.9 112 320 112zM390.7 233.9C398.5 223.2 413.5 220.8 424.2 228.6C434.9 236.4 437.3 251.4 429.5 262.1L307.4 430.1C303.3 435.8 296.9 439.4 289.9 439.9C282.9 440.4 276 437.9 271.1 433L215.2 377.1C205.8 367.7 205.8 352.5 215.2 343.2C224.6 333.9 239.8 333.8 249.1 343.2L285.1 379.2L390.7 234z"></path>
          </svg>
          <h2 class="LeadFormSent__title">Заявка отправлена</h2>
          <h3 class="LeadFormSent__subTitle">Мы свяжемся с вами в ближайшее время</h3>
          <div class="LeadFormSent__continueElementWrap">
            <button class="Button Button--primary Button--defaultSize" @click="continueAction">Продолжить</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  methods: {
    close() {
      this.$emit('close');
    },
    continueAction() {
      this.$emit('close');
    },
  },
};
</script>
<style scoped lang="scss">
.LeadFormSentModal {
  &__leadFormWrap {
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    display: flex;
    align-items: center;
  }
}

.LeadFormSent {
  text-align: center;
  padding: 0 24px;
  margin: 0 auto;

  &__successAnimation {
    margin: 0 auto 24px;
    color: $checkbox_color;
  }

  &__title {
    color: $block_title;
    font-size: 20px;
    line-height: 28px;
  }

  &__subTitle {
    padding-top: 6px;
    font-size: 14px;
    line-height: 18px;
    color: $block_sub_title;
  }

  &__continueElementWrap {
    margin: 32px auto 16px;
  }
}

@media (min-width: 1000px) {
  .LeadFormSentModal {
    &__leadFormWrap {
      position: static;
      display: block;
    }
  }

  .LeadFormSent {
    padding: 99px 40px 95px;

    &__continueElementWrap {
      max-width: 320px;
    }
  }
}
</style>