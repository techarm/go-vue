import security from "./security";

const API_URL = process.env.VUE_APP_API_URL;

let requests = {
    get: async function(uri) {
        const res = await fetch(API_URL + uri, security.getOptions());
        return await res.json();
    },

    post: async function(uri, payload) {
        const res = await fetch(API_URL + uri, security.postOptions(payload));
        return await res.json();
    }
}

export default requests;