<template>
    <side-menu selectedMenu="bookList" />
    <main class="main-content position-relative max-height-vh-100 h-100 border-radius-lg ">
        <!-- Navbar -->
        <nav class="navbar navbar-main navbar-expand-lg px-0 mx-4 shadow-none border-radius-xl" id="navbarBlur" navbar-scroll="true">
            <div class="container-fluid py-1 px-3">
                <nav aria-label="breadcrumb">
                    <ol class="breadcrumb bg-transparent mb-0 pb-0 pt-1 px-0 me-sm-6 me-5">
                        <li class="breadcrumb-item text-sm"><a class="opacity-5 text-dark" href="javascript:;">書籍管理</a></li>
                        <li class="breadcrumb-item text-sm text-dark active" aria-current="page">書籍一覧</li>
                    </ol>
                </nav>
            </div>
        </nav>
        <!-- End Navbar -->
        <div class="container-fluid pt-4">
            <div class="row">
                <div class="col-12">
                    <div class="card mt-4">
                        <div class="card-header p-0 position-relative mt-n4 mx-3 z-index-2">
                            <div class="bg-gradient-primary shadow-primary border-radius-lg pt-4 pb-3">
                                <h6 class="text-white text-capitalize ps-3">書籍一覧</h6>
                            </div>
                        </div>
                        <div class="card-body">
                            <div class="table-responsive">
                                <table class="table table-hover mb-0">
                                    <thead>
                                        <tr>
                                            <th class="px-2">ID</th>
                                            <th>名前</th>
                                            <th>作者</th>
                                            <th class="text-secondary opacity-7"></th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        <tr v-for="book in books" v-bind:key="book.id">
                                            <td>{{book.id}}</td>
                                            <td>{{book.title}}</td>
                                            <td>{{book.author.author_name}}</td>
                                            <td>
                                                <router-link class="btn btn-link py-0 my-0 text-secondary font-weight-bold text-xs" :to="`/admin/books/${book.id}`">
                                                    <i class="material-icons text-sm me-2">edit</i>編集
                                                </router-link>
                                            </td>
                                        </tr>
                                    </tbody>
                                </table>
                            </div>
                        </div>
                    </div>
                </div>
            </div>        
        </div>
    </main>
</template>

<script>
import requests from '@/components/request.js';
import security from '@/components/security.js';
import SiderMenu from '@/components/admin/SiderMenu.vue';

export default {
    data() {
        return {
            books: [],
            ready: false
        }
    },
    components: {
        "side-menu": SiderMenu
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