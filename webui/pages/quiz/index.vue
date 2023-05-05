<script setup lang="ts">
const nuxtApp = useNuxtApp()
let origin: string | undefined = ''
if (process.server) {
    origin = nuxtApp.ssrContext?.event.node.req.headers.origin
} else {
    origin = window.location.origin
}
const url = origin + '/mock_data/questions.json'
const showAnswer = ref(false)

const currentNumberQuestion = ref(1)

const { data: questions } = await useFetch<any[]>(url)

const quizTitle = computed(() => `Question ${currentNumberQuestion.value} of ${questions.value?.length}`)
const question = computed(() => questions.value![currentNumberQuestion.value - 1])
const answers = computed(() => question?.value?.answer_list && question?.value?.answer_list![0].answers)
const disableBtnNext = computed(() => currentNumberQuestion.value === questions.value?.length)
const disableBtnPrev = computed(() => currentNumberQuestion.value <= 1)
const next = () => {
    if (currentNumberQuestion.value < questions.value!.length) {
        currentNumberQuestion.value++;
    }
}

const prev = () => {
    if (currentNumberQuestion.value > 1) {
        currentNumberQuestion.value--;
    }
}

</script>

<template>
    <div class="flex flex-col items-center mt-5">
        <a-card class="w-full max-w-screen-lg" :title="quizTitle">
            <template #extra>
                <a-button type="primary">Exit Quiz</a-button>
            </template>
            <div class="mb-4">
                <QuizQuestion :question="question.question_text" />
            </div>
            <div class="mb-4">
                <QuizAnswer :answers="answers" :key="currentNumberQuestion" />
            </div>
            <div class="mb-4">
                <a-button type="primary" class="!flex" @click="showAnswer = !showAnswer">
                    <template #icon>
                        <component :is="iconMap.eye" />
                    </template>
                    <span>Show answer</span>
                </a-button>
            </div>
            <QuizExplanation v-if="showAnswer" :explaination="question.general_feedback" />
            <div class="flex justify-between mt-8">
                <a-button class="!flex" @click="prev()" :disabled="disableBtnPrev">
                    <template #icon>
                        <component :is="iconMap.prev" />
                    </template>
                    <span>Prev</span>
                </a-button>
                <a-button class="!flex" @click="next()" :disabled="disableBtnNext">
                    <span>Next</span>
                    <template #icon>
                        <component :is="iconMap.next" />
                    </template>
                </a-button>
            </div>
        </a-card>
    </div>
</template>