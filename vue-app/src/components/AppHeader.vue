<template>
    <!-- Navbar -->
    <div class="container position-sticky z-index-sticky top-0">
        <div class="row">
            <div class="col-12">
                <!-- Navbar -->
                <nav class="navbar navbar-expand-lg blur border-radius-xl top-0 z-index-3 shadow position-absolute my-3 py-2 start-0 end-0 mx-4">
                    <div class="container-fluid ps-2 pe-0">
                        <router-link class="navbar-brand font-weight-bolder ms-lg-0 ms-3" to="/">Techarm Book Store</router-link>
                        <button class="navbar-toggler shadow-none ms-2" type="button" data-bs-toggle="collapse" data-bs-target="#navigation" aria-controls="navigation" aria-expanded="false" aria-label="Toggle navigation">
                            <span class="navbar-toggler-icon mt-2">
                                <span class="navbar-toggler-bar bar1"></span>
                                <span class="navbar-toggler-bar bar2"></span>
                                <span class="navbar-toggler-bar bar3"></span>
                            </span>
                        </button>
                        <div class="collapse navbar-collapse" id="navigation">
                            <ul class="navbar-nav mx-auto">
                                <li class="nav-item">
                                    <router-link class="nav-link d-flex align-items-center me-2" v-if="store.token !== ''" to="/admin">
                                        <i class="fa fa-chart-pie opacity-6 text-dark me-1"></i>管理
                                    </router-link>
                                </li>
                            </ul>
                            <ul class="navbar-nav d-lg-block d-none">
                                <li class="nav-item">
                                    <router-link v-if="store.token === ''" class="nav-link me-2" to="/login">
                                        <i class="fas fa-user-circle opacity-6 text-dark me-1"></i>ログイン
                                    </router-link>
                                    <a href="javascript:void(0);" v-else class="nav-link me-2" @click="logout">
                                        <i class="fas fa-user-slash opacity-6 text-dark me-1"></i>ログアウト
                                    </a>
                                </li>
                            </ul>
                        </div>
                    </div>
                </nav>
                <!-- End Navbar -->
            </div>
        </div>
    </div>
</template>

<script>
import {store, clearStoreAndCookie} from '@/components/store.js';
import router from '@/router/index.js';
import request from '@/components/request.js';

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
                clearStoreAndCookie();
                router.push("/login")
            });
        }
    }
}
</script>
