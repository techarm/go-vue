<template>
    <div class="bg-gray-200">
        <app-header />
        <main class="main-content mt-7">
            <section>
                <div class="page-header align-items-start min-vh-100">
                    <div class="container">
                        <div v-if="ready" class="row justify-content-md-center my-3">
                            <div class="col-md-2">
                                <img class="img-fluid img-thumbnail" :src="`${imgPath}/${book.slug}.jpg`" alt="book cover">
                            </div>
                            <div class="col-md-6">
                                <h3 class="mt-3">{{book.title}}</h3>
                                <hr>
                                <p><strong>作者:</strong> {{book.author.author_name}}</p>
                                <p><strong>公開年:</strong> {{book.publication_year}}</p>
                                <p>{{book.description}}</p>
                            </div>
                        </div>
                        <div v-else class="row">Loading...</div>
                    </div>
                    <AppFooter />
                </div>
            </section>
        </main>
    </div>
</template>
<script>
import requests from '@/components/request.js';
import AppHeaderVue from '../AppHeader.vue';

export default {
    data() {
        return {
            book: {},
            imgPath: process.env.VUE_APP_IMAGE_URL + '/covers',
            ready: false
        }
    },
    components: {
        "app-header": AppHeaderVue
    },
    beforeMount() {
        requests.get("/books/" + this.$route.params.bookName).then(res => {
            if (res.error) {
                this.$emit("error", res.message);
            } else {
                this.book = res.data.book;
                this.ready = true;
            }
        }).catch(e => {
            console.log(e);
            this.$emit("error", "本データが取得できません");
        })
    }
}
</script>