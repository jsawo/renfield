<script lang="ts" setup>
import { reactive } from 'vue'
import { PrettifyJSON } from '../../wailsjs/go/main/App'

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

</script>

<template>
  <main>
    <div>JSON formatter</div>

    <div class="input-wrapper">
      <textarea id="input" class="text-box" v-model="data.input"
                placeholder="Paste JSON here">
      </textarea>

      <textarea id="output" class="text-box" v-model="data.output"></textarea>
    </div>

    <div class="controls">
      <label>Indentation<va-counter class="mx-4 my-2" v-model="data.indent" /></label>
      <va-button :rounded="false" class="" @click="runFormatter" size="large">Prettify</va-button>
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
