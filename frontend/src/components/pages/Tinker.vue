<script lang="ts" setup>
import { reactive, onMounted, watch, ref } from 'vue'
import { Splitpanes, Pane } from 'splitpanes'
import 'splitpanes/dist/splitpanes.css'
import { ExecuteCommand, GetLastCode } from '@wails/go/tinker/Tinker'
import Editor from '@/components/Editor.vue'
import Button from '@/components/Button.vue'
import EditorTabBar from '@/components/EditorTabBar.vue'
import { registerPHPSnippetLanguage } from '@/registerPHPSnippetLanguage'

const props = withDefaults(defineProps<{
  lineWrap?: boolean,
  project?: Project,
}>(), {
  lineWrap: false,
})

const data = reactive({
  input: "",
  output: "",
})

const busy = ref(false)

function runTinker() {
  busy.value = true
  ExecuteCommand(data.input).then(result => {
    busy.value = false
    console.log('got result:', result)
    data.output = result
  }).catch(error => {
    busy.value = false
    console.log('got error when executing tinker command:', error)
    data.output = error
  })
}

function handleKeyboardShortcuts(e: KeyboardEvent): void {
  if (e.ctrlKey && e.code === 'KeyR') {
    runTinker()
  }
}

watch(() => props.project, (newValue, oldValue) => {
  if(oldValue && oldValue.Tinker && newValue.Tinker.ActiveTab !== oldValue.Tinker.ActiveTab) {
    refreshEditor()
  }
})

onMounted(() => {
  refreshEditor()
})

function refreshEditor() {
  busy.value = true
  let tabId = ""
  if (props.project && props.project.Tinker) {
    tabId = props.project.Tinker.ActiveTab
  }
  GetLastCode(tabId).then((editorContent) => {
    data.input = editorContent.Input
    data.output = editorContent.Output
    busy.value = false
  })
}

const handleMonacoBeforeMount = function (monaco) {
  registerPHPSnippetLanguage(monaco.languages)
}
</script>

<template>
  <main class="h-full pb-48" @keypress="handleKeyboardShortcuts">
    <EditorTabBar 
      :tabs="project.Tinker?.Tabs" 
      :activeTab="project.Tinker?.ActiveTab"
      module="tinker"
      :busy="busy"
    />
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
