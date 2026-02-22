import HomePage from "@/components/HomePage.vue";
import ProductPage from "@/components/ProductPage.vue";
import CategoryPage from "@/components/CategoryPage.vue";
import DeliveryPage from "@/components/DeliveryPage.vue";
import PrivacyPage from "@/components/PrivacyPage.vue";
import ContactsPage from "@/components/ContactsPage.vue";
import CategoriesPage from "@/components/CategoriesPage.vue";
import NotFound from '@/components/errors/404.vue';
import {createRouter, createWebHistory} from "vue-router";


const routes = [
    {
        path: "/",
        name: "Home",
        component: HomePage,
        meta: { isHomePage: true }
    },
    {
        "path": "/products/:id",
        "name": "Product",
        "component": ProductPage,
        meta: { isHomePage: false },
    },
    {
        "path": "/categories/:id/products",
        "name": "Category",
        "component": CategoryPage,
        meta: { isHomePage: false },
    },
    {
        "path": "/categories",
        "name": "Categories",
        "component": CategoriesPage,
        meta: { isHomePage: false },
    },
    {
        "path": "/delivery",
        "name": "Delivery",
        "component": DeliveryPage,
        meta: { isHomePage: false },
    },
    {
        "path": "/privacy",
        "name": "Privacy",
        "component": PrivacyPage,
        meta: { isHomePage: false },
    },
    {
        "path": "/contacts",
        "name": "Contacts",
        "component": ContactsPage,
        meta: { isHomePage: false },
    },
    {
        path: '/:pathMatch(.*)*',
        name: 'not-found',
        component: NotFound,
        meta: { isHomePage: false },
    }
];

const router = createRouter({
    history: createWebHistory('/'),
    routes
});
router.afterEach((to) => {
    const baseUrl = 'https://shariki-v-permi.ru'
    const path = to.path

    let canonical = document.querySelector("link[rel='canonical']")
    if (!canonical) {
        canonical = document.createElement('link')
        canonical.setAttribute('rel', 'canonical')
        document.head.appendChild(canonical)
    }

    canonical.setAttribute('href', baseUrl + path)
})

export default router;