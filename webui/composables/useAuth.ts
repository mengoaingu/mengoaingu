import { FacebookAuthProvider, getAuth, getRedirectResult, signInWithPopup, GoogleAuthProvider, signOut as firebaseSignOut } from 'firebase/auth'

export const useAuth = () => {
    const user = ref<any>()
    const auth = getAuth()
    const signInWithFacebook = () => {
        const provider = new FacebookAuthProvider();
        signInWithPopup(auth, provider).then((result) => {
            user.value = result.user
            console.log(user.value)
        })
    }

    const signInWithGoogle = () => {
        const provider = new GoogleAuthProvider();
        signInWithPopup(auth, provider).then((result) => {
            user.value = result.user
            console.log(user.value)
        })
    }

    const signOut = () => {
        firebaseSignOut(auth).then(() => {
            user.value = null
        })
    }

    return {
        user,
        signInWithFacebook,
        signInWithGoogle,
        signOut,
    }
}