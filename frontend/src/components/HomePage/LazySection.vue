<template>
  <div ref="container">
    <template v-for="categoryData in categories" :key="categoryData.id">
      <component v-if="visible"
                 :is="component"
                 :categoryData="categoryData"/>
    </template>
  </div>
</template>

<script setup>
import {onMounted, ref} from 'vue';
import api from '@/services/api';

const container = ref(null);
const visible = ref(false);
const component = ref(null);
const categories = ref([]);

const loadComponent = async () => {
  component.value = (await import('./ProductsSection.vue')).default;

  try {
    const response = await api.get(`/homepage`);
    categories.value = response.data.categories;
  } catch (error) {
    console.error('Ошибка при получении категории для lazy:', error);
  }
  visible.value = true;
}
onMounted(() => {
  const observer = new IntersectionObserver(
      ([entry]) => {
        if (entry.isIntersecting) {
          loadComponent();
          observer.disconnect();
        }
      },
      {threshold: 0.1}
  );
  observer.observe(container.value);
});
</script>