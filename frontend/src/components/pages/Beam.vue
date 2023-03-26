<script lang="ts" setup>
import { computed } from 'vue'
import chroma from 'chroma-js'
import { CopyToClipboard } from '@wails/go/main/App'
import { notify } from "notiwind"
import Button from '@/components/Button.vue'
import Card from '@/components/Card.vue'

const port = '3333'

const props = defineProps<{
  messages: Array<BeamMessage>,
}>()

const messagesToShow = computed(() => {
  return props.messages.map((msg) => {
    msg.colorGroup = getGroupColor(msg.group)
    return msg
  })
})

function getGroupColor(group: string): string | null {
  var s = new Option().style
  s.color = group

  if (s.color !== '' && (s.color === group || s.color.startsWith('rgb'))) {
    return s.color
  }

  return null
}

function doCopy(data) {
  CopyToClipboard(data)
  notify({
    title: '',
    text: 'Copied to clipboard.',
    type: 'info',
    group: 'top-right',
  }, 3000)
}

function copyLaravelString() {
  const data = `\\Illuminate\\Support\\Facades\\Http::post("http://localhost:${port}/beam", ["payload"=> json_encode(["message" => "Hello World"])]);`
  doCopy(data)
}

function copyPHPString() {
  const data = `file_get_contents('http://localhost:${port}/beam', false, stream_context_create([
        'http' => ['method' => 'POST', 'header' => 'Content-Type: application/json',
            'content' => json_encode(['payload' => json_encode([
                'message' => 'Hello World',
        ])]),
    ]]));`
  doCopy(data)
}

function copyGOString() {
  const data = `client := &http.Client{}
    req, _ := http.NewRequest("POST", "http://localhost:${port}/beam", bytes.NewBuffer([]byte(\`{
	      "Payload": "{ \\"message\\": \\"Hello World\\"}"
    }\`)))
    res, _ := client.Do(req)
    defer res.Body.Close()
  ;`
  doCopy(data)
}

function copyCurlString() {
  const data = `curl -X POST -d '{"payload": "{ \\\"message\\\": \\\"Hello World\\\" }"}' http://localhost:${port}/beam`
  doCopy(data)
}

function getPayload(payload: string): string {
  try {
    return JSON.stringify(JSON.parse(payload), null, 2)
  } catch (e) {
    return payload
  }
}
</script>

<template>
  <main class="flex flex-col m-0 gap-2 overflow-auto h-full p-4  pb-24">
    <div class="pt-1 pr-1">
      Beam (running on port :{{ port }})
      <Button @click="copyLaravelString" class="mr-2">Laravel</Button>
      <Button @click="copyPHPString" class="mr-2">PHP</Button>
      <Button @click="copyGOString" class="mr-2">GO</Button>
      <Button @click="copyCurlString" class="">CURL</Button>

      <div style="float: right;">
        <Button @click="$emit('clearBeamMessages')">Clear</Button>
      </div>
    </div>

    <Card id="empty-state" v-if="messagesToShow.length === 0">
      <div>No messages received yet…</div>
    </Card>

    <ul class="log_box p-1 h-full overflow-y-scroll scroll-auto">
      <li v-for="(msg, index) in messagesToShow" :key="index" 
        class="first:mb-3 first:shadow-xl first:mx-0 mx-2 first:p-4 p-3 mb-3 bg-white shadow-sm"
      >
        <div class="flex">
          <div class="w-24 shrink-0 text-slate-600 text-sm overflow-scroll">
            <div class="pl-1">{{ msg.timestamp }}</div>
            <div v-if="msg.group">
              <div v-if="msg.colorGroup" 
                class="w-20 h-4 px-2 text-xs text-center mt-1 rounded-full border overflow-hidden" 
                :style="`background-color: ${msg.colorGroup}; 
                         border-color: ${chroma(msg.colorGroup).darken().hex()};
                         color: ${chroma(msg.colorGroup).luminance() > 0.4 ? '#000' : '#fff'};
                        `"
              >
                <span>{{ msg.group }}</span>
              </div>
              <div v-else class="pl-1">{{ msg.group }}</div>
            </div>
          </div>
          <pre class="w-full">{{ getPayload(msg.payload) }}</pre>
        </div>
      </li>
    </ul>
  </main>
</template>

<style scoped>
#empty-state {
  text-align: center;
  margin-top: 2rem;
}
</style>
