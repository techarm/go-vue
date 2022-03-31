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

export default {
    name: 'AppLogin',
    data() {
        return {
            email: "",
            password: ""
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

            fetch('http://localhost:8081/users/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data)
            })
            .then(response => response.json())
            .then(data => {
                console.log(data);
            });
        }
    }
}
</script>