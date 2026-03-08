<template>
  <Header @open-modal="openLeadModal"/>
  <router-view></router-view>
  <LeadPopUp @open-modal="openLeadModal"/>
  <LeadModal v-if="isModalShown" @close="closeLeadModal" @form-submitted="showSuccessModal"/>
  <LeadModalSent v-if="isLeadModalSentShown" @close="closeLeadSentModal"/>
  <Footer />
</template>

<script setup>
import Header from './components/Common/SiteHeader.vue';
import LeadPopUp from './components/Modal/LeadPopUp.vue';
import LeadModal from './components/Modal/LeadModal.vue';
import LeadModalSent from './components/Modal/LeadModalSent.vue';
import Footer from "@/components/Common/SiteFooter.vue";
import { ref, provide } from 'vue';

const isModalShown = ref(false);
const isLeadModalSentShown = ref(false);

function openLeadModal() {
  isModalShown.value = true;
}

function closeLeadModal() {
  isModalShown.value = false;
}

function showSuccessModal() {
  isModalShown.value = false;
  isLeadModalSentShown.value = true;
}

function closeLeadSentModal() {
  isLeadModalSentShown.value = false;
}

provide('openLeadModal', openLeadModal);
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
footer {
  margin-top: auto;
}
</style>
