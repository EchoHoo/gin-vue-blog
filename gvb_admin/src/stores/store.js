import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useStore = defineStore('gvb', {
  state:()=>{
    return {
      theme: true
    }
  },
  actions:{
    // 切换主题
    setTheme(){
      this.theme = !this.theme
    },
    // 加载主题
    loadTheme(){

    }
  }
})
