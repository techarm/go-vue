<template>
    <side-menu selectedMenu="userAdd" />
    <main class="main-content position-relative max-height-vh-100 h-100 border-radius-lg ">
        <!-- Navbar -->
        <nav class="navbar navbar-main navbar-expand-lg px-0 mx-4 shadow-none border-radius-xl" id="navbarBlur" navbar-scroll="true">
            <div class="container-fluid py-1 px-3">
                <nav aria-label="breadcrumb">
                    <ol class="breadcrumb bg-transparent mb-0 pb-0 pt-1 px-0 me-sm-6 me-5">
                        <li class="breadcrumb-item text-sm"><a class="opacity-5 text-dark" href="javascript:;">ユーザー管理</a></li>
                        <li class="breadcrumb-item text-sm text-dark active" aria-current="page">ユーザー登録</li>
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
                                <h6 class="text-white text-capitalize ps-3">ユーザー登録</h6>
                            </div>
                        </div>
                        <div class="card-body pb-0">
                            <div class="row">
                                <div class="col-md-8">
                                    <form-tag v-if="this.ready" @userEditEvent="submitHandler" name="userform" event="userEditEvent" class="p-3">
                                        <text-input
                                            v-model="user.last_name"
                                            type="text"
                                            required="true"
                                            label="氏名(姓)"
                                            :value="user.last_name"
                                            name="last-name"></text-input>

                                        <text-input
                                            v-model="user.first_name"
                                            type="text"
                                            required="true"
                                            label="氏名(名)"
                                            :value="user.first_name"
                                            name="first-name"></text-input>

                                        <text-input
                                            v-model="user.email"
                                            type="email"
                                            required="true"
                                            label="メールアドレス"
                                            :value="user.email"
                                            name="email"></text-input>

                                        <text-input
                                            v-if="this.user.id === 0"
                                            v-model="user.password"
                                            type="password"
                                            required="true"
                                            label="パスワード"
                                            :value="user.password"
                                            name="password"></text-input>

                                        <text-input
                                            v-else
                                            v-model="user.password"
                                            type="password"
                                            label="Password"
                                            help="パスワード変更しない場合は、入力する必要はありません。"
                                            :value="user.password"
                                            name="password"></text-input>

                                        <div class="row mb-5 input-group input-group-static">
                                            <label class="mb-2">ユーザー状態</label>
                                            <div>
                                                <div class="form-check form-check-inline px-0" v-for="item in userActiveRadios" v-bind:key="item.value">
                                                    <input class="form-check-input"
                                                        type="radio"
                                                        :id="'user-active' + item.value"
                                                        name="user-active"
                                                        :value="item.value"
                                                        v-model="user.active">
                                                    <label class="mx-2" :for="'user-active' + item.value">{{item.text}}</label>
                                                </div>
                                            </div>
                                        </div>

                                        <div class="float-start">
                                            <input type="submit" class="btn btn-success me-2" value="保存">
                                            <router-link to="/admin/users" class="btn btn-outline-secondary">キャンセル</router-link>
                                        </div>
                                        <div class="float-end">
                                            <a v-if="(this.$route.params.userId > 0) && (parseInt(String(this.$route.params.userId), 10) !== store.user.id)"
                                                class="btn btn-danger" href="javascript:void(0);" @click="confirmDelete(this.user.id)">削除</a>
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
import FormTag from '@/components/form/FormTag.vue';
import TextInput from '@/components/form/TextInput.vue';
import SiderMenu from '@/components/admin/SiderMenu.vue';
import security from '@/components/security.js';
import requests from '@/components/request.js';
import { store } from '@/components/store.js';
import router from '@/router/index.js';
import notie from 'notie';

export default {
    data() {
        return {
            userActiveRadios: [
                {value: "1", text: "有効"},
                {value: "0", text: "無効"},
            ],
            user: {
                id: 0,
                first_name: "",
                last_name: "",
                email: "",
                password: "",
                active: "0"
            },
            store,
            ready: false,
            editMode: false
        }
    },
    components: {
        'form-tag': FormTag,
        'text-input': TextInput,
        "side-menu": SiderMenu
    },
    beforeMount() {
        if (!security.requireToken()) {
            return;
        }
        
        if (parseInt(String(this.$route.params.userId), 10) > 0) {
            requests.get(`/admin/users/${this.$route.params.userId}`).then(res => {
                if (res.error) {
                    this.$emit("error", res.message);
                } else {
                    this.user = {
                        id: res.data.user.id,
                        first_name: res.data.user.first_name,
                        last_name: res.data.user.last_name,
                        email: res.data.user.email,
                        active: res.data.user.active ? "1" : "0",
                    };
                    this.ready = true;
                    this.editMode = true;
                }
            })
        } else {
            this.ready = true;
            this.editMode = false;
        }
    },
    methods: {
        submitHandler() {
            const payload = {
                id: parseInt(String(this.$route.params.userId), 10),
                first_name: this.user.first_name,
                last_name: this.user.last_name,
                email: this.user.email,
                password: this.user.password,
                active: this.user.active === "1"
            };

            let response;
            if (this.editMode) {
                response = requests.put(`/admin/users/${payload.id}`, payload)
            } else {
                response = requests.post("/admin/users", payload)
            }

            response.then(res => {
                if (res.error) {
                    this.$emit("error", res.message);
                } else {
                    this.$emit("success", res.message);
                    router.push("/admin/users");
                }
            });
        },
        confirmDelete(id) {
            notie.confirm({
                text: "ユーザーを削除しますが、よろしいですか？",
                submitText: "削除",
                submitCallback: () => {
                    requests.delete(`/admin/users/${id}`).then(res => {
                        if (res.error) {
                            this.$emit("error", res.message);
                        } else {
                            this.$emit("success", res.message);
                            router.push("/admin/users");
                        }
                    });
                }
            });
        }
    }
}
</script>
