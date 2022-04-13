<template>
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
        <div class="container-fluid">
            <a class="navbar-brand" href="#">Navbar</a>
                <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
            </button>
            <div class="collapse navbar-collapse" id="navbarNav">
                <ul class="navbar-nav me-auto mb-2 mb-lg-0">
                    <li class="nav-item">
                        <router-link class="nav-link active" aria-current="page" to="/">Home</router-link>
                    </li>
                    <li class="nav-item">
                        <router-link class="nav-link active" to="/books">本</router-link>
                    </li>
                    <li v-if="store.token !== ''">
                        <a class="nav-link dropdown-toggle" href="#" id="navbarDropdown" role="button" data-bs-toggle="dropdown" aria-expanded="false">
                            管理
                        </a>
                        <ul class="dropdown-menu" aria-labelledby="navbarDropdown">
                            <li>
                                <router-link class="dropdown-item" to="/admin/books">本一覧</router-link>
                            </li>   
                            <li>
                                <router-link class="dropdown-item" to="/admin/users">ユーザー一覧</router-link>
                            </li>
                            <li>
                                <router-link class="dropdown-item" :to="{name: 'UserEdit', params: {userId : 0}}">ユーザー登録</router-link>
                            </li>
                        </ul>
                    </li>
                    <li class="nav-item ">
                        <router-link v-if="store.token === ''" class="nav-link" to="/login">ログイン</router-link>
                        <a href="javascript:void(0);" v-else class="nav-link" @click="logout">ログアウト</a>
                    </li>
                </ul>
                <span class="navbar-text">
                   {{ store.user.first_name ?? '' }}
                   {{ store.user.last_name ?? '' }}
                </span>
            </div>
        </div>
    </nav>
</template>

<script>
import {store, clearStoreAndCookie} from './store.js';
import router from './../router/index.js';
import request from './request.js';

export default {
    name: 'AppHeader',
    data() {
        return {
            store
        }
    },
    methods: {
        logout() {
            const payload = {
                token: store.token
            };
            request.post("/users/logout", payload).then(res => {
                console.log(res);
                // store.token = "";
                // store.user = {};
                // document.cookie = "_site_data=; Path=/; SameSite=Strict; Secure; Expires=Thu, 01 Jan 1970 00:00:01 GMT;";
                clearStoreAndCookie();
                router.push("/login")
            });
        }
    }
}
</script>
