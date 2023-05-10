<script setup>
const radioStyle = reactive({
    display: 'flex',
    height: '40px',
    lineHeight: '40px',
});

const answers = [
    {
        title: 'It can be used to understand the current state',

    },
    {
        title: 'It can be used to understand the current state',

    },
    {
        title: 'It can be used to understand the current state',

    },
    {
        title: 'It can be used to understand the current state',

    }
]
</script>

<template>
    <div class="flex flex-col items-center mt-5">
        <a-card class="w-full max-w-screen-lg" title="Question 1 of 20">
            <template #extra>
                <a-button type="primary">Exit Quiz</a-button>
            </template>
            <a-typography-title :level="5">Which of the following statements are true with respect to the Business
                Analysis?</a-typography-title>
            <div class="flex flex-col bg-gray-100 px-4 py-5">
                <a-radio-group v-model:value="value">
                    <a-radio :style="radioStyle" :value="1" v-for="answer in answers">
                        <a-typography-text>A. It can be used to understand the current state</a-typography-text>
                    </a-radio>
                </a-radio-group>
            </div>
            <div class="mt-5">
                <a-button type="primary" class="!flex">
                    <template #icon>
                        <component :is="iconMap.eye" />
                    </template>
                    <span>Show answer</span>
                </a-button>
            </div>
            <div class="mt-5 flex flex-col">
                <a-typography-text>Explanation:</a-typography-text>
                <div class="mt-4">
                    <a-typography-text>Choice D is the correct answer.</a-typography-text>
                    <a-typography-text>Business analysis may be performed within the boundaries of a project or throughout
                        enterprise evolution and continuous improvement. It can be used to</a-typography-text>
                </div>
            </div>
        </a-card>
    </div>
</template>

<style lang="scss">
.custom {
    @apply !flex !flex-col;
}
</style>
