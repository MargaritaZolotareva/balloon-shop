<template>
  <TopSection @open-modal="openLeadModal"/>
  <AdvantagesMobile/>
  <ProductsSection :categoryId="51" :section="section"/>
  <ContentSection/>
  <PhotosSection/>
  <ReviewsSection/>
</template>

<script>
import AdvantagesMobile from './LandingPage/AdvantagesMobile.vue';
import TopSection from './LandingPage/TopSection.vue';
import ProductsSection from './LandingPage/ProductsSection.vue';
import ContentSection from './LandingPage/ContentSection.vue';
import PhotosSection from './LandingPage/PhotosSection.vue';
import ReviewsSection from './LandingPage/ReviewsSection.vue';

export default {
  name: 'App',
  components: {
    AdvantagesMobile,
    TopSection,
    ProductsSection,
    ContentSection,
    PhotosSection,
    ReviewsSection,
  },
  data() {
    return {
      section: {
        title: 'Шары на выписку'
      },
    };
  },
}
</script>
