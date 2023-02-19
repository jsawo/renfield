<script lang="ts" setup>
import Editor from '@guolao/vue-monaco-editor'

withDefaults(defineProps<{
  value: string,
  language: string,
  formatOnPaste?: boolean,
}>(), {
  formatOnPaste: true,
})
</script>

<template>
  <Editor
      height="100%"
      :value="value"
      theme='vs-dark'
      :defaultLanguage="language"
      :options="{
        lineNumbers: 'off',
        wordWrap: 'off',
        minimap: {
          enabled: false,
        },
        scrollBeyondLastLine: false,
        fontSize: 16,
        padding: {top: 10, bottom: 10},
        mouseWheelZoom: true,
        formatOnPaste: formatOnPaste,
      }"
  />
</template>
