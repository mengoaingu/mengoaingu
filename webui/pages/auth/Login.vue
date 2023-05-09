<script setup lang="ts">
import { FacebookAuthProvider, getAuth, getRedirectResult, signInWithPopup, GoogleAuthProvider } from 'firebase/auth'
const { $firebaseApp } = useNuxtApp()
const user = ref<any>()
const loginWithFacebook = () => {
    const provider = new FacebookAuthProvider();
    const auth = getAuth()
    signInWithPopup(auth, provider).then((result) => {
        user.value = result.user
        console.log(result.user)
    })
}

const loginWithGoogle = () => {
    const provider = new GoogleAuthProvider();
    const auth = getAuth()
    signInWithPopup(auth, provider).then((result) => {
        user.value = result.user
        console.log(result.user)
    })
}

</script>

<template>
    <a-button @click="loginWithFacebook">Login with facebook</a-button>
    <a-button @click="loginWithGoogle">Login with google</a-button>
    <div class="px-5" v-if="user">
        <a-typography-title>User Info</a-typography-title>
        <div class="flex flex-col">
            <a-typography-text>Uid: {{ user.uid }}</a-typography-text>
            <a-typography-text>Email: {{ user.email }}</a-typography-text>
            <a-typography-text>Display Name: {{ user.displayName }}</a-typography-text>
            
        </div>
    </div>
</template>