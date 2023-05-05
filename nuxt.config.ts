// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    app: {
        head: {
            title: 'NocoAPI',
            link: [
                { rel: 'icon', type: 'image/x-icon', href: '/favicon.png' }
            ]
        }
    },
    ssr: false,
    srcDir: 'webui',
    devServer: {
        port: 3333
    },

    plugins: [
        '@/plugins/antd',
        '@/plugins/api',
    ],
    modules: [
        '@vueuse/nuxt',
        'nuxt-windicss',
        '@pinia/nuxt',
    ],
    windicss: {
        config: './windi.config.ts',
    },

})
