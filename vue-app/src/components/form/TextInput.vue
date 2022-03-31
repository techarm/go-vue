<template>
    <div class="row mb-3">
        <label :for="name" class="col-sm-3 col-form-label">{{label}}</label>
        <div class="col-sm-9">
            <input
                :type="type"
                :name="name"
                :label="label"
                :placeholder="placeholder"
                :required="required"
                :min="min"
                :max="max"
                :value="modelValue"
                :autocomplete="name + '-new'"
                @input="$emit('update:modelValue', $event.target.value)"
                class="form-control">
        </div>
    </div>
</template>

<script>
export default {
    name: "TextInput",
    props: {
        name: String,
        type: String,
        label: String,
        placeholder: String,
        required: String,
        min: String,
        max: String,
        modelValue: String,
    }
}
</script>
