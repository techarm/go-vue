<template>
    <div class="row mb-3">
        <div class="col-12">
            <div class="input-group input-group-static mb-4">
                <label :for="name">{{label}}</label>
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
                <p class="form-text">{{help}}</p>
            </div>
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
        help: String,
        modelValue: String,
    }
}
</script>
