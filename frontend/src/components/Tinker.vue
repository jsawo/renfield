<script lang="ts" setup>
import { reactive, ref } from 'vue'
import { ExecuteTinkerCommand, OpenDirectoryDialog } from '../../wailsjs/go/main/App'

const data = reactive({
  projectDir: "",
  input: "User::factory()->make()",
  output: "",
})

// const flag

function runTinker() {
  ExecuteTinkerCommand(data.projectDir, data.input).then(result => {
    data.output = result
  })
}

function openDirectoryDialog(): void {
  OpenDirectoryDialog().then((value: string) => {
    data.projectDir = value
  })
}

function handleKeyboardShortcuts(e: KeyboardEvent): void {
  if (e.ctrlKey && (e.code === 'Enter' || e.code === 'KeyR')) {
    runTinker()
  }
}

</script>

<template>
  <main @keypress="handleKeyboardShortcuts">
    <div>Tinker {{ data.projectDir }}</div>

    <div class="input-wrapper">
      <textarea id="input" class="text-box" v-model="data.input" placeholder="Code here, eg: User::factory()->make()"></textarea>

      <textarea id="output" class="text-box" v-model="data.output"></textarea>
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
  padding: 1.5rem;
}

.input-wrapper {
  display: flex;
  flex-grow: 1;
  gap: 1rem;
}

.text-box {
  box-sizing: border-box;
  border-radius: 3px;
  outline: none;
  font-size: 1rem;
  line-height: 1.5rem;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
  height: 100%;
  flex-grow: 1;
  padding: 0.5rem;
  font-family: 'Courier New', Courier, monospace;
  white-space: pre;
  overflow-wrap: normal;
  overflow-x: scroll;
}

.text-box:hover {
  background-color: rgba(255, 255, 255, 1);
}
</style>
