<template>
  <div class="Modal Modal" aria-modal="true">
    <div class="Modal__backdrop" @click="close"></div>
    <div class="Modal__content">
      <button class="Button Button--rounded Button--defaultSize Modal__closeButton"
              @click="close"
              aria-label="Закрыть модальное окно">
        <svg
            class="Modal__closeButton"
            viewBox="0 0 512 512"
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            fill="currentColor">
          <path fill="currentColor" d="M183.1 137.4C170.6 124.9 150.3 124.9 137.8 137.4C125.3 149.9 125.3 170.2 137.8 182.7L275.2 320L137.9 457.4C125.4 469.9 125.4 490.2 137.9 502.7C150.4 515.2 170.7 515.2 183.2 502.7L320.5 365.3L457.9 502.6C470.4 515.1 490.7 515.1 503.2 502.6C515.7 490.1 515.7 469.8 503.2 457.3L365.8 320L503.1 182.6C515.6 170.1 515.6 149.8 503.1 137.3C490.6 124.8 470.3 124.8 457.8 137.3L320.5 274.7L183.1 137.4z"></path>
        </svg>
      </button>
      <div class="LeadFormModal__leadFormWrap">
        <div class="Block LeadForm LeadFormModal__leadForm">
          <span class="Block__anchor" aria-hidden="true"></span>
          <h2 class="Block__title Block__title--center LeadForm__title">
            <div class="Block__titleWrap">Свяжемся и уточним все детали</div>
          </h2>
          <h3 class="Block__subTitle Block__subTitle--center">Оставьте свои контакты, мы ответим в&nbsp;ближайшее
            время</h3>
          <div class="Block__content Block__content--center LeadForm__content">
            <form class="LeadForm__form" @submit.prevent="submitForm">
              <label class="FormLayoutGroup">
                <div class="FormLayoutGroup__label">Номер телефона<span
                    class="FormLayoutGroup__labelRequiredChar"> *</span></div>
                <div class="FormLayoutGroup__content">
                  <input type="tel" placeholder=" " autocomplete="tel" class="Input"
                                                             value="" v-model="formData.phone">
                  <div v-if="errors.phone" class="form-error">{{ errors.phone }}</div>
                </div>
              </label>
              <label class="FormLayoutGroup">
                <div class="FormLayoutGroup__label">Имя<span
                    class="FormLayoutGroup__labelRequiredChar"> *</span></div>
                <div class="FormLayoutGroup__content">
                  <input type="text" placeholder=" " autocomplete="given-name"
                                                             class="Input" value="" v-model="formData.name">
                  <div v-if="errors.name" class="form-error">{{ errors.name }}</div>
                </div>
              </label>
              <label class="FormLayoutGroup">
                <div class="FormLayoutGroup__label">Сообщение</div>
                <div class="FormLayoutGroup__content">
                  <textarea placeholder=" " class="Textarea"
                                                                v-model="formData.message"></textarea>
                  <div v-if="errors.message" class="form-error">{{ errors.message }}</div>
                </div>
              </label>
              <label class="CheckBox CheckBox--default LeadForm__agreeCheckBox">
                <input class="CheckBox__input" type="checkbox"
                       v-model="formData.isChecked">
                <div class="CheckBox__indicator" aria-hidden="true">
                  <svg
                      v-if="formData.isChecked"
                      class="CheckBox__indicatorIcon"
                      viewBox="0 0 512 512"
                      xmlns="http://www.w3.org/2000/svg"
                      width="24"
                      height="24"
                      fill="currentColor">
                    <path fill="currentColor" d="M480 96C515.3 96 544 124.7 544 160L544 480C544 515.3 515.3 544 480 544L160 544C124.7 544 96 515.3 96 480L96 160C96 124.7 124.7 96 160 96L480 96zM160 144C151.2 144 144 151.2 144 160L144 480C144 488.8 151.2 496 160 496L480 496C488.8 496 496 488.8 496 480L496 160C496 151.2 488.8 144 480 144L160 144zM390.7 233.9C398.5 223.2 413.5 220.8 424.2 228.6C434.9 236.4 437.3 251.4 429.5 262.1L307.4 430.1C303.3 435.8 296.9 439.4 289.9 439.9C282.9 440.4 276 437.9 271.1 433L215.2 377.1C205.8 367.7 205.8 352.5 215.2 343.2C224.6 333.9 239.8 333.8 249.1 343.2L285.1 379.2L390.7 234z"></path>
                  </svg>
                </div>
                <!--                    TODO: добавить ссылку на политику конфиденциальности-->
                <span>Отправляя форму, вы соглашаетесь с <a href="privacy" target="_blank" rel="noopener noreferrer"
                                                            class="Link Link--default">политикой конфиденциальности</a></span>
                <div v-if="errors.isChecked" class="form-error">{{ errors.isChecked }}</div>
              </label>
              <button type="submit" :disabled="!formData.isChecked"
                      class="SectionButton"
                      aria-label="Отправить заявку">
                <div class="SectionButton__border">
                  <div class="SectionButton__background">
                    <div class="SectionButton__text">Оставить заявку</div>
                  </div>
                </div>
              </button>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import api from '@/services/api'

export default {
  name: 'LeadModal',
  data() {
    return {
      formData: {
        name: '',
        phone: '',
        message: '',
        isChecked: false
      },
      isSubmitting: false,
      errors: {},
      statusMessage: '',
      statusClass: '',
    };
  },
  methods: {
    close() {
      this.$emit('close');
    },
    async submitForm() {
      console.log("Метод обработчик вызван")
      if (!this.validateForm()) {
        return;
      }

      this.isSubmitting = true;
      try {
        const response = await api.post(`/lead-form`, this.formData);

        if (response.status === 200) {
          if (typeof window.ym === 'function') {
            window.ym(90974648, 'reachGoal', 'submit_form')
          }
          this.$emit('form-submitted');
        }
      } catch (error) {
        this.statusMessage = 'Произошла ошибка при отправке формы. Попробуйте снова.';
        this.statusClass = 'error';
      } finally {
        this.isSubmitting = false;
      }
    },
    validateForm() {
      this.errors = {};

      if (!this.formData.name) {
        this.errors.name = 'Имя обязательно';
      }

      const phoneRegex = /^\+?[0-9]{10,15}$/;
      if (!this.formData.phone || !phoneRegex.test(this.formData.phone)) {
        this.errors.phone = 'Введите корректный номер телефона';
      }

      console.log(this.errors)
      if (!this.formData.isChecked) {
        return false;
      }

      return Object.keys(this.errors).length === 0;
    },
  }
};
</script>
<style scoped lang="scss">
.form-error {
  color: red;
  font-size: 0.85rem;
  margin-top: 4px;
  line-height: 25px;
}
.LeadForm {
  padding: 80px 24px;
  margin-top: 0;
  background-color: $white;
  border-top: 1px solid $block_border;

  &__title {
    font-weight: 500;
    -webkit-font-smoothing: subpixel-antialiased;
    -moz-osx-font-smoothing: auto;
    line-height: 30px;
  }

  &__content {
    padding-top: 26px;
  }

  &__form {
    display: grid;
    gap: 26px;
    width: 100%;
    max-width: 600px;
    margin: 0 auto;
    padding: 0 24px 24px;
  }

  &__agreeCheckBox {
    align-items: flex-start;
    text-align: left;
    color: $gray_dark;
    font-size: 14px;
    line-height: 18px;
  }
}

.LeadFormModal {
  &__leadForm {
    padding: 0;
    border-top-width: 0;
  }
}

.Block {
  &__title {
    font-size: calc($section_font_size_offset + 24px);
    font-weight: 500;
    -webkit-font-smoothing: subpixel-antialiased;
    -moz-osx-font-smoothing: auto;
    color: $block_title;
    line-height: 30px;
    text-align: center;
    padding: 0 24px;
  }

  &__titleWrap, &__title {
    font-weight: $section_font_weight;
  }

  &__subTitle {
    display: none;
  }

  &__content {
    padding-top: 32px;
  }
}

.FormLayoutGroup {
  display: grid;
  gap: 8px;

  &__label {
    text-align: left;
    font-size: 14px;
    line-height: 18px;
    color: $form_group_label;
  }

  &__labelRequiredChar {
    color: $red;
  }
}

@media (min-width: 1000px) {
  .Block {
    padding: 0 34px;
    max-width: 1200px;
    margin: 0 auto;
    padding: 0;
    font-size: calc($section_font_size_offset + 34px);
    line-height: 44px;
    text-align: initial;

    &__title {
      max-width: 1200px;
      margin: 0 auto;
      padding: 0;
      font-size: calc($section_font_size_offset + 34px);
      line-height: 44px;
      text-align: initial;

      &--center {
        font-size: calc($section_font_size_offset + 24px);
        max-width: 540px;
        text-align: center;
        margin: 0 auto;
      }
    }

    &__subTitle {
      display: block;
      padding-top: 8px;
      font-size: 15px;
      line-height: 20px;
      color: $block_sub_title;

      &--center {
        text-align: center;
      }
    }

    &__content {
      padding-top: 40px;

      &--center {
        text-align: center;
      }
    }
  }

  .LeadFormModal {
    &__leadFormWrap {
      padding: 32px 24px 24px;
      width: 448px;
    }
  }

  .LeadForm {
    &__form {
      padding: 0;
    }
  }
}
</style>