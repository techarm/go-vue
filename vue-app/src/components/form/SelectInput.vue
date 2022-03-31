<template>
    <div class="row mb-3">
        <label :for="name" class="col-sm-3 col-form-label">{{label}}</label>
        <div class="col-sm-9">
            <select class="form-select" :id="name" :name="name" :required="required">
                <option v-for="option in items" :value="option.value" :key="option.value">{{option.text}}</option>
            </select>
        </div>
    </div>
</template>

<script>
export default {
    props: ["items", "name", "label", "required"]
}
</script>