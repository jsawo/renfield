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
        <w-card class="mt2">
          <div class="title5">Date here</div>
          <pre class="payload">{{ JSON.stringify(JSON.parse(msg.Payload), null, 2) }}</pre>
        </w-card>
      </li>
    </ul>
  </main>
</template>

<style scoped>
main {
  display: flex;
  height: auto;
  flex-direction: column;
}
</style>
