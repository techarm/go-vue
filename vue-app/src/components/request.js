import security from "./security";

const API_URL = process.env.VUE_APP_API_URL;

let requests = {
    get: async function(uri) {
        const res = await fetch(API_URL + uri, security.requestOptions("GET"));
        return await res.json();
    },
    post: async function(uri, payload) {
        const res = await fetch(API_URL + uri, security.requestOptions("POST", payload));
        return await res.json();
    },
    put: async function(uri, payload) {
        const res = await fetch(API_URL + uri, security.requestOptions("PUT", payload));
        return await res.json();
    },
    delete: async function(uri) {
        const res = await fetch(API_URL + uri, security.requestOptions("DELETE"));
        return await res.json();
    }
}

export default requests;