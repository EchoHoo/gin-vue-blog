<template>
    <i v-if="theme" class="fa fa-sun-o" @click="setTheme"></i>
   <i v-else class="fa fa-moon-o" @click="setTheme"></i>

</template>

<script setup>
import {ref} from "vue";
import {useStore} from "@/stores/store.js";


const store = useStore()
const theme = store.theme// true代表白天 false代表黑夜
function setTheme(){
  store.setTheme()
}

</script>
<style>
</style>