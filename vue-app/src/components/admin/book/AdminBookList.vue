<template>
    <div class="container">
        <div class="row">
            <div class="col" v-if="this.ready">
                <h1 class="mt-3">本情報一覧</h1>
                <hr>
                <table class="table table-hover table-compact">
                    <thead>
                        <tr>
                            <th>ID</th>
                            <th>名前</th>
                            <th>作者</th>
                            <th>操作</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="book in books" v-bind:key="book.id">
                            <td>{{book.id}}</td>
                            <td>{{book.title}}</td>
                            <td>{{book.author.author_name}}</td>
                            <td>
                                <router-link class="btn btn-primary btn-sm" :to="`/admin/book/${book.id}`">編集</router-link>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
            <div class="col" v-else>Loading...</div>
        </div>
    </div>
</template>

<script>
import requests from '@/components/request.js';
import security from '@/components/security.js';

export default {
    data() {
        return {
            books: [],
            ready: false
        }
    },
    beforeMount() {
        if (security.requireToken()) {
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
    },
}
</script>