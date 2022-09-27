<script lang="ts" setup>
import { reactive } from 'vue'

defineProps<{
  messages: Array<BeamMessage>,
}>()
</script>

<template>
  <main>
    <div>Beam (:3333)</div>
    <ul class="log_box">
      <li v-for="msg in messages">
        <va-card square outlined>
          <va-card-title>Date here</va-card-title>
          <va-card-content>
            <pre class="payload">{{ JSON.stringify(JSON.parse(msg.Payload), null, 2) }}</pre>
          </va-card-content>
        </va-card>
      </li>
    </ul>
  </main>
</template>

<style scoped>
main {
  display: flex;
  height: auto;
  overflow: scroll;
  flex-direction: column;
}

.payload {
  text-align: left;
}
</style>
