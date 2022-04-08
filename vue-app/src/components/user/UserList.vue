<template>
    <div class="container">
        <div class="row">
            <div class="col">
                <h1 class="mt-3">User List</h1>
            </div>
        </div>
    </div>
</template>

<script>
import security from '../security.js';
import requests from '../request.js';

export default {
    data() {
        return {
            users: []
        }
    },
    beforeMount() {
        if (security.requireToken()) {
            requests.get("/admin/users").then(res => {
                if (res.error) {
                    alert(res.message);
                } else {
                    this.users = res.data.users;
                }
            }).catch(error => {
                alert(error);
            });
        }
    }
}
</script>
