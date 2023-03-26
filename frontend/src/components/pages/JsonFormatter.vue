<script lang="ts" setup>
import { reactive, onMounted } from 'vue'
import { Splitpanes, Pane } from 'splitpanes'
import 'splitpanes/dist/splitpanes.css'
import { PrettifyJSON, GetLastCode } from '@wails/go/json/JSONFormatter'
import Editor from '@/components/Editor.vue'
import Button from '@/components/Button.vue'

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
  <main class="h-full pb-40" @keypress="handleKeyboardShortcuts">
    <div class="input-wrapper | h-full">
      <splitpanes class="default-theme">
        <pane>
          <Editor class="code-editor"
            :value="data.input"
            :format-on-paste="false"
            language="json"
            @Change="(val, event) => data.input = val"
          />
        </pane>
        <pane>
          <Editor class="code-editor"
            :value="data.output"
            language="json"
          />
        </pane>
      </splitpanes>
    </div>

    <div class="text-center pt-2">
      <Button @click="runFormatter">Format</Button>
    </div>
  </main>
</template>
