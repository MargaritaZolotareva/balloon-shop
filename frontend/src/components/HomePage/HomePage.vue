<template>
  <TopSection/>
  <AdvantagesMobile v-if="showAdvantages"/>
  <LazySection v-for="cat in categories" :key="cat.id" :categoryId="cat.id" />

  <ContentSection/>
<!--  <PhotosSection/>-->
<!--  <ReviewsSection/>-->
</template>

<script>
import TopSection from './TopSection.vue';
import ContentSection from './ContentSection.vue';
import LazySection from './LazySection.vue';
import {defineAsyncComponent} from "vue";
// import PhotosSection from './PhotosSection.vue';
// import ReviewsSection from './ReviewsSection.vue';

export default {
  name: 'App',
  components: {
    AdvantagesMobile: defineAsyncComponent(() => import('./AdvantagesMobile.vue')),
    TopSection,
    ContentSection,
    LazySection
    // PhotosSection,
    // ReviewsSection,
  },
  mounted() {
    requestAnimationFrame(() => {
      this.showAdvantages = true
    })
  },
  data() {
    return {
      categories: [
        { id: 48 }, { id: 49 }, { id: 51 }, { id: 45 },
        { id: 46 }, { id: 50 }, { id: 44 }, { id: 47 },
      ],
      showAdvantages: false
    };
  },
}
</script>
