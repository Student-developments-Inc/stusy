import {createRouter, createWebHistory} from 'vue-router';
import HomeView from '../views/HomeView.vue';

const routes = [
    {
        path: '/',
        name: 'Главная',
        component: HomeView,
    },
    {
        path: '/auth',
        name: 'Вход',
        component: () => import(/* webpackChunkName: "about" */ '../views/AuthView.vue'),
    },
    {
        path: '/:pathMatch(.*)*',
        name: 'Не найдено',
        component: () => import(/* webpackChunkName: "about" */ '../views/404View.vue'),
    },
];

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes,
});

router.beforeEach((to, from, next) => {
    document.title = to.name;
    next();
});

export default router;
