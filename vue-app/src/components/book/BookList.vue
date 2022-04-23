<template>
    <div class="bg-gray-200">
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
        <main class="main-content mt-0">
            <section>
                <div class="page-header align-items-start min-vh-100">
                    <div class="container">
                        <div class="row mt-7">
                            <div class="col">
                                <div class="text-center mb-3">
                                    <span class="filter" v-bind:class="{active: currentFilter === 0}" v-on:click="setFilter(0)">All</span>
                                    <span class="filter" v-bind:class="{active: currentFilter === 1}" v-on:click="setFilter(1)">Science Fiction</span>
                                    <span class="filter" v-bind:class="{active: currentFilter === 2}" v-on:click="setFilter(2)">Fanstasy</span>
                                    <span class="filter" v-bind:class="{active: currentFilter === 3}" v-on:click="setFilter(3)">Romance</span>
                                    <span class="filter" v-bind:class="{active: currentFilter === 4}" v-on:click="setFilter(4)">Thriller</span>
                                    <span class="filter" v-bind:class="{active: currentFilter === 5}" v-on:click="setFilter(5)">Mastery</span>
                                    <span class="filter" v-bind:class="{active: currentFilter === 6}" v-on:click="setFilter(6)">Horror</span>
                                    <span class="filter" v-bind:class="{active: currentFilter === 7}" v-on:click="setFilter(7)">Classic</span>
                                </div>
                            </div>
                        </div>
                        <div class="row">
                            <div class="col">
                                <div class="card-group">
                                <transition-group class="p-3 d-flex flex-warp" tag="div" appear name="books">
                                    <div v-for="book in books" :key="book.id">
                                            <div class="card me-2 ms-1 mb-3" style="width: 13rem;"
                                                v-if="book.genre_ids.includes(currentFilter) || currentFilter === 0">
                                                <router-link :to="`/books/${book.slug}`">
                                                    <img class="card-img-top" :src="`${this.imgPath}/${book.slug}.jpg`" :alt="book.name"/>
                                                </router-link>
                                                <div class="card-body text-center">
                                                    <h5 class="card-title">{{book.title}}</h5>
                                                    <p class="card-text">{{book.author.author_name}}</p>
                                                </div>
                                                <div class="card-footer text-center">
                                                    <small class="text-muted" v-for="(g, index) in book.genres" :key="g.id">
                                                        <em class="em-1">{{g.genre_name}}</em>
                                                        <template v-if="index !== (book.genres.length - 1)">, </template>
                                                    </small>
                                                </div>
                                            </div>
                                    </div>
                                </transition-group>
                                </div>
                            </div>
                        </div>
                    </div>
                    <AppFooter />
                </div>
            </section>
        </main>
    </div>
</template>

<script>
import {store, clearStoreAndCookie} from '@/components/store.js';
import requests from '@/components/request.js';
import AppFooter from '../AppFooter.vue';
import router from '@/router/index.js';

export default {
    name: "BookList",
    data() {
        return {
            store,
            books: [],
            imgPath: process.env.VUE_APP_IMAGE_URL + "/covers",
            currentFilter: 0,
            ready: false,
        }
    },
    components: {
        AppFooter,
    },
    emits: ['error'],
    beforeMount() {
        requests.get("/books").then(res => {
            if (res.error) {
                this.$emit('error', res.message);
            } else {
                this.books = res.data.books;
                this.ready = true;
            }
        }).catch(e => {
            console.log(e);
            this.$emit('error', "データ取得失敗しました。");
        })
    },
    methods: {
        setFilter(filter) {
            this.currentFilter = filter;
        },
        logout() {
            const payload = {
                token: store.token
            };
            requests.post("/users/logout", payload).then(res => {
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

<style scoped>
.filter {
    padding: 6px;
    cursor: pointer;
    border-radius: 6px;
    transition: all 0.35s;
    border: 1px solid silver;
    margin-right: 5px;
    line-height: 2.5rem;
}

.filter.active {
    background: lightgreen;
}

.filter:hover {
    background: lightgray;
}

/* transition styles */
.books-move {
    transition: all 0.5s ease-in-out 50ms;
}

.books-enter-active, .books-leave-active {
    transition: all 1.5s ease;
}

.books-enter-from, .books-leave-to {
    transform: translateX(-80px);
    opacity: 0;
}

</style>