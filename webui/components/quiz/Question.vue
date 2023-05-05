<script setup lang="ts">

const props = defineProps<{
    question: string;
}>()

</script>

<template>
    <a-typography-title :level="5">
        <div v-html="question"></div>
    </a-typography-title>
</template>