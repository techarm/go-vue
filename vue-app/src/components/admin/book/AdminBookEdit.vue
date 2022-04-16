<template>
    <div class="container">
        <div class="row p-3">
            <div class="col">
                <h1 class="mt-3">本編集</h1>
                <hr>

                <form-tag @bookEditEvent="submitHandler" name="bookForm" event="bookEditEvent">

                    <div v-if="this.book.slug !== ''" class="mb-3">
                        <img :src="`${this.imgPath}/covers/${this.book.slug}.jpg`" class="img-fluid img-thumbnail book-cover" alt="cover">
                    </div>

                    <div class="mb-3">
                        <label for="formFile" class="form-label">カバー画像</label>
                        <input v-if="this.book.id === 0" ref="coverInput" class="form-control" type="file" id="formFile"
                            required accept="image/jpeg" @change="loadCoverImage">
                        <input v-else ref="coverInput" class="form-control" type="file" id="formFile"
                            accept="image/jpeg" @change="loadCoverImage">
                    </div>

                    <text-input
                        v-model="book.title"
                        type="text"
                        required="true"
                        label="タイトル"
                        :value="book.title"
                        name="title"></text-input>

                    <select-input
                        name="author-id"
                        v-model="this.book.author_id"
                        :items="this.authors"
                        required="required"
                        label="作者"></select-input>

                    <text-input
                        v-model="book.publication_year"
                        type="number"
                        required="true"
                        label="発売日"
                        :value="book.publication_year"
                        name="publication-year"></text-input>

                    <div class="mb-3">
                        <label for="description" class="form-label">説明</label>
                        <textarea required v-model="book.description" class="form-control" id="description" rows="3"></textarea>
                    </div>

                    <div class="mb-3">
                        <label for="genres" class="form-label">ジャンル</label>
                        <select ref="genres"
                            id="genres"
                            class="form-select"
                            required
                            size="7"
                            v-model="this.book.genre_ids"
                            multiple>
                            <option v-for="g in this.genres" :value="g.value" :key="g.value">
                                {{g.text}}
                            </option>
                        </select>
                    </div>

                    <hr>

                    <div class="float-start">
                        <input type="submit" class="btn btn-primary me-2" value="保存" />
                        <router-link to="/admin/books" class="btn btn-outline-secondary">キャンセル</router-link>
                    </div>
                    <div class="float-end">
                        <a v-if="this.book.id > 0"
                            class="btn btn-danger" href="javascript:void(0);" @click="confirmDelete(this.book.id)">
                            削除
                        </a>
                    </div>
                    <div class="clearfix"></div>

                </form-tag>
            </div>
        </div>
    </div>
</template>

<script>
import security from '@/components/security.js';
import requests from '@/components/request.js';
import FormTag from '@/components/form/FormTag.vue';
import TextInput from '@/components/form/TextInput.vue';
import SelectInput from '@/components/form/SelectInput.vue';

export default {
    name: "AdminBookEdit",
    data() {
        return {
            book: {
                id: 0,
                title: "",
                author_id: 0,
                publication_year: null,
                description: "",
                cover: "",
                slug: "",
                genres: [],
                genre_ids: [],
            },
            authors: [],
            imgPath: process.env.VUE_APP_IMAGE_URL,
            genres: []
        }
    },
    components: {
        'form-tag': FormTag,
        'text-input': TextInput,
        'select-input': SelectInput
    },
    beforeMount() {
        if (security.requireToken()) {
            // 作者情報リストを取得
            requests.get("/admin/authors").then(res => {
                if (res.error) {
                    this.$emit("error", res.message);
                } else {
                    this.authors = res.data.authors.map(a => {
                        return {value: a.id, text: a.author_name}
                    })
                }
            });

            // ジャンル情報リストを取得
            requests.get("/admin/genres").then(res => {
                if (res.error) {
                    this.$emit("error", res.message);
                } else {
                    this.genres = res.data.genres.map(g => {
                        return {value: g.id, text: g.genre_name}
                    });
                }
            });

            if (this.$route.params.bookId > 0) {
                requests.get(`/admin/books/${this.$route.params.bookId}`).then(res => {
                    if (res.error) {
                        this.$emit("error", res.message);
                    } else {
                        this.book = res.data.book;
                        let genreArray = [];
                        for (let i = 0; i < this.book.genres.length; i++) {
                            genreArray.push(this.book.genres[i].id);
                        }
                        this.book.genres_ids = genreArray;
                    }
                })
            }
        }
    }
}
</script>

<style scoped>
.book-cover {
    max-width: 10em;
}
</style>