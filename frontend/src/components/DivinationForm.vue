<template>
  <form class="form" @submit.prevent="handleSubmit">
    <label class="form__group">
      <span>占卦主题</span>
      <input v-model="form.subject" type="text" placeholder="请输入问题或主题" required />
    </label>
    <label class="form__group">
      <span>起卦方式</span>
      <select v-model="form.method">
        <option v-for="option in methodOptions" :key="option" :value="option">{{ option }}</option>
      </select>
    </label>
    <label class="form__group">
      <span>备注</span>
      <textarea v-model="form.notes" rows="3" placeholder="可记录背景、时间或灵感"></textarea>
    </label>
    <button class="form__submit" type="submit" :disabled="loading">
      {{ loading ? '占卜中…' : '立即起卦' }}
    </button>
  </form>
</template>

<script setup>
import { reactive } from 'vue'

defineProps({
  loading: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['submit'])

const methodOptions = ['随意起卦', '时间起卦', '数字起卦', '心念起卦']

const form = reactive({
  subject: '',
  method: methodOptions[0],
  notes: ''
})

const handleSubmit = () => {
  if (!form.subject.trim()) {
    return
  }
  emit('submit', { ...form })
  form.subject = ''
  form.notes = ''
}
</script>

<style scoped lang="scss">
.form {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.form__group {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  font-weight: 600;
  color: #374151;

  input,
  select,
  textarea {
    padding: 0.75rem 0.9rem;
    border: 1px solid rgba(15, 23, 42, 0.1);
    border-radius: 0.75rem;
    font-size: 1rem;
    background: rgba(255, 255, 255, 0.7);
    transition: border-color 0.2s ease, box-shadow 0.2s ease;

    &:focus {
      outline: none;
      border-color: #6366f1;
      box-shadow: 0 0 0 4px rgba(99, 102, 241, 0.1);
    }
  }
}

.form__submit {
  padding: 0.85rem;
  border: none;
  border-radius: 999px;
  font-weight: 600;
  background: linear-gradient(120deg, #6366f1, #8b5cf6);
  color: #fff;
  transition: transform 0.2s ease, box-shadow 0.2s ease;

  &:hover:enabled {
    transform: translateY(-1px);
    box-shadow: 0 12px 24px rgba(99, 102, 241, 0.24);
  }

  &:disabled {
    opacity: 0.6;
    cursor: not-allowed;
  }
}
</style>
