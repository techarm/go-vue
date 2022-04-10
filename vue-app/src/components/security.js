import { store, clearStoreAndCookie } from './store'
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
    },

    checkToken: async function() {
        if (store.token !== '') {
            const payload = {
                token: store.token
            };

            await fetch(process.env.VUE_APP_API_URL + "/validate-token", this.requestOptions("POST", payload))
            .then(res => res.json())
            .then(res => {
                if (res.error) {
                    console.error(res.message);
                } else {
                    if (!res.data) {
                        clearStoreAndCookie();
                        return false;
                    }
                }
                return true;
            });
        }
        return true;
    }
}

export default security;