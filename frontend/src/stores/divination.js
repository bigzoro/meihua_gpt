import { defineStore } from 'pinia'
import client from '../api/client'

export const useDivinationStore = defineStore('divination', {
  state: () => ({
    history: [],
    currentDivination: null,
    loading: false,
    error: ''
  }),
  actions: {
    async fetchHistory() {
      this.loading = true
      this.error = ''
      try {
        const { data } = await client.get('/divinations')
        this.history = data
        if (data.length) {
          this.currentDivination = data[0]
        }
      } catch (error) {
        this.error = error.message || '获取历史失败'
      } finally {
        this.loading = false
      }
    },
    async createDivination(payload) {
      this.loading = true
      this.error = ''
      try {
        const { data } = await client.post('/divinations', payload)
        this.currentDivination = data
        this.history = [data, ...this.history]
      } catch (error) {
        this.error = error.message || '起卦失败'
        throw error
      } finally {
        this.loading = false
      }
    },
    selectFromHistory(id) {
      const found = this.history.find((item) => item.id === id)
      if (found) {
        this.currentDivination = found
      }
    }
  }
})
