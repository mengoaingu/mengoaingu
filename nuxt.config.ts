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
    ssr: true,
    srcDir: 'webui',
    devServer: {
        port: 3333,
        https: true,
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
