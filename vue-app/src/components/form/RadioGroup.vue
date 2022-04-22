<template>
    <div class="row mb-3">
        <label :for="name" class="col-sm-4 col-form-label">{{label}}</label>
        <div class="col-sm-8">
            <div class="form-check form-check-inline" v-for="item in items" v-bind:key="item.value">
                <input class="form-check-input"
                    type="radio"
                    :id="name + item.value"
                    :name="name"
                    :value="item.value"
                    :checked="modelValue"
                    @input="$emit('update:modelValue', $event.target.value)">
                <label class="form-check-label" :for="name + item.value">{{item.text}}</label>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "RadioGroup",
    props: {
        name: String,
        required: String,
        label: String,
        items: Array,
        modelValue: String,
    }
}
</script>