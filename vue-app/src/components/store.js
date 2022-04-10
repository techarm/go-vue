import { reactive } from 'vue'

export const store = reactive({
    token: '',
    user: {}
});

export const clearStoreAndCookie = function() {
    store.token = "";
    store.user = {};
    document.cookie = "_site_data=; Path=/; SameSite=Strict; Secure; Expires=Thu, 01 Jan 1970 00:00:01 GMT;";
}