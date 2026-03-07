<template>
  <TopSection/>
  <AdvantagesMobile v-if="showAdvantages"/>
  <LazySection v-for="cat in categories" :key="cat.id" :categoryId="cat.id" />

  <ContentSection/>
<!--  <PhotosSection/>-->
<!--  <ReviewsSection/>-->
</template>

<script>
import TopSection from './LandingPage/TopSection.vue';
import ContentSection from './LandingPage/ContentSection.vue';
import LazySection from './LandingPage/LazySection.vue';
import {defineAsyncComponent} from "vue";
// import PhotosSection from './LandingPage/PhotosSection.vue';
// import ReviewsSection from './LandingPage/ReviewsSection.vue';

export default {
  name: 'App',
  components: {
    AdvantagesMobile: defineAsyncComponent(() => import('./LandingPage/AdvantagesMobile.vue')),
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
