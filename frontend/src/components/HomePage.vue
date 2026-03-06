<template>
  <TopSection/>
  <AdvantagesMobile/>
  <LazySection v-for="cat in categories" :key="cat.id" :categoryId="cat.id" />

  <ContentSection/>
<!--  <PhotosSection/>-->
<!--  <ReviewsSection/>-->
</template>

<script>
import AdvantagesMobile from './LandingPage/AdvantagesMobile.vue';
import TopSection from './LandingPage/TopSection.vue';
import ContentSection from './LandingPage/ContentSection.vue';
import LazySection from './LandingPage/LazySection.vue';
// import PhotosSection from './LandingPage/PhotosSection.vue';
// import ReviewsSection from './LandingPage/ReviewsSection.vue';

export default {
  name: 'App',
  components: {
    AdvantagesMobile,
    TopSection,
    ContentSection,
    LazySection
    // PhotosSection,
    // ReviewsSection,
  },
  data() {
    return {
      categories: [
        { id: 48 }, { id: 49 }, { id: 51 }, { id: 45 },
        { id: 46 }, { id: 50 }, { id: 44 }, { id: 47 },
      ]
    };
  },
}
</script>
