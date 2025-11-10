<template>
  <section class="result">
    <header class="result__header">
      <h2>{{ divination.primaryName }}</h2>
      <p class="result__subtitle">变卦：{{ divination.secondaryName }}</p>
    </header>
    <div class="result__content">
      <div class="result__hexagrams">
        <div class="hexagram">
          <h3>本卦</h3>
          <ul>
            <li v-for="line in formattedLines" :key="line.position">
              <span class="hexagram__index">第{{ line.position }}爻</span>
              <span class="hexagram__symbol" :class="{ 'hexagram__symbol--changing': line.changing }">
                {{ line.symbol }}
              </span>
              <span class="hexagram__note">{{ line.note }}</span>
            </li>
          </ul>
        </div>
        <div class="hexagram__trigrams">
          <div>
            <h4>上卦</h4>
            <p>{{ divination.upperTrigram }}</p>
          </div>
          <div>
            <h4>下卦</h4>
            <p>{{ divination.lowerTrigram }}</p>
          </div>
        </div>
      </div>
      <article class="result__commentary">
        <pre>{{ divination.commentary }}</pre>
      </article>
    </div>
  </section>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  divination: {
    type: Object,
    required: true
  }
})

const formattedLines = computed(() => {
  const changing = new Set(props.divination.changingLines || [])
  return (props.divination.lines || []).map((value, index) => {
    const position = index + 1
    return {
      position,
      symbol: describeLine(value),
      note: lineNotes[value] || '—',
      changing: changing.has(position)
    }
  }).reverse()
})

const lineNotes = {
  6: '老阴，柔中带变',
  7: '少阳，刚健而升',
  8: '少阴，柔顺守中',
  9: '老阳，刚极生变'
}

const describeLine = (value) => {
  switch (value) {
    case 6:
      return '⚋'
    case 7:
      return '⚊'
    case 8:
      return '⚋'
    case 9:
      return '⚊'
    default:
      return '—'
  }
}
</script>

<style scoped lang="scss">
.result {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.result__header {
  h2 {
    margin: 0;
    font-size: 1.6rem;
  }

  .result__subtitle {
    margin: 0;
    color: #6b7280;
  }
}

.result__content {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.result__hexagrams {
  display: flex;
  flex-wrap: wrap;
  gap: 1.5rem;
  align-items: flex-start;
}

.hexagram {
  flex: 1;
  min-width: 220px;
}

.hexagram ul {
  list-style: none;
  margin: 0;
  padding: 0;
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.hexagram__index {
  font-weight: 600;
  color: #4b5563;
  width: 4.5rem;
}

.hexagram__symbol {
  font-size: 1.4rem;
  display: inline-block;
  width: 2rem;
}

.hexagram__symbol--changing {
  color: #d97706;
}

.hexagram__note {
  color: #6b7280;
  font-size: 0.9rem;
}

.hexagram__trigrams {
  display: grid;
  gap: 0.75rem;
}

.result__commentary {
  background: #f9fafb;
  border-radius: 14px;
  padding: 1rem;
  white-space: pre-wrap;
  line-height: 1.6;
  color: #374151;
}

pre {
  margin: 0;
  font-family: inherit;
}
</style>
