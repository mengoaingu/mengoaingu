<script setup lang="ts">
const { user, signInWithFacebook, signInWithGoogle, signOut } = useAuth()
</script>

<template>
    <div class="shadow h-16 flex items-center justify-between px-5">
        <div>
            <AppLogo />
        </div>
        <div>
            <div v-if="user" class="flex items-center">
                <a-typography-text class="mr-3">{{ user?.displayName }}</a-typography-text>
                <a-dropdown :trigger="['click']">
                    <div class="flex items-center cursor-pointer">
                        <a-avatar :src="user?.photoURL" />
                        <component :is="iconMap.arrowDown" />
                    </div>
                    <template #overlay>
                        <a-menu>
                            <a-menu-item @click="signOut">
                                <template #icon>
                                    <component :is="iconMap.signout" />
                                </template>
                                Logout
                            </a-menu-item>
                        </a-menu>
                    </template>
                </a-dropdown>


            </div>
            <div v-else class="flex">
                <a-button @click="signInWithFacebook" class="!flex items-center mr-4">
                    <template #icon>
                        <img src="https://www.whizlabs.com/images/filled-facebook.svg" />
                    </template>
                    <span class="ml-2">Continue with facebook</span>
                </a-button>
                <a-button @click="signInWithGoogle" class="!flex items-center">
                    <template #icon>
                        <img src="https://www.whizlabs.com/images/filled-google.svg" />
                    </template>
                    <span class="ml-2">Continue with google</span>
                </a-button>

            </div>
        </div>
    </div>
</template>