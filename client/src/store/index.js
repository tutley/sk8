import Vuex from 'vuex'
import Vue from 'vue'

const state = {
  viewHeight: 0,
  isLoggedIn: false
}

const actions = {
  setViewHeight({ commit }, height) {
    commit('setViewHeight', height)
  },
  logout({ commit }) {
    localStorage.removeItem('token')
    commit('setIsLoggedIn', false)
  },
  setIsLoggedIn({ commit }, payload) {
    commit('setIsLoggedIn', payload)
  }
}

const mutations = {
  setIsLoggedIn(state, payload) {
    state.isLoggedIn = payload
  },
  setViewHeight(state, height) {
    state.viewHeight = height
  }
}

const getters = {
  isLoggedIn(state) {
    return state.isLoggedIn
  },
  viewHeight(state) {
    return state.viewHeight
  }
}

Vue.use(Vuex)

export default new Vuex.Store({
  state,
  actions,
  mutations,
  getters
})
