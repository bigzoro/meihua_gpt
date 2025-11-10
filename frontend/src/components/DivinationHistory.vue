<template>
  <div class="history">
    <p v-if="!items.length" class="history__empty">暂无记录，先起一卦试试吧。</p>
    <ul v-else class="history__list">
      <li
        v-for="item in items"
        :key="item.id"
        :class="['history__item', { 'history__item--active': item.id === selectedId }]"
        @click="() => emit('select', item.id)"
      >
        <div class="history__subject">{{ item.subject }}</div>
        <div class="history__meta">
          <span>{{ item.primaryName }}</span>
          <time>{{ formatDate(item.createdAt) }}</time>
        </div>
      </li>
    </ul>
  </div>
</template>

<script setup>
defineProps({
  items: {
    type: Array,
    default: () => []
  },
  selectedId: {
    type: Number,
    default: null
  }
})

const emit = defineEmits(['select'])

const formatDate = (value) => {
  if (!value) return ''
  const date = new Date(value)
  return date.toLocaleString('zh-CN', {
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit'
  })
}
</script>

<style scoped lang="scss">
.history {
  max-height: 520px;
  overflow-y: auto;
}

.history__empty {
  margin: 0;
  padding: 1rem;
  color: #6b7280;
  text-align: center;
}

.history__list {
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.history__item {
  border: 1px solid transparent;
  border-radius: 14px;
  padding: 0.85rem 1rem;
  background: rgba(248, 250, 252, 0.9);
  cursor: pointer;
  transition: border-color 0.2s ease, transform 0.2s ease;

  &:hover {
    border-color: rgba(99, 102, 241, 0.4);
    transform: translateY(-1px);
  }
}

.history__item--active {
  border-color: #6366f1;
  box-shadow: 0 12px 24px rgba(99, 102, 241, 0.2);
}

.history__subject {
  font-weight: 600;
  margin-bottom: 0.25rem;
}

.history__meta {
  display: flex;
  justify-content: space-between;
  font-size: 0.85rem;
  color: #6b7280;
}
</style>
