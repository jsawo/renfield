<script lang="ts" setup>
import { reactive, onMounted } from 'vue'
import { Splitpanes, Pane } from 'splitpanes'
import 'splitpanes/dist/splitpanes.css'
import { PrettifyJSON, GetLastCode } from '@wails/go/json/JSONFormatter'
import Editor from '@/components/Editor.vue'
import Button from 'primevue/button'
import Slider from 'primevue/slider'

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
  if (e.ctrlKey && e.code === 'KeyR') {
    runFormatter()
  }
}

onMounted(() => {
  GetLastCode().then((editorContent) => {
    data.input = editorContent.Input
    data.output = editorContent.Output
  })
})
</script>

<template>
  <main class="content-wrapper" @keypress="handleKeyboardShortcuts">
    <div>JSON formatter</div>

    <div class="input-wrapper">
      <splitpanes class="default-theme">
        <pane>
          <Editor class="code-editor text-box"
            :value="data.input"
            :format-on-paste="false"
            language="json"
            @Change="(val, event) => data.input = val"
          />
        </pane>
        <pane>
          <Editor class="code-editor text-box"
            :value="data.output"
            language="json"
          />
        </pane>
      </splitpanes>
    </div>

    <div class="controls">
      <div class="input_indentation">
        <div class="flex flex-column">
          <span class="pb-1">Indent&nbsp;{{ data.indent }}</span> <Slider v-model="data.indent" :min="0" :max="8" />
        </div>
      </div>

      <Button label="Format" @click="runFormatter" />
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

.input_indentation {
  width: 6em;
  display: inline-block;
  padding: 0 1.5rem .5rem 0;
}
</style>
