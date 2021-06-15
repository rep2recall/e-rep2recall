import { api, initAPI } from '@/assets/api'
import Init from '@/components/Init.vue'
import { createApp } from 'vue'

async function main() {
  const { ok } = await initAPI()
  if (!ok) {
    return
  }

  const { searchParams } = new URL(location.href)
  const q = searchParams.get('q') || ''
  const files = searchParams.get('files') || ''

  const { data } = await api.post('/api/quiz/init', undefined, {
    params: {
      q,
      files,
      state: 'new,learning,due',
    },
  })

  createApp(Init, {
    type: 'Quiz',
    session: data.id,
    standalone: true,
  }).mount('#Quiz')
}

main()
