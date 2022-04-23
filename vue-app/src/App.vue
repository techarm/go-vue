<template>
  <div>
    <router-view v-slot="{Component}" :key="componentKey" @success="success" @error="error" @warning="warning" @forceUpdate="forceUpdate">
      <keep-alive include="['BookList']">
        <component :is="Component" />
      </keep-alive>
    </router-view>
  </div>
</template>

<script>
import AppHeader from './components/AppHeader.vue';
import AppFooter from './components/AppFooter.vue';
import { store } from './components/store.js';
import notie from 'notie';
import BookList from './components/book/BookList.vue';

const getCookie = (name) => {
  return document.cookie.split("; ").reduce((r, v) => {
    const parts = v.split("=");
    return parts[0] === name ? decodeURIComponent(parts[1]) : r;
  }, "");
}

export default {
  name: 'App',
  components: {
    AppHeader,
    AppFooter,
    BookList
  },
  data() {
    return {
      store,
      componentKey: 0
    }
  },
  beforeMount() {
    let data = getCookie("_site_data");
    if (data !== "") {
      let cookieData = JSON.parse(data);
      
      // restore data to store
      store.token = cookieData.user.token.token;
      store.user = {
        id: cookieData.user.id,
        first_name: cookieData.user.first_name,
        last_name: cookieData.user.last_name,
        email: cookieData.user.email
      }
    }
  },
  methods: {
    success(msg) {
      notie.alert({
        type: "success",
        text: msg
      })
    },
    error(msg) {
      notie.alert({
        type: "error",
        text: msg
      })
    },
    warning(msg) {
      notie.alert({
        type: "warning",
        text: msg
      })
    },
    forceUpdate() {
      this.componentKey += 1;
      console.log("componentKey", this.componentKey)
    }
  }
}
</script>

<style>

</style>