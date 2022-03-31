import { createRouter, createWebHistory } from 'vue-router';
import AppBody from './../components/AppBody';
import AppLogin from './../components/AppLogin';

const routes = [
    {
        path: '/',
        name: 'Home',
        component: AppBody
    },
    {
        path: '/login',
        name: 'Login',
        component: AppLogin
    }
];

const router = createRouter({history: createWebHistory(), routes});
export default router;