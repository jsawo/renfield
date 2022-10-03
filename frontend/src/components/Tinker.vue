<script lang="ts" setup>
import { reactive, onMounted } from 'vue'
import Editor from '@guolao/vue-monaco-editor'
import {
  ExecuteTinkerCommand,
  OpenDirectoryDialog,
  SetProjectDir,
  GetProjectDir,
  GetLastCode
} from '../../wailsjs/go/main/App'

const data = reactive({
  projectDir: "",
  input: "",
  output: "",
})

function runTinker() {
  ExecuteTinkerCommand(data.input).then(result => {
    data.output = result
  })
}

function openDirectoryDialog(): void {
  OpenDirectoryDialog().then((value: string) => {
    data.projectDir = value
    SetProjectDir(value)
  })
}

function handleKeyboardShortcuts(e: KeyboardEvent): void {
  if (e.ctrlKey && (e.code === 'Enter' || e.code === 'KeyR')) {
    runTinker()
  }
}

onMounted(() => {
  GetProjectDir().then((dir) => data.projectDir = dir)
  GetLastCode().then((code) => data.input = code)
})
</script>

<template>
  <main @keypress="handleKeyboardShortcuts">
    <div>Tinker ({{ data.projectDir }})</div>

    <div class="input-wrapper">
    <Editor
      id="input" class="code-editor text-box"
      :value="data.input"
      theme='vs-dark'
      defaultLanguage="php"
      @Change="(val, event) => data.input = val"
    />

    <Editor
      id="output" class="code-editor text-box"
      :value="data.output"
      theme='vs-dark'
      defaultLanguage="php"
    />

    </div>

    <div class="controls">
      <label>Project dir
        <va-button outline :rounded="false" class="mr-4" @click="openDirectoryDialog" size="small">Select</va-button>
      </label>
      <va-button :rounded="false" class="" @click="runTinker" size="large">Execute</va-button>
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
