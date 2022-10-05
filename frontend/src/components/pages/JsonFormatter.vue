<script lang="ts" setup>
import { reactive, onMounted } from 'vue'
import { Splitpanes, Pane } from 'splitpanes'
import 'splitpanes/dist/splitpanes.css'
import { PrettifyJSON, GetLastCode } from '@wails/go/json/JSONFormatter'
import Editor from '@/components/Editor.vue'

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
  GetLastCode().then((code) => data.input = code)
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
        <w-input type="number" v-model="data.indent" outline min="0" max="8">Indentation</w-input>
      </div>

      <w-button xl class="ma1" bg-color="primary" color="white" @click="runFormatter">Format</w-button>
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
}
</style>
