import { IQuiz } from "~/lib/types"

export const useQuiz = () => {
    const quizList = ref<IQuiz[]>()

    return {
        quizList,
    }
}