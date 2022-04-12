import { createRouter, createWebHistory } from 'vue-router';
import AppBody from '../components/AppBody.vue';
import AppLogin from '../components/AppLogin.vue';
import BookList from '../components/book/BookList.vue';
import BookDetail from '../components/book/BookDetail.vue';
import UserList from '../components/user/UserList.vue';
import UserEdit from '../components/user/UserEdit.vue';
import security from '@/components/security';

const routes = [
    {
        path: '/',
        name: 'Home',
        component: AppBody
    },
    {
        path: '/books',
        name: 'BookList',
        component: BookList
    },
    {
        path: '/books/:bookName',
        name: 'BookDetail',
        component: BookDetail
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

router.beforeEach(async () => {
    const canAccess = await security.checkToken();
    if (!canAccess) return '/login';
})
export default router;