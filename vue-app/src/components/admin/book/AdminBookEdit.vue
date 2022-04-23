<template>
    <side-menu selectedMenu="bookAdd" />
    <main class="main-content position-relative max-height-vh-100 h-100 border-radius-lg ">
        <!-- Navbar -->
        <nav class="navbar navbar-main navbar-expand-lg px-0 mx-4 shadow-none border-radius-xl" id="navbarBlur" navbar-scroll="true">
            <div class="container-fluid py-1 px-3">
                <nav aria-label="breadcrumb">
                    <ol class="breadcrumb bg-transparent mb-0 pb-0 pt-1 px-0 me-sm-6 me-5">
                        <li class="breadcrumb-item text-sm"><a class="opacity-5 text-dark" href="javascript:;">書籍管理</a></li>
                        <li class="breadcrumb-item text-sm text-dark active" aria-current="page">書籍登録</li>
                    </ol>
                </nav>
            </div>
        </nav>
        <!-- End Navbar -->
        <div class="container-fluid pt-4">
            <div class="row">
                <div class="col-12">
                    <div class="card my-4">
                        <div class="card-header p-0 position-relative mt-n4 mx-3 z-index-2">
                            <div class="bg-gradient-primary shadow-primary border-radius-lg pt-4 pb-3">
                                <h6 class="text-white text-capitalize ps-3">書籍登録</h6>
                            </div>
                        </div>
                        <div class="card-body pb-2">
                            <div class="row">
                                <div class="col-md-8">
                                    <form-tag @bookEditEvent="submitHandler" name="bookForm" event="bookEditEvent">
                                        <div v-if="this.book.slug !== ''" class="mb-3">
                                            <img :src="`${this.imgPath}/covers/${this.book.slug}.jpg`" class="img-fluid img-thumbnail book-cover" alt="cover">
                                        </div>

                                        <div class="mb-5 form-group form-file-upload form-file-multiple">
                                            <label for="formFile" class="ms-0">カバー画像</label><br>
                                            <input v-if="this.book.id === 0" ref="coverInput" type="file" id="formFile"
                                                required accept="image/jpeg" @change="loadCoverImage">
                                            <input v-else ref="coverInput" type="file" id="formFile"
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

                                        <div class="input-group input-group-static mb-4">
                                            <label for="description">説明</label>
                                            <textarea required v-model="book.description" class="form-control" id="description" rows="3"
                                                placeholder="書籍の説明内容を入力してください。"></textarea>
                                        </div>

                                        <div class="input-group input-group-static mb-4">
                                            <label for="genres">ジャンル</label>
                                            <select ref="genres"
                                                id="genres"
                                                class="form-control"
                                                required
                                                size="7"
                                                v-model="this.book.genre_ids"
                                                multiple>
                                                <option v-for="g in this.genres" :value="g.value" :key="g.value">
                                                    {{g.text}}
                                                </option>
                                            </select>
                                        </div>

                                        <div class="float-start">
                                            <input type="submit" class="btn btn-success me-2" value="保存" />
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
                    </div>
                </div>
            </div>        
        </div>
    </main>
</template>

<script>
import security from '@/components/security.js';
import requests from '@/components/request.js';
import FormTag from '@/components/form/FormTag.vue';
import TextInput from '@/components/form/TextInput.vue';
import SelectInput from '@/components/form/SelectInput.vue';
import SiderMenu from '@/components/admin/SiderMenu.vue';
import router from '@/router/index.js';
import notie from 'notie';

export default {
    name: "AdminBookEdit",
    data() {
        return {
            editMode: false,
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
        'select-input': SelectInput,
         "side-menu": SiderMenu
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
                this.editMode = true;
            }
        }
    },
    methods: {
        submitHandler() {
            const payload = {
                id: this.book.id,
                title: this.book.title,
                author_id : parseInt(this.book.author_id, 10),
                publication_year: parseInt(this.book.publication_year, 10),
                description: this.book.description,
                cover: this.book.cover,
                slug: this.book.slug,
                genre_ids: this.book.genre_ids,
            }
            
            let response;
            if (this.editMode) {
                response = requests.put(`/admin/books/${payload.id}`, payload);
            } else {
                response = requests.post("/admin/books", payload);
            }
            
            response.then(res => {
                if (res.error) {
                    this.$emit('error', res.message);
                } else {
                    this.$emit('success', res.message);
                    router.push('/admin/books');
                }
            })
        },
        loadCoverImage() {
            const file = this.$refs.coverInput.files[0];
            const reader = new FileReader();
            reader.onloadend = () => {
                const base64String = reader.result;
                this.book.cover = base64String;
            }
            reader.readAsDataURL(file);
        },
        confirmDelete(id) {
            notie.confirm({
                text: "書籍情報を削除しましたが、よろしいですか？",
                submitText: "削除",
                submitCallback: () => {
                    requests.delete(`/admin/books/${id}`).then((res) => {
                        if (res.error) {
                            this.$emit('error', res.message);
                        } else {
                            this.$emit('success', res.message);
                            router.push("/admin/books");
                        }
                    })
                }
            })
        }
    }
}
</script>

<style scoped>
.book-cover {
    max-width: 10em;
}
</style>