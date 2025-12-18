import HomePage from "@/components/HomePage.vue";
import ProductPage from "@/components/ProductPage.vue";
import CategoryPage from "@/components/CategoryPage.vue";
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
    }
];

const router = createRouter({
    history: createWebHistory('/'),
    routes
});

export default router;