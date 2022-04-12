<template>
    <div class="container">
        <div class="row">
            <div class="col">
                <h1 class="mt-3">本一覧</h1>                
            </div>
            <hr>
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
            <hr>
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
</template>

<script>
import {store} from '@/components/store.js';
import requests from '@/components/request.js';

export default {
    name: "book-list",
    data() {
        return {
            store,
            books: {},
            imgPath: process.env.VUE_APP_IMAGE_URL + "/covers",
            currentFilter: 0,
            ready: false,
        }
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