<script setup lang="ts">
const props = defineProps<{
    explaination: any
}>()
</script>

<template>
    <div class="mt-5 flex flex-col">
        <a-typography-text>Explanation:</a-typography-text>
        <div v-html="explaination" class="mt-4"></div>

    </div>
</template>