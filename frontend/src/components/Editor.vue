<script lang="ts" setup>
import Editor from '@guolao/vue-monaco-editor'

defineProps<{
  value: string,
  language: string,
}>()

</script>

<template>
  <Editor
    :value="value"
    theme='vs-dark'
    :defaultLanguage="language"
    :options="{
      lineNumbers: 'off',
      wordWrap: 'on',
      minimap: {
        enabled: false,
      },
      scrollBeyondLastLine: false,
      renderWhitespace: 'none',
      fontSize: 16,
    }"
  />
</template>
