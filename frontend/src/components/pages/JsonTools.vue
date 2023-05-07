<script lang="ts" setup>
import { reactive, ref, onMounted } from 'vue'
import { Splitpanes, Pane } from 'splitpanes'
import 'splitpanes/dist/splitpanes.css'
import { SetActiveJSONTool } from '@wails/go/main/App'
import { PrettifyJSON, GetLastCode, JSONToPHP, PHPToJSON } from '@wails/go/json/JSONTools'
import Editor from '@/components/Editor.vue'
import Button from '@/components/Button.vue'
import Toolbar from '@/components/Toolbar.vue'
import ToolbarButton from '@/components/ToolbarButton.vue'

const props = withDefaults(defineProps<{
  lineWrap?: boolean,
  project?: Project,
}>(), {
  lineWrap: false,
})

const data = reactive({
  input: "",
  output: "",
  indent: 2,
})

function runTool() {
  switch (props.project.JSONTools.ActiveTool) {
    case 'format':
      PrettifyJSON(data.indent, data.input).then(result => {
        data.output = result
      })
      break
    case 'jsontophp':
      JSONToPHP(data.input).then(result => {
        data.output = result
      })
      break
    case 'phptojson':
      PHPToJSON(data.input).then(result => {
        data.output = result
      })
      break
  }
}

function handleKeyboardShortcuts(e: KeyboardEvent): void {
  if (e.ctrlKey && e.code === 'KeyR') {
    runTool()
  }
}

function switchTool(newTool: string) {
  SetActiveJSONTool(newTool)
}

onMounted(() => {
  GetLastCode().then((editorContent) => {
    data.input = editorContent.Input
    data.output = editorContent.Output
  })
})
</script>

<template>
  <main class="h-full pb-48" @keypress="handleKeyboardShortcuts">
    <Toolbar>
      <span class="text-xs mr-2">Tool:</span>
      <ToolbarButton @click="switchTool('format')" :class="[project.JSONTools.ActiveTool === 'format' ? '' : '!bg-transparent']">Prettify</ToolbarButton>
      <ToolbarButton @click="switchTool('jsontophp')" :class="[project.JSONTools.ActiveTool === 'jsontophp' ? '' : '!bg-transparent']">JSON to PHP</ToolbarButton>
      <ToolbarButton @click="switchTool('phptojson')" :class="[project.JSONTools.ActiveTool === 'phptojson' ? '' : '!bg-transparent']">PHP to JSON</ToolbarButton>
    </Toolbar>
    <div class="input-wrapper | h-full">
      <splitpanes class="default-theme">
        <pane>
          <Editor class="code-editor"
            :value="data.input"
            :format-on-paste="false"
            language="text"
            @Change="(val, event) => data.input = val"
            :wrap="lineWrap"
          />
        </pane>
        <pane>
          <Editor class="code-editor"
            :value="data.output"
            language="php"
            :wrap="lineWrap"
          />
        </pane>
      </splitpanes>
    </div>

    <div class="text-center pt-2">
      <Button @click="runTool">Format</Button>
    </div>
  </main>
</template>
