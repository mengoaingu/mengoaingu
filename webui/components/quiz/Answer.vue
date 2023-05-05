<script setup lang="ts">
const props = defineProps<{
    answers: any[];
    modelValue?: any;
}>()

const vModel = useVModel(props, 'modelValue')

const value = ref<number | null>(null)

watch(value, () => {
    vModel.value = value;
})

watch(() => props.answers, () => {
    value.value = null
})

</script>

<template>
    <div class="flex flex-col bg-gray-100 px-4 py-5">
        <a-radio-group v-model:value="value">
            <a-radio :style="{ display: 'flex' }" v-for="(answer, index) in answers" :value="index" class="!mb-6">
                <a-typography-text>
                    <div v-html="answer.choice"></div>
                </a-typography-text>
            </a-radio>
        </a-radio-group>
    </div>
</template>