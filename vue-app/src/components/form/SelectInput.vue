<template>
    <div class="row mb-3">
        <label :for="name" class="col-sm-4 col-form-label">{{label}}</label>
        <div class="col-sm-8">
            <select 
                class="form-select"
                :id="name"
                :name="name"
                :required="required"
                :multiple="multiple"
                :value="modelValue"
                @input="$emit('update:modelValue', $event.target.value)">
                <option v-for="option in items" :value="option.value" :key="option.value">{{option.text}}</option>
            </select>
        </div>
    </div>
</template>

<script>
export default {
    name: "SelectInput",
    props: ["items", "name", "label", "required", "modelValue", "multiple"],
    emits: ['update:modelValue']
}
</script>