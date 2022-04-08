<template>
    <div class="container">
        <div class="row">
            <div class="col">
                <h1 class="mt-3">User List</h1>
                <table class="table table-hover">
                    <thead>
                        <tr>
                            <th>User ID</th>
                            <th>User Name</th>
                            <th>Email</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="user in users" v-bind:key="user.id">
                            <td>{{ user.id }}</td>
                            <td>{{ user.first_name + ', ' + user.last_name }}</td>
                            <td>{{ user.email }}</td>
                        </tr>
                    </tbody>
                </table>
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
