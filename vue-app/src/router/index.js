import { createRouter, createWebHistory } from 'vue-router';
import AppBody from '../components/AppBody.vue';
import AppLogin from '../components/AppLogin.vue';
import UserList from '../components/user/UserList.vue';

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
    },
    {
        path: "/users",
        name: 'Users',
        component: UserList
    }
];

const router = createRouter({history: createWebHistory(), routes});
export default router;