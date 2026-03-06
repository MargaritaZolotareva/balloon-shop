<template>
  <div ref="container">
    <component :is="component" v-if="visible" :categoryId="categoryId" />
  </div>
</template>

<script>
export default {
  props: ['categoryId', 'section'],
  data() {
    return { visible: false, component: null };
  },
  mounted() {
    const observer = new IntersectionObserver(
        ([entry]) => {
          if (entry.isIntersecting) {
            this.loadComponent();
            observer.disconnect();
          }
        },
        { threshold: 0.1 }
    );
    observer.observe(this.$refs.container);
  },
  methods: {
    async loadComponent() {
      this.component = (await import('./ProductsSection.vue')).default;
      this.visible = true;
    }
  }
}
</script>