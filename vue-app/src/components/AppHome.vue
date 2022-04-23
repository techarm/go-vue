<template>
    <div class="bg-gray-200">
        <app-header />
        <main class="main-content mt-0">
            <section>
                <div class="page-header align-items-center min-vh-100">
                    <div class="container my-auto">
                        <div class="row">
                            <div class="col">
                                <div class="card-group">
                                    <div class="p-3 d-flex flex-warp">
                                        <div v-for="book in books" :key="book.id">
                                            <div class="card me-2 ms-1 mb-3" style="width: 13rem;">
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
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <app-footer />
                </div>
            </section>
        </main>
    </div>
</template>

<script>
import AppFooter from '@/components/AppFooter.vue';
import AppHeader from '@/components/AppHeader.vue';
import requests from '@/components/request.js';

export default {
    name: "AppHome",
    components: {
        'app-header': AppHeader,
        'app-footer': AppFooter,
    },
    data() {
        return {
            books: [],
            imgPath: process.env.VUE_APP_IMAGE_URL + "/covers",
        }
    },
    emits: ['error'],
    beforeMount() {
        requests.get("/books").then(res => {
            if (res.error) {
                this.$emit('error', res.message);
            } else {
                this.books = res.data.books;
            }
        }).catch(e => {
            console.log(e);
            this.$emit('error', "データ取得失敗しました。");
        })
    }
}
</script>