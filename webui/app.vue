<script setup lang="ts">
import { Header } from 'ant-design-vue/lib/layout/layout';
import './main.scss'
</script>

<template>
  <div>
    <NuxtLayout>
      <LayoutHeader />
      <div class="px-5">
        <NuxtPage />
      </div>
    </NuxtLayout>
  </div>
</template>
