import { message } from 'ant-design-vue'
import { FacebookAuthProvider, getAuth, getRedirectResult, signInWithPopup, GoogleAuthProvider, signOut as firebaseSignOut, } from 'firebase/auth'

export const useAuth = () => {
    const user = ref<any>()
    const auth = getAuth()
    const signInWithFacebook = () => {
        const provider = new FacebookAuthProvider();
        signInWithPopup(auth, provider).then((result) => {
            user.value = result.user
            console.log(user.value)
            message.success('Login success!')
        }).catch((err) => {
            if (err.code === FirebaseErrorCode.ACCOUNT_EXISTS) {
                message.error(FIREBASE_ERROR_MESSAGE[FirebaseErrorCode.ACCOUNT_EXISTS])
            }
        })
    }

    const signInWithGoogle = () => {
        const provider = new GoogleAuthProvider();
        signInWithPopup(auth, provider).then((result) => {
            user.value = result.user
            user.value.getIdToken().then((token: any) => {
                console.log(user.value, token)
                message.success('Login success!')
            })
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