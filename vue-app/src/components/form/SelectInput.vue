<template>
    <div class="row mb-5">
        <div class="col-12">
            <div class="input-group input-group-static">
                <label :for="name" class="ms-0">{{label}}</label>
                <select 
                    class="form-control"
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
    </div>
</template>

<script>
export default {
    name: "SelectInput",
    props: ["items", "name", "label", "required", "modelValue", "multiple"],
    emits: ['update:modelValue']
}
</script>