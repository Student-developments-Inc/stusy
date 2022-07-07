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
        component: () => import('../views/AuthView.vue'),
    },
    {
        path:'/profile',
        name: 'Профиль',
        component: () => import('../views/ProfileView'),
    },
    {
        path:'/courses',
        name: 'Курсы',
        component: () => import('../views/MyCourses'),
    },
    {
        path:'/creatingCourses',
        name: 'Создание курса',
        component: () => import('../views/CreatingCourses'),
    },
    {
        path: '/:pathMatch(.*)*',
        name: 'Не найдено',
        component: () => import('../views/404View.vue'),
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
