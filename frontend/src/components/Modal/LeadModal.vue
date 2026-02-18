<template>
  <div class="Modal Modal" aria-modal="true">
    <div class="Modal__backdrop" @click="close"></div>
    <div class="Modal__content">
      <button class="Button Button--rounded Button--defaultSize Modal__closeButton" @click="close">
        <i class="far fa-xmark Modal__closeButton"></i>
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
                  <div v-if="errors.phone" class="error">{{ errors.phone }}</div>
                </div>
              </label>
              <label class="FormLayoutGroup">
                <div class="FormLayoutGroup__label">Имя<span
                    class="FormLayoutGroup__labelRequiredChar"> *</span></div>
                <div class="FormLayoutGroup__content">
                  <input type="text" placeholder=" " autocomplete="given-name"
                                                             class="Input" value="" v-model="formData.name">
                  <div v-if="errors.name" class="error">{{ errors.name }}</div>
                </div>
              </label>
              <label class="FormLayoutGroup">
                <div class="FormLayoutGroup__label">Сообщение</div>
                <div class="FormLayoutGroup__content">
                  <textarea placeholder=" " class="Textarea"
                                                                v-model="formData.message"></textarea>
                  <div v-if="errors.message" class="error">{{ errors.message }}</div>
                </div>
              </label>
              <label class="CheckBox CheckBox--default LeadForm__agreeCheckBox">
                <input class="CheckBox__input" type="checkbox"
                       v-model="formData.isChecked">
                <div class="CheckBox__indicator" aria-hidden="true">
                  <i v-if="formData.isChecked" class="far fa-square-check CheckBox__indicatorIcon"></i>
                </div>
                <!--                    TODO: добавить ссылку на политику конфиденциальности-->
                <span>Отправляя форму, вы соглашаетесь с <a href="privacy" target="_blank" rel="noopener noreferrer"
                                                            class="Link Link--default">политикой конфиденциальности</a></span>
                <div v-if="errors.isChecked" class="error">{{ errors.isChecked }}</div>
              </label>
              <button type="submit" :disabled="!formData.isChecked" class="SectionButton">
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
        const apiUrl = process.env.VUE_APP_API_URL;
        const response = await api.post(`${apiUrl}/lead-form`, this.formData);

        if (response.status === 200) {
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
.error {
  color: red;
  font-size: 0.85rem;
  margin-top: 4px;
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