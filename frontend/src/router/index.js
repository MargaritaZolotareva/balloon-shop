import HomePage from "@/components/HomePage.vue";
import ProductPage from "@/components/ProductPage.vue";
import CategoryPage from "@/components/CategoryPage.vue";
import DeliveryPage from "@/components/DeliveryPage.vue";
import PrivacyPage from "@/components/PrivacyPage.vue";
import ContactsPage from "@/components/ContactsPage.vue";
import CategoriesPage from "@/components/CategoriesPage.vue";
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
    }
];

const router = createRouter({
    history: createWebHistory('/'),
    routes
});

export default router;