<script lang="ts" setup>
import { reactive, onMounted } from 'vue'
import { Splitpanes, Pane } from 'splitpanes'
import 'splitpanes/dist/splitpanes.css'
import {
  ExecuteTinkerCommand,
  OpenDirectoryDialog,
  SetProjectDir,
  GetProjectDir,
  GetLastCode
} from '../../../wailsjs/go/main/App'
import Editor from '../Editor.vue'
import { registerPHPSnippetLanguage } from '../../registerPHPSnippetLanguage'


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
  if (e.ctrlKey && e.code === 'KeyR') {
    runTinker()
  }
}

onMounted(() => {
  GetProjectDir().then((dir) => data.projectDir = dir)
  GetLastCode().then((code) => data.input = code)
})

const handleMonacoBeforeMount = function (monaco) {
  registerPHPSnippetLanguage(monaco.languages)
}
</script>

<template>
  <main class="contentwrapper" @keypress="handleKeyboardShortcuts">
    <div>Tinker ({{ data.projectDir }})</div>

    <div class="input-wrapper">
      <splitpanes class="default-theme">
        <pane>
          <Editor class="code-editor text-box"
            :value="data.input"
            language="php-snippet"
            @Change="(val, event) => data.input = val"
            :onBeforeMount="handleMonacoBeforeMount"
          />
        </pane>
        <pane>
          <Editor class="code-editor text-box"
            :value="data.output"
            language="php"
          />
        </pane>
      </splitpanes>
    </div>

    <div class="controls">
      <label>Project dir:
        <w-button class="ma1" color="primary" @click="openDirectoryDialog" outline md>Select</w-button>
      </label>
      <w-button xl class="ma1" bg-color="primary" color="white" @click="runTinker">Execute</w-button>
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
  padding-bottom: 1.5rem;
}

.input-wrapper {
  flex-grow: 1;
  height: 50%;
}

.controls {
  text-align: center;
}
</style>
