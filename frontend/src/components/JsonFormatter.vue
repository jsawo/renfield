<script lang="ts" setup>
import { reactive, onMounted } from 'vue'
import Editor from '@guolao/vue-monaco-editor'
import { PrettifyJSON, GetLastJSON } from '../../wailsjs/go/main/App'

const data = reactive({
  input: "",
  output: "",
  indent: 2,
})

function runFormatter() {
  PrettifyJSON(data.indent, data.input).then(result => {
    data.output = result
  })
}

function handleKeyboardShortcuts(e: KeyboardEvent): void {
  if (e.ctrlKey && (e.code === 'Enter' || e.code === 'KeyR')) {
    runFormatter()
  }
  if (e.ctrlKey && e.code === 'KeyZ') {
    document.execCommand("undo")
  }
  if (e.ctrlKey && e.shiftKey && e.code === 'KeyZ') {
    document.execCommand("redo")
  }
}

onMounted(() => {
  GetLastJSON().then((code) => data.input = code)
})
</script>

<template>
  <main @keypress="handleKeyboardShortcuts">
    <div>JSON formatter</div>

    <div class="input-wrapper">
      <Editor
        id="input" class="code-editor text-box"
        :value="data.input"
        theme='vs-dark'
        defaultLanguage="json"
        @Change="(val, event) => data.input = val"
      />

      <Editor
        id="input" class="code-editor text-box"
        :value="data.output"
        theme='vs-dark'
        defaultLanguage="json"
      />
    </div>

    <div class="controls">
      <label>Indentation<va-counter class="mx-4 my-2" v-model="data.indent" /></label>
      <va-button :rounded="false" class="" @click="runFormatter" size="large">Format</va-button>
    </div>
  </main>
</template>

<style scoped>
main {
  box-sizing: border-box;
  display: flex;
  gap: 1rem;
  flex-direction: column;
  height: 100%;
  padding: 1.5rem .5rem;
}

.input-wrapper {
  display: flex;
  flex-grow: 1;
  gap: .5rem;
}

.controls {
  text-align: center;
}
</style>
