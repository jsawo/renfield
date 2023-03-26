<script lang="ts" setup>
import { reactive, onMounted } from 'vue'
import { Splitpanes, Pane } from 'splitpanes'
import 'splitpanes/dist/splitpanes.css'
import { ExecuteCommand, GetLastCode } from '@wails/go/tinker/Tinker'
import Editor from '@/components/Editor.vue'
import Button from '@/components/Button.vue'
import { registerPHPSnippetLanguage } from '@/registerPHPSnippetLanguage'

withDefaults(defineProps<{
  lineWrap?: boolean,
}>(), {
  lineWrap: false,
})

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
  GetLastCode().then((editorContent) => {
    data.input = editorContent.Input
    data.output = editorContent.Output
  })
})

const handleMonacoBeforeMount = function (monaco) {
  registerPHPSnippetLanguage(monaco.languages)
}
</script>

<template>
  <main class="h-full pb-40" @keypress="handleKeyboardShortcuts">
    <div class="input-wrapper | h-full">
      <splitpanes class="default-theme">
        <pane>
          <Editor class="code-editor text-box"
            :value="data.input"
            language="php-snippet"
            @Change="(val, event) => data.input = val"
            :onBeforeMount="handleMonacoBeforeMount"
            :wrap="lineWrap"
          />
        </pane>
        <pane>
          <Editor class="code-editor text-box"
            :value="data.output"
            language="php"
            :wrap="lineWrap"
          />
        </pane>
      </splitpanes>
    </div>

    <div class="text-center pt-2">
      <Button @click="runTinker">Execute</Button>
    </div>
  </main>
</template>
