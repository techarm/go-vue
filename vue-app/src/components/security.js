import { store } from './store'
import router from '../router/index'

let security = {
    requireToken: function() {
        if (store.token === '') {
            router.push("/login");
            return false;
        }
        return true;
    },

    getOptions: function() {
        const headers = new Headers();
        headers.append("Content-Type", "application/json");
        headers.append("Authorization", "Bearer " + store.token);
        console.log(JSON.stringify(store));

        return {
            method: "GET",
            headers: headers
        }
    },

    postOptions: function(payload) {
        const headers = new Headers();
        headers.append("Content-Type", "application/json");
        headers.append("Authorization", "Bearer " + store.token);

        return {
            method: "POST",
            body: JSON.stringify(payload),
            headers: headers
        }
    }
}

export default security;