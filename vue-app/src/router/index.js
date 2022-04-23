import { createRouter, createWebHistory } from 'vue-router';
import AppLogin from '@/components/AdminLogin.vue';
import BookList from '@/components/book/BookList.vue';
import BookDetail from '@/components/book/BookDetail.vue';
import DashBoard from '@/components/admin/DashBoard.vue';
import UserList from '@/components/admin/user/UserList.vue';
import UserEdit from '@/components/admin/user/UserEdit.vue';
import AdminBookList from '@/components/admin/book/AdminBookList.vue';
import AdminBookEdit from '@/components/admin/book/AdminBookEdit.vue';
import security from '@/components/security';

const routes = [
    {
        path: '/',
        name: 'Home',
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
        path: "/admin",
        name: 'DashBoard',
        component: DashBoard
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
    },
    {
        path: "/admin/books",
        name: "AdminBookList",
        component: AdminBookList
    },
    {
        path: "/admin/books/:bookId",
        name: "AdminBookEdit",
        component: AdminBookEdit
    }
];

const router = createRouter({history: createWebHistory(), routes});

router.beforeEach(async () => {
    const canAccess = await security.checkToken();
    if (!canAccess) return '/login';
})
export default router;