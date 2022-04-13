<template>
    <div class="container">
        <div class="row">
            <div class="col">
                <h1 class="mt-3">
                    <span v-if="editMode">ユーザー編集</span>
                    <span v-else>ユーザー登録</span>
                </h1>
                <hr>
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

                    <div class="row mb-3">
                        <label class="col-sm-4 col-form-label">ユーザー状態</label>
                        <div class="col-sm-8">
                            <div class="form-check form-check-inline" v-for="item in userActiveRadios" v-bind:key="item.value">
                                <input class="form-check-input"
                                    type="radio"
                                    :id="'user-active' + item.value"
                                    name="user-active"
                                    :value="item.value"
                                    v-model="user.active">
                                <label class="form-check-label" :for="'user-active' + item.value">{{item.text}}</label>
                            </div>
                        </div>
                    </div>

                    <hr>
                    <div class="float-start">
                        <input type="submit" class="btn btn-primary me-2" value="保存">
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
</template>

<script>
import FormTag from '@/components/form/FormTag.vue';
import TextInput from '@/components/form/TextInput.vue';
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
