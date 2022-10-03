<script lang="ts" setup>
import { reactive, onMounted } from 'vue'
import Editor from '@guolao/vue-monaco-editor'
import { Splitpanes, Pane } from 'splitpanes'
import 'splitpanes/dist/splitpanes.css'
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
  <main class="contentwrapper" @keypress="handleKeyboardShortcuts">
    <div>Tinker ({{ data.projectDir }})</div>

    <div class="input-wrapper">
      <splitpanes class="default-theme">
        <pane>
          <Editor
            id="input" class="code-editor text-box"
            :value="data.input"
            theme='vs-dark'
            defaultLanguage="php"
            @Change="(val, event) => data.input = val"
          />
        </pane>
        <pane>
          <Editor
            id="output" class="code-editor text-box"
            :value="data.output"
            theme='vs-dark'
            defaultLanguage="php"
          />
        </pane>
      </splitpanes>
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
.contentwrapper {
  box-sizing: border-box;
  display: flex;
  gap: 1rem;
  flex-direction: column;
  height: 99%;
  padding: 1.5rem .5rem;
}

.input-wrapper {
  flex-grow: 1;
  height: 50%;
}

.controls {
  text-align: center;
}
</style>
