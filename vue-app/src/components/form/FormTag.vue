<template>
    <form
        :ref="name"
        :method="method"
        :action="action"
        :event="event" 
        @submit.prevent="submit"
        autocomplete="off"
        class="needs-validation w-50"
        novalidate>
        <slot></slot>
    </form>
</template>

<script>
export default {
    name: 'FormTag',
    props: ["method", "action", "name", "event"],
    methods: {
        submit() {
            let form = this.$refs[this.$props.name];
            if (!form.checkValidity()) {
                form.classList.add('was-validated');
                console.log("form data is invaliadte.");
                return
            }
            this.$emit(this.$props.event);
        }
    }
}
</script>