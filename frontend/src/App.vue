<template>
  <Header @open-modal="openLeadModal"/>
  <router-view></router-view>
  <LeadPopUp @open-modal="openLeadModal"/>
  <LeadModal v-if="isModalShown" @close="closeLeadModal" @form-submitted="showSuccessModal"/>
  <LeadModalSent v-if="isLeadModalSentShown" @close="closeLeadSentModal"/>
</template>

<script>
import Header from './components/SiteHeader.vue';
import LeadPopUp from './components/Modal/LeadPopUp.vue';
import LeadModal from './components/Modal/LeadModal.vue';
import LeadModalSent from './components/Modal/LeadModalSent.vue';
import '@fortawesome/fontawesome-free/css/all.min.css';

export default {
  name: 'App',
  components: {
    Header,
    LeadPopUp,
    LeadModal,
    LeadModalSent
  },
  data() {
    return {
      isModalShown: false,
      isLeadModalSentShown: false,
      isProductModalShown: false
    };
  },
  methods: {
    openLeadModal() {
      this.isModalShown = true;
    },
    closeLeadModal() {
      this.isModalShown = false;
    },
    showSuccessModal() {
      this.isModalShown = false;
      this.isLeadModalSentShown = true;
    },
    closeLeadSentModal() {
      this.isLeadModalSentShown = false;
    },
  },
  provide() {
    return {
      openLeadModal: this.openLeadModal
    };
  },
}
</script>

<style>
*, ::before, ::after {
  font-family: 'Roboto', 'Open Sans', 'Noto Sans', 'PT Serif', 'Lato', 'Merriweather', sans-serif;
  -webkit-tap-highlight-color: transparent;
  box-sizing: border-box;
}

body, h1, h2, h3, h4, h5, h6, p, ul, li {
  font-weight: 400;
  -webkit-font-smoothing: subpixel-antialiased;
  -moz-osx-font-smoothing: auto;
  padding: 0;
  margin: 0;
}
.far, .far:before {
  font-family: 'FontAwesome', sans-serif;
}
</style>
