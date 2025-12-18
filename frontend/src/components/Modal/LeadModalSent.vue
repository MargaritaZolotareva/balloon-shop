<template>
  <div class="Modal Modal" aria-modal="true">
    <div class="Modal__backdrop" @click="close"></div>
    <div class="Modal__content">
      <button class="Button Button--rounded Button--defaultSize Modal__closeButton" @click="close">
        <i class="far fa-circle-xmark Modal__closeButton"></i>

      </button>
      <div class="LeadFormSentModal__leadFormWrap">
        <div class="LeadFormSent">
          <i class="far fa-circle-check LeadFormSent__successAnimation"></i>
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
    font-size: 56px;
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