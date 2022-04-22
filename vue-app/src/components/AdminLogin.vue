<template>
    <main class="main-content mt-0">
        <div class="page-header align-items-start min-vh-100" style="background-image: url('https://images.unsplash.com/photo-1497294815431-9365093b7331?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1950&q=80');">
            <span class="mask bg-gradient-dark opacity-6"></span>
            <div class="container my-auto">
                <div class="row">
                    <div class="col-lg-4 col-md-8 col-12 mx-auto">
                        <div class="card z-index-0 fadeIn3 fadeInBottom">
                            <div class="card-header p-0 position-relative mt-n4 mx-3 z-index-2">
                                <div class="bg-gradient-primary shadow-primary border-radius-lg py-3 pe-1">
                                    <h4 class="text-white font-weight-bolder text-center mt-2 mb-0">Techarm Admin</h4>
                                </div>
                            </div>
                            <div class="card-body">
                                <form-tag role="form" class="text-start" @loginevent="submitHandler" name="loginForm" event="loginevent">
                                    <div class="input-group input-group-outline my-3">
                                        <label class="form-label">Email</label>
                                        <input type="email" class="form-control" required="true" v-model="email">
                                    </div>
                                    <div class="input-group input-group-outline mb-3">
                                        <label class="form-label">Password</label>
                                        <input type="password" class="form-control" required="true" v-model="password">
                                    </div>
                                    <div class="text-center">
                                        <button type="submit" class="btn bg-gradient-primary w-100 my-4 mb-2">ログイン</button>
                                    </div>
                                </form-tag>
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
import router from '@/router/index.js';
import { store } from '@/components/store.js';
import requests from '@/components/request.js';

export default {
    name: 'AppLogin',
    data() {
        return {
            email: "",
            password: "",
            store
        }
    },
    components: {
        "form-tag": FormTag
    },
    methods: {
        submitHandler() {
            console.log(this.email, this.password);
            const data = {
                email: this.email,
                password: this.password
            };

            requests.post('/users/login', data).then(res => {
                if (res.error) {
                    alert(res.message)
                } else {
                    // save data to store
                    store.token = res.data.user.token.token;
                    store.user = {
                        id: res.data.user.id,
                        first_name: res.data.user.first_name,
                        last_name: res.data.user.last_name,
                        email: res.data.user.email
                    };

                    // save store data to cookie
                    let date = new Date();
                    let expDays = 1;
                    date.setTime(date.getTime() + (expDays * 24 * 60 * 60 * 1000));
                    const expires = "expires=" + date.toUTCString();
                    document.cookie = `_site_data=${JSON.stringify(res.data)}; Expires=${expires}; Path=/; SameSites=struct; Secure;`;

                    // redirect to index page
                    router.push("/");
                }
            });
        }
    }
}
</script>