<template>
    <div class="container p-5">
        <div class="d-flex justify-content-center">
            <form-tag @loginevent="submitHandler" name="loginForm" event="loginevent">
                <text-input type="email" name="email" label="メールアドレス" required="true" v-model="email" />
                <text-input type="password" name="password" label="パスワード" required="true" v-model="password" />
                <div class="text-center mt-5">
                    <button type="submit" class="btn btn-success">ログイン</button>     
                </div>
            </form-tag>
        </div>
    </div>
</template>

<script>
import TextInput from './form/TextInput.vue';
import FormTag from './form/FormTag.vue';
import router from './../router/index.js';
import { store } from './store.js';

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
        TextInput,
        FormTag
    },
    methods: {
        submitHandler() {
            console.log(this.email, this.password);
            const data = {
                email: this.email,
                password: this.password
            };

            fetch(process.env.VUE_APP_API_URL + '/users/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data)
            })
            .then(res => res.json())
            .then(res => {
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