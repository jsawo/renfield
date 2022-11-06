<script lang="ts" setup>
import { reactive, onMounted } from 'vue'
import { Splitpanes, Pane } from 'splitpanes'
import 'splitpanes/dist/splitpanes.css'
import { ExecuteCommand, GetLastCode } from '@wails/go/tinker/Tinker'
import Editor from '@/components/Editor.vue'
import { registerPHPSnippetLanguage } from '@/registerPHPSnippetLanguage'
import Button from 'primevue/button'

const data = reactive({
  projectDir: "",
  input: "",
  output: "",
})

function runTinker() {
  ExecuteCommand(data.input).then(result => {
    data.output = result
  })
}

function handleKeyboardShortcuts(e: KeyboardEvent): void {
  if (e.ctrlKey && e.code === 'KeyR') {
    runTinker()
  }
}

onMounted(() => {
  GetLastCode().then((code) => data.input = code)
})

const handleMonacoBeforeMount = function (monaco) {
  registerPHPSnippetLanguage(monaco.languages)
}
</script>

<template>
  <main class="content-wrapper" @keypress="handleKeyboardShortcuts">
    <div>Tinker</div>

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
      <Button label="Execute" @click="runTinker" />
    </div>
  </main>
</template>

<style scoped>
.content-wrapper {
  box-sizing: border-box;
  display: flex;
  gap: 1rem;
  flex-direction: column;
  height: 100%;
}

.input-wrapper {
  flex-grow: 1;
  height: 50%;
}

.controls {
  text-align: center;
}
</style>
