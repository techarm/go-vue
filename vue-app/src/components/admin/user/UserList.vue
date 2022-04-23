<template>
    <side-menu selectedMenu="userList" />
    <main class="main-content position-relative max-height-vh-100 h-100 border-radius-lg">
        <!-- Navbar -->
        <nav class="navbar navbar-main navbar-expand-lg px-0 mx-4 shadow-none border-radius-xl" id="navbarBlur" navbar-scroll="true">
            <div class="container-fluid py-1 px-3">
                <nav aria-label="breadcrumb">
                    <ol class="breadcrumb bg-transparent mb-0 pb-0 pt-1 px-0 me-sm-6 me-5">
                        <li class="breadcrumb-item text-sm"><a class="opacity-5 text-dark" href="javascript:;">ユーザー管理</a></li>
                        <li class="breadcrumb-item text-sm text-dark active" aria-current="page">ユーザー一覧</li>
                    </ol>
                </nav>
            </div>
        </nav>
        <!-- End Navbar -->
        <div class="container-fluid py-4">
            <div class="row">
                <div class="col-12">
                    <div class="card my-4">
                        <div class="card-header p-0 position-relative mt-n4 mx-3 z-index-2">
                            <div class="bg-gradient-primary shadow-primary border-radius-lg pt-4 pb-3">
                                <h6 class="text-white text-capitalize ps-3">ユーザー一覧</h6>
                            </div>
                        </div>
                        <div class="card-body">
                            <div class="table-responsive">
                                <table class="table table-hover mb-0">
                                    <thead>
                                        <tr>
                                            <th class="px-2">ID</th>
                                            <th>氏名</th>
                                            <th>メールアドレス</th>
                                            <th class="text-center">有効／有無</th>
                                            <th class="text-center">状態</th>
                                            <th class="text-center"></th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        <tr v-for="user in users" v-bind:key="user.id">
                                            <td>{{ user.id }}</td>
                                            <td>{{ user.first_name + ', ' + user.last_name }}</td>
                                            <td>{{ user.email }}</td>
                                            <td class="align-middle text-center">
                                                <span v-if="user.active" class="badge bg-gradient-success">有効</span>
                                                <span v-else class="badge bg-gradient-danger">無効</span>
                                            </td>
                                            <td class="align-middle text-center">
                                                <a v-if="user.has_token > 0" href="javascript:void(0);">
                                                    <span class="badge bg-gradient-success" @click="logUserOut(user.id)">ログイン中</span>
                                                </a>
                                                <span v-else class="badge bg-gradient-secondary">未ログイン</span>
                                            </td>
                                            <td class="align-middle text-center">
                                                <router-link class="btn btn-link py-0 my-0 text-secondary font-weight-bold text-xs" :to="`/admin/users/${user.id}`">
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
import security from '@/components/security.js';
import requests from '@/components/request.js';
import { store } from '@/components/store.js';
import SiderMenu from '@/components/admin/SiderMenu.vue';
import notie from 'notie';

export default {
    data() {
        return {
            users: [],
            ready: false,
            store
        }
    },
    components: {
        "side-menu": SiderMenu
    },
    beforeMount() {
        if (security.requireToken()) {
            requests.get("/admin/users").then(res => {
                if (res.error) {
                    this.$emit("error", res.message)
                } else {
                    this.users = res.data.users;
                    this.ready = true;
                }
            }).catch(error => {
                this.$emit("error", error)
            });
        }
    },
    methods: {
        logUserOut(id) {
            console.log(id);
            if (id !== store.user.id) {
                notie.confirm({
                    text: "対象ユーザーをログアウトさせますが、よろしいですか？",
                    submitText: "確定",
                    submitCallback: () => {
                        requests.post(`/admin/log-user-out/${id}`).then(res => {
                            if (res.error) {
                                this.$emit("error", res.message);
                            } else {
                                this.$emit("success", res.message);
                                this.$emit("forceUpdate");
                            }
                        })
                    }
                })
            } else {
                this.$emit("error", "自分をログアウトさせることができません");
            }
        }
    }
}
</script>
