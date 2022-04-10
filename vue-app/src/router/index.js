import { createRouter, createWebHistory } from 'vue-router';
import AppBody from '../components/AppBody.vue';
import AppLogin from '../components/AppLogin.vue';
import UserList from '../components/user/UserList.vue';
import UserEdit from '../components/user/UserEdit.vue';

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
        path: "/admin/users",
        name: 'UserList',
        component: UserList
    },
    {
        path: "/admin/users/:userId",
        name: 'UserEdit',
        component: UserEdit
    }
];

const router = createRouter({history: createWebHistory(), routes});
export default router;