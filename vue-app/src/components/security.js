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

    requestOptions: function(method, payload) {
        const headers = new Headers();
        headers.append("Content-Type", "application/json");
        headers.append("Authorization", "Bearer " + store.token);

        return {
            method: method,
            body: payload ? JSON.stringify(payload) : null,
            headers: headers
        }
    }
}

export default security;