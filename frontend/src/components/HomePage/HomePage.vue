<template>
  <TopSection/>
  <AdvantagesMobile v-if="showAdvantages"/>
  <LazySection />

  <ContentSection/>
<!--  <PhotosSection/>-->
<!--  <ReviewsSection/>-->
</template>

<script setup>
import TopSection from './TopSection.vue';
import ContentSection from './ContentSection.vue';
import LazySection from './LazySection.vue';
import {ref, onMounted, defineAsyncComponent} from "vue";
// import PhotosSection from './PhotosSection.vue';
// import ReviewsSection from './ReviewsSection.vue';

const showAdvantages = ref(false);
const AdvantagesMobile = defineAsyncComponent(() => import('./AdvantagesMobile.vue'));

onMounted(() => {
  requestAnimationFrame(() => {
    showAdvantages.value = true;
  });
});
</script>
