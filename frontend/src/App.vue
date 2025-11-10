<template>
  <div class="page">
    <header class="hero">
      <h1>梅花易数助手</h1>
      <p>根据主题快速生成梅花易数卦象，记录每一次灵感。</p>
    </header>
    <main class="layout">
      <section class="column column--primary">
        <div class="card">
          <h2>起卦</h2>
          <DivinationForm :loading="loading" @submit="onCreate" />
          <p v-if="error" class="error">{{ error }}</p>
        </div>
        <HexagramDisplay v-if="currentDivination" class="card" :divination="currentDivination" />
      </section>
      <aside class="column column--secondary">
        <div class="card">
          <h2>占卦记录</h2>
          <DivinationHistory :items="history" :selected-id="currentDivination?.id" @select="onSelect" />
        </div>
      </aside>
    </main>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'
import { storeToRefs } from 'pinia'
import DivinationForm from './components/DivinationForm.vue'
import DivinationHistory from './components/DivinationHistory.vue'
import HexagramDisplay from './components/HexagramDisplay.vue'
import { useDivinationStore } from './stores/divination'

const store = useDivinationStore()
const { currentDivination, history, loading, error } = storeToRefs(store)

const onCreate = async (payload) => {
  await store.createDivination(payload)
}

const onSelect = (id) => {
  store.selectFromHistory(id)
}

onMounted(() => {
  store.fetchHistory()
})
</script>

<style scoped lang="scss">
.page {
  max-width: 1200px;
  margin: 0 auto;
  padding: 2rem 1.5rem 4rem;
}

.hero {
  text-align: center;
  margin-bottom: 2rem;

  h1 {
    font-size: 2.4rem;
    margin-bottom: 0.5rem;
  }

  p {
    margin: 0;
    color: #4b5563;
  }
}

.layout {
  display: flex;
  flex-wrap: wrap;
  gap: 1.5rem;
}

.column {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.column--primary {
  flex: 2;
  min-width: 320px;
}

.column--secondary {
  flex: 1;
  min-width: 260px;
}

.card {
  background: rgba(255, 255, 255, 0.9);
  border-radius: 16px;
  padding: 1.5rem;
  box-shadow: 0 20px 40px rgba(31, 41, 55, 0.08);
  backdrop-filter: blur(4px);
}

.error {
  color: #dc2626;
  margin-top: 1rem;
}

@media (max-width: 900px) {
  .layout {
    flex-direction: column;
  }
}
</style>
