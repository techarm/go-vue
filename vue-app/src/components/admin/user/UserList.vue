<template>
    <div class="container">
        <div class="row">
            <div class="col">
                <h1 class="mt-3">ユーザー情報一覧</h1>
                <hr>
                <table v-if="this.ready" class="table table-hover table-compact">
                    <thead>
                        <tr>
                            <th>ID</th>
                            <th>氏名</th>
                            <th>メールアドレス</th>
                            <th>有効／有無</th>
                            <th>状態</th>
                            <th>操作</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="user in users" v-bind:key="user.id">
                            <td>{{ user.id }}</td>
                            <td>{{ user.first_name + ', ' + user.last_name }}</td>
                            <td>{{ user.email }}</td>
                            <td v-if="user.active">
                                <span class="badge bg-success">有効</span>
                            </td>
                            <td v-else>
                                <span class="badge bg-danger">無効</span>
                            </td>
                            
                            <td v-if="user.has_token > 0">
                                <a href="javascript:void(0);">
                                    <span class="badge bg-success" @click="logUserOut(user.id)">ログイン中</span>
                                </a>
                            </td>
                            <td v-else>
                                <span class="badge bg-danger">未ログイン</span>
                            </td>
                            <td>
                                <router-link class="btn btn-primary btn-sm" :to="`/admin/users/${user.id}`">編集</router-link>
                            </td>
                        </tr>
                    </tbody>
                </table>

                 <p v-else>Loading...</p>
            </div>
        </div>
    </div>
</template>

<script>
import security from '@/components/security.js';
import requests from '@/components/request.js';
import { store } from '@/components/store.js';
import notie from 'notie';

export default {
    data() {
        return {
            users: [],
            ready: false,
            store
        }
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
