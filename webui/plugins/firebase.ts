import { initializeApp } from 'firebase/app'
import { getAnalytics } from "firebase/analytics";

export default defineNuxtPlugin(() => {
    const firebaseConfig = {
        apiKey: "AIzaSyD7isT38DUJ8V1PlTA0E7H0gQu_CzeMrNE",
        authDomain: "mengoaingu-vn.firebaseapp.com",
        projectId: "mengoaingu-vn",
        storageBucket: "mengoaingu-vn.appspot.com",
        messagingSenderId: "594058127517",
        appId: "1:594058127517:web:5f4c11f70708a8db648c74",
        measurementId: "G-J8NTG211PB"
    };
    const app = initializeApp(firebaseConfig)
    const analytics = getAnalytics(app);
    return {
        provide: {
            firebaseApp: app,
            firebaseAnalytics: analytics,
        }
    }
})