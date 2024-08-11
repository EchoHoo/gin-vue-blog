<template>
    <MdEditor v-model="text" />
</template>

<script setup>
import { reactive, ref } from 'vue';
import { MdEditor } from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';


</script>