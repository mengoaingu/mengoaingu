import { defineStore } from "pinia";

export const useQuizStore = defineStore('quiz', () => {
    const { $api } = useNuxtApp()
    const fetchQuizList = () => {
        $api.get('/quiz')
    }

    return {
        fetchQuizList,
    }
})