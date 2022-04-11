<template>
    <div class="container">
        <div class="row">
            <div class="col">
                <h1 class="mt-3">本一覧</h1>                
            </div>
            <hr>
            <div class="col">
                <div class="card-group">
                   <div class="p-3 d-flex flex-warp">
                       <div v-for="book in books" :key="book.id">
                            <div class="card me-2 ms-1 mb-3" style="width: 13rem;"
                                v-if="book.genre_ids.includes(currentFilter) || currentFilter === 0">
                                <img class="card-img-top" :src="`${this.imgPath}/${book.slug}.jpg`" :alt="book.name"/>
                                <div class="card-body text-center">
                                    <h5 class="card-title">{{book.title}}</h5>
                                    <p class="card-text">{{book.author.author_name}}</p>
                                </div>
                            </div>
                       </div>
                   </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import {store} from '@/components/store.js';
import requests from '@/components/request.js';

export default {
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
    }
}
</script>